#!/bin/bash

go tool pprof -http=localhost:8057 http://localhost:6060/debug/pprof/profile?seconds=3

# Function to execute the curl command
execute_curl() {
    curl --location --request POST 'http://127.0.0.1/timesvc/time' --header 'Content-Type: application/json' --data '{"twentyfourhour": true}'
}

# Run the curl commands in parallel
for i in {1..10}
do
    execute_curl &
done

# Wait for all background processes to finish
wait
