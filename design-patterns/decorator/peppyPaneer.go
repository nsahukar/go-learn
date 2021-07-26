package decorator

type peppyPaneer struct {
}

func (p *peppyPaneer) getPrice() int {
	return 20
}
