package oopsample

import "fmt"

/*
Internal names (lowerCase) can be accessed from anywhere in the same package,
even if the package contains multiple structs across multiple files. If you find this unsettling, packages can also be as small as you need.

Composition, embedding and interfaces provide powerful tools for Object-Oriented Design in Go.

While Idiomatic Go requires a change in thinking, I am pleasantly surprised with how simple and concise Go code can be when playing to its strengths.
 */

type Part struct {
	name        string
	Description string
	NeedsSpare  bool
}

/*
Go uses packages for namespacing. Exported identifiers begin with a capital letter.
 To make an identifier internal to a package, we start it with a lowercase letter:
 */
/*
Notice that we don’t prefix getters with Get (eg. GetName). Getters aren’t strictly necessary either,
especially with strings. When the need arises, we can always change the Name field to use a custom type that satisfies the Stringer interface.
 */
func (part Part) Name() string {
	return part.name
}

func (part *Part) SetName(name string) {
	part.name = name
}



type Parts []Part

/*
We can declare methods on any user-defined type,
so Parts can have all the behavior of slices, plus our own custom behavior.
 */
func (parts Parts) Spares() (spares Parts) {
	for _, part := range parts {
		if part.NeedsSpare {
			spares = append(spares, part)
		}
	}
	return spares
}

type Bicycle struct {
	Size string
	Parts
}

var (
	RoadBikeParts = Parts{
		{"chain", "10-speed", true},
		{"tire_size", "23", true},
		{"tape_color", "red", true},
	}

	MountainBikeParts = Parts{
		{"chain", "10-speed", true},
		{"tire_size", "2.1", true},
		{"front_shock", "Manitou", false},
		{"rear_shock", "Fox", true},
	}

	RecumbentBikeParts = Parts{
		{"chain", "9-speed", true},
		{"tire_size", "28", true},
		{"flag", "tall and orange", true},
	}
)

func PartBikeSample() {

	roadBike := Bicycle{Size: "L", Parts: RoadBikeParts}
	mountainBike := Bicycle{Size: "L", Parts: MountainBikeParts}
	recumbentBike := Bicycle{Size: "L", Parts: RecumbentBikeParts}

	fmt.Println(roadBike.Spares())
	fmt.Println(mountainBike.Spares())
	fmt.Println(recumbentBike.Spares())

	/*
	Parts behaves like a slice. Getting the length, slicing the slice, or combining multiple slices all works as usual.
	 */
	comboParts := Parts{}
	comboParts = append(comboParts, mountainBike.Parts...)
	comboParts = append(comboParts, roadBike.Parts...)
	comboParts = append(comboParts, recumbentBike.Parts...)

	fmt.Println(len(comboParts), comboParts[9:])
	fmt.Println(comboParts.Spares())

}

/*
Polymorphism in Go is provided by interfaces. They are satisfied implicitly, unlike Java or C#, so interfaces can be defined for code we don’t own.
 */

func (part Part) String() string {
	return fmt.Sprintf("%s: %s", part.Name, part.Description)
}

/*
Interface types can be used in the same places as other types.
Variables and arguments can take a Stringer, which accepts anything that implements the String() string method signature.
 */
type Stringer interface {
	String() string
}







