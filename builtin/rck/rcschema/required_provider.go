package rcschema

import (
	"fmt"

	"github.com/4rchr4y/goray/kernel/hcllang"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// Reserved for future expansion
var (
	requiredProviderBlockReservedAttributeList = [...]string{
		"aliases",
	}
	requiredProviderBlockReservedBlockList = [...]string{}
)

var requiredProviderBlockSchema = &hcl.BodySchema{
	Attributes: NewAttributeList(
		hcl.AttributeSchema{
			Name:     "source",
			Required: true,
		},
		hcl.AttributeSchema{
			Name:     "path",
			Required: false,
		},
		hcl.AttributeSchema{
			Name:     "version",
			Required: false,
		},
	)(requiredProviderBlockReservedAttributeList[:]...),
	Blocks: NewBlockList(
		hcl.BlockHeaderSchema{
			Type:       "_",
			LabelNames: []string{},
		},
	)(requiredProviderBlockReservedBlockList[:]...),
}

type RequiredProvider struct {
	_       [0]int
	Name    string
	Source  string
	Path    string
	Version string
	Body    hcl.Body
	Content *hcl.BodyContent
}

func DecodeRequiredProviderBlock(block *hcl.Block) (requiredProvider *RequiredProvider, diagnostics hcl.Diagnostics) {
	content, body, partialContentDiag := block.Body.PartialContent(requiredProviderBlockSchema)
	diagnostics = append(diagnostics, partialContentDiag...)
	if diagnostics.HasErrors() {
		return nil, diagnostics
	}

	// existence of a label is checked when a block is detected
	if !hclsyntax.ValidIdentifier(block.Labels[0]) {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Invalid provider name",
			Detail:   fmt.Sprintf("Provider name is invalid. %s", hcllang.BadIdentDetail),
		})
	}

	requiredProvider = &RequiredProvider{
		Name:    block.Labels[0],
		Content: content,
		Body:    body,
	}

	if attr, exists := content.Attributes["source"]; exists {
		diags := gohcl.DecodeExpression(attr.Expr, nil, &requiredProvider.Source)
		diagnostics = append(diagnostics, diags...)
		if diagnostics.HasErrors() {
			return nil, diagnostics
		}

		// TODO: source validation
	}

	if attr, exists := content.Attributes["path"]; exists {
		diags := gohcl.DecodeExpression(attr.Expr, nil, &requiredProvider.Path)
		diagnostics = append(diagnostics, diags...)
		if diagnostics.HasErrors() {
			return nil, diagnostics
		}

		// TODO: source validation
	}

	if attr, exists := content.Attributes["version"]; exists {
		diags := gohcl.DecodeExpression(attr.Expr, nil, &requiredProvider.Version)
		diagnostics = append(diagnostics, diags...)
		if diagnostics.HasErrors() {
			return nil, diagnostics
		}

		// TODO: source validation
	}

	return requiredProvider, diagnostics
}

func NewAttributeList(attributes ...hcl.AttributeSchema) func(reserved ...string) []hcl.AttributeSchema {
	return func(reserved ...string) []hcl.AttributeSchema {
		for i := range reserved {
			attributes = append(attributes, hcl.AttributeSchema{
				Name:     reserved[i],
				Required: false,
			})
		}

		return attributes
	}
}

func NewBlockList(blocks ...hcl.BlockHeaderSchema) func(reserved ...string) []hcl.BlockHeaderSchema {
	return func(reserved ...string) []hcl.BlockHeaderSchema {
		for i := range reserved {
			blocks = append(blocks, hcl.BlockHeaderSchema{
				Type: reserved[i],
			})
		}

		return blocks
	}
}