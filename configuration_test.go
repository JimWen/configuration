package configuration

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestParseKeyOrder(t *testing.T) {

	wg := &sync.WaitGroup{}

	fn := func() {

		defer func() {
			wg.Done()
		}()

		for i := 0; i < 100000; i++ {
			conf := LoadConfig("tests/configs.conf")
			for g := 1; g < 3; g++ {
				for i := 1; i < 4; i++ {
					key := fmt.Sprintf("test.out.a.b.c.d.groups.g%d.o%d.order", g, i)
					order := conf.GetInt32(key, -1)

					if order != int32(i) {
						fmt.Println(conf)
						t.Fatalf("order not match,group %d, except: %d, real order: %d", g, i, order)
						return
					}
				}
			}
			conf = nil
			runtime.Gosched()
		}
	}

	wg.Add(2)

	go fn()
	go fn()

	wg.Wait()
}

func TestArrayParse(t *testing.T) {
	conf := LoadConfig("tests/t3.conf")

	fmt.Println(conf.GetString("test.out.a.b.c.d.groups.g2.o1.order"))

	arr1 := conf.GetObjectArray("test.out.a.b.c.d.groups.g2.o3")
	fmt.Println(arr1[0].GetString("key1"))
	fmt.Println(arr1[0].GetString("key2"))

	l := arr1[2].GetStringList("key222")
	fmt.Println(l)

	fmt.Println(arr1[0].GetStringMapString(""))

	map2 := conf.GetMapValue("test.out.a.b.c.d.groups.g2")
	for k, v := range map2 {
		fmt.Printf("key:%s -> value:%v\n", k, v)
	}
}
