package decorator

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "I am component!"
}

type ConcreteDecorator struct {
	component Component
}

func (d *ConcreteDecorator) Operation() string {
	return "<strong>" + d.component.Operation() + "</strong>"
}
