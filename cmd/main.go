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
	app            config.AppConfig
	staticHandler  http.Handler
	th             handlers.ThoughtHandler
	rd             renderer.Renderer
	sessionManager *scs.SessionManager
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
		mux := setupRouter()
		log.Fatal(http.ListenAndServe(port, mux))
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

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.IsProd
	app.SessionManager = sessionManager

	staticHandler = http.FileServer(http.Dir("./static/"))
	rd = renderer.NewRenderer(&app)
	tcache, err := rd.CreateTemplateCache()
	if err != nil {
		return err
	}
	app.TemplCache = tcache
	th = handlers.NewThoughtHandler(&app, rd)
	return nil
}
