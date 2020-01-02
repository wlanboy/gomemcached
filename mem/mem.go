package main

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	/*DefaultTimeout set default timeout*/
	DefaultTimeout = 10 * time.Second
)

func main() {
	mc := memcache.New("nuc:11211")
	mc.Timeout = DefaultTimeout

	mc.Set(&memcache.Item{Key: "key1", Value: []byte("hello")})
	mc.Set(&memcache.Item{Key: "key2", Value: []byte("world")})

	val, err := mc.Get("key1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(val.Key)
	fmt.Println(string(val.Value))

	it, err := mc.GetMulti([]string{"key1", "key2"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range it {
		fmt.Println(string(k))
		fmt.Println(string(v.Value))
	}

}
