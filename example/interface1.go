/* package main

import "fmt"

func main() {
	x := 10
	y := (&x)
	*y = 20
	fmt.Println(x)
} */
package main

import (
	"fmt"
)

type rectangle struct {
	a, b float64
}

func (r rectangle) display() {
	fmt.Println("Kenar Uzunlukları: ", r.a, "ve", r.b, "olan dikdörgenin...")
}
func (r rectangle) area() float64 {
	return r.a * r.b
}
func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

type shape interface {
	area() float64
	circumference() float64
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())
}

func main() {
	r1 := rectangle{3, 8}
	r1.display()
	fmt.Println("Alanı: ", r1.area())
	fmt.Println("Çevre: ", r1.circumference())
	fmt.Println("-------------------------")
	interfaceFunc(r1)
}

/*
package main

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}
*/
/*
func showFamily() []family {
	family1 := family{
		name:      "Richard",
		age:       26,
		isMarried: false,
	}
	family2 := family{
		name:      "John",
		age:       29,
		isMarried: true,
	}
	family3 := family{
		name:      "Elis",
		age:       18,
		isMarried: false,
	}
	family4 := family{
		name:      "Donald",
		age:       35,
		isMarried: true,
	}
	return []family{family1, family2, family3, family4}

}

func main() {
	families := showFamily()
	fmt.Println(families)
}
*/
