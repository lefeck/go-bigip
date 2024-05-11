#!/bin/bash

#filename="data.txt"
#
#while IFS= read -r line; do
#  touch ${line}.go
##  url="${line}"
##substring=$(echo "${url}" | awk -F 'profile/' '{print $2}' | cut -d'?' -f1)
##echo "${substring}"
#done < "$filename"
#

for file in profile_*.go; do
    mv "$file" "${file/profile_/}"
done