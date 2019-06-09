package main

import (
  faktory "github.com/contribsys/faktory/client"
)

func main() {

	client, _ := faktory.Open()
	job := faktory.NewJob("SomeJob", 1, 2, 3)
	_ = client.Push(job)
}

