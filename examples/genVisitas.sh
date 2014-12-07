#!/bin/bash

echo "year,month,day,visits"
for ddd in $(seq $1)
do
	fecha=$(date +"%Y,%m,%d" -d "-$ddd days")
	echo $fecha,$RANDOM
done
