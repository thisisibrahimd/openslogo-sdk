package v2alpha

import (
	"github.com/invopop/jsonschema"
	orderedmap "github.com/pb33f/ordered-map/v2"
)

// JSONSchema returns a custom JSON Schema for DurationShorthand.
func (DurationShorthand) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "string",
		Description: "A shorthand representation of time duration, e.g. '1m', '10d', '2w'.",
		Pattern:     `^[0-9]+[mhdwMQY]$`,
	}
}

// JSONSchema returns a custom JSON Schema for SLOAlertPolicy.
func (SLOAlertPolicy) JSONSchema() *jsonschema.Schema {
	inlineProps := orderedmap.New[string, *jsonschema.Schema]()
	inlineProps.Set("kind", &jsonschema.Schema{Type: "string"})
	inlineProps.Set("metadata", &jsonschema.Schema{Ref: "#/$defs/Metadata"})
	inlineProps.Set("spec", &jsonschema.Schema{Ref: "#/$defs/AlertPolicySpec"})

	refProps := orderedmap.New[string, *jsonschema.Schema]()
	refProps.Set("alertPolicyRef", &jsonschema.Schema{Type: "string"})

	return &jsonschema.Schema{
		Description: "An alert policy that can be provided inline or as a reference.",
		OneOf: []*jsonschema.Schema{
			{
				Type:                 "object",
				Properties:           inlineProps,
				Required:             []string{"kind", "metadata", "spec"},
				AdditionalProperties: jsonschema.FalseSchema,
			},
			{
				Type:                 "object",
				Properties:           refProps,
				Required:             []string{"alertPolicyRef"},
				AdditionalProperties: jsonschema.FalseSchema,
			},
		},
	}
}

// JSONSchema returns a custom JSON Schema for AlertPolicyCondition.
func (AlertPolicyCondition) JSONSchema() *jsonschema.Schema {
	inlineProps := orderedmap.New[string, *jsonschema.Schema]()
	inlineProps.Set("kind", &jsonschema.Schema{Type: "string"})
	inlineProps.Set("metadata", &jsonschema.Schema{Ref: "#/$defs/Metadata"})
	inlineProps.Set("spec", &jsonschema.Schema{Ref: "#/$defs/AlertConditionSpec"})

	refProps := orderedmap.New[string, *jsonschema.Schema]()
	refProps.Set("conditionRef", &jsonschema.Schema{Type: "string"})

	return &jsonschema.Schema{
		Description: "An alert condition that can be provided inline or as a reference.",
		OneOf: []*jsonschema.Schema{
			{
				Type:                 "object",
				Properties:           inlineProps,
				Required:             []string{"kind", "metadata", "spec"},
				AdditionalProperties: jsonschema.FalseSchema,
			},
			{
				Type:                 "object",
				Properties:           refProps,
				Required:             []string{"conditionRef"},
				AdditionalProperties: jsonschema.FalseSchema,
			},
		},
	}
}

// JSONSchema returns a custom JSON Schema for AlertPolicyNotificationTarget.
func (AlertPolicyNotificationTarget) JSONSchema() *jsonschema.Schema {
	inlineProps := orderedmap.New[string, *jsonschema.Schema]()
	inlineProps.Set("kind", &jsonschema.Schema{Type: "string"})
	inlineProps.Set("metadata", &jsonschema.Schema{Ref: "#/$defs/Metadata"})
	inlineProps.Set("spec", &jsonschema.Schema{Ref: "#/$defs/AlertNotificationTargetSpec"})

	refProps := orderedmap.New[string, *jsonschema.Schema]()
	refProps.Set("targetRef", &jsonschema.Schema{Type: "string"})

	return &jsonschema.Schema{
		Description: "A notification target that can be provided inline or as a reference.",
		OneOf: []*jsonschema.Schema{
			{
				Type:                 "object",
				Properties:           inlineProps,
				Required:             []string{"kind", "metadata", "spec"},
				AdditionalProperties: jsonschema.FalseSchema,
			},
			{
				Type:                 "object",
				Properties:           refProps,
				Required:             []string{"targetRef"},
				AdditionalProperties: jsonschema.FalseSchema,
			},
		},
	}
}
