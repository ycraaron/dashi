package schema

import (
	"strings"
	"text/template"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/generators/scaffold"
)

func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()
	helpers := template.FuncMap{
		"getFieldName": func(field string) string {
			return strings.Split(field, ":")[0]
		},
		"getFieldType": func(field string) string {
			return strings.Split(field, ":")[1]
		},
		"pascalize":  flect.Pascalize,
		"camelize":   flect.Camelize,
		"underscore": flect.Underscore,
	}
	//Change to camel
	data := map[string]interface{}{
		"opts": opts,
	}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)
	g.Transformer(genny.Replace("-entity-", opts.Name.Underscore().String()))
	g.Transformer(genny.Replace("-path-", "graphql/schema"))
	if err := g.Box(packr.New("scaffold:schema:templates", "../schema/templates")); err != nil {
		return g, err
	}

	return g, nil
}
