// Define data sources
package datasource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lc "github.com/ibrahim925/LogiCore"
)

// Defines schema for data source response body
// This data source gets all the services from LSBS
func Services() *schema.Resource {
	return &schema.Resource{
		ReadContext: servicesRead,
		Schema: map[string]*schema.Schema{
			"tracking_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true},
						"service_type_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_type_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_tax_exempt": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_inclusive_taxes": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"default_service_status_type_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"default_service_status_type_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_category_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_category_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_base_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_base_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"usage_frequency": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"usage_frequency_type_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"usage_frequency_type_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Get the data from LSBS through the API client and output the data to the TF file 
func servicesRead(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Get API client instance that was defined in providerConfigure
	client := m.(*lc.Client)

	var diags diag.Diagnostics

	// Use API Client's GetServices method to get all ther services
	getServicesResponse, err := client.GetServices()
	if err != nil {
		return diag.FromErr(err)
	}

	setServices(d, getServicesResponse)

	return diags
}

// Set the terraform schema keys to their corresponding API response values
func setServices(d *schema.ResourceData, data *lc.ServicesGetResponse) {
	d.SetId(data.TrackingID)
	d.Set("tracking_id", data.TrackingID)
	d.Set("total_count", data.TotalCount)

	// Put all items in API response struct into a map[string]interface{} so that they can be read by Terraform
	transformedItems := transformItems(&data.Items)

	// Set the map[string]interface{} representation of all the services to "items"
	d.Set("items", *transformedItems)
}

func transformItems(items *[]lc.Item) *[]map[string]interface{} {
	transformed := make([]map[string]interface{}, len(*items))

	for i, item := range *items {
		transformed[i] = make(map[string]interface{})
		transformed[i]["identity"] = item.Identity
		transformed[i]["name"] = item.Name
		transformed[i]["description"] = item.Description
		transformed[i]["service_type_id"] = item.ServiceTypeID
		transformed[i]["service_type_name"] = item.ServiceTypeName
		transformed[i]["created"] = item.Created
		transformed[i]["is_active"] = item.IsActive
		transformed[i]["is_tax_exempt"] = item.IsTaxExempt
		transformed[i]["is_inclusive_taxes"] = item.IsInclusiveTaxes
		transformed[i]["default_service_status_type_id"] = item.DefaultServiceStatusTypeID
		transformed[i]["default_service_status_type_name"] = item.DefaultServiceStatusTypeName
		transformed[i]["service_category_id"] = item.ServiceCategoryID
		transformed[i]["service_category_name"] = item.ServiceCategoryName
		transformed[i]["service_base_id"] = item.ServiceBaseID
		transformed[i]["service_base_name"] = item.ServiceBaseName
		transformed[i]["usage_frequency"] = item.UsageFrequency
		transformed[i]["usage_frequency_type_id"] = item.UsageFrequencyTypeID
		transformed[i]["usage_frequency_type_name"] = item.UsageFrequencyTypeName
	}

	return &transformed
}
