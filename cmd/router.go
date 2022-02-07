package main

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/IroNEDR/thoughts/internals/middleware"
	"github.com/gorilla/csrf"
)

type route struct {
	regex   *regexp.Regexp
	handler http.Handler
	method  string
}

type ctxKey struct{}

func GetField(r *http.Request, idx int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[idx]
}

func newRoute(method, pattern string, handler http.Handler) route {
	return route{regexp.MustCompile("^" + pattern + "$"), handler, method}
}

func setupRoutes() []route {
	routes := []route{
		newRoute(http.MethodGet, "/", http.HandlerFunc(lh.Index)),
		newRoute(http.MethodGet, "/thoughts", http.HandlerFunc(th.List)),
		newRoute(http.MethodPost, "/thoughts", http.HandlerFunc(th.Create)),
		newRoute(http.MethodGet, "/thoughts/([^/]+)", http.HandlerFunc(th.Get)),
		newRoute(http.MethodPut, "/thoughts/([^/]+)", http.HandlerFunc(th.Update)),
		newRoute(http.MethodDelete, "/thoughts/([^/]+)", http.HandlerFunc(th.Delete)),
		newRoute(http.MethodGet, "/static/(.*)+", http.StripPrefix("/static", staticHandler)),
	}
	return routes
}

func serve(w http.ResponseWriter, r *http.Request) {
	routes := setupRoutes()
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := rd.LoadTemplate("notfound.page.tmpl")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func setupRouter() http.Handler {
	mux := middleware.RequestLogger(http.HandlerFunc(serve))
	mux = csrf.Protect(app.CSRFkey)(mux)
	mux = app.SessionManager.LoadAndSave(mux)
	return mux
}
