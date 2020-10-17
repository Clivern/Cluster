// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"
)

// TestCluster test cases
func TestCluster(t *testing.T) {
	// TestCluster
	t.Run("TestCluster", func(t *testing.T) {
		configs := NewNodeConfig()
		configs.Format = "test-123-45-67-89"
		clusterNode := NewClusterNode(configs)

		node, err := clusterNode.Init([]string{})

		Expect(t, err, nil)
		Expect(t, node.Name, "test-123-45-67-89")
	})
}
