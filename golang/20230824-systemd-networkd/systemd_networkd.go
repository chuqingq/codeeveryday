package main

import (
	"fmt"
	"os/exec"

	"gopkg.in/ini.v1"
)

type SystemdNetwork struct {
	Match struct {
		Name string
	}
	Network struct {
		DHCP    string
		Address []string
		Gateway string
		DNS     []string
	}
}

var confPathFmt = "/etc/systemd/network/%s.network"

func confPath(inf string) string {
	return fmt.Sprintf(confPathFmt, inf)
}

func (n *SystemdNetwork) Load(inf string) error {
	f, err := ini.ShadowLoad(confPath(inf))
	if err != nil {
		return err
	}
	err = f.MapTo(n)
	if err != nil {
		return err
	}
	n.Network.Address = f.Section("Network").Key("Address").ValueWithShadows()
	n.Network.DNS = f.Section("Network").Key("DNS").ValueWithShadows()
	return nil
}

func (n *SystemdNetwork) SaveTo(inf string) error {
	file := ini.Empty(ini.LoadOptions{AllowShadows: true})

	// Match.Name
	file.Section("Match").Key("Name").SetValue(n.Match.Name)
	// Network
	if n.Network.DHCP != "" {
		// DHCP
		file.Section("Network").Key("DHCP").SetValue(n.Network.DHCP)
	} else {
		// static
		// Address
		file.Section("Network").DeleteKey("Address")
		for _, addr := range n.Network.Address {
			file.Section("Network").Key("Address").AddShadow(addr)
		}
		// Gateway
		file.Section("Network").Key("Gateway").SetValue(n.Network.Gateway)
		// DNS
		for _, addr := range n.Network.DNS {
			file.Section("Network").Key("DNS").AddShadow(addr)
		}
	}
	return file.SaveTo(confPath(inf))
}

func StopService() error {
	return exec.Command("systemctl", "stop", "systemd-networkd.service").Run()
}

func StartService() error {
	return exec.Command("systemctl", "start", "systemd-networkd.service").Run()
}

func EnableService() error {
	return exec.Command("systemctl", "enable", "systemd-networkd.service").Run()
}

