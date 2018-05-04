.PHONY: hashing http queues json

projectpath = ${PWD}
glidepath = ${PWD}/vendor/github.com/Masterminds/glide
easyjsonpath = ${PWD}/vendor/github.com/mailru/easyjson

gobench2csv:
	@go build -o build/gobench2csv cmd/gobench2csv/main.go

hashing: results
	@rm -rf ./results/hashing.*
	@go test ./hashing -benchmem -bench=. | tee ./results/hashing.log
	@Rscript plotting/gobench_multi_nsop.r ./results/hashing.log ./results/hashing-multi.png
	@Rscript plotting/gobench_histogram_nsop.r ./results/hashing.log ./results/hashing-histogram.png

queues: results
	@rm -rf ./results/queues.*
	@go test ./queues -benchmem -bench=. | tee ./results/queues.log
	@Rscript plotting/gobench_single_nsop.r ./results/queues.log ./results/queues.png

json: generate results
  @rm -rf ./results/json.*
	@go test ./json -benchmem -bench=. | tee ./results/json.log
  @Rscript plotting/gobench_single_nsop.r ./results/json.log ./results/json.png

plot: results
	@Rscript plotting/gobench_multi_nsop.r ./results/hashing.log ./results/hashing-multi.png
	@Rscript plotting/gobench_histogram_nsop.r ./results/hashing.log ./results/hashing-histogram.png
	@Rscript plotting/gobench_single_nsop.r ./results/queues.log ./results/queues.png

http:
	@go build -o build/http/evio http/evio.go

target:
	@go build

test:
	@go test

results:
	@mkdir results

$(glidepath)/glide:
	@git clone https://github.com/Masterminds/glide.git $(glidepath)
	@cd $(glidepath);make build
	@cp $(glidepath)/glide .

$(easyjsonpath)/.root/bin/easyjson:
	@cd $(easyjsonpath); make build

easyjson: $(easyjsonpath)/.root/bin/easyjson
	@cp $(easyjsonpath)/.root/bin/easyjson .

libs: $(glidepath)/glide
	$(glidepath)/glide install
	sudo Rscript plotting/setup.r

deps: libs

generate: easyjson
	@go generate ./...
