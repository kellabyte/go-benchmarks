# Rscript gobench.r gobench.out output.png

# Parse the command line args.
args = commandArgs(trailingOnly=TRUE) # this line only works when you run this script from the command line.

library(tidyverse)
library(drlib)

plot = readr::read_delim(args[1], "\t", skip = 3, col_names = FALSE) %>%
  set_names("Name", "Ops", "NsPerOp", "MBPerS", "AllocatedBytesPerOp", "AllocationsPerOp") %>%

  # Clean up the text so we get numbers
  mutate(NsPerOp = as.numeric(str_replace(NsPerOp, " ns/op", "")),
         MBPerS = as.numeric(str_replace(MBPerS, " MB/s", "")),
         AllocatedBytesPerOp = as.numeric(str_replace(AllocatedBytesPerOp, " B/op", "")),
         AllocationsPerOp = as.numeric(str_replace(AllocationsPerOp, " allocs/op", "")),
         Ops = as.numeric(Ops)) %>%
  filter(!is.na(NsPerOp)) %>%

  # Strip out the "Benchmark" prefix from the benchmark name.
  mutate (Name = str_replace(Name, "Benchmark", "")) %>%

  # Split out the cores.
  separate(Name, c("Name", "Cores"), "-") %>%
  mutate(Cores = as.numeric(Cores)) %>%

  # Generate the plots.
  ggplot(aes(Name, NsPerOp)) + 
  theme(text = element_text(size=24)) +
  geom_col() + 
  rotate_x_labels(vjust = .5) +
  # labs(title="Title", subtitle="Sub title") +
  xlab("name") +
  ylab("nanoseconds per operation")
  
ggsave(args[2], plot, width = 16, height = 9)