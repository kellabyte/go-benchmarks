This repository includes various Go benchmarks of data structures and services split into sub directories. Please feel free to contribute your own. This repo is meant to be a collection for the community to work on together.

# Prerequisites
### macOS
```
brew install r
```

### Ubuntu
```
sudo apt install curl libcurl4-openssl-dev libxml2-dev r-base
```

### Dependencies
```
make deps
```

# Benchmarks
[Hashing algorithms](https://github.com/kellabyte/go-benchmarks/tree/master/hashing)
```
make hashing
```

[Json](https://github.com/kellabyte/go-benchmarks/tree/master/json)
```
make json
```

[Queue data structures](https://github.com/kellabyte/go-benchmarks/tree/master/queues)
```
make queues
```
# Setup
#### Hardware
```
[2]  Intel(R) Xeon(R) CPU E5-2670 v2 @ 2.50GHz
[16] 1600 MHz 8GB DDR-3 Single Rank DIMM
[1]  Raid Controller
[8]  Intel S3520 SSD 240GB
[1]  Broadcom 5720 1Gbps Ethernet Adapter
[2]  Intel Ethernet Server Adapter X520 10GbE Bonded (20GbE)
```

#### Network Performance
```
Interval       Transfer     Bandwidth
0.0-10.0 sec  21.3 GBytes  18.3 Gbits/sec
0.0-10.0 sec  21.2 GBytes  18.2 Gbits/sec
0.0-10.0 sec  21.4 GBytes  18.3 Gbits/sec
```

#### Disk Performance
```
sudo hdparm -Tt /dev/sda

/dev/sda:
  Timing cached reads:   21800 MB in  1.99 seconds = 10933.80 MB/sec
  Timing buffered disk reads: 2828 MB in  3.00 seconds = 942.53 MB/sec
```

#### Operating System
```
Distributor ID:	Ubuntu
Description:	Ubuntu 16.04.4 LTS
Release:	16.04
Codename:	xenial

Linux version 4.4.0-121-generic 
(buildd@lcy01-amd64-004) 
(gcc version 5.4.0 20160609 (Ubuntu 5.4.0-6ubuntu1~16.04.9) ) 
#145-Ubuntu SMP Fri Apr 13 13:47:23 UTC 2018
```
