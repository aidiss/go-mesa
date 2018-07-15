package main

import (
	"math/rand"
	"sync"

	"github.com/icrowley/fake"
)

func main() {
	agentCount := 5
	stepCount := 5

	var agents []Agent
	model := Model{"Model name", agents, agentCount}

	model.createAgents()

	// Step model
	for i := 0; i < stepCount; i++ {
		model.step()
	}

	// Print out a report
	for _, agent := range model.agents {
		println(agent.name, agent.wealth)
	}
}

// Model contains agents, their context and scheduler
type Model struct {
	name       string
	agents     []Agent
	agentCount int
}

func (model *Model) step() {
	var wg sync.WaitGroup
	for _, agent := range model.agents {
		new_agent := agent
		wg.Add(1)
		new_agent.step(&wg)
	}
	wg.Wait()
}

func (model *Model) createAgents() {
	for i := 0; i < model.agentCount; i++ {
		agent := Agent{i, model, fake.FirstName(), 5}
		model.agents = append(model.agents, agent)
		// fmt.Println(agent)
	}
}

/// Agent has to move
type Agent struct {
	unique_id int
	model     *Model
	name      string
	wealth    int
}

func (agent *Agent) decreaseWealth(n int) {
	agent.wealth -= n
}
func (agent *Agent) increaseWealth(n int) {

	agent.wealth += n
}

func (agent *Agent) step(wg *sync.WaitGroup) {
	// fmt.Println("Stepping", agent)
	// agentCount := len(agent.model.agents)
	// println(agentCount)
	if agent.wealth == 0 {
		return
	}
	agentCount := len(agent.model.agents)
	randomIndex := rand.Intn(agentCount)
	otherAgent := agent.model.agents[randomIndex]
	otherAgent.wealth += 1
	agent.wealth -= 5
	otherAgent.increaseWealth(1)
	agent.decreaseWealth(1)

	wg.Done()
}
