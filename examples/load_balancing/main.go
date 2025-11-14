// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package main demonstrates load balancing across cluster nodes using round-robin distribution.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/clivern/cluster"
	"github.com/gin-gonic/gin"
)

// //////////////////////////////////
// To run first node in the cluster
// $ go run main.go 8001
//
// To run other nodes in the cluster, you have to provide first node address
// you can see the address in the first node logs, something like (Node Name: Clivern-2.local--2aabd8d1-dc11-4e5f-acff-86f64a525ff7 Addr: 192.168.2.2:54968)
// $ go run main.go 8002 192.168.2.2:54968
// //////////////////////////////////
func main() {
	clus, _ := LaunchNode(os.Args[2:])

	address := fmt.Sprintf(
		"%s:%d",
		clus.GetLocalNode().Addr,
		clus.GetLocalNode().Port,
	)

	fmt.Printf("Node Address is %s\n", address)

	rr := cluster.NewRoundRobinBalancer(clus)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/get_node", func(c *gin.Context) {
		node := rr.Get()
		c.JSON(200, gin.H{
			"nodeAddr": fmt.Sprintf(
				"%s:%d",
				node.Addr,
				node.Port,
			),
		})
	})

	r.Run(fmt.Sprintf(":%s", os.Args[1]))
}

func LaunchNode(members []string) (*cluster.Cluster, *cluster.Delegate) {
	delegate := &cluster.Delegate{}

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

	return clus, delegate
}
