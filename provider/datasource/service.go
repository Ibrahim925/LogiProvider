// Define data sources
package datasource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lc "github.com/ibrahim925/LogiCore"
)

// Defines schema for data source response body
// This data source gets the data for one specific service
func Service() *schema.Resource {
	return &schema.Resource{
		ReadContext: serviceRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tracking_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
				Type:     schema.TypeInt,
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
	}
}

// Get the data from LSBS through the API client and output the data to the TF file 
func serviceRead(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Get API client instance that was defined in providerConfigure
	client := m.(*lc.Client)

	// Get service id defined by client in the .tf file
	id := d.Get("id").(int)

	// Warnings are stored here
	var diags diag.Diagnostics

	// Use API Client's GetService method to retrieve a specific service from LSBS
	getServiceResponse, err := client.GetService(id)
	if err != nil {
		return diag.FromErr(err)
	}

	setService(d, getServiceResponse)

	return diags
}

// Set the terraform schema keys to their corresponding API response values
func setService(d *schema.ResourceData, data *lc.ServiceGetResponse) {
	d.Set("tracking_id", data.TrackingID)
	d.Set("identity", data.Instance.Identity)
	d.Set("name", data.Instance.Name)
	d.Set("description", data.Instance.Description)
	d.Set("service_type_id", data.Instance.ServiceTypeID)
	d.Set("service_type_name", data.Instance.ServiceTypeName)
	d.Set("created", data.Instance.Created)
	d.Set("is_active", data.Instance.IsActive)
	d.Set("is_tax_exempt", data.Instance.IsTaxExempt)
	d.Set("is_inclusive_taxes", data.Instance.IsInclusiveTaxes)
	d.Set("default_service_status_type_id", data.Instance.DefaultServiceStatusTypeID)
	d.Set("default_service_status_type_name", data.Instance.DefaultServiceStatusTypeName)
	d.Set("service_category_id", data.Instance.ServiceCategoryID)
	d.Set("service_category_name", data.Instance.ServiceCategoryName)
	d.Set("service_base_id", data.Instance.ServiceBaseID)
	d.Set("service_base_name", data.Instance.ServiceBaseName)
	d.Set("usage_frequency", data.Instance.UsageFrequency)
	d.Set("usage_frequency_type_id", data.Instance.UsageFrequencyTypeID)
	d.Set("usage_frequency_type_name", data.Instance.UsageFrequencyTypeName)
}
