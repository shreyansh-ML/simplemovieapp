package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"example.com/model"
)

type Movie struct {
	l *log.Logger
}

func NewMovie(l *log.Logger) *Movie {
	return &Movie{l}
}

func (p *Movie) getMovies(w http.ResponseWriter) {

	p.l.Println("Handle get Movies")

	lp := model.Get()
	//err := lp.ToJSON(w)
	d, err := json.Marshal(lp)
	//err:= json.NewEncoder(w).Encode(lp)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	w.Write(d)
}

func (p *Movie) addMovie(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle add Movie")

	movie := &model.Metadata{}

	err := json.NewDecoder(r.Body).Decode(movie)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
		return
	}

	model.Add(movie)

	w.WriteHeader(http.StatusCreated)
}
func (p *Movie) updateMovie(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle update Movie")

	movie := &model.Metadata{}
	in_ur, _ := url.Parse(r.URL.Path)
	//err := json.NewDecoder(r.Body).Decode(movie)
	fmt.Printf("in_ur: %v\n", in_ur)
	id1, _ := strconv.Atoi(path.Base(r.URL.Path))
	id2 := path.Base(in_ur.Path)
	fmt.Printf("id1: %v\n", id1)
	fmt.Printf("id2: %v\n", id2)
	fmt.Printf("r.url %v\n", r.URL)
	fmt.Printf("r.url.Path %v\n", r.URL.Path)
	//err := json.NewDecoder(r.Body).Decode(movie)
	err := json.NewDecoder(r.Body).Decode(movie)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
		return
	}
	//err := movie.FromJSON(r.Body)
	fmt.Printf("movie hahhahah: %v\n", movie)
	d := model.Update(movie, int32(id1))
	d.ToJSON(w)
	//w.WriteHeader(http.StatusNoContent)
}
func (p *Movie) deleteMovie(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle delete Movie")

	movie := &model.Metadata{}
	var id int32 = movie.ID
	err := json.NewDecoder(r.Body).Decode(movie)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
		return
	}

	model.DeleteMovie(id)

	w.WriteHeader(http.StatusNoContent)
}

// ServeHTTP is the main handler for the movie service
// It will handle all the requests for the movie service
// It will call the appropriate handler based on the method
// It will return the appropriate response based on the method
// It will return the appropriate status code based on the method
func (m *Movie) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		m.l.Println("Handle get Movies")
		m.getMovies(w)
		return
	case http.MethodPut:
		m.l.Println("Handle update Movie")
		m.updateMovie(w, r)
		return
	case http.MethodDelete:
		m.l.Println("Handle delete Movie")
		m.deleteMovie(w, r)
		return
	//case http.MethodPost:
	case http.MethodPost:
		m.l.Println("Handle add Movie")
		m.addMovie(w, r)
		return
		//if r.Method == http.MethodGet {
		//	m.getMovies(w, r)
		//	return
	}

	m.l.Println("Invalid method")
	//w.WriteHeader(http.StatusMethodNotAllowed)

	w.WriteHeader(http.StatusMethodNotAllowed)
}
