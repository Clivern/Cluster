// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBroadcast test cases
func TestUnitBroadcast(t *testing.T) {
	// TestMessage
	t.Run("TestMessage", func(t *testing.T) {
		m := &Message{
			Key:   "key",
			Value: "value",
		}
		err := m.Load(m.Bytes())
		assert.NoError(t, err)
		assert.Equal(t, "key", m.Key)
		assert.Equal(t, "value", m.Value)
	})
}
