package env

import (
	"os"
	"sync"
)

// environmentMiddleware contains the environment variables for the protected APIs
type environmentMiddleware struct {
	user       string
	userOk     bool
	password   string
	passwordOk bool
}

// InitApp initializes the environment variables for the application
//
// if the variables are not set default values will be used
func (e *environmentMiddleware) initMiddleware(wg *sync.WaitGroup) {
	MuApp.Lock()
	e.user, e.userOk = os.LookupEnv("API_USER")
	e.password, e.passwordOk = os.LookupEnv("API_PASSWORD")
	MuApp.Unlock()
	wg.Done()
}

func (e *environmentMiddleware) GetMiddlewareUser() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.userOk {
		val = "default_user"
	} else {
		val = e.user
	}
	return
}

func (e *environmentMiddleware) GetMiddlewarePassword() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.passwordOk {
		val = "default_password"
	} else {
		val = e.password
	}
	return
}
