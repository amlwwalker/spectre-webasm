package custom_renderers

import (
	"encoding/base64"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"io"
)

//todo - consider goldmark for commonmark support - https://github.com/yuin/goldmark

type CustomRenderer struct {
	// other renderer components...
	baseRenderer markdown.Renderer
	Component    string
}

func NewCustomRenderer(component string) *CustomRenderer {
	opts := html.RendererOptions{
		Flags: html.CommonFlags,
	}
	return &CustomRenderer{
		Component:    component,
		baseRenderer: html.NewRenderer(opts),
	}
}

func (r *CustomRenderer) RenderHeader(w io.Writer, ast ast.Node) {
	// For the sake of simplicity, we aren't adding any custom behavior here.
	// However, you could write some custom header content to 'w' if needed.
}

func (r *CustomRenderer) RenderFooter(w io.Writer, ast ast.Node) {
	// Similar to RenderHeader, for the sake of simplicity, we aren't adding
	// any custom behavior here. However, you could write some custom footer
	// content to 'w' if needed.
}

// Define custom rendering for code blocks:
func (r *CustomRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	if code, ok := node.(*ast.CodeBlock); ok {
		// Check if it's a JSON block:
		if string(code.Info) == "json" {
			// Convert to a placeholder for React:
			placeholder := createReactPlaceholder(string(code.Literal), r.Component)
			w.Write([]byte(placeholder))
			return ast.GoToNext
		} else {
			//for code languages we should apply a prettifier component
		}
	}
	return r.baseRenderer.RenderNode(w, node, entering)
}

func createReactPlaceholder(jsonContent, component string) string {
	// Convert the JSON content to a base64 encoded string
	encodedJSON := base64.StdEncoding.EncodeToString([]byte(jsonContent))

	fmt.Println("encodedJSON", encodedJSON, "component", component)
	// Return a placeholder that will be picked up by your JSX mounter

	return fmt.Sprintf(`<div data-react-component="%s" data-react-props="%s">HELLO WORLD</div>`, component, encodedJSON)
}
