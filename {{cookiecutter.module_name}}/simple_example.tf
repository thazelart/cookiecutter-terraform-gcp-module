// This is a simple example 
resource "google_project_service" "compute" {
  project = "${var.gcp_project_id}"
  service = "compute.googleapis.com"
}

resource "google_compute_instance" "example" {
  project = "${var.gcp_project_id}"
  name = "simple-example-compute"
  machine_type = "f1-micro"
  zone = "europe-west4-a"

  labels = {"simple"= "example"}

  boot_disk {
    initialize_params {
      image = "centos-7"
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }
}
