package actions

import (
	"fmt"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"strings"
)

func ExtractModelToBase(model *v3.Document, baseNamePrefix string) map[string]string {
	extractedModels := make(map[string]string)
	for sn, sp := range model.Components.Schemas {
		if len(sp.Schema().OneOf) > 0 && len(sp.Schema().Properties) > 0 {
			bi := createBaseItem(sp)
			baseSchemaName := fmt.Sprintf("%s%s", baseNamePrefix, sn)
			model.Components.Schemas[baseSchemaName] = bi
			extractedModels[sn] = baseSchemaName
			oldSchema := sp.Schema()
			oneOfSchema := extractOneOf(oldSchema)
			oneOfSchemaProxy := base.CreateSchemaProxy(oneOfSchema)
			refToBase := base.CreateSchemaProxyRef(fmt.Sprintf("#/components/schemas/%s", baseSchemaName))
			ns := &base.Schema{
				AllOf: []*base.SchemaProxy{
					oneOfSchemaProxy,
					refToBase,
				},
			}
			newProxy := base.CreateSchemaProxy(ns)
			model.Components.Schemas[sn] = newProxy
		}
	}
	return extractedModels
}

func createBaseItem(mbase *base.SchemaProxy) *base.SchemaProxy {
	item := base.CreateSchemaProxy(&base.Schema{
		Type:       []string{"object"},
		Properties: mbase.Schema().Properties,
	})
	return item
}

func extractOneOf(schema *base.Schema) *base.Schema {
	oneOfSchema := &base.Schema{
		OneOf:         schema.OneOf,
		Discriminator: schema.Discriminator,
	}
	return oneOfSchema
}

func ReplaceAllOfReference(model *v3.Document, replaceModels map[string]string) {
	for _, sp := range model.Components.Schemas {
		sch := sp.Schema()
		for i, ao := range sch.AllOf {
			if ao.IsReference() {
				ref := ao.GetReference()
				lastPart := strings.Split(ref, "/")
				if to, ok := replaceModels[lastPart[len(lastPart)-1]]; ok {
					newAo := base.CreateSchemaProxyRef(fmt.Sprintf("#/components/schemas/%s", to))
					sch.AllOf[i] = newAo
				}
			}
		}
	}
}
