terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.8.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

resource "google_compute_network" "vpc_network" {
  name = "poc-network"
}

resource "google_compute_instance" "vm_instance" {
  name = "poc-instance"
  machine_type = var.machine_type
  tags = ["poc", "k8s"]

  boot_disk {
    initialize_params {
      image = var.machine_image
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.name
    access_config {
      
    }
  }

  metadata = {
    ssh-keys = "ubuntu:${file("~/workspace/keys/k8s-sandbox.pub")}"
    startup-script = file("userdata.tpl")
  }
}

resource "google_compute_firewall" "rules" {
  name        = "my-firewall-rule"
  network     = google_compute_network.vpc_network.name
  description = "Creates firewall rule targeting tagged instances"

  allow {
    protocol  = "tcp"
    ports     = ["22"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags = ["poc", "k8s"]
}