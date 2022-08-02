// Define resources 
package resource

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lc "github.com/ibrahim925/LogiCore"
)

// Defines schema for resource request body
// This resource allows customers to perform Create, Update, and Delete operations on services in LSBS
func Service() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceCreate,
		ReadContext:   resourceServiceRead,
		UpdateContext: resourceServiceUpdate,
		DeleteContext: resourceServiceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created": {
				Type:	  schema.TypeString,
				Computed: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_tax_exempt": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_inclusive_taxes": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"default_service_status_type_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_category_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_base_type_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// Terraform calls this function when creating a new service
func resourceServiceCreate(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Get API client instance that was defined in providerConfigure
	client := m.(*lc.Client)

	// Warnings and errors will be stored here
	var diags diag.Diagnostics

	// Request body will be stored within this struct
	service := lc.ServicePostStruct{}

	service.Name = d.Get("name").(string)
	service.ServiceTypeName = d.Get("service_type_name").(string)
	service.IsActive = d.Get("is_active").(bool)
	service.IsTaxExempt = d.Get("is_tax_exempt").(bool)
	service.IsInclusiveTaxes = d.Get("is_inclusive_taxes").(bool)
	service.DefaultServiceStatusTypeName = d.Get("default_service_status_type_name").(string)
	service.Description = d.Get("description").(string)
	service.ServiceCategoryName = d.Get("service_category_name").(string)
	service.ServiceBaseTypeName = d.Get("service_base_type_name").(string)

	postResponse, err := client.CreateService(service)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set resource ID to a string represention of new service's "identity"
	serviceIDStr := strconv.Itoa(postResponse.Results.Items[0].Identity)
	d.SetId(serviceIDStr)

	// Set resource "created" to new service's "created"
	timeCreated := postResponse.Results.Items[0].Instance.Created
	d.Set("created", timeCreated)

	return diags
}

// Terraform will call this function to check the state of the world (LSBS) and compare it to the state of the 
// configuration file (desired state)
func resourceServiceRead(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Get API client instance that was defined in providerConfigure
	client := m.(*lc.Client)

	// Warnings and errors will be stored here
	var diags diag.Diagnostics

	// Retrieve service id and convert it to an int
	serviceIDStr := d.Id()
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		return diag.FromErr(err)
	}

	// Use service id to get a specific service use the API client's GetService method
	service, err := client.GetService(serviceID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set schema values
	setService(d, service)

	return diags

}

func resourceServiceUpdate(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Get API client instance that was defined in providerConfigure
	client := m.(*lc.Client)

	// Warnings and errors will be stored here
	var diags diag.Diagnostics

	// Retrieve service id and convert it to an int
	serviceIDStr := d.Id()
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if any of the values have been changed
	if d.HasChanges("name", "service_type_name", "created", "is_active", "is_tax_exempt", "is_inclusive_taxes", "default_service_status_type_name", "description") {
		// Create patch request body
		updatedService := lc.ServicePatchStruct{}

		updatedService.Name = d.Get("name").(string)
		updatedService.ServiceTypeName = d.Get("service_type_name").(string)
		updatedService.IsActive = d.Get("is_active").(bool)
		updatedService.IsTaxExempt = d.Get("is_tax_exempt").(bool)
		updatedService.IsInclusiveTaxes = d.Get("is_inclusive_taxes").(bool)
		updatedService.DefaultServiceStatusTypeName = d.Get("default_service_status_type_name").(string)
		updatedService.Description = d.Get("description").(string)

		// Call API Client's UpdateService method
		if _, err := client.UpdateService(serviceID, updatedService); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

// Terraform will call this function when deleting a service
func resourceServiceDelete(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Get API client instance that was defined in providerConfigure
	client := m.(*lc.Client)

	// Warnings and errors will be stored here
	var diags diag.Diagnostics

	// Retrieve service id and convert it to an int
	serviceIDStr := d.Id()
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call API Client's DeleteService method
	if _, err = client.DeleteService(serviceID); err != nil {
		return diag.FromErr(err)
	}

	// Remove resource id 
	d.SetId("")

	return diags
}

// Used by get request to set schema values
func setService(d *schema.ResourceData, data *lc.ServiceGetResponse) {
	d.Set("name", data.Instance.Name)
	d.Set("service_type_name", data.Instance.ServiceTypeName)
	d.Set("is_active", data.Instance.IsActive)
	d.Set("is_tax_exempt", data.Instance.IsTaxExempt)
	d.Set("is_inclusive_taxes", data.Instance.IsInclusiveTaxes)
	d.Set("default_service_status_type_name", data.Instance.DefaultServiceStatusTypeName)
	d.Set("description", data.Instance.Description)
	d.Set("service_category_name", data.Instance.ServiceCategoryName)
	d.Set("service_base_name", data.Instance.ServiceBaseName)
}
