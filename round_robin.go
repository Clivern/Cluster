// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"sync"
)

// RoundRobin ...
type RoundRobin struct {
	sync.Mutex

	current int
	pool    []int
}

// NewRoundRobin ...
func NewRoundRobin(size int) *RoundRobin {
	a := make([]int, size)

	for i := range a {
		a[i] = i
	}

	return &RoundRobin{
		current: 0,
		pool:    a,
	}
}

// Get ...
func (r *RoundRobin) Get() int {
	r.Lock()
	defer r.Unlock()

	if r.current >= len(r.pool) {
		r.current = r.current % len(r.pool)
	}

	result := r.pool[r.current]
	r.current++
	return result
}
