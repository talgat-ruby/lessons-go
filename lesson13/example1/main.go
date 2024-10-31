package main

import (
	"fmt"
	"sync"
)

func main() {
	scoreboard := NewMutexScoreboardManager()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		scoreboard.Update("Grifindor", 10)
		scoreboard.Update("Hufflepuff", 20)
		scoreboard.Update("Slytherin", 30)
		scoreboard.Update("Ravenclaw", 40)
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if gr, ok := scoreboard.Read("Grifindor"); ok {
			fmt.Println("Grifindor:", gr)
		}
		if hf, ok := scoreboard.Read("Hufflepuff"); ok {
			fmt.Println("Hufflepuff:", hf)
		}
		if sl, ok := scoreboard.Read("Slytherin"); ok {
			fmt.Println("Slytherin:", sl)
		}
		if rv, ok := scoreboard.Read("Ravenclaw"); ok {
			fmt.Println("Ravenclaw:", rv)
		}
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		scoreboard.Update("Grifindor", 1000)
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if gr, ok := scoreboard.Read("Grifindor"); ok {
			fmt.Println("Grifindor:", gr)
		}
	}()

	wg.Wait()
}

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}
func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}
func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
