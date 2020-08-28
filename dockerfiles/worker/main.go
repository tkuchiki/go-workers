package main

import (
	"fmt"

	"github.com/tkuchiki/go-workers"
)

func perform(msg *workers.Msg) {
	fmt.Println(msg.Jid())
	a, _ := msg.Args().Array()
	fmt.Println(a)
}

func main() {
	workers.Configure(map[string]string{
		"server":  "redis:6379",
		"process": "1",
	})

	workers.Process("default", perform, 1)
	workers.Run()
}
