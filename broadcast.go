// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/memberlist"
)

// Broadcast struct
type Broadcast struct {
	msg    []byte
	notify chan<- struct{}
}

// Invalidates ..
func (b *Broadcast) Invalidates(other memberlist.Broadcast) bool {
	return false
}

// Message ..
func (b *Broadcast) Message() []byte {
	return b.msg
}

// Finished ..
func (b *Broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}

// Delegate struct
type Delegate struct {
	State      []byte
	Broadcasts *memberlist.TransmitLimitedQueue
	Cluster    *Cluster
}

// NotifyMsg ..
func (d *Delegate) NotifyMsg(msg []byte) {
	fmt.Printf(" === Received Broadcast of Remote State %s === \n", string(msg))

	d.State = msg
}

// NodeMeta ..
func (d *Delegate) NodeMeta(limit int) []byte {
	return []byte{}
}

// LocalState ..
func (d *Delegate) LocalState(join bool) []byte {
	fmt.Println(" === Sharing Remote State for push/pull sync === ")

	return d.State
}

// GetBroadcasts ..
func (d *Delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return d.Broadcasts.GetBroadcasts(overhead, limit)
}

// MergeRemoteState ..
func (d *Delegate) MergeRemoteState(buf []byte, join bool) {
	fmt.Printf(" === Merging Remote State %s for push/pull sync === \n", string(buf))

	d.State = buf
}

// SetCluster ..
func (d *Delegate) SetCluster(cluster *Cluster) {
	d.Cluster = cluster

	d.Broadcasts = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return d.Cluster.Memlist.NumMembers()
		},
		RetransmitMult: 3,
	}
}

// UpdateState ..
func (d *Delegate) UpdateState(data []byte) {
	d.Broadcasts.QueueBroadcast(&Broadcast{
		msg:    data,
		notify: nil,
	})
}

// Message struct
type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Bytes ..
func (m *Message) Bytes() []byte {
	data, err := json.Marshal(m)

	if err != nil {
		return []byte("")
	}

	return data
}

// Load ..
func (m *Message) Load(data []byte) error {

	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	return nil
}
