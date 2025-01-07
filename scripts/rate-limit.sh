#!/bin/bash

set -euo pipefail

temp_file=$(mktemp)

# Send 20 concurrent GET requests and store status codes
for _ in {1..20}; do
  curl -s -o /dev/null -w "%{http_code}\n" https://osscontribute.com/repos >> "$temp_file" &
done
wait

echo "Summary of HTTP Status Codes:"
sort "$temp_file" | uniq -c | awk '{print "Status Code " $2 ": " $1}'

rm "$temp_file"
