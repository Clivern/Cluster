// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"
	"os"

	"github.com/hashicorp/memberlist"
)

// Cluster struct
type Cluster struct {
	// https://github.com/hashicorp/memberlist/blob/master/config.go#L350
	Config  *memberlist.Config
	Memlist *memberlist.Memberlist
}

// GetConfig
func (c *Cluster) GetConfig() *memberlist.Config {
	return memberlist.DefaultLocalConfig()
}

// SetConfig
func (c *Cluster) SetConfig(config *memberlist.Config) {
	c.Config = config
}

// AddLocalNode
func (c *Cluster) AddLocalNode(members []string) (int, error) {
	var err error

	c.Memlist, err = memberlist.Create(c.Config)

	if err != nil {
		return 0, err
	}

	if len(members) > 0 {
		count, err := c.Memlist.Join(members)

		if err != nil {
			return count, err
		}
	}

	return 0, nil
}

// GetNodeName
func (c *Cluster) GetNodeName() string {
	hostname, _ := os.Hostname()

	return fmt.Sprintf("%s--%s", hostname, GenerateUUID4())
}

// GetLocalNode
func (c *Cluster) GetLocalNode() *memberlist.Node {
	return c.Memlist.LocalNode()
}
