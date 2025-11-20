package main

import (
	"fmt"

)

	type car struct{
		make string
		model string
		tyre wheel
	}

	type wheel struct{
		radius int
		material string
	}
func main(){
	toyota := car{
		make: "toyota",
		model: "toyota probox",
		tyre: wheel{
			radius:20,
			material:"steel",

		},
	}

	make := toyota.make
	material := toyota.tyre.material
	fmt.Println(make)
	fmt.Println(material)

}


