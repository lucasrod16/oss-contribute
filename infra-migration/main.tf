terraform {
  required_providers {
    hcloud = {
      source = "hetznercloud/hcloud"
    }
  }
}

provider "hcloud" {}

data "http" "current_ip" {
  url = "https://checkip.amazonaws.com"
}

resource "hcloud_ssh_key" "default" {
  name       = "oss-projects-key"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "hcloud_firewall" "web_firewall" {
  name = "oss-projects-firewall"

  rule {
    direction  = "in"
    port       = "22"
    protocol   = "tcp"
    source_ips = ["${trimspace(data.http.current_ip.response_body)}/32"]
  }

  rule {
    direction  = "in"
    port       = "8080"
    protocol   = "tcp"
    source_ips = ["0.0.0.0/0", "::/0"]
  }
}

resource "hcloud_server" "web_server" {
  name        = "oss-projects-server"
  image       = "docker-ce"
  server_type = "cpx21"
  location    = "hil"

  ssh_keys     = [hcloud_ssh_key.default.id]
  firewall_ids = [hcloud_firewall.web_firewall.id]

  labels = {
    project = "oss-projects"
  }
}

output "server_ip" {
  value       = hcloud_server.web_server.ipv4_address
  description = "Public IP address of the server"
}

output "ssh_command" {
  value       = "ssh root@${hcloud_server.web_server.ipv4_address}"
  description = "SSH command to connect to the server"
}
