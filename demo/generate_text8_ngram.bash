#!/bin/bash

cat text8 | tr ' ' '\n' > text8_0
cat text8 | sed -s 's/^\s//' | tr ' ' '\n' > text8_1
cat text8 | sed -s 's/^\s anarchism//' | tr ' ' '\n' > text8_2

paste text8_0 text8_1 -d "_" > text8_bigram
paste text8_0 text8_1 text8_2 -d "_" > text8_trigram
