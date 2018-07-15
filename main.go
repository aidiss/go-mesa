package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"

	"github.com/icrowley/fake"
)

func main() {
	agentCount := 100
	stepCount := 100

	var agents []Agent
	model := Model{"Model name", agentCount}

	// Create agents
	for i := 0; i < model.agentCount; i++ {
		agent := Agent{i, fake.FirstName(), 5}
		fmt.Println(agent)
		agents = append(agents, agent)
	}

	// Print out all created agents
	for _, agent := range agents {
		fmt.Println(agent)
	}

	// Run model
	for i := 0; i < stepCount; i++ {
		// Run each agent
		for i := range agents {
			go agents[i].step(agents)
			fmt.Println(agents[i])
		}
	}

	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Wealth < agents[j].Wealth
	})

	// Print out a report
	for _, agent := range agents {
		println(agent.Name, agent.Wealth)
	}

	a, err := json.Marshal(agents)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", a)

	err = ioutil.WriteFile("output1.json", a, 0644)
}

// Model contains agents, their context and scheduler
type Model struct {
	name       string
	agentCount int
}

/// Agent has to move
type Agent struct {
	UniqueID int    `uniqueId`
	Name     string `name`
	Wealth   int    `wealth`
}

func (agent *Agent) step(agents []Agent) {
	if agent.Wealth <= 0 {
		println("Done")
		return
	}
	agentCount := len(agents)
	randomIndex := rand.Intn(agentCount)
	agents[randomIndex].Wealth += 1
	agent.Wealth -= 1
}
