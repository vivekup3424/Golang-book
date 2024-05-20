#!bin/bash
for i in {1...50000}
do
    echo "Executing command $i"
    nc localhost 8080
done