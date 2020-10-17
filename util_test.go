// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"
)

// TestInArray test cases
func TestInArray(t *testing.T) {
	// TestInArray
	t.Run("TestInArray", func(t *testing.T) {
		Expect(t, InArray("A", []string{"A", "B", "C", "D"}), true)
		Expect(t, InArray("B", []string{"A", "B", "C", "D"}), true)
		Expect(t, InArray("H", []string{"A", "B", "C", "D"}), false)
		Expect(t, InArray(1, []int{2, 3, 1}), true)
		Expect(t, InArray(9, []int{2, 3, 1}), false)
	})
}
