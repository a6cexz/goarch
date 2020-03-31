package main

import (
	"event/channel"
	"fmt"
	"log"
)

func main() {
	user1 := channel.NewUser("user1")
	user2 := channel.NewUser("user2")

	ch1 := channel.NewChannel()
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := channel.NewChannel()
	ch2.Subscribe(user2)

	pub := channel.NewPublisher()
	pub.AddChannel("ch1", ch1)
	pub.AddChannel("ch2", ch2)

	if err := pub.Send("HELLO!", "ch1"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "ch2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("To All"); err != nil {
		fmt.Printf("%v\n", err)
	}

	chs := pub.GetChannels()
	fmt.Printf("%v channels \n", len(chs))

	if err := pub.DeleteChannel("ch3"); err != nil {
		fmt.Printf("%v\n", err)
	}

	if err := pub.DeleteChannel("ch1"); err != nil {
		fmt.Printf("%v\n", err)
	}

	chs = pub.GetChannels()
	fmt.Printf("%v channels \n", len(chs))
}
