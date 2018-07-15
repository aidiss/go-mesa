package main

import (
	"sync"
	"testing"
)

func TestStep(t *testing.T) {
	var agents []Agent
	agent := Agent{1, "My name is", 5}
	agents = append(agents, agent)
	var wg sync.WaitGroup
	wg.Add(1)
	go agent.step(agents)

}
