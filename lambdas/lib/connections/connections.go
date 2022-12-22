package connections

import (
	"errors"
	"sync"
)

type Connections struct {
	mu            sync.Mutex
	ConnectionIds []string
}

func (stack *Connections) Pop() (string, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	count := len(stack.ConnectionIds)
	if count == 0 {
		return "", errors.New("no more elements")
	}

	connectionId := stack.ConnectionIds[count-1]
	stack.ConnectionIds = stack.ConnectionIds[:count-1]
	return connectionId, nil
}

func (stack *Connections) Len() int {
	stack.mu.Lock()
	defer stack.mu.Unlock()
	return len(stack.ConnectionIds)
}
