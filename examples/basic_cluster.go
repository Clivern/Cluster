// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/clivern/cluster"
)

func main() {
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

	clus.AddLocalNode([]string{})

	fmt.Println(clus.GetLocalNode())

	for {

	}
}
