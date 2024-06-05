#!/bin/bash


for file in wideip_*; do
  mv "$file" "${file/wideip_/}"
done