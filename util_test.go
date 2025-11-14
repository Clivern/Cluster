// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInArray test cases
func TestUnitInArray(t *testing.T) {
	// TestInArray
	t.Run("TestInArray", func(t *testing.T) {
		assert.True(t, InArray("A", []string{"A", "B", "C", "D"}))
		assert.True(t, InArray("B", []string{"A", "B", "C", "D"}))
		assert.False(t, InArray("H", []string{"A", "B", "C", "D"}))
		assert.True(t, InArray(1, []int{2, 3, 1}))
		assert.False(t, InArray(9, []int{2, 3, 1}))
	})
}

// TestGenerateUUID4 test cases
func TestUnitGenerateUUID4(t *testing.T) {
	t.Run("TestGenerateUUID4", func(t *testing.T) {
		uuid := GenerateUUID4()
		assert.NotEmpty(t, uuid, "Generated UUID should not be empty")
		assert.Equal(t, 36, len(uuid), "UUID should be 36 characters long")
	})
}

// TestUnset test cases
func TestUnitUnset(t *testing.T) {
	t.Run("TestUnset", func(t *testing.T) {
		slice := []string{"A", "B", "C", "D"}
		result := Unset(slice, 1)
		assert.Equal(t, 3, len(result))
		assert.False(t, InArray("B", result))
		assert.True(t, InArray("A", result))
		assert.True(t, InArray("C", result))
		assert.True(t, InArray("D", result))
	})
}
