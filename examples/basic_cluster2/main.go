// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package main demonstrates a basic cluster setup with node membership information display.
package main

import (
	"fmt"

	"github.com/clivern/cluster"
)

func LaunchNode(members []string) string {
	clus := &cluster.Cluster{}

	// Generate a unique name
	nodeName := clus.GetNodeName()

	// Get a default configs
	config := clus.GetConfig()
	config.Name = nodeName
	config.BindPort = 0 // assign a free port
	config.Events = cluster.NewNodeEvents(nil)

	// Override configs
	clus.SetConfig(config)

	clus.AddLocalNode(members)

	fmt.Printf(
		"Node Name: %s Addr: %s:%d\n",
		clus.GetLocalNode().Name,
		clus.GetLocalNode().Addr,
		clus.GetLocalNode().Port,
	)

	fmt.Printf(
		"Live Nodes Count: %d\n",
		clus.Memlist.NumMembers(),
	)

	fmt.Printf(
		"Live Nodes List: %v\n",
		clus.Memlist.Members(),
	)

	return fmt.Sprintf(
		"%s:%d",
		clus.GetLocalNode().Addr,
		clus.GetLocalNode().Port,
	)
}

func main() {
	node := LaunchNode([]string{})

	// Run cluster nodes
	for i := 0; i < 5; i++ {
		go func() {
			LaunchNode([]string{node})
			select {} // Keep goroutine alive
		}()
	}

	select {} // Keep main goroutine alive
}
