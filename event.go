// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

// NodeEvents struct
// implements https://github.com/hashicorp/memberlist/blob/master/event_delegate.go#L7
type NodeEvents struct{}

// NotifyJoin
func (n *NodeEvents) NotifyJoin(node *memberlist.Node) {
	fmt.Println("A node has joined: " + node.String())
}

// NotifyLeave
func (n *NodeEvents) NotifyLeave(node *memberlist.Node) {
	fmt.Println("A node has left: " + node.String())
}

// NotifyUpdate
func (n *NodeEvents) NotifyUpdate(node *memberlist.Node) {
	fmt.Println("A node was updated: " + node.String())
}
