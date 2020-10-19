// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"encoding/json"
)

// Broadcast struct
type Broadcast struct {
}

// Delegate struct
type Delegate struct {
	msgCh chan []byte
}

// Message struct
type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NotifyMsg ..
func (d *Delegate) NotifyMsg(msg []byte) {
	d.msgCh <- msg
}

// NodeMeta ..
func (d *Delegate) NodeMeta(limit int) []byte {
	return []byte("")
}

// LocalState ..
func (d *Delegate) LocalState(join bool) []byte {
	return []byte("")
}

// GetBroadcasts ..
func (d *Delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return nil
}

// MergeRemoteState ..
func (d *Delegate) MergeRemoteState(buf []byte, join bool) {

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
