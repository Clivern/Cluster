// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"sync"

	"github.com/hashicorp/memberlist"
)

// RoundRobinBalancer ...
type RoundRobinBalancer struct {
	sync.Mutex

	cluster *Cluster
	current int
	pool    []*memberlist.Node
}

// NewRoundRobinBalancer ...
func NewRoundRobinBalancer(clus *Cluster) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		cluster: clus,
		current: 0,
		pool:    clus.Memlist.Members(),
	}
}

// Get ...
func (r *RoundRobinBalancer) Get() *memberlist.Node {
	r.Lock()
	defer r.Unlock()

	r.pool = r.cluster.Memlist.Members()

	if r.current >= len(r.pool) {
		r.current = r.current % len(r.pool)
	}

	result := r.pool[r.current]
	r.current++
	return result
}
