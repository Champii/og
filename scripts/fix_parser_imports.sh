#!/bin/bash

sed -i -e 's/"strconv"/"strconv"\n\t"strings"/g' parser/golang_parser.go