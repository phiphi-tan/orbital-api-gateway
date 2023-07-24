#!/bin/bash


go tool pprof -http=localhost:8055 http://localhost:6060/debug/pprof/profile?seconds=5

# Run the curl commands in parallel
for i in {1..10}
do
    curl --location --request POST 'http://127.0.0.1/mathsvc/add' --header 'Content-Type: application/json' --data '{"first": i, "second": i}' &
done

# Wait for all background processes to finish
wait
