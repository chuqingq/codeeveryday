package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSave(t *testing.T) {
	confPathFmt = "./%s.conf"

	var wg WireGuard
	err := wg.Load("wg1")
	assert.Nil(t, err)
	assert.Equal(t, "10.11.0.2/24", wg.Interface.Address)
	assert.Equal(t, 10, wg.Peer.PersistentKeepalive)

	// modify
	const newAddr = "10.11.0.1/24"
	wg.Interface.Address = newAddr

	// save
	err = wg.SaveTo("wg2")
	defer os.Remove(confPath("wg2"))
	assert.Nil(t, err)

	// load
	var wg2 WireGuard
	err = wg2.Load("wg2")
	assert.Nil(t, err)
	assert.Equal(t, newAddr, wg.Interface.Address)
	assert.Equal(t, 10, wg.Peer.PersistentKeepalive)
}

func TestGetAllowedIps(t *testing.T) {
	confPathFmt = "./%s.conf"

	var wg WireGuard
	err := wg.Load("wg1")
	assert.Nil(t, err)
	assert.Equal(t, 6, len(wg.GetAllowedIPs()))
}
