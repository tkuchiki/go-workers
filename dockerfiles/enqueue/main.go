package main

import "github.com/tkuchiki/go-workers"

func main() {
	workers.Configure(map[string]string{
		"server":  "redis:6379",
		"process": "1",
	})
	workers.Enqueue("default", "TestWorker", []int{1, 2})
}
