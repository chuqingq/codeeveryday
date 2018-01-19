package main

import(
	"math/rand"
	rand2 "crypto/rand"
	"encoding/binary"
	"time"
	// "strconv"
	"fmt"
	"sync"

	"github.com/satori/go.uuid"
)

// 根据纳秒时间戳设置随机种子
func getRandom() uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint32()
}
// $ time go run random_sample.go 
// 1998834
// 
// real 0m9.565s
// user 0m57.312s
// sys  0m0.540s

// 安全随机数
func getRandom2() uint32 {
	var b [4]byte
	_, err := rand2.Read(b[:])
	if err != nil {
		fmt.Println("error:", err)
		return 0
	}
	return binary.BigEndian.Uint32(b[:])
}
// $ time go run random_sample2.go 
// 1998736
// 
// real 0m9.659s
// user 0m57.136s
// sys  0m0.624s

// 第三方UUID
func getRandom3() string {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error uuid.NewV4():", err)
		return ""
	}
	return u.String()
}
// chuqq@chuqq-hp:~/gopath/src/test$ time go run random_sample3.go 
// 2000000
// 
// real	0m5.025s
// user	0m6.796s
// sys	0m12.204s
// chuqq@chuqq-hp:~/gopath/src/test$ time go run random_sample3.go 
// 2000000
// 
// real	0m4.716s
// user	0m6.084s
// sys	0m11.724s

func main(){
	var m sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 20; i++{
		wg.Add(1)
		go func(){
			for j := 0; j < 100000; j++ {
				m.Store(getRandom3(), true)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	count := 0
	m.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	fmt.Println(count)
}

