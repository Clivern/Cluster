// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package cluster provides a simple interface for creating and managing distributed clusters
// using the hashicorp/memberlist library. It includes features for node discovery, state
// synchronization, and load balancing across cluster members.
package cluster

import (
	"fmt"
	"os"

	"github.com/hashicorp/memberlist"
)

// Cluster represents a cluster node that can join and communicate with other nodes.
// It wraps the hashicorp/memberlist library to provide a simpler interface.
type Cluster struct {
	// Config holds the memberlist configuration.
	// See https://github.com/hashicorp/memberlist/blob/master/config.go#L350
	Config *memberlist.Config

	// Memlist is the underlying memberlist instance.
	Memlist *memberlist.Memberlist
}

// GetConfig returns a default local memberlist configuration.
// The returned config can be modified before calling SetConfig.
func (c *Cluster) GetConfig() *memberlist.Config {
	return memberlist.DefaultLocalConfig()
}

// SetConfig sets the memberlist configuration for this cluster.
// The configuration must be set before calling AddLocalNode.
func (c *Cluster) SetConfig(config *memberlist.Config) {
	c.Config = config
}

// AddLocalNode creates a new memberlist instance and optionally joins existing members.
// If members is empty, it creates a standalone node. Otherwise, it attempts to join
// the cluster using the provided member addresses.
// Returns the number of nodes successfully joined and any error that occurred.
func (c *Cluster) AddLocalNode(members []string) (int, error) {
	var err error

	c.Memlist, err = memberlist.Create(c.Config)
	if err != nil {
		return 0, fmt.Errorf("failed to create memberlist: %w", err)
	}

	if len(members) > 0 {
		count, err := c.Memlist.Join(members)
		if err != nil {
			return count, fmt.Errorf("failed to join cluster: %w", err)
		}
		return count, nil
	}

	return 0, nil
}

// GetNodeName generates a unique node name by combining the hostname with a UUID.
// If hostname cannot be determined, it uses "unknown" as the hostname.
func (c *Cluster) GetNodeName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	uuid := GenerateUUID4()
	if uuid == "" {
		// Fallback if UUID generation fails
		uuid = "fallback"
	}

	return fmt.Sprintf("%s--%s", hostname, uuid)
}

// GetLocalNode returns the local node information from the memberlist.
// Returns nil if the cluster has not been initialized with AddLocalNode.
func (c *Cluster) GetLocalNode() *memberlist.Node {
	if c.Memlist == nil {
		return nil
	}
	return c.Memlist.LocalNode()
}
