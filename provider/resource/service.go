package resource

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lc "github.com/ibrahim925/LogiCore"
)

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

func resourceServiceCreate(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*lc.Client)

	var diags diag.Diagnostics

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

	serviceIDStr := strconv.Itoa(postResponse.Results.Items[0].Identity)

	d.SetId(serviceIDStr)

	timeCreated := postResponse.Results.Items[0].Instance.Created
	d.Set("created", timeCreated)

	return diags
}

func resourceServiceRead(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*lc.Client)

	var diags diag.Diagnostics

	serviceIDStr := d.Id()
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		return diag.FromErr(err)
	}

	service, err := client.GetService(serviceID)
	if err != nil {
		return diag.FromErr(err)
	}

	setService(d, service)

	return diags

}

func resourceServiceUpdate(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*lc.Client)

	var diags diag.Diagnostics

	serviceIDStr := d.Id()
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("name", "service_type_name", "created", "is_active", "is_tax_exempt", "is_inclusive_taxes", "default_service_status_type_name", "description") {
		updatedService := lc.ServicePatchStruct{}

		updatedService.Name = d.Get("name").(string)
		updatedService.ServiceTypeName = d.Get("service_type_name").(string)
		updatedService.IsActive = d.Get("is_active").(bool)
		updatedService.IsTaxExempt = d.Get("is_tax_exempt").(bool)
		updatedService.IsInclusiveTaxes = d.Get("is_inclusive_taxes").(bool)
		updatedService.DefaultServiceStatusTypeName = d.Get("default_service_status_type_name").(string)
		updatedService.Description = d.Get("description").(string)

		if _, err := client.UpdateService(serviceID, updatedService); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceServiceDelete(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*lc.Client)

	var diags diag.Diagnostics

	serviceIDStr := d.Id()
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err = client.DeleteService(serviceID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

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
