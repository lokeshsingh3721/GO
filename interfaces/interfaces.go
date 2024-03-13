package main

import (
	"fmt"
)

// // Shape interface defines methods for common shapes
// type Shape interface {
//   Area() float64
//   Perimeter() float64
// }

// // Square struct implements Shape
// type Square struct {
//   side float64
// }

// func (s Square) Area() float64 {
//   return s.side * s.side
// }

// func (s Square) Perimeter() float64 {
//   return 4 * s.side
// }

// // Rectangle struct also implements Shape
// type Rectangle struct {
//   length float64
//   width  float64
// }

// func (r Rectangle) Area() float64 {
//   return r.length * r.width
// }

// func (r Rectangle) Perimeter() float64 {
//   return 2 * (r.length + r.width)
// }

// // Circle struct can also implement Shape with its own calculations
// type Circle struct {
//   radius float64
// }

// func (c Circle) Area() float64 {
//   return 3.14 * c.radius * c.radius
// }

// func (c Circle) Perimeter() float64 {
//   return 2 * 3.14 * c.radius
// }

// func main() {
//   square := Square{5}
//   rectangle := Rectangle{4, 6}
//   // Both square and rectangle can be used with calculateArea because they implement Shape
//   fmt.Println("Square Area:", calculateArea(square))
//   fmt.Println("Rectangle Area:", calculateArea(rectangle))

//   //  can add a Circle struct later and use it with calculateArea seamlessly
//   // as long as it implements the Shape interface
// }

// func calculateArea(s Shape) float64 {
//   return s.Area()
// }

type area struct {
	width int 
	height int 
}

type address struct {
	city string 
	state string 
	country string 
}

type house struct{
	area area
	address address
}

type getArea interface {
	calcArea() int 
}

type getAddress interface {
	getYourAddress()  
}

func (a area) calcArea() int {
	return a.height * a.width
}

func (a address) getYourAddress()  {
 fmt.Printf("you live in city %s \n",a.city)
}

func main(){

	h1:=house{
		area{
			height: 20,
			width: 20,
		},
	address{
		city: "faridabad",
		state: "haryana",
		country: "India",
	},
	}

	area:=h1.area.calcArea()
	h1.address.getYourAddress()

	fmt.Printf("area is %d ",area)

}