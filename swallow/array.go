package swallow


import (
"sort"
"sync"
)

type ArrayCustom struct {
	sync.RWMutex
	m map[int]bool
}

// 新建集合对象
func New(items ...int) *ArrayCustom {
	s := &ArrayCustom{
		m: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}

// 添加元素
func (s *ArrayCustom) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.m[v] = true
	}
}

// 删除元素
func (s *ArrayCustom) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		delete(s.m, v)
	}
}

// 判断元素是否存在
func (s *ArrayCustom) Has(items ...int) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

// 元素个数
func (s *ArrayCustom) Count() int {
	return len(s.m)
}

// 清空集合
func (s *ArrayCustom) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

// 空集合判断
func (s *ArrayCustom) Empty() bool {
	return len(s.m) == 0
}

// 无序列表
func (s *ArrayCustom) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 排序列表
func (s *ArrayCustom) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// 并集
func (s *ArrayCustom) Merge(sets ...*ArrayCustom) *ArrayCustom {
	r := New(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			r.m[e] = true
		}
	}
	return r
}

// 差集
func (s *ArrayCustom) Diff(sets ...*ArrayCustom) *ArrayCustom {
	r := New(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			if _, ok := s.m[e]; ok {
				delete(r.m, e)
			}
		}
	}
	return r
}

// 交集
func (s *ArrayCustom) Intersect(sets ...*ArrayCustom) *ArrayCustom {
	r := New(s.List()...)
	for _, set := range sets {
		for e := range s.m {
			if _, ok := set.m[e]; !ok {
				delete(r.m, e)
			}
		}
	}
	return r
}

// 补集
func (s *ArrayCustom) Complement(full *ArrayCustom) *ArrayCustom {
	r := New()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			r.Add(e)
		}
	}
	return r
}
