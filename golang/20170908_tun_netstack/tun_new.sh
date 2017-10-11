sudo ip tuntap add user chuqq mode tun tun0
sudo ip link set tun0 up
sudo ip addr add 10.0.0.1/24 dev tun0

sudo sysctl net.ipv4.ip_forward=1
// sudo iptables -P INPUT ACCEPT
sudo iptables -t nat -A PREROUTING -p tcp -i eth0 --dport 8888 -j DNAT --to 10.0.0.2:8888
sudo iptables -t nat -A POSTROUTING -j MASQUERADE

tun_tcp_echo tun0 10.0.0.2 8888

