package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func main() {

	fmt.Println(Gallons(12.23))
	fmt.Println(Liters(1.2))
	fmt.Println(Milliliters(90.67123))

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

}
