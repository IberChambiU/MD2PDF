package env

import "sync"

var (
	e    *environment = &environment{}
	once sync.Once
	wg   sync.WaitGroup
)

// Centralizer for environment variables
type environment struct {
	environmentApp
	environmentMiddleware
}

func Init() {
	once.Do(func() {
		e.init()
	})
}

func (e *environment) init() {
	wg.Add(2)
	go e.initApp(&wg)
	go e.initMiddleware(&wg)
	wg.Wait()
}

func Env() *environment {
	return e
}
