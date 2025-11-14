// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCluster test cases
func TestUnitCluster(t *testing.T) {
	// TestCluster
	t.Run("TestCluster", func(t *testing.T) {
		cluster := &Cluster{}
		nodeName := cluster.GetNodeName()

		config := cluster.GetConfig()
		config.Name = nodeName
		cluster.SetConfig(config)

		count, err := cluster.AddLocalNode([]string{})

		assert.Equal(t, 0, count)
		assert.NoError(t, err)
		assert.Equal(t, nodeName, cluster.GetLocalNode().Name)
	})
}
