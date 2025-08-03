#!/usr/bin/env bash

set -euo pipefail

server_ip=$(terraform output -raw server_ip)
ssh_opts="-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null"

ssh "${ssh_opts}" root@"${server_ip}" -- mkdir -p /root/oss-projects
scp "${ssh_opts}" docker-compose.yml root@"${server_ip}":/root/oss-projects/docker-compose.yml
scp "${ssh_opts}" -r conf root@"${server_ip}":/root/oss-projects/
ssh "${ssh_opts}" root@"${server_ip}" -- docker compose -f /root/oss-projects/docker-compose.yml up -d
