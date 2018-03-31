package map1


import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

type CompareFunction func(interface{}, interface{}) int8

type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	CompareFunc() CompareFunction
	ElemType() reflect.Type
}

// compareFunc的结果值：
//   小于0: 第一个参数小于第二个参数
//   等于0: 第一个参数等于第二个参数
//   大于1: 第一个参数大于第二个参数
type myKeys struct {
	/*
	由于Go语言本身并没有对自定义泛型提供支持，所以只有这样我们才能够用这个字段的值
	存储某一个数据类型的元素值
	 */
	container   []interface{}
	/*
	接口类型不具备有序性，不可能比较它们的大小。不过，也许把这个问题抛出去并让使用这个Keys的实现类型
	的编程人员来解决它是一个可行的方案。因为他们应该知道添加到Keys类型值中的元素值的实际类型并知道应该
	怎样比较它们。
	通过把比较两个元素值大小的问题抛给使用者，我们既解决动态确定元素类型的问题，又明确比较两个元素值大小的解决方式。
	 */
	compareFunc CompareFunction
	/*
	由于container字段是[]interface{}类型，所以我们常常不能够很方便地在运行时获取到它的实际元素类型
	（比如在它的值中还没有任何元素值的时候）。因此，我们需要一个明确container字段的实际元素类型字段
	 */
	elemType    reflect.Type
}

func (keys *myKeys) Len() int {
	return len(keys.container)
}

func (keys *myKeys) Less(i, j int) bool {
	return keys.compareFunc(keys.container[i], keys.container[j]) == -1
}

func (keys *myKeys) Swap(i, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	/*
	由于reflect.Type时一个接口类型，所以我们使用比较操作符!=来判定他们的相等性是否合法的
	 */
	if reflect.TypeOf(k) != keys.elemType {
		return false
	}
	return true
}

func (keys *myKeys) Add(k interface{}) bool {
	/*
	在真正向字段container的值添加元素值之前，我们应该先判断这个元素值的类型是否符合要求
	 */
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}
	keys.container = append(keys.container, k)
	/*
	sort.Sort函数使用的排序算法时一种由三向切分的快速排序算法、堆排序算法和插入排序算法组成的混合算法。
	虽然快速排序时最快的通用排序算法，但在元素值很少的情况下它比插入排序要慢一些。
	而堆排序的空间复杂度时常数级别的，且它的时间复杂度在大多数情况下只略逊于其他两种排序算法。
	所以在快速排序中的递归达到一定深度的时候，切换至堆排序来节约空间时值得的。
	 */
	sort.Sort(keys)
	return true
}


/*
从切片值中删除一个元素值由很多方式，比如使用for语句、copy函数或append函数等等。
我们在这里选择用append函数来实现，因为他可以在不增加时间复杂度和空间复杂度的情况下
使用更少的代码来完成功能，且不降低可读性。
 */
func (keys *myKeys) Remove(k interface{}) bool {
	index, contains := keys.Search(k)
	if !contains {
		return false
	}
	/*
	append是一个可变参函数。所以，我们可以在第二个参数值之后添加"..."以表示把第二个参数值中的每个元素值都作为传给
	 append函数的独立参数.
	 */
	keys.container = append(keys.container[0:index], keys.container[index+1:]...)
	return true
}

func (keys *myKeys) Clear() {
	keys.container = make([]interface{}, 0)
}

func (keys *myKeys) Get(index int) interface{} {
	if index >= keys.Len() {
		return nil
	}
	return keys.container[index]
}

func (keys *myKeys) GetAll() []interface{} {
	initialLen := len(keys.container)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for _, key := range keys.container {
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
}

func (keys *myKeys) Search(k interface{}) (index int, contains bool) {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return
	}
	/*
	由于sort.Search函数使用二分查找算法在切片值中搜索制定的元素值。
	这种搜索算法有着稳定的O(logN)的时间复杂度，但它要求被搜索的数组（这里时切片值）必须时有序的。
	 */
	index = sort.Search(
		keys.Len(),
		func(i int) bool { return keys.compareFunc(keys.container[i], k) >= 0 })
	/*
	sort.Search函数的结果值总会在[0,n]的范围内，但结果值并不一定就是欲查找的元素值所对应的索引值。
	因此，我们还需要在得到调用sort.Search函数的结果值之后在进行一次判断。
	 */
	if index < keys.Len() && keys.container[index] == k {
		contains = true
	}
	return
}

func (keys *myKeys) ElemType() reflect.Type {
	return keys.elemType
}

func (keys *myKeys) CompareFunc() CompareFunction {
	return keys.compareFunc
}

func (keys *myKeys) String() string {
	var buf bytes.Buffer
	buf.WriteString("Keys<")
	buf.WriteString(keys.elemType.Kind().String())
	buf.WriteString(">{")
	first := true
	buf.WriteString("[")
	for _, key := range keys.container {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("]")
	buf.WriteString("}")
	return buf.String()
}

func NewKeys(
compareFunc func(interface{}, interface{}) int8,
elemType reflect.Type) Keys {
	return &myKeys{
		container:   make([]interface{}, 0),
		compareFunc: compareFunc,
		elemType:    elemType,
	}
}

