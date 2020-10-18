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