.PHONY: hashing http queues json time

r: reporting/setup.r
	@Rscript reporting/setup.r

$(GOPATH)/bin/glide:
	@echo "Install Glide..."
	@curl https://glide.sh/get | sh

$(GOPATH)/bin/easyjson: 
	@echo "Install Easyjson..."
	@go get -u github.com/mailru/easyjson/...

vendor: $(GOPATH)/bin/glide glide.yaml glide.lock
	@echo "Install deps..."
	@glide install

generate: $(GOPATH)/bin/easyjson
	@echo "Generate go files..."
	@go generate ./...

results:
	@echo "Create results folder..."
	@mkdir results

calls: results
	@rm -rf ./results/calls.*
	@go test ./calls -benchmem -bench=. | tee ./results/calls.log

calls-report: calls
	@Rscript reporting/gobench_single_nsop.r ./results/calls.log ./results/calls.png

hashing: results
	@rm -rf ./results/hashing.*
	@go test ./hashing -benchmem -bench=. | tee ./results/hashing.log

hashing-report: hashing
	@Rscript reporting/gobench_multi_nsop.r ./results/hashing.log ./results/hashing-multi.png
	@Rscript reporting/gobench_histogram_nsop.r ./results/hashing.log ./results/hashing-histogram.png

queues: results
	@rm -rf ./results/queues.*
	@go test ./queues/one_to_one -benchmem -bench=. | tee ./results/queues.log

queues-report: queues
	@Rscript reporting/gobench_single_nsop.r ./results/queues.log ./results/queues.png

json: vendor generate results
	@rm -rf ./results/json.*
	@go test ./json -benchmem -bench=. | tee ./results/json.log

json-report: json
	@Rscript reporting/gobench_multi_nsop.r ./results/json.log ./results/json-multi.png

time: results
	@rm -rf ./results/time.*
	@go test ./time -benchmem -bench=. | tee ./results/time.log

time-report: time
	@Rscript ./reporting/gobench_single_nsop.r ./results/time.log ./results/time.png
	@Rscript ./reporting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 1 results/time_p90.png
	@Rscript ./reporting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 2 results/time_p99.png
	@Rscript ./reporting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 3 results/time_p999.png
	@Rscript ./reporting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 4 results/time_p9999.png
	@Rscript ./reporting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 5 results/time_p99999.png
	@Rscript ./reporting/hdr_histogram.r ./results/nanotime.histogram ./results/hrtime.histogram 6 results/time_p999999.png

benchmark: hashing queues json time

reports: hashing-report queues-report json-report time-report calls-report

gobench2csv: cmd/gobench2csv/main.go
	@go build -o build/gobench2csv cmd/gobench2csv/main.go

plot: results
	@Rscript ./reporting/gobench_multi_nsop.r ./results/hashing.log ./results/hashing-multi.png
	@Rscript ./reporting/gobench_histogram_nsop.r ./results/hashing.log ./results/hashing-histogram.png
	@Rscript ./reporting/gobench_single_nsop.r ./results/queues.log ./results/queues.png
	@Rscript ./reporting/gobench_multi_nsop.r ./results/json.log ./results/json-multi.png

http: http/evio.go
	@go build -o build/http/evio http/evio.go
