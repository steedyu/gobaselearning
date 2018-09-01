package oopsample

import "fmt"



//Inheritance and Composition in Go
/*
Go中实现组合是一件十分容易的事情。简单组合两个结构体就能够构造一个新的数据类型。
这就是Go中实现代码和数据共享的常用方式
 */
type TokenType uint16

/*
const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LBRACKET
	RBRACKET
	INT
)

type Token struct {
	Type   TokenType
	Lexeme string
}

type IntegerConstant struct {
	Token *Token
	Value uint64
}
*/

//这就是Go中实现代码和数据共享的常用方式。然而如果你想实现继承机制，我们该如何去做？
type Token interface {
	Type() TokenType
	Lexeme() string
}

type Match struct {
	toktype TokenType
	lexeme  string
}

//type IntegerConstant struct {
//	token Token
//	value uint64
//}

func (m *Match) Type() TokenType {
	return m.toktype
}

func (m *Match) Lexeme() string {
	return m.lexeme
}

/*
type IntegerConstant struct {
	token Token
	value uint64
}


func (i *IntegerConstant) Type() TokenType {
	return i.token.Type()
}

func (i *IntegerConstant) Lexeme() string {
	return i.token.Lexeme()
}
*/

func (i *IntegerConstant) Value() uint64 {
	return i.value
}

/*
继承机制的简化版 上面的实现方案的一个问题是*IntegerConstant的方法调用中，出现了重复造轮子的问题。
但是我们可以使用Go内建的嵌入机制来避免此类情况的出现。嵌入机制(匿名嵌入)允许类型之前共享代码和数据。
 */
type IntegerConstant struct {
	Token
	value uint64
}

/*
在以上的方案中，你不能嵌入与嵌入类型相同的方法名。
例如结构体Bar匿名嵌入结构体Foo后，就不能拥有名称为Foo的方法，同样也不能实现type Fooer interface { Foo() }接口类型。
.\golangoopsample.go:87: type IntegerConstant has both field and method named Token
func (i *IntegerConstant) Token()  {

}
*/


/*
继承自其他结构体的struct类型可以直接访问父类结构体的字段和方法
*/
//type Pet struct {
//	name string
//}
//
//type Dog struct {
//	Pet
//	Breed string
//}
//
//func (p *Pet) Speak() string {
//	return fmt.Sprintf("my name is %v", p.name)
//}
//
//func (p *Pet) Name() string {
//	return p.name
//}
//
//func (d *Dog) Speak() string {
//	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
//}
//
//
//func DoPetDog1() {
//	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
//	fmt.Println(d.Name())
//	fmt.Println(d.Speak())
//}

/*
嵌入式继承机制的的局限
Overriding Methods 上面的Pet例子中，Dog类型重载了Speak()方法。
然而如果Pet有另外一个方法Play()被调用，但是Dog没有实现Play()的时候，Dog类型的Speak()方法则不会被调用。
 */
type Pet struct {
	name string
}

type Dog struct {
	Pet
	Breed string
}

func (p *Pet) Play() {
	fmt.Println(p.Speak())
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}

func DoPetDog2() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	d.Play()
}

//type Pet struct {
//	speaker func() string
//	name    string
//}
//
//type Dog struct {
//	Pet
//	Breed string
//}
//
//func NewPet(name string) *Pet {
//	p := &Pet{
//		name: name,
//	}
//	p.speaker = p.speak
//	return p
//}
//
//func (p *Pet) Play() {
//	fmt.Println(p.Speak())
//}
//
//func (p *Pet) Speak() string {
//	return p.speaker()
//}
//
//func (p *Pet) speak() string {
//	return fmt.Sprintf("my name is %v", p.name)
//}
//
//func (p *Pet) Name() string {
//	return p.name
//}
//
//func NewDog(name, breed string) *Dog {
//	d := &Dog{
//		Pet:   Pet{name: name},
//		Breed: breed,
//	}
//	d.speaker = d.speak
//	return d
//}
//
//func (d *Dog) speak() string {
//	return fmt.Sprintf("%v and I am a %v", d.Pet.speak(), d.Breed)
//}
//
//func DoPetDog3() {
//	d := NewDog("spot", "pointer")
//	fmt.Println(d.Name())
//	fmt.Println(d.Speak())
//	d.Play()
//}


//https://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html