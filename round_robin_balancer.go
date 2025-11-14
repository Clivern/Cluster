// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"sync"

	"github.com/hashicorp/memberlist"
)

// RoundRobinBalancer distributes requests across cluster nodes using a round-robin algorithm.
type RoundRobinBalancer struct {
	mu      sync.Mutex
	cluster *Cluster
	current int
	pool    []*memberlist.Node
}

// NewRoundRobinBalancer creates a new round-robin balancer for the given cluster.
func NewRoundRobinBalancer(clus *Cluster) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		cluster: clus,
		current: 0,
		pool:    clus.Memlist.Members(),
	}
}

// Get returns the next node in the round-robin sequence.
// It automatically refreshes the node pool from the cluster membership.
func (r *RoundRobinBalancer) Get() *memberlist.Node {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cluster == nil || r.cluster.Memlist == nil {
		return nil
	}

	r.pool = r.cluster.Memlist.Members()

	if len(r.pool) == 0 {
		return nil
	}

	// Normalize current index to valid range
	if r.current >= len(r.pool) {
		r.current = r.current % len(r.pool)
	}

	result := r.pool[r.current]
	r.current = (r.current + 1) % len(r.pool)
	return result
}
