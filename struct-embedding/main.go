package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v:", b.num)
}

// Struct embedding (looks like a field without a name)
type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "str",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	// Call the method of the embedded struct as container embeds base
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// The embedded struct with methods can be passed onto other structs interface implemenations
	var d describer = co
	fmt.Println("describer:", d.describe())
}
