package main

import (
	"fmt"
	"os/exec"
	"strings"

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
		AllowedIPs          string
		Endpoint            string
		PersistentKeepalive int
	}
}

var confPathFmt = "/etc/wireguard/%s.conf"

func confPath(inf string) string {
	return fmt.Sprintf(confPathFmt, inf)
}

func (w *WireGuard) Load(inf string) error {
	err := ini.MapTo(w, confPath(inf))
	if err != nil {
		return err
	}
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

func (w *WireGuard) GetAllowedIPs() []string {
	allowedIPs := strings.Split(w.Peer.AllowedIPs, ",")
	// 去掉AllowedIPs中的空格
	for i, v := range allowedIPs {
		allowedIPs[i] = strings.TrimSpace(v)
	}
	return allowedIPs
}

func (w *WireGuard) Stop(inf string) error {
	return exec.Command("wg-quick", "down", inf).Run()
}

func (w *WireGuard) Start(inf string) error {
	return exec.Command("wg-quick", "up", inf).Run()
}
