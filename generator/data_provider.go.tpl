// !!! This file has been automatically generated !!!

package routeros

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func {{ .NameGolang }}Exists(d *schema.ResourceData, m interface{}) (bool, error) {
	c, err := m.(Config).Client()
	if err != nil {
		return false, err
	}
	id := d.Id()
	reply, err := c.Run("{{ .APIPath }}/print", "?.id="+id)
	if err != nil {
		return false, err
	}
	if len(reply.Re) <= 0 {
		return false, fmt.Errorf(`no {{ .NameGolang }} with ID "%s" found`, id)
	}
	if len(reply.Re) > 1 {
		return false, fmt.Errorf(`more than one (%d) {{ .NameGolang }} with ID "%s" found`, len(reply.Re), id)
	}

	return true, nil
}

func {{ .NameGolang }}Read(d *schema.ResourceData, m interface{}) error {
	c, err := m.(Config).Client()
	if err != nil {
		return err
	}
	apiIdentifier := d.Get("{{ .APIIdentifierField }}").(string)
	reply, err := c.Run("{{ .APIPath }}/print", "?{{ .APIIdentifierField }}="+apiIdentifier)
	if err != nil {
		return err
	}

	if len(reply.Re) <= 0 {
		return fmt.Errorf(`no {{ .NameGolang }} with {{ .APIIdentifierField }} "%s" found`, apiIdentifier)
	}
	if len(reply.Re) > 1 {
		return fmt.Errorf(`more than one {{ .NameGolang }} (%d) with {{ .APIIdentifierField }} "%s" found`, len(reply.Re), apiIdentifier)
	}

	re := reply.Re[0]
	d.SetId(re.Map[".id"])
	{{- range $schema := .Schema }}
	{{- if eq $schema.ValueType "bool" }}
	if err := d.Set("{{ $schema.Name }}", mustParseBool(re.Map["{{ Replace $schema.Name "_" "-" -1}}"])); err != nil {
	    return err
	}
	{{- else if eq $schema.ValueType "int" }}
	if err := d.Set("{{ $schema.Name }}", mustAtoi(re.Map["{{ Replace $schema.Name "_" "-" -1}}"])); err != nil {
	    return err
	}
	{{- else if eq $schema.ValueType "map" }}
	{{- if eq $schema.Elem "bool" }}
	if err := d.Set("{{ $schema.Name }}", stringToBoolMap(re.Map["{{ ReplaceAll $schema.Name "_" "-"}}"])); err != nil {
		return err
	}
	{{- else }}
	panic("Unsupported elem type: {{ $schema.Elem }}")
	{{- end }}
	{{- else }}
	if err := d.Set("{{ $schema.Name }}", re.Map["{{ Replace $schema.Name "_" "-" -1}}"]); err != nil {
	    return err
	}
	{{- end }}
	{{- end }}

	return nil
}

func Data{{ CapitalizeFirstLetter .NameGolang }}() *schema.Resource {
	return &schema.Resource{
	    Exists: {{ .NameGolang }}Exists,
		Read: {{ .NameGolang }}Read,
		Schema: map[string]*schema.Schema{
			{{- range $schema := .Schema }}
			"{{ $schema.Name }}": {
				Description: "{{ ReplaceAll $schema.Description "\"" "\\\"" }}",
				Type: schema.{{ ToValueType $schema.ValueType }},{{- if $schema.Required }}
				Required: {{ $schema.Required }},{{- end }}{{- if $schema.Optional }}
				Optional: {{ $schema.Optional }},{{- end }}{{- if $schema.Computed }}
				Computed: {{ $schema.Computed }},{{- end }}{{- if $schema.Elem }}
				Elem: schema.{{ ToElemType $schema.Elem }},{{- end }}{{- if $schema.ForceNew }}
				ForceNew: {{ $schema.ForceNew }},{{- end }}
			},
			{{- end }}
		},
	}
}
