package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	//"log"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session{
	session , err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return  session
}

var collection = getSession().DB("curso_go").C("movies")


var movies = Movies{
	Movie{"sin limite",2013,"Desconocido"},
	Movie{"Batman",1999,"Desco  sdo"},
	Movie{"A todo gas",2015,"Duan Anton"},
}


func Index( w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hola koli <a href='/contacto'></a>")
}

func MovieList	( w http.ResponseWriter, r *http.Request){

	json.NewEncoder(w).Encode(movies)

	//fmt.Fprintln(w,"Lista pelicula")
}

func MovieShow ( w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintln(w,"Has cargado la pelicula "+movie_id )
}

func MovieAdd ( w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if(err != nil){
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(movie_data)
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(200)

}