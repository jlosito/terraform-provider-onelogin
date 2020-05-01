package parameters

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
)

// AppParameter returns a key/value map of the various fields that make up
// the AppParameter field for a OneLogin App.
func ParameterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"param_key_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"param_id": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"label": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"user_attribute_mappings": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"user_attribute_macros": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"attributes_transformations": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"default_values": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"skip_if_blank": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"values": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"provisioned_entitlements": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"safe_entitlements_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
	}
}

// InflateAppParameter takes a key/value map of interfaces and uses the fields to construct
// an AppParameter struct, a sub-field of a OneLogin App.
func Inflate(s map[string]interface{}) models.AppParameters {
	out := models.AppParameters{}
	var b, notNil bool = false, false
	var d int
	var st string

	if st, notNil = s["label"].(string); notNil {
		out.Label = oltypes.String(st)
	}

	if st, notNil = s["user_attribute_mappings"].(string); notNil {
		out.UserAttributeMappings = oltypes.String(st)
	}

	if st, notNil = s["user_attribute_macros"].(string); notNil {
		out.UserAttributeMacros = oltypes.String(st)
	}

	if st, notNil = s["attributes_transformations"].(string); notNil {
		out.AttributesTransformations = oltypes.String(st)
	}

	if st, notNil = s["values"].(string); notNil {
		out.Values = oltypes.String(st)
	}

	if st, notNil = s["default_values"].(string); notNil {
		out.DefaultValues = oltypes.String(st)
	}

	if b, notNil = s["skip_if_blank"].(bool); notNil {
		out.SkipIfBlank = oltypes.Bool(b)
	}

	if b, notNil = s["provisioned_entitlements"].(bool); notNil {
		out.ProvisionedEntitlements = oltypes.Bool(b)
	}

	if b, notNil = s["safe_entitlements_enabled"].(bool); notNil {
		out.SafeEntitlementsEnabled = oltypes.Bool(b)
	}

	if d, notNil = s["param_id"].(int); notNil {
		out.ID = oltypes.Int32(int32(d))
	}

	return out
}

func Flatten(params map[string]models.AppParameters) []map[string]interface{} {
	out := make([]map[string]interface{}, 0)
	for k, v := range params {
		param := map[string]interface{}{
			"param_key_name":             k,
			"param_id":                   v.ID,
			"label":                      v.Label,
			"user_attribute_mappings":    v.UserAttributeMappings,
			"user_attribute_macros":      v.UserAttributeMacros,
			"attributes_transformations": v.AttributesTransformations,
			"skip_if_blank":              v.SkipIfBlank,
			"values":                     v.Values,
			"default_values":             v.DefaultValues,
			"provisioned_entitlements":   v.ProvisionedEntitlements,
			"safe_entitlements_enabled":  v.SafeEntitlementsEnabled,
		}
		out = append(out, param)
	}
	return out
}