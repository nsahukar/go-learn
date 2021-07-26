package decorator

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzePrice := c.pizza.getPrice()
	return pizzePrice + 10
}
