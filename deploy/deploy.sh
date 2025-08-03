#!/usr/bin/env bash

set -euo pipefail

server_ip=$(terraform output -raw server_ip)

ssh root@"${server_ip}" -- mkdir -p /root/oss-projects
scp docker-compose.yml root@"${server_ip}":/root/oss-projects/docker-compose.yml
scp -r conf root@"${server_ip}":/root/oss-projects/
ssh root@"${server_ip}" -- docker compose -f /root/oss-projects/docker-compose.yml up -d
