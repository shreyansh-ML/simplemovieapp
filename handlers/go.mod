module example.com/handlers

go 1.22.2

replace example.com/model => ../model

require example.com/model v0.0.0-00010101000000-000000000000

require github.com/gorilla/mux v1.8.0
