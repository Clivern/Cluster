// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"log"
	"os"

	"github.com/hashicorp/memberlist"
)

// Logger is an interface for logging cluster events.
// Users can implement this interface to provide custom logging behavior.
type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// NodeEvents handles cluster membership events (join, leave, update).
// It implements the memberlist.EventDelegate interface.
// See https://github.com/hashicorp/memberlist/blob/master/event_delegate.go#L7
type NodeEvents struct {
	Logger Logger
}

// NewNodeEvents creates a new NodeEvents instance.
func NewNodeEvents(logger Logger) *NodeEvents {
	if logger == nil {
		logger = log.New(os.Stdout, "[cluster] ", log.LstdFlags)
	}
	return &NodeEvents{
		Logger: logger,
	}
}

// NotifyJoin is called when a node joins the cluster.
func (n *NodeEvents) NotifyJoin(node *memberlist.Node) {
	n.Logger.Printf("A node has joined: %s", node.String())
}

// NotifyLeave is called when a node leaves the cluster.
func (n *NodeEvents) NotifyLeave(node *memberlist.Node) {
	n.Logger.Printf("A node has left: %s", node.String())
}

// NotifyUpdate is called when a node's metadata is updated.
func (n *NodeEvents) NotifyUpdate(node *memberlist.Node) {
	n.Logger.Printf("A node was updated: %s", node.String())
}
