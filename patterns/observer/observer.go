package observer

import "errors"

type Publisher interface {
	Subscribe(o Observer) (bool, error)
	Unsubscribe(o Observer) (bool, error)
	SetState(state string)
	Notify()
}

type Observer interface {
	Update(state string)
}

type ConcretePublisher struct {
	observers []Observer
	state     string
}

func NewPublisher() Publisher {
	return &ConcretePublisher{}
}

func (s *ConcretePublisher) Subscribe(observer Observer) (bool, error) {
	for _, o := range s.observers {
		if o == observer {
			return false, errors.New("observer already exists")
		}
	}
	s.observers = append(s.observers, observer)
	return true, nil
}

func (s *ConcretePublisher) Unsubscribe(observer Observer) (bool, error) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("observer not found")
}

func (s *ConcretePublisher) SetState(state string) {
	s.state = state
	s.Notify()
}

func (s *ConcretePublisher) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

type ConcreteObserver struct {
	state string
}

func (o *ConcreteObserver) Update(state string) {
	o.state = state
}
