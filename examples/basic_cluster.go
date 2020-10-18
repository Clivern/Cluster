// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/clivern/cluster"
)

func LauchNode(members []string) string {
	clus := &cluster.Cluster{}

	// Generate a unique name
	nodeName := clus.GetNodeName()

	// Get a default configs
	config := clus.GetConfig()
	config.Name = nodeName
	config.BindPort = 0 // assign a free port
	config.Events = &cluster.NodeEvents{}

	// Override configs
	clus.SetConfig(config)

	clus.AddLocalNode(members)

	fmt.Printf(
		"Node Name: %s Addr: %s:%d\n",
		clus.GetLocalNode().Name,
		clus.GetLocalNode().Addr,
		clus.GetLocalNode().Port,
	)

	return fmt.Sprintf(
		"%s:%d",
		clus.GetLocalNode().Addr,
		clus.GetLocalNode().Port,
	)
}

func main() {
	node := LauchNode([]string{})

	// Run cluster nodes
	for i := 0; i < 5; i++ {
		go func() {
			LauchNode([]string{node})
			for {

			}
		}()
	}

	for {

	}
}
