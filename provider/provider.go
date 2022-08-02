package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lc "github.com/ibrahim925/LogiCore"
	"terraform-logiprovider/provider/datasource"
	"terraform-logiprovider/provider/resource"
)

//Define the provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOGISENSE_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOGISENSE_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOGISENSE_PASSWORD", nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOGISENSE_CLIENT_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"logiprovider_service": resource.Service(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"logiprovider_service":  datasource.Service(),
			"logiprovider_services": datasource.Services(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	client_id := d.Get("client_id").(string)

	var diags diag.Diagnostics

	client, err := lc.NewClient(host, username, password, client_id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create LogiSense API client",
			Detail:   "Unable to authenticate user for LogiSense API client",
		})
		return nil, diags
	}

	return client, diags
}
