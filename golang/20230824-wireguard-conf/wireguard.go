package main

import (
	"fmt"
	"os/exec"

	"gopkg.in/ini.v1"
)

func main() {

}

type WireGuard struct {
	Interface struct {
		PrivateKey string
		Address    string
	}
	Peer struct {
		PublicKey           string
		PresharedKey        string
		AllowedIPs          []string
		Endpoint            string
		PersistentKeepalive int
	}
}

var confPathFmt = "/etc/wireguard/%s.conf"

func confPath(inf string) string {
	return fmt.Sprintf(confPathFmt, inf)
}

func (w *WireGuard) Load(inf string) error {
	f, err := ini.Load(confPath(inf))
	if err != nil {
		return err
	}
	err = f.MapTo(w)
	if err != nil {
		return err
	}
	w.Peer.AllowedIPs = f.Section("Peer").Key("AllowedIPs").Strings(",")
	return nil
}

func (w *WireGuard) SaveTo(inf string) error {
	file := ini.Empty()
	err := ini.ReflectFrom(file, w)
	if err != nil {
		return err
	}
	return file.SaveTo(confPath(inf))
}

func (w *WireGuard) Stop(inf string) error {
	return exec.Command("wg-quick", "down", inf).Run()
}

func (w *WireGuard) Start(inf string) error {
	return exec.Command("wg-quick", "up", inf).Run()
}
