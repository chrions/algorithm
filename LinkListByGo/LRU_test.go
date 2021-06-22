package LinkListByGo

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3) // 该操作会使得关键字 2 作废
	fmt.Println(cache.Get(2))
	cache.Put(4, 4) // 该操作会使得关键字 3 作废
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
