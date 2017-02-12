package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=%s&supervoteid=%s&uin=MTMwMzUxMjg3Mw%3D%3D&key=82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&wxtoken=543112670&mid=2650886522&idx=1"
	u, _ := url.Parse(s)
	fmt.Printf("u: %+v\n", u)

	v := u.Query()
	fmt.Printf("v: %+v\n", v)
	v.Set("item", `{"name":"value"}`)
	fmt.Printf("v2: %+v\n", v)
	fmt.Printf("v.Encode(): %v\n", v.Encode())

	s2 := `{"name":"value"}`
	fmt.Printf("s2: %+v\n", url.QueryEscape(s2))
}

// u: https://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=%s&supervoteid=%s&uin=MTMwMzUxMjg3Mw%3D%3D&key=82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&wxtoken=543112670&mid=2650886522&idx=1
// v: map[action:[show] uin:[MTMwMzUxMjg3Mw==] key:[82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5] pass_ticket:[WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W] wxtoken:[543112670] mid:[2650886522] idx:[1]]
// v2: map[action:[show] uin:[MTMwMzUxMjg3Mw==] key:[82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5] pass_ticket:[WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W] wxtoken:[543112670] mid:[2650886522] idx:[1] item:[{"name":"value"}]]
// v.Encode(): action=show&idx=1&item=%7B%22name%22%3A%22value%22%7D&key=82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5&mid=2650886522&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&uin=MTMwMzUxMjg3Mw%3D%3D&wxtoken=543112670
// s2: %7B%22name%22%3A%22value%22%7D
