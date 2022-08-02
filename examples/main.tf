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

resource "logiprovider_service" "all" {
  name = "KBTEST57" 
  created = timestamp()
  service_type_name = "Recurring Service"
  is_active = true
  is_tax_exempt = false
  is_inclusive_taxes = false
  default_service_status_type_name = "Active"
  description = "Test"
  service_category_name = "Default"
  service_base_type_name = "Recurring Charge"
}

output "all" {
  value=resource.logiprovider_service.all
}
