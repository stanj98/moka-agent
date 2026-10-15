package agent

import (
	"fmt"
	"strings"

	"github.com/stanj98/moka-agent/internal/memory"
	"github.com/stanj98/moka-agent/internal/provider"
)

//This mirrors Waku's core reasoning engine: it retrieves local context from SQLite, stitches history, interacts with the LLM provider, and evaluates whether new info should be memorized.

type Agent struct {
	Store    *memory.Store
	Provider *provider.Client
}

func NewAgent(s *memory.Store, p *provider.Client) *Agent {
	return &Agent{Store: s, Provider: p}
}

func (a *Agent) Run(userInput string) (string, error) {
	//1. Fetch relevant memory context
	memContext, err := a.Store.GetRelevantMemories()
	if err != nil {
		memContext = ""
	}

	systemPrompt := "You are Moka, a concise helpful local AI assistant.\n"
	if memContext != "" {
		systemPrompt += fmt.Sprintf("\nRelevant content from memory:\n%s", memContext)
	}

	//2. Build conversation payload
	messages := []provider.Message{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: userInput},
	}

	//3. Call provider loop
	reply, err := a.Provider.Chat(messages)
	if err != nil {
		return "", err
	}

	//4. Simple memory gate rule (e.g. if user says that "remember that....")
	if strings.Contains(strings.ToLower(userInput), "remember that") {
		_ = a.Store.AddMemory(userInput, "fact")
	}
	return reply, nil

}
