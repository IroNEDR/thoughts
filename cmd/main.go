package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/handlers"
	"github.com/IroNEDR/thoughts/internals/renderer"
	"github.com/gorilla/csrf"
)

const port = ":9090"

var app config.AppConfig
var th handlers.ThoughtHandler
var rd renderer.Renderer

func main() {
	err := AppSetup()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("running on port %s", port)
		log.Fatal(http.ListenAndServe(port, csrf.Protect(app.CSRFkey)(http.HandlerFunc(Serve))))
		wg.Done()
	}()

	wg.Wait()
}

func AppSetup() error {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Println("Warning: server is running in Development mode")
		environment = "dev"
		csrf.Secure(false)
	}
	file, err := os.Open(".env." + environment + ".json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	env := config.Environment{}
	err = decoder.Decode(&env)
	if err != nil {
		return err
	}
	app.CSRFkey = []byte(env.CSRF_KEY)
	rd = renderer.NewRenderer(&app)
	tcache, err := rd.CreateTemplateCache()
	if err != nil {
		return err
	}
	app.TemplCache = tcache
	th = handlers.NewThoughtHandler(&app, &rd)
	return nil
}
