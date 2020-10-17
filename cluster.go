// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"os"
	"strings"

	"github.com/hashicorp/memberlist"
)

// NodeConfigs struct
type NodeConfigs struct {
	Format string
}

// ClusterNode struct
type ClusterNode struct {
	Config *memberlist.Config
}

// NewNodeConfig
func NewNodeConfig() *NodeConfigs {
	return &NodeConfigs{
		Format: "{host}-{uuid}",
	}
}

// NewClusterNode
func NewClusterNode(config *NodeConfigs) *ClusterNode {
	node := &ClusterNode{}

	// https://github.com/hashicorp/memberlist/blob/master/config.go#L350
	node.Config = memberlist.DefaultLocalConfig()

	hostname, _ := os.Hostname()
	name := strings.Replace(config.Format, "{host}", hostname, -1)
	name = strings.Replace(name, "{uuid}", GenerateUUID4(), -1)

	node.Config.Name = name

	return node
}

// Init
func (c *ClusterNode) Init(members []string) (*memberlist.Node, error) {
	m, err := memberlist.Create(c.Config)

	if err != nil {
		return &memberlist.Node{}, err
	}

	if len(members) > 0 {
		_, err := m.Join(members)

		if err != nil {
			return &memberlist.Node{}, err
		}
	}

	return m.LocalNode(), nil
}
