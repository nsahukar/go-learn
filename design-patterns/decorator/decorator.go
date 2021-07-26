package decorator

import "fmt"

func veggePizzaWithCheeseAndTomato() {
	veggePizza := &veggeMania{}
	
	// Add cheese topping
	veggePizzaWithCheese := &cheeseTopping{
		pizza: veggePizza,
	}
	// Add tomato topping
	veggePizzaWithCheeseAndTomato := &tomotoTopping{
		pizza: veggePizzaWithCheese,
	}

	fmt.Printf("VeggeMania pizza with tomato and cheese topping: $%d\n", veggePizzaWithCheeseAndTomato.getPrice())
}

func peppyPaneerPizzaWithCheese() {
	peppyPaneerPizza := &peppyPaneer{}

	// Add cheese topping
	peppyPaneerPizzaWithCheese := &cheeseTopping{
		pizza: peppyPaneerPizza,
	}

	fmt.Printf("PeppyPaneer Pizza with cheese topping: $%d\n", peppyPaneerPizzaWithCheese.getPrice())
}

func ExamplePizza() {
	veggePizzaWithCheeseAndTomato()
	peppyPaneerPizzaWithCheese()
}
