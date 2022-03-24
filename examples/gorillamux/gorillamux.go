package gorillamux

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	regexstruct "github.com/paul-nelson-baker/golib-regexstruct"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.MatcherFunc(matcher).Handler(http.HandlerFunc(handler))
	apiRouter.HandleFunc("/", handler).MatcherFunc(matcher)
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	pathNamespace := mux.Vars(r)["namespace"]
	pathType := mux.Vars(r)["type"]
	fmt.Fprintln(w, pathNamespace)
	fmt.Fprintln(w, pathType)
}

func matcher(r *http.Request, rm *mux.RouteMatch) bool {
	// Handle awkward URL paths, such as:
	// - https://www.terraform.io/internals/provider-registry-protocol#list-available-versions
	regex := regexp.MustCompile(`api/(?P<namespace>(?:\w|-|\/)+)/(?P<type>(?:\w+|-)+)/versions/?`)
	matches, isMatch := regexstruct.RegexMatchMap(regex, r.URL.Path)
	if !isMatch {
		return false
	}
	rm.Vars = make(map[string]string)
	for key, value := range matches {
		rm.Vars[key] = value
	}
	return true
}
