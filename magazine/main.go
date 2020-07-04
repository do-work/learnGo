package magazine

import (
	"fmt"
)

func main() {
	subscriber := Subscriber{Name: "John Doe"}
	subscriber.Street = "123 Test Street"

	fmt.Println("Subscriber home address", subscriber.Street)

	employee := Employee{Name: "Jane Doe"}
	employee.Street = "321 Test Street"

	fmt.Println("Employee home address", employee.Street)
}
