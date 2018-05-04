.PHONY: hashing http queues

projectpath = ${PWD}
glidepath = ${PWD}/vendor/github.com/Masterminds/glide

gobench2csv:
	go build -o build/gobench2csv cmd/gobench2csv/main.go

hashing: results
	@rm -rf ./results/hashing.*
	@go test ./hashing -benchmem -bench=. | tee ./results/hashing.log
	@Rscript plotting/gobench_multi_nsop.r ./results/hashing.log ./results/hashing.png

queues: results
	@rm -rf ./results/queues.*
	@go test ./queues -benchmem -bench=. | tee ./results/queues.log
	@Rscript plotting/gobench_single_nsop.r ./results/queues.log ./results/queues.png

http:
	go build -o build/http/evio http/evio.go

target:
	@go build

test:
	@go test

results:
	@mkdir results

$(glidepath)/glide:
	git clone https://github.com/Masterminds/glide.git $(glidepath)
	cd $(glidepath);make build
	cp $(glidepath)/glide .

libs: $(glidepath)/glide
	$(glidepath)/glide install
	#R CMD BATCH plotting/setup.r
	Rscript plotting/setup.r

deps: libs
