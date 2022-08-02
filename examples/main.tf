terraform {
  required_providers {
    logiprovider = {
      version = "0.2"
      source  = "logisense.com/service/logiprovider"
    }
  }
}

// TODO: ENTER ACCESS DETAILS HERE (PREFERABLY THROUGH ENVIRONMENT VARIABLES)
provider "logiprovider" {
  username = "EXAMPLE" 
  password = "EXAMPLE"
  client_id = "EXAMPLE"
  host = "EXAMPLE"
}

resource "logiprovider_service" "all" {
  name = "IKBTEST105" 
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

