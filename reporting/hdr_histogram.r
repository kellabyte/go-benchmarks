#Rscript hdr_histogram.r histogramfile1 histogramfile2 ... histogramfileN NumberOfNines output.png
args = commandArgs(trailingOnly=TRUE)
output = args[length(args)]
nines = as.numeric(args[length(args) -1])
files = args[1:(length(args)-2)]

library(tidyverse)
library(drlib)
library(stringr)
library(tidyr)

plot = data_frame(FileName = files) %>%
  mutate(lines = map(FileName, ~ data.frame(line = readr::read_lines(.x, skip = 2)))) %>%
  unnest(lines) %>%
  mutate(line = trimws(line)) %>%
  separate(line, c("Value", "Percentile", "TotalCount", "OneOnOneMinusPercentile"), "\\s+") %>%
  mutate(Value = as.numeric(Value),
         Percentile = as.numeric(Percentile),
         TotalCount = as.numeric(TotalCount),
         OneOnOneMinusPercentile = as.numeric(OneOnOneMinusPercentile)) %>%
  mutate(Nines = log10(1-Percentile) / log10(.1)) %>%
  filter(!is.na(OneOnOneMinusPercentile)) %>%
  filter(Nines <= nines + .02) %>%
  ggplot(aes(Nines, Value, color=FileName)) +
  geom_line() + 
  scale_x_continuous(breaks = 0:(nines), labels=paste0((1-(.1^(0:(nines)))) * 100, "%"), limits = c(0, nines + .02)) + 
  labs(x="Percentile",
       y="Latency",
       title="Latency by Percentile Distribution") +
  theme_minimal() +
  theme(legend.position = "bottom") +
  theme(text = element_text(size=20))
  

ggsave(output, plot, width = 16, height = 4)
