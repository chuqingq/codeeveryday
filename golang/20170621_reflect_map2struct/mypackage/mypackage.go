package mypackage

type MyConfig struct {
	ListenIP   string `cfg:"listenip" json:"listenip"`
	ListenPort int    `cfg:"listenport" json:"listenport"`
	Keepalive  struct {
		Idle  int `cfg:"idle" json:"idle"`
		count int `cfg:"count" json:"count"`
		Abc   struct {
			AbcValue string `cfg:"abcvalue" json:"abcvalue"`
		} `cfg:"abc" json:"abc"`
	} `cfg:"keepalive" json:"keepalive"`
}
