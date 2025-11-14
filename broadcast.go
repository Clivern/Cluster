// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/memberlist"
)

// Broadcast implements the memberlist.Broadcast interface for state synchronization.
// It wraps a message and optional notification channel.
type Broadcast struct {
	msg    []byte
	notify chan<- struct{}
}

// Invalidates determines if this broadcast invalidates another broadcast.
// Returns false to allow all broadcasts to be processed.
func (b *Broadcast) Invalidates(_ memberlist.Broadcast) bool {
	return false
}

// Message returns the broadcast message bytes.
func (b *Broadcast) Message() []byte {
	return b.msg
}

// Finished is called when the broadcast has been sent to all nodes.
// It closes the notification channel if one was provided.
func (b *Broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}

// Delegate implements the memberlist.Delegate interface for cluster state management.
// It handles state synchronization, broadcasting, and merging remote state.
type Delegate struct {
	State      []byte
	Broadcasts *memberlist.TransmitLimitedQueue
	Cluster    *Cluster
	Logger     Logger
}

// NewDelegate creates a new Delegate instance with default logging.
func NewDelegate() *Delegate {
	return &Delegate{
		Logger: log.New(os.Stdout, "[cluster] ", log.LstdFlags),
	}
}

// NotifyMsg is called when a broadcast message is received from another node.
func (d *Delegate) NotifyMsg(msg []byte) {
	if d.Logger != nil {
		d.Logger.Printf("Received broadcast of remote state: %s", string(msg))
	} else {
		fmt.Printf(" === Received Broadcast of Remote State %s === \n", string(msg))
	}

	d.State = msg
}

// NodeMeta returns metadata about this node.
// Returns empty bytes by default, can be overridden for custom metadata.
func (d *Delegate) NodeMeta(_ int) []byte {
	return []byte{}
}

// LocalState returns the local state to be shared during push/pull synchronization.
func (d *Delegate) LocalState(_ bool) []byte {
	if d.Logger != nil {
		d.Logger.Println("Sharing remote state for push/pull sync")
	} else {
		fmt.Println(" === Sharing Remote State for push/pull sync === ")
	}

	return d.State
}

// GetBroadcasts returns queued broadcast messages to be sent to other nodes.
func (d *Delegate) GetBroadcasts(overhead, limit int) [][]byte {
	if d.Broadcasts == nil {
		return [][]byte{}
	}
	return d.Broadcasts.GetBroadcasts(overhead, limit)
}

// MergeRemoteState merges remote state received during push/pull synchronization.
func (d *Delegate) MergeRemoteState(buf []byte, _ bool) {
	if d.Logger != nil {
		d.Logger.Printf("Merging remote state for push/pull sync: %s", string(buf))
	} else {
		fmt.Printf(" === Merging Remote State %s for push/pull sync === \n", string(buf))
	}

	d.State = buf
}

// SetCluster initializes the delegate with a cluster instance and sets up the broadcast queue.
// This must be called before using UpdateState or GetBroadcasts.
func (d *Delegate) SetCluster(cluster *Cluster) {
	d.Cluster = cluster

	if cluster != nil && cluster.Memlist != nil {
		d.Broadcasts = &memberlist.TransmitLimitedQueue{
			NumNodes: func() int {
				return d.Cluster.Memlist.NumMembers()
			},
			RetransmitMult: 3,
		}
	}
}

// UpdateState updates the local state and queues a broadcast to all cluster nodes.
func (d *Delegate) UpdateState(data []byte) {
	if d.Broadcasts == nil {
		return
	}
	d.Broadcasts.QueueBroadcast(&Broadcast{
		msg:    data,
		notify: nil,
	})
}

// Message represents a key-value message that can be serialized to JSON.
type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Bytes serializes the message to JSON bytes.
// Returns an empty byte slice if marshaling fails.
func (m *Message) Bytes() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		return []byte("")
	}
	return data
}

// Load deserializes JSON bytes into the message.
// Returns an error if unmarshaling fails.
func (m *Message) Load(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	return nil
}
