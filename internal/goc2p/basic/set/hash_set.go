package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}


/*
编写实现清除所有元素值功能的方法会用到一个小技巧。最干脆和简洁的方法就是为字段m重新赋值。
 */
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

/*
GO语言的源代码获知，当我们把一个interface{}类型作为键添加到一个字典值的时候，Go语言会先获取这个interface{}
类型值的实际类型（即动态类型），然后再使用与之相对应的hash函数对该值进行hash运算。
 */
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

/*
由于HashSet类型值的元素迭代顺序的不确定性。这种不确定性会使我们无法通过索引值获取某一个元素值。
一个简单可行的解决方案就是先生成一个它的快照，然后再在这个快照之上进行迭代操作。所谓快照，就是目标值在某一个时刻的映像。
 */
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot

	/*
	之所以我们使用这么多条语句来实现这个方法是因为需要考虑道在从获取字段m的值的长度道对m的值迭代完成的这个时间段内，
	m的值中的元素数量可能会发生变化。
	 */
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("HashSet{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}
