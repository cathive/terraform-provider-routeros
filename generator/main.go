package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/yaml.v2"
)

type DataDefinition struct {
	NameTerraform      string             `yaml:"name_terraform"`
	NameGolang         string             `yaml:"name_golang"`
	APIIdentifierField string             `yaml:"api_identifier_field"`
	APIPath            string             `yaml:"api_path"`
	Schema             []SchemaDefinition `yaml:"schema"`
}

type SchemaDefinition struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	ForceNew    bool   `yaml:"force_new"`
	ValueType   string `yaml:"value_type"`
	Required    bool   `yaml:"required,omitempty"`
	Optional    bool   `yaml:"optional,omitempty"`
	Computed    bool   `yaml:"computed,omitempty"`
	Elem        string `yaml:"elem,omitempty"`
}

func CapitalizeFirstLetter(str string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(str[:1]), str[1:])
}

var valueTypes = map[string]schema.ValueType{
	"string": schema.TypeString,
	"int":    schema.TypeInt,
	"bool":   schema.TypeBool,
	"map":    schema.TypeMap,
}

func ToValueType(str string) schema.ValueType {
	return valueTypes[str]
}

func ToElemType(str string) schema.ValueType {
	return valueTypes[str]
}

type MetaSchema struct {
	*schema.Schema
	FromTFFunc *func(interface{}) interface{}
}

func main() {
	if _, err := fmt.Fprintf(os.Stdout, "Generating sources...\n"); err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/data_provider.go.tpl", os.Args[1]))
	if err != nil {
		panic(err)
	}

	tpl, err := template.New("data").Funcs(template.FuncMap{
		"CapitalizeFirstLetter": CapitalizeFirstLetter,
		"ToElemType":            ToElemType,
		"ToValueType":           ToValueType,
		"Replace":               strings.Replace,
		"ReplaceAll":            strings.ReplaceAll,
	}).Parse(string(data))
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(os.Args[2])
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range files {
		name := fileInfo.Name()
		if !strings.HasSuffix(name, ".yaml") {
			continue // If it ain't has a .yaml ending, we are not interested.
		}
		filename := fmt.Sprintf("%s/%s", os.Args[2], name)
		file, err := os.Open(filename)
		if err != nil {
			panic(fmt.Errorf("unable to open file: %s: %v", filename, err))
		}
		defer file.Close()
		var dd DataDefinition
		if _, err = fmt.Fprintf(os.Stdout, "* %s ...\n", filename); err != nil {
			panic(err)
		}
		if err := yaml.NewDecoder(file).Decode(&dd); err != nil {
			panic(err)
		}
		f, err := os.Create(fmt.Sprintf("%s/zzz_%s.go", os.Args[3], dd.NameTerraform))
		if err != nil {
			panic(err)
		}
		if err := tpl.Execute(f, dd); err != nil {
			panic(err)
		}
	}

}
