package componentGenerator

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"golang.org/x/net/html"
	"strings"
)

type Text struct {
	Content string
}

func (t Text) String() string {
	return t.Content
}
func (t Text) Generate() app.UI {
	if strings.TrimSpace(t.Content) == "" {
		return nil
	}
	res := app.Text(t.String())
	fmt.Println("text object returning ", res, " for ", t.String())
	return res
}

// An adapter to treat raw HTML string as an Element
type HTMLStringElement string

func (h HTMLStringElement) String() string {
	return string(h)
}
func (h HTMLStringElement) Generate() app.UI {
	fmt.Println("htmlStringElement ", string(h))
	return app.Raw(string(h))
}

// A mock function for urlLink, replace with the actual logic if it's different
func urlLink(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}

type Link struct {
	Href string
	Text string
}

type HTMLElement struct {
	TagName  string
	Attrs    []html.Attribute
	Children []Element
}

func (e HTMLElement) String() string {
	attributes := ""
	for _, attr := range e.Attrs {
		attributes += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}

	content := "<" + e.TagName + attributes + ">"
	for _, child := range e.Children {
		content += child.String()
	}
	content += "</" + e.TagName + ">"
	return content
}
func (e HTMLElement) Generate() app.UI {
	// Create the raw start tag with attributes
	attributes := ""
	for _, attr := range e.Attrs {
		attributes += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}
	startTag := app.Raw("<" + e.TagName + attributes + ">")

	// Convert child elements
	var childComponents []app.UI
	for _, child := range e.Children {
		if comp, ok := child.(Element); ok {
			if e.TagName == "p" {
				fmt.Println("paragraph found")
			} else if e.TagName == "section" {
				fmt.Println("section found")
			}
			childComponents = append(childComponents, comp.Generate())
		} else {
			childComponents = append(childComponents, app.Raw(child.String()))
		}
	}
	if len(childComponents) == 0 {
		return nil // Or whatever represents an empty component in your context
	}
	fmt.Println("childs ", app.HTMLString(app.Div().Body(childComponents...)))
	switch e.TagName {
	case "body":
		//this should be a body tag
		return app.Div().Class("normalize").Body(childComponents...)
	case "div":
		fmt.Println("---- div")
		return app.Div().Body(childComponents...)
	case "p":
		fmt.Println("------- p")
		return app.P().Body(childComponents...)
	case "section":
		fmt.Println("------ section")
		return app.Section().Body(childComponents...)
	}

	// Create the raw end tag
	endTag := app.Raw("</" + e.TagName + ">")

	// Combine the start tag, child components, and end tag into a single div to keep them together
	combinedComponents := append([]app.UI{startTag}, childComponents...)
	combinedComponents = append(combinedComponents, endTag)
	fmt.Println("html element returning ", e.TagName)
	return app.Div().Body(combinedComponents...)
}

type Div struct {
	BaseElement
	Children []Element
}

func (d Div) String() string {
	attributes := ""
	if d.Class != "" {
		attributes += fmt.Sprintf(" class=\"%s\"", d.Class)
	}
	if d.ID != "" {
		attributes += fmt.Sprintf(" id=\"%s\"", d.ID)
	}

	content := "<div" + attributes + ">"
	for _, child := range d.Children {
		content += child.String()
	}
	content += "</div>"

	return content
}
func (d Div) Generate() app.UI {
	// Start with an empty div
	divComponent := app.Div()

	// Add class if present
	if d.Class != "" {
		divComponent = divComponent.Class(d.Class)
	}
	//fmt.Println("added classes ", app.HTMLString(divComponent))
	// Add ID if present
	if d.ID != "" {
		divComponent = divComponent.ID(d.ID)
	}

	// Convert child elements
	var childComponents []app.UI
	for _, child := range d.Children {
		if comp, ok := child.(Element); ok {
			s := comp.String()
			runes := []rune(s)
			if len(runes) > 10 {
				s = string(runes[:10])
			}
			fmt.Println("processing ", s)
			childComponents = append(childComponents, comp.Generate())
		} else {
			fmt.Println("app.Raw called - ", child.String())
			childComponents = append(childComponents, app.Raw(child.String()))
		}
	}

	// Add child components to div
	divComponent = divComponent.Body(childComponents...)
	//fmt.Println("generated div ", app.HTMLString(divComponent))
	return divComponent
}
