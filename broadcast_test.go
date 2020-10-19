// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"
)

// TestBroadcast test cases
func TestBroadcast(t *testing.T) {
	// TestMessage
	t.Run("TestMessage", func(t *testing.T) {
		m := &Message{
			Key:   "key",
			Value: "value",
		}
		m.Load(m.Bytes())
		Expect(t, m.Key, "key")
		Expect(t, m.Value, "value")
	})
}
