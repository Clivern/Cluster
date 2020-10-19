// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"
)

// TestNewRoundRobin test cases
func TestNewRoundRobin(t *testing.T) {
	// TestNewRoundRobin
	t.Run("TestNewRoundRobin", func(t *testing.T) {
		rb := NewRoundRobin(10)
		for i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			Expect(t, rb.Get(), i)
		}
		for i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			Expect(t, rb.Get(), i)
		}
	})
}
