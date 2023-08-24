package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSave(t *testing.T) {
	confPathFmt = "./20-%s.network"

	var net SystemdNetwork
	err := net.Load("wired-dhcp")
	assert.Nil(t, err)
	assert.Equal(t, "ipv4", net.Network.DHCP)

	err = net.Load("wired-static")
	assert.Nil(t, err)
	assert.Equal(t, []string{"10.1.10.9/24"}, net.Network.Address)
}

// TestLoadMultiAddress 测试加载含多个Address的配置
func TestLoadMultiAddress(t *testing.T) {
	confPathFmt = "./20-%s.network"

	var net SystemdNetwork
	err := net.Load("wired-static-multi-address")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(net.Network.Address))
	assert.Equal(t, "10.1.10.9/24", net.Network.Address[0])
	assert.Equal(t, "10.1.10.10/24", net.Network.Address[1])
}

// TestSaveMultiAddress 测试保存含多个Address的配置
func TestSaveMultiAddress(t *testing.T) {
	confPathFmt = "./20-%s.network"

	var net SystemdNetwork
	net.Match.Name = "eth0"
	net.Network.Address = []string{"10.1.10.9/24", "10.1.10.10/24"}
	net.Network.DNS = []string{"10.1.1.1"}

	dev := "test1"
	err := net.SaveTo(dev)
	defer os.Remove(confPath(dev))
	assert.Nil(t, err)

	var read SystemdNetwork
	err = read.Load(dev)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(read.Network.Address))
}
