# Rscript gobench.r gobench.out output.png

# Parse the command line args.
args = commandArgs(trailingOnly=TRUE) # this line only works when you run this script from the command line.

library(tidyverse)
library(drlib)

plot = readr::read_delim(args[1], "\t", skip = 5, col_names = FALSE) %>%
  set_names("Name", "Ops", "NsPerOp", "MBPerS", "AllocatedBytesPerOp", "AllocationsPerOp") %>%

  # Clean up the text so we get numbers
  mutate(NsPerOp = as.numeric(str_replace(NsPerOp, " ns/op", "")),
         MBPerS = as.numeric(str_replace(MBPerS, " MB/s", "")),
         AllocatedBytesPerOp = as.numeric(str_replace(AllocatedBytesPerOp, " B/op", "")),
         AllocationsPerOp = as.numeric(str_replace(AllocationsPerOp, " allocs/op", "")),
         Ops = as.numeric(Ops)) %>%
  filter(!is.na(NsPerOp)) %>%

  # Split out the cores.
  separate(Name, c("Name", "Cores"), "-") %>%
  mutate(Cores = as.numeric(Cores)) %>%
  
  # Split out the bytes.
  separate(Name, c("Name", "Bytes"), "/") %>%
  mutate(Bytes = as.numeric(Bytes)) %>%
  
  ggplot(aes(x=reorder(Name,-NsPerOp), y=NsPerOp,
                 group=as.factor(Bytes), fill=as.factor(Bytes))) + 
  geom_histogram(stat="identity") +
  theme(axis.text.x = element_text(angle = 45, hjust = 1),
        axis.title.x=element_blank())

ggsave(args[2], plot, width = 16, height = 9)