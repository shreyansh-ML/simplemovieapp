package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product APi", log.LstdFlags)
	//hh:= handlers.NewHello(l)
	hh := handlers.NewMovie(l)
	//hh1 := handlers.NewMovie(l)
	//sm := http.NewServeMux()
	sm := mux.NewRouter()
	//getRouter := sm.Methods(http.MethodGet).Subrouter()
	//	sm.Handle("/", hh)
	//sm.Handle("/{id}", hh1)
	sm.Methods(http.MethodGet).Path("/movies").HandlerFunc(hh.GetMovies)
	//sm.Methods(http.MethodGet).Path("/movies/{id:[0-9]+}").HandlerFunc(hh.GetMovie)
	sm.Methods(http.MethodGet).Path("/{id:[0-9]{1,}}").HandlerFunc(hh.GetMovie)
	sm.Methods(http.MethodPut).Path("/{id:[0-9]{1,}}").Handler(http.HandlerFunc(hh.UpdateMovie))
	sm.Use(hh.MiddlewareValidateProduct)
	sm.Methods(http.MethodPost).Path("/").Handler(http.HandlerFunc(hh.AddMovie))
	sm.Methods(http.MethodDelete).Path("/{id:[0-9]{1,}}").Handler(http.HandlerFunc(hh.DeleteMovie))
	createServer := func() *http.Server {
		s := &http.Server{
			Addr:         ":9090",
			Handler:      sm,
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 3 * time.Second,
		}
		return s
	}
	ss := createServer()
	go func() {
		err := ss.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Received termination", sig)
	// http.HandleFunc("/hello" func(w http.ResponseWriter,r *http.Request){
	//   log.Println("inside handlefunc handler of helll0"
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//ctx,_:=context.WithDeadline(context.Background(),30)
	ss.Shutdown(ctx)

}
