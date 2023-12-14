package main

import (
	"fmt"
	"math/rand"
	"time"
)

func eat(num uint8) {
	time.Sleep(1 * time.Second)
	fmt.Println("I have eaten", num, "times")
}

func produce(num uint8) {
	time.Sleep(2 * time.Second)
	fmt.Println("I have produced", num, "times")
}

func main() {
	var resources = 10

	for resources > 0 {
		//Production
		var prod_quantity = rand.Intn(10)
		go produce(uint8(prod_quantity))
		resources += prod_quantity

		//Consumption
		var con_quantity = rand.Intn(resources)
		go eat(uint8(con_quantity))
		resources -= con_quantity

		fmt.Println("######################################")
		fmt.Println("There is ", resources, "left      #")
		fmt.Println("######################################")
		time.Sleep(2 * time.Second)
	}
}
