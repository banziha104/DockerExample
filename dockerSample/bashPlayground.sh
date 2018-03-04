#!/bin/bash

echo "test"
hello="hello"
echo $hello
temp=50
if [ $temp -le 100 ]; then
    echo $temp
fi

for i in {1..100}
do
if [ $i -le 50 ]; then
    echo "$i is less Than 50"
fi
done
