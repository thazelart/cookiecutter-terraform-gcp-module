provider "google-beta" {
  project       = "${var.gcp_project_id}"
}

provider "google" {
  project = "${var.gcp_project_id}"
}

provider "github" {
  organization = "dktunited"
}
