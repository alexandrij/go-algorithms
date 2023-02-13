package observer

func ExampleObserver() {
	publisher := NewPublisher()

	publisher.Subscribe(&ConcreteObserver{
		state: "Observer A",
	})
	publisher.Subscribe(&ConcreteObserver{
		state: "Observer B",
	})

	publisher.SetState("New State")
}
