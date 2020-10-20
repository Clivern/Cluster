// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/clivern/cluster"
)

var (
	delegate = &cluster.Delegate{}
)

func LaunchNode(members []string) *cluster.Cluster {
	delegate = &cluster.Delegate{}

	clus := &cluster.Cluster{}

	// Generate a unique name
	nodeName := clus.GetNodeName()

	// Get a default configs
	config := clus.GetConfig()
	config.Name = nodeName
	config.BindPort = 0 // assign a free port
	config.Events = &cluster.NodeEvents{}
	config.Delegate = delegate
	config.PushPullInterval = time.Second * 10 // to make it demonstrable
	config.ProbeInterval = time.Second * 1     // to make failure demonstrable

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

	delegate.SetCluster(clus)

	return clus
}

func main() {
	clus := LaunchNode([]string{})

	address := fmt.Sprintf(
		"%s:%d",
		clus.GetLocalNode().Addr,
		clus.GetLocalNode().Port,
	)

	// Run cluster nodes
	for i := 0; i < 5; i++ {
		go func() {
			LaunchNode([]string{address})
			for {

			}
		}()
	}

	for {
		delegate.UpdateState([]byte("Hello"))

		time.Sleep(5 * time.Second)
	}
}
