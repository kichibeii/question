#!/bin/bash

# building apps
go build

# extract document per line
sed 's/ /\n/g' docs.txt  > docs_extract.txt

# run app 
./number1
