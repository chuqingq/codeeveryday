package main

// Command 命令。一般是工具发给设备。
type Command struct {
	Cmd    string                 // 命令
	Device Device                 // 设备。用于指定命令是发给哪个设备的，或者从哪个设备回来的；以及网络参数。
	Info   map[string]interface{} // 命令的其他属性
}

// Device 设备信息
type Device struct {
	SerialNumber  string // 设备序列号。设备的唯一标识
	AdminPassword string // 管理员密码

	DHCP          bool   // 开启DHCP
	IP            string // IP地址
	Port          int    // 端口
	Mask          string // 子网掩码
	Gateway       string // 网关
	IPv6          string // IPv6地址
	IPv6Gateway   string // IPv6网关
	IPv6PrefixLen int    // IPv6子网前缀长度

	Info map[string]interface{} // 其他属性

	Socket   *MulticastUDP `json:"-"` // 多播socket
	RecvChan chan *Command `json:"-"` // 命令接收通道
}

const (
	CmdDiscover string = "_discover" // 发现设备
)

func Req(cmd string) string {
	return cmd + ".req"
}

func Res(cmd string) string {
	return cmd + ".res"
}
