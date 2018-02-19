package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Router

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes{
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).HandlerFunc(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Router{
		"Index",
		"GET",
		"/",
		Index,
	},
	Router{
		"MovieList",
		"GET",
		"/peliculas",
		MovieList,
	},
	Router{
		"MovieShow",
		"GET",
		"/pelicula/{id}",
		MovieShow,
	},
	Router{
		"MovieAdd",
		"POST",
		"/pelicula",
		MovieAdd,
	},
	Router{
		"MovieUpdate",
		"PUT",
		"/pelicula/{id}",
		MovieUpdate,
	},
	Router{
		"MovieRemove",
		"DELETE",
		"/pelicula/{id}",
		MovieRemove,
	},
}
