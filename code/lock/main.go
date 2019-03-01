package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	userAges := &UserAges{
		ages: make(map[string]int),
	}
	var wg *sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			userAges.Add("oldboy", i)
		}(i)
		wg.Done()
	}
	wg.Wait()

	for i := 0; i < 10000; i++ {
		go func(i int) {
			age := userAges.Get("oldboy")
			fmt.Println(age)
		}(i)
	}
}
