package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {

	mc := memcache.New("nuc:11211")

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
