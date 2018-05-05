.PHONY: hashing http queues json time

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
	@go test ./queues -benchmem -bench=. -run=none -benchtime=3s | tee ./results/queues.log
	@Rscript plotting/gobench_single_nsop.r ./results/queues.log ./results/queues.png

json: generate results
	@rm -rf ./results/json.*
	@go test ./json -benchmem -bench=. | tee ./results/json.log
	@Rscript plotting/gobench_multi_nsop.r ./results/json.log ./results/json-multi.png

plot: results
	@Rscript plotting/gobench_multi_nsop.r ./results/hashing.log ./results/hashing-multi.png
	@Rscript plotting/gobench_histogram_nsop.r ./results/hashing.log ./results/hashing-histogram.png
	@Rscript plotting/gobench_single_nsop.r ./results/queues.log ./results/queues.png
	@Rscript plotting/gobench_multi_nsop.r ./results/json.log ./results/json-multi.png

http:
	@go build -o build/http/evio http/evio.go

time: results
	@rm -rf ./results/time.*
	@go test ./time -benchmem -bench=. | tee ./results/time.log
	@Rscript plotting/gobench_single_nsop.r ./results/time.log ./results/time.png
	@Rscript ./plotting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 1 results/time_p90.png
	@Rscript ./plotting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 2 results/time_p99.png
	@Rscript ./plotting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 3 results/time_p999.png
	@Rscript ./plotting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 4 results/time_p9999.png
	@Rscript ./plotting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 5 results/time_p99999.png
	@Rscript ./plotting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 6 results/time_p999999.png

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
	@cd $(easyjsonpath); go build -i -o .root/bin/easyjson ./easyjson

easyjson: $(easyjsonpath)/.root/bin/easyjson
	@cp $(easyjsonpath)/.root/bin/easyjson .

libs: $(glidepath)/glide
	$(glidepath)/glide install
	sudo Rscript plotting/setup.r

deps: libs

generate: easyjson
	@go generate ./...
