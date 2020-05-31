package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	sub1 := defaultSubscriber("John Doe")
	applyDiscount(&sub1)
	printInfo(sub1)

	sub2 := defaultSubscriber("Jane Doe")
	printInfo(sub2)
}
