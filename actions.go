package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session{
	session , err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return  session
}

func responseMovie(w http.ResponseWriter, status int, results Movie){
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie){
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession().DB("curso_go").C("movies")


func Index( w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hola koli <a href='/contacto'></a>")
}

func MovieList	( w http.ResponseWriter, r *http.Request){

	var results []Movie
	err :=collection.Find(nil).Sort("year").All(&results)

	if(err != nil){
		log.Fatal(err)
	}else{
		fmt.Println("Resultados: ",results)
	}
	responseMovies(w,200,results)

}

func MovieShow ( w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if(err != nil){
		w.WriteHeader(404)
		return
	}

	responseMovie(w,200,results)

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

	responseMovie(w,200,movie_data)

}