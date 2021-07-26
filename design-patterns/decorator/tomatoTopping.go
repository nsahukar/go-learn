package decorator

type tomotoTopping struct {
	pizza pizza
}

func (c *tomotoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
