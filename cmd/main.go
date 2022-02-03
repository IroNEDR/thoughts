package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/handlers"
	"github.com/IroNEDR/thoughts/internals/renderer"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
)

const port = ":9090"

var (
	app     config.AppConfig
	th      handlers.ThoughtHandler
	rd      renderer.Renderer
	session *scs.SessionManager
)

func main() {
	err := AppSetup()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("running on port %s", port)
		log.Fatal(
			http.ListenAndServe(port, csrf.Protect(app.CSRFkey)(
				app.Session.LoadAndSave(
					http.HandlerFunc(Serve)))))
		wg.Done()
	}()

	wg.Wait()
}

func AppSetup() error {
	app.Env = os.Getenv("ENVIRONMENT")
	app.IsProd = true
	if app.Env == "" {
		log.Println("Warning: server is running in Development mode")
		app.Env = "dev"
		app.IsProd = false
	}
	file, err := os.Open(".env." + app.Env + ".json")
	if err != nil {
		return err
	}
	defer file.Close()
	csrf.Secure(app.IsProd)
	decoder := json.NewDecoder(file)
	env := config.Environment{}
	err = decoder.Decode(&env)
	if err != nil {
		return err
	}
	app.CSRFkey = []byte(env.CSRF_KEY)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProd
	app.Session = session

	rd = renderer.NewRenderer(&app)
	tcache, err := rd.CreateTemplateCache()
	if err != nil {
		return err
	}
	app.TemplCache = tcache
	th = handlers.NewThoughtHandler(&app, &rd)
	return nil
}
