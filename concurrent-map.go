/*
@Time : 2023/2/6 12:00
@Author : kyle
@Mail : kyle.wdw@gmail.com
*/
package concurrent_map

import (
	"sync"
)

type ConcurrentMap[key comparable, value any] struct {
	oriMap sync.Map
}

func (c *ConcurrentMap[mapKey, mapValue]) Store(key mapKey, value mapValue) {
	c.oriMap.Store(key, value)
}

func (c *ConcurrentMap[mapKey, mapValue]) Load(key mapKey) (ret mapValue) {
	b, ok := c.oriMap.Load(key)
	if ok {
		ret = b.(mapValue)
	}
	return
}

func (c *ConcurrentMap[mapKey, mapValue]) Delete(key mapKey) {
	c.oriMap.Delete(key)
}

func (c *ConcurrentMap[mapKey, mapValue]) Range(fn func(key mapKey, value mapValue) bool) {
	c.oriMap.Range(func(key, value any) bool {
		k := key.(mapKey)
		v := value.(mapValue)
		if fn(k, v) {
			return true
		}
		return false
	})
}

// Len 采用遍历的方式获取数量，性能消耗较大
func (c *ConcurrentMap[mapKey, mapValue]) Len() int32 {
	var count int32
	c.Range(func(key mapKey, value mapValue) bool {
		count++
		return true
	})
	return count
}

func CreateMap[mapKey comparable, mapValue any]() *ConcurrentMap[mapKey, mapValue] {
	return &ConcurrentMap[mapKey, mapValue]{oriMap: sync.Map{}}
}
