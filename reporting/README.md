# Reporting
This directory contains the R scripts that generate the benchmark reports.

# Generating reports
You can see examples of how the reports are generated in the `Makefile` using R scripts. You can find some sample result files in `results/samples`.

```
Rscript plotting/gobench_multi_nsop.r ./results/samples/hashing.log ./results/hashing-multi.png
```
_Figure 1_. Example command to generate a multi-graph report using R.

```
Rscript plotting/gobench_single_nsop.r ./results/samples/queues.log ./results/queues.png
```
_Figure 2_. Example command to generate a single graph report using R.

```
Rscript plotting/gobench_histogram_nsop.r ./results/samples/hashing.log ./results/hashing-histogram.png
```
_Figure 3_. Example command to generate a histogram graph using R.

```
Rscript ./plotting/hdr_histogram.r ./results/samples/nanotime.histogram ./results/hrtime.histogram 6 results/time_p999999.png
```
_Figure 4_. Example command to generate a latency distribution graph using HdrHistogram log file with R.
