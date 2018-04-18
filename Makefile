projectpath = ${PWD}
glidepath = ${PWD}/vendor/github.com/Masterminds/glide

hashing: deps
	cd hashing;go test -benchmem -bench .

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
