package api

import (
	"net/http"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
var Routes = []Route{
	Route{
		Name:        "GetIndex",
		Method:      http.MethodGet,
		Pattern:     "/",
		HandlerFunc: GetIndex,
	},
	Route{
		Name:        "PostMine",
		Method:      http.MethodPost,
		Pattern:     "/mine",
		HandlerFunc: PostMine,
	},
	Route{
		Name:        "PostTransaction",
		Method:      http.MethodPost,
		Pattern:     "/transactions",
		HandlerFunc: PostTransaction,
	},
}
