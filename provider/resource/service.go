package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//	lc "github.com/ibrahim925/LogiCore"
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
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServiceCreate(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceServiceRead(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceServiceUpdate(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceServiceDelete(context context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
