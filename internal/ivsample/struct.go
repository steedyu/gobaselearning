package ivsample

import (
	"fmt"
	"reflect"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

/*
这个例子可以对比和.net的不同，但是golang结构体 匿名字段的理解不要光看这个理解
 */
func EmbeddedStructMethod() {
	t := Teacher{}
	//继承：如果 struct 中的一个匿名段实现了一个 method，那么包含这个匿名段的 struct 也能调用该 method。
	//这里并不存在 类似.NET override
	t.ShowA()

	//重写：如果 struct 中的一个匿名段实现了一个 method，包含这个匿名段的 struct 是可以重写匿名字段的方法的。
	t.ShowB()
}



// 示例1。
// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

// 示例2。
type Animal1 struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (a Animal1) String() string {
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}

// 示例3。
type Cat struct {
	name string
	Animal1
}

// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal1.AnimalCategory, cat.name)
}

func (cat *Cat)SetName(in string) {
	cat.name = in
}

/*
1 结构体同名的值方法和指针方法不能同时存在
2 比较定义两者以及结构体指针变量和值变量在调用它们各自异同


1 值方法的接收者是该方法所属的那个类型值的一个副本。我们在该方法内对该副本的修改一般都不会体现在原值上，除非这个类型本身是某个引用类型（比如切片或字典）的别名类型。
而指针方法的接收者，是该方法所属的那个基本类型值的指针值的一个副本。我们在这样的方法内对该副本指向的值进行修改，却一定会体现在原值上。

2 一个自定义数据类型的方法集合中仅会包含它的所有值方法，而该类型的指针类型的方法集合却囊括了前者的所有方法，包括所有值方法和所有指针方法。
严格来讲，我们在这样的基本类型的值上只能调用到它的值方法。但是，Go 语言会适时地为我们进行自动地转译，使得我们在这样的值上也能调用到它的指针方法。
比如，在Cat类型的变量cat之上，之所以我们可以通过cat.SetName("monster")修改猫的名字，是因为 Go 语言把它自动转译为了(&cat).SetName("monster")，
即：先取cat的指针值，然后在该指针值上调用SetName方法。

3 在后边你会了解到，一个类型的方法集合中有哪些方法与它能实现哪些接口类型是息息相关的。
如果一个基本类型和它的指针类型的方法集合是不同的，那么它们具体实现的接口类型的数量就也会有差异，除非这两个数量都是零。
比如，一个指针类型实现了某某接口类型，但它的基本类型却不一定能够作为该接口的实现类型。

 */
//func (cat Cat)SetName(in string) {
//	cat.name = in
//}


func EmbeddedStructMethod2() {

	// 示例1。
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)

	// 示例2。
	animal := Animal1{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)

	// 示例3。
	cat := Cat{
		name:   "little pig",
		Animal1: animal,
	}
	fmt.Printf("The cat: %s\n", cat)
}

func StructValueandPointMethod() {

	cat := Cat{
		name: "little pig",
	}
	cat.SetName("little fee")

	cat1 := &Cat{
		name : "little cat",
	}
	cat1.SetName("little ppe")

	fmt.Println(cat, cat1)
}



func CompareStruct() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	/*
	进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
	invalid operation: sn1 != sn3 (mismatched types struct { age int; name string } and struct { name string; age int })
	  */
	//sn3:= struct {
	//	name string
	//	age  int
	//}{age:11,name:"qq"}
	//if sn1 != sn3 {
	//	fmt.Println("sn1 != sn3")
	//}


	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11,
		m: map[string]string{"a": "1"},
	}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11,
		m: map[string]string{"a": "1"},
	}

	/*
	结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
	.\struct.go:74: invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
	 */
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}else {
		fmt.Println("sm1 != sm2")
	}
}