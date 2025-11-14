// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"
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

// NewNodeEvents creates a new NodeEvents instance with default logging.
// Users can set a custom logger by modifying the Logger field.
func NewNodeEvents() *NodeEvents {
	return &NodeEvents{
		Logger: log.New(os.Stdout, "[cluster] ", log.LstdFlags),
	}
}

// NotifyJoin is called when a node joins the cluster.
func (n *NodeEvents) NotifyJoin(node *memberlist.Node) {
	if n.Logger != nil {
		n.Logger.Printf("A node has joined: %s", node.String())
	} else {
		fmt.Println("A node has joined: " + node.String())
	}
}

// NotifyLeave is called when a node leaves the cluster.
func (n *NodeEvents) NotifyLeave(node *memberlist.Node) {
	if n.Logger != nil {
		n.Logger.Printf("A node has left: %s", node.String())
	} else {
		fmt.Println("A node has left: " + node.String())
	}
}

// NotifyUpdate is called when a node's metadata is updated.
func (n *NodeEvents) NotifyUpdate(node *memberlist.Node) {
	if n.Logger != nil {
		n.Logger.Printf("A node was updated: %s", node.String())
	} else {
		fmt.Println("A node was updated: " + node.String())
	}
}
