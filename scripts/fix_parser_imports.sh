#!/bin/bash

sed -i -e 's/"strconv"/"strconv"\n\t"strings"/g' parser/og_parser.go