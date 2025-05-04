module example.com

go 1.22.2

replace example.com/handlers => ./handlers

replace example.com/model => ./model

require example.com/handlers v0.0.0-00010101000000-000000000000

require example.com/model v0.0.0-00010101000000-000000000000 // indirect

require github.com/gorilla/mux v1.8.0
