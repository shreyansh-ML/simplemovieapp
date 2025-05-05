all: install-swagger generate-specs

install-swagger:
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
generate-specs:
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
	GO111MODULE=off swagger generate spec -o ./swagger.json --scan-models
