.PHONY: hashing http json

projectpath = ${PWD}
glidepath = ${PWD}/vendor/github.com/Masterminds/glide
easyjsonpath = ${PWD}/vendor/github.com/mailru/easyjson

gobench2csv:
	@go build -o build/gobench2csv cmd/gobench2csv/main.go

hashing:
	@cd hashing;go test -benchmem -bench . > hashing.results

json: generate
	@cd json;go test -benchmem -bench . > json.results

http:
	@go build -o build/http/evio http/evio.go

target:
	@go build

test:
	@go test

$(glidepath)/glide:
	@git clone https://github.com/Masterminds/glide.git $(glidepath)
	@cd $(glidepath);make build
	@cp $(glidepath)/glide .

$(easyjsonpath)/.root/bin/easyjson:
	@cd $(easyjsonpath); make build

easyjson: $(easyjsonpath)/.root/bin/easyjson
	@cp $(easyjsonpath)/.root/bin/easyjson .

libs: $(glidepath)/glide
	@$(glidepath)/glide install
	@R CMD BATCH plotting/setup.r

deps: libs

generate: easyjson
	@go generate ./...
