package main

import "./pkg"

func main() {

	sub1 := &pkg.Subscriber{"Sub"}
	sub2 := &pkg.Subscriber{"Sub1"}
	sub3 := &pkg.Subscriber{"Sub2"}
	sub4 := &pkg.Subscriber{"Sub3"}

	chanel := pkg.Publisher {
		Name: "Publisher channel",
		Consumers: map[string]pkg.Consumer{},
	}

	chanel.Subscribe(sub1)
	chanel.Subscribe(sub2)
	chanel.Subscribe(sub3)
	chanel.Subscribe(sub4)
	chanel.Notify()

	chanel.UnSubscribe(sub3)
	chanel.Notify()
}
