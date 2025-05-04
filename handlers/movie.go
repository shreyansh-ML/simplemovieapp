package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/model"
	"github.com/gorilla/mux"
)

type Movie struct {
	l *log.Logger
}

func NewMovie(l *log.Logger) *Movie {
	return &Movie{l}
}

func (p *Movie) GetMovies(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle get Movies")

	lp := model.GetMovies()
	//err := lp.ToJSON(w)
	d, err := json.Marshal(lp)
	//err:= json.NewEncoder(w).Encode(lp)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	w.Write(d)
}
func (p *Movie) GetMovie(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle get Movies")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	lp := model.GetMovie(int32(id))
	//err := lp.ToJSON(w)
	d, err := json.Marshal(lp)
	//err:= json.NewEncoder(w).Encode(lp)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	w.Write(d)
}

func (p *Movie) AddMovie(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle add Movie")

	//movie := &model.Metadata{}
	//fmt.Printf("movie: %v\n", r.Body)
	movie := r.Context().Value(k).(model.Metadata)

	//movie.FromJSON(r.Body)
	fmt.Printf("movie hahhahah: %v\n", movie)

	// err := json.NewDecoder(r.Body).Decode(movie)
	// if err != nil {
	// 	http.Error(w, "Unable to decode json", http.StatusBadRequest)
	// 	return
	// }

	model.Add(&movie)

	w.WriteHeader(http.StatusCreated)
}
func (p *Movie) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle update Movie")

	//movie := &model.Metadata{}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	// in_ur, _ := url.Parse(r.URL.Path)
	// //err := json.NewDecoder(r.Body).Decode(movie)
	// fmt.Printf("in_ur: %v\n", in_ur)
	// id1, _ := strconv.Atoi(path.Base(r.URL.Path))
	// id2 := path.Base(in_ur.Path)
	// fmt.Printf("id1: %v\n", id1)
	// fmt.Printf("id2: %v\n", id2)
	// fmt.Printf("r.url %v\n", r.URL)
	// fmt.Printf("r.url.Path %v\n", r.URL.Path)
	//err := json.NewDecoder(r.Body).Decode(movie)
	//err = json.NewDecoder(r.Body).Decode(movie)
	movie := r.Context().Value(k).(model.Metadata)

	//err := movie.FromJSON(r.Body)
	fmt.Printf("movie hahhahah: %v\n", movie)
	d := model.Update(&movie, int32(id))
	d.ToJSON(w)
	//w.WriteHeader(http.StatusNoContent)
}
func (p *Movie) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle delete Movie")

	//movie := &model.Metadata{}
	//var id int32 = movie.ID
	//err := json.NewDecoder(r.Body).Decode(movie)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	model.DeleteMovie(int32(id))

	w.WriteHeader(http.StatusNoContent)
}

type KeyProduct string

var k = KeyProduct("validjson")

func (p *Movie) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := model.Metadata{}
		p.l.Println("Handle add Movie0")
		if r.ContentLength > 0 {
			fmt.Println("inside if middleware")
			p.l.Println("Handle add Movie0")
			err := prod.FromJSON(r.Body)
			if err != nil {
				p.l.Println("[ERROR] deserializing product", err)
				http.Error(rw, "Error reading product", http.StatusBadRequest)
				return
			}
		}
		// add the product to the context
		fmt.Printf("inside middleware:prod: %v\n", prod)
		p.l.Println("Handle add Movie2")
		ctx := context.WithValue(r.Context(), k, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

// ServeHTTP is the main handler for the movie service
// It will handle all the requests for the movie service
// It will call the appropriate handler based on the method
// It will return the appropriate response based on the method
// It will return the appropriate status code based on the method
// func (m *Movie) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
// 	case http.MethodGet:
// 		m.l.Println("Handle get Movies")
// 		m.getMovies(w)
// 		return
// 	case http.MethodPut:
// 		m.l.Println("Handle update Movie")
// 		m.updateMovie(w, r)
// 		return
// 	case http.MethodDelete:
// 		m.l.Println("Handle delete Movie")
// 		m.deleteMovie(w, r)
// 		return
// 	//case http.MethodPost:
// 	case http.MethodPost:
// 		m.l.Println("Handle add Movie")
// 		m.addMovie(w, r)
// 		return
// 		//if r.Method == http.MethodGet {
// 		//	m.getMovies(w, r)
// 		//	return
// 	}

// 	m.l.Println("Invalid method")
// 	//w.WriteHeader(http.StatusMethodNotAllowed)

// 	w.WriteHeader(http.StatusMethodNotAllowed)
// }
