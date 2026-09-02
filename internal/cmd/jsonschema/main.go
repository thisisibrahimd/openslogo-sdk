package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/OpenSLO/go-sdk/pkg/openslo/v1alpha"
	v2alpha "github.com/OpenSLO/go-sdk/pkg/openslo/v2alpha"
)

type objectGenerator struct {
	versionDir string
	generate   func() any
}

func main() {
	outDir := "docs/jsonschema"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	generators := []objectGenerator{
		// v1alpha
		{"v1alpha", func() any { return &v1alpha.Service{} }},
		{"v1alpha", func() any { return &v1alpha.SLO{} }},
		// v1
		{"v1", func() any { return &v1.Service{} }},
		{"v1", func() any { return &v1.SLO{} }},
		{"v1", func() any { return &v1.SLI{} }},
		{"v1", func() any { return &v1.DataSource{} }},
		{"v1", func() any { return &v1.AlertPolicy{} }},
		{"v1", func() any { return &v1.AlertCondition{} }},
		{"v1", func() any { return &v1.AlertNotificationTarget{} }},
		// v2alpha
		{"v2alpha", func() any { return &v2alpha.Service{} }},
		{"v2alpha", func() any { return &v2alpha.SLO{} }},
		{"v2alpha", func() any { return &v2alpha.SLI{} }},
		{"v2alpha", func() any { return &v2alpha.DataSource{} }},
		{"v2alpha", func() any { return &v2alpha.AlertPolicy{} }},
		{"v2alpha", func() any { return &v2alpha.AlertCondition{} }},
		{"v2alpha", func() any { return &v2alpha.AlertNotificationTarget{} }},
	}

	r := &jsonschema.Reflector{
		ExpandedStruct: true,
		AssignAnchor:   true,
	}
	if err := r.AddGoComments("github.com/OpenSLO/go-sdk", "./pkg"); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to add Go comments: %v\n", err)
	}

	for _, gen := range generators {
		obj := gen.generate()
		schema := r.Reflect(obj)

		// Set $id for the schema
		kind := getKind(obj)
		schema.ID = jsonschema.ID(fmt.Sprintf("https://openslo.com/schemas/%s/%s.json", gen.versionDir, kind))

		data, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to marshal schema for %s/%s: %v\n", gen.versionDir, kind, err)
			os.Exit(1)
		}

		dir := filepath.Join(outDir, gen.versionDir)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "failed to create directory %s: %v\n", dir, err)
			os.Exit(1)
		}

		path := filepath.Join(dir, kind+".json")
		if err := os.WriteFile(path, data, 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write %s: %v\n", path, err)
			os.Exit(1)
		}

		fmt.Printf("Generated %s\n", path)
	}
}

func getKind(obj any) string {
	switch obj.(type) {
	case *v1alpha.Service:
		return "service"
	case *v1alpha.SLO:
		return "slo"
	case *v1.Service:
		return "service"
	case *v1.SLO:
		return "slo"
	case *v1.SLI:
		return "sli"
	case *v1.DataSource:
		return "data-source"
	case *v1.AlertPolicy:
		return "alert-policy"
	case *v1.AlertCondition:
		return "alert-condition"
	case *v1.AlertNotificationTarget:
		return "alert-notification-target"
	case *v2alpha.Service:
		return "service"
	case *v2alpha.SLO:
		return "slo"
	case *v2alpha.SLI:
		return "sli"
	case *v2alpha.DataSource:
		return "data-source"
	case *v2alpha.AlertPolicy:
		return "alert-policy"
	case *v2alpha.AlertCondition:
		return "alert-condition"
	case *v2alpha.AlertNotificationTarget:
		return "alert-notification-target"
	default:
		return "unknown"
	}
}
