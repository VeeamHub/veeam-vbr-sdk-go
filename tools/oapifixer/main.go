package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/veeamhub/veeam-vbr-sdk-go/v2/tools/oapifixer/actions"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
)

func main() {

	inputFilePtr := flag.String("input", "", "openapi specification to process")
	outputFilePtr := flag.String("output", "", "output file name")
	flag.Parse()

	if inputFilePtr == nil || *inputFilePtr == "" {
		log.Fatal("input file name is required")
	}

	if outputFilePtr == nil || *outputFilePtr == "" {
		log.Fatal("output file name is required")
	}

	spec, err := os.ReadFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}
	doc, err := libopenapi.NewDocument(spec)
	if err != nil {
		panic(err)
	}
	docModel, errors := doc.BuildV3Model()
	handleErrors(errors)

	extractedModels := actions.ExtractModelToBase(&docModel.Model, "Base")
	actions.ReplaceAllOfReference(&docModel.Model, extractedModels)

	file, doc, newModel, errors := doc.RenderAndReload()
	handleErrors(errors)

	outputFormat := filepath.Ext(*outputFilePtr)
	if outputFormat == ".yaml" && doc.GetSpecInfo().SpecFileType == datamodel.JSONFileType {
		file = newModel.Model.RenderWithIndention(doc.GetSpecInfo().OriginalIndentation)
	} else if outputFormat == ".json" && doc.GetSpecInfo().SpecFileType == datamodel.YAMLFileType {
		jsonIndent := "  "
		i := doc.GetSpecInfo().OriginalIndentation
		if i > 2 {
			for l := 0; l < i-2; l++ {
				jsonIndent += " "
			}
		}
		file = newModel.Model.RenderJSON(jsonIndent)
	}

	err = os.WriteFile(*outputFilePtr, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func handleErrors(errors []error) {
	if len(errors) > 0 {
		for i := range errors {
			log.Printf("error: %e\n", errors[i])
		}
		log.Printf("cannot create v3 model from document: %d errors reported", len(errors))
	}
}
