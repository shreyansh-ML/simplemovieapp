package handlers

import(
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct{
  l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
   return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
  h.l.Println("Inside Hello ServerHttp")
  d,err := io.ReadAll(r.Body)
  if err != nil{
    h.l.Println(" Not able to read the body")
    http.Error(w," not able to read body",http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w,"Hello %s",d)

}


