package main

import (
	"context"
	"net/http"
	"regexp"
	"strings"
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
		newRoute(http.MethodGet, "/", LoggingMiddleware(http.HandlerFunc(th.List))),
		newRoute(http.MethodPost, "/", LoggingMiddleware(http.HandlerFunc(th.Create))),
		newRoute(http.MethodGet, "/thoughts/([^/]+)", LoggingMiddleware(http.HandlerFunc(th.Get))),
		newRoute(http.MethodPut, "/thoughts/([^/]+)", LoggingMiddleware(http.HandlerFunc(th.Update))),
		newRoute(http.MethodDelete, "/thoughts/([^/]+)", LoggingMiddleware(http.HandlerFunc(th.Delete))),
	}
	return routes
}

func Serve(w http.ResponseWriter, r *http.Request) {
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
	http.NotFound(w, r)
}
