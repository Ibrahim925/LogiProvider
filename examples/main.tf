terraform {
  required_providers {
    logiprovider = {
      version = "0.2"
      source  = "logisense.com/service/logiprovider"
    }
  }
}

provider "logiprovider" {
  username = "admin"
  password = "admin"
  client_id = "044b8ad6006845c29446b2f18e5b5909"
  host = "https://vnexttrainingps.dev.logisensebilling.com"
}

data "logiprovider_services" "all" {

}

output "all" {
  value=data.logiprovider_services.all
}
