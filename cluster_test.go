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
		cluster := &Cluster{}
		nodeName := cluster.GetNodeName()

		config := cluster.GetConfig()
		config.Name = nodeName
		cluster.SetConfig(config)

		count, err := cluster.AddLocalNode([]string{})

		Expect(t, count, 0)
		Expect(t, err, nil)
		Expect(t, nodeName, cluster.GetLocalNode().Name)
	})
}
