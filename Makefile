.PHONY: hashing http

projectpath = ${PWD}
glidepath = ${PWD}/vendor/github.com/Masterminds/glide

gobench2csv:
	go build -o build/gobench2csv cmd/gobench2csv/main.go

hashing: 
	cd hashing;go test -benchmem -bench . > hashing.results

http:
	go build -o build/http/evio http/evio.go

target:
	@go build

test:
	@go test

integration: test
	@go test -tags=integration

linux:
	env GOOS=linux GOARCH=amd64 go build -o benchmarks-linux-amd64

mac:
	env GOOS=darwin GOARCH=amd64 go build -o benchmarks-mac-amd64

windows:
	env GOOS=windows GOARCH=amd64 go build -o benchmarks-windows-amd64.exe

$(glidepath)/glide:
	git clone https://github.com/Masterminds/glide.git $(glidepath)
	cd $(glidepath);make build
	cp $(glidepath)/glide .

libs: $(glidepath)/glide
	$(glidepath)/glide install
	R CMD BATCH plotting/setup.r

deps: libs
