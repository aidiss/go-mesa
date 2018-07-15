package main

import (
	"sync"
	"testing"
)

func TestStep(t *testing.T) {
	var agents []Agent
	model := Model{"a", agents, 5}
	agent := Agent{1, &model, "My name is", 5}
	model.agents = append(model.agents, agent)
	var wg sync.WaitGroup
	wg.Add(1)
	go agent.step(&wg)

}

func TestAgentIncreaseWealth(t *testing.T) {
	var agents []Agent
	model := Model{"a", agents, 5}
	agent := Agent{1, &model, "My name is", 5}
	agent.increaseWealth(1)
	if agent.wealth != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", agent.wealth, 6)
	}

}
