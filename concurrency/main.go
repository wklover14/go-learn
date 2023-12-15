package main

// first Go routine example
import (
	"fmt"
	"math/rand"
	"time"
)

var chanResources = make(chan int, 2)
var resources int = 10


func eat(num int) {
	time.Sleep(1 * time.Second)
	chanResources <- ( resources - num )
	resources -= num
	fmt.Println("I have eaten", num, "times")
}

func produce(num int) {
	time.Sleep(2 * time.Second)
	chanResources <- ( resources  + num )
	resources += num
	fmt.Println("I have produced", num, "times")
}

func main() {

	for resources > 0 {
		//Production
		var prod_quantity = rand.Intn(10)
		go produce(prod_quantity)

		//Consumption
		var con_quantity = rand.Intn(resources)
		go eat(con_quantity)

		resources = <-chanResources
		fmt.Println("########## There is ", resources, "left ########")
		time.Sleep(4 * time.Second)
	}
}
