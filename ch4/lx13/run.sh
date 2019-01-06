#!/bin/bash
key_word=China
for y in $(seq 1986 2019)
do	
	go run poster.go $y $key_word
done
