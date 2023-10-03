package main

import (
	"fmt"
	"github.com/amlwwalker/spectre-webasm/cmd/templator/pkg/componentGenerator"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
	"log"
	"strings"
)

const template = `<div class="container">
		<navbar data-brand="Spectre.css" data-doc-link="Docs" data-git-link="Github" data-img-src="https://picturepan2.github.io/spectre/img/spectre-logo.svg"></navbar>
		<sidebar data-link1-href="#link1" data-link1-text="Link 1 Text" data-link2-href="#link2" data-link2-text="Link 2 Text">
			<accordionmenu></accordionmenu>
			<hero data-title="hello world" data-description="the description"></hero>
			<p>content</p>
		</sidebar>
	</div>`

// ... [Other Element Structs like Navbar, Sidebar, etc. here]

func getAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func processNode(n *html.Node) componentGenerator.Element {
	if n.Type == html.DocumentNode {
		// For the root document node, just process its children
		fmt.Println("processing document node", n.Data)
		return componentGenerator.HTMLElement{
			TagName:  n.Data,
			Attrs:    n.Attr,
			Children: processChildren(n),
		}
	} else if n.Type == html.ElementNode {
		base := componentGenerator.BaseElement{
			Class: getAttribute(n, "class"),
			ID:    getAttribute(n, "id"),
		}
		fmt.Println("processing ", n.Data)
		switch n.Data {
		case "div":
			fmt.Println("found div")
			return componentGenerator.Div{BaseElement: base, Children: processChildren(n)}
		case "navbar":
			return componentGenerator.Navbar{
				BaseElement: base,
				Brand:       getAttribute(n, "data-brand"),
				DocLink:     getAttribute(n, "data-doc-link"),
				GitLink:     getAttribute(n, "data-git-link"),
				ImageSrc:    getAttribute(n, "data-img-src"),
			}
		case "sidebar":
			// Extract links from data attributes (you can extend this logic for more links)
			link1 := componentGenerator.Link{Href: getAttribute(n, "data-link1-href"), Text: getAttribute(n, "data-link1-text")}
			link2 := componentGenerator.Link{Href: getAttribute(n, "data-link2-href"), Text: getAttribute(n, "data-link2-text")}
			// ... add more links as necessary

			return componentGenerator.Sidebar{
				BaseElement: componentGenerator.BaseElement{
					Class: getAttribute(n, "class"),
					ID:    getAttribute(n, "id"),
				},
				Links: []componentGenerator.Link{link1, link2}, // ... add more links as necessary
				Body:  processChildren(n),
			}
		case "accordionmenu":
			return componentGenerator.Accordion{
				Menus: []componentGenerator.AccordionMenu{
					{
						PathID:       "getting-started",
						MenuTitle:    "Getting Started",
						MenuElements: []string{"Installation", "Custom Version", "Browser-Support"},
					},
					//... other predefined menus
				},
			}
		case "hero":
			title := getAttribute(n, "data-title")
			desc := getAttribute(n, "data-description")
			return componentGenerator.Hero{
				BaseElement: componentGenerator.BaseElement{
					Class: getAttribute(n, "class"),
					ID:    getAttribute(n, "id"),
				},
				Title:       title,
				Description: desc,
			}
			// Add cases for other tags (AccordionMenu, Hero, etc.)
			// Add cases for other specific tags (AccordionMenu, Hero, etc.)
		default:
			fmt.Println("default ", n.Data)
			// Fallback to the generic HTMLElement
			return componentGenerator.HTMLElement{
				TagName:  n.Data,
				Attrs:    n.Attr,
				Children: processChildren(n),
			}
		}

	} else if n.Type == html.TextNode {
		return componentGenerator.Text{Content: n.Data}
	}
	return nil
}

func processChildren(n *html.Node) []componentGenerator.Element {
	var children []componentGenerator.Element
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childElement := processNode(c)
		if childElement != nil {
			children = append(children, childElement)
		}
	}
	return children
}

// findBody will traverse the HTML node tree to find the body node
func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if body := findBody(c); body != nil {
			return body
		}
	}
	return nil
}

func main() {

	var input = `
<div class="container">
	<navbar data-brand="Spectre.css" data-doc-link="Docs" data-git-link="Github" data-img-src="https://picturepan2.github.io/spectre/img/spectre-logo.svg"></navbar>
	<sidebar data-link1-href="#link1" data-link1-text="Link 1 Text" data-link2-href="#link2" data-link2-text="Link 2 Text">
	<accordionmenu></accordionmenu>
	<hero data-title="hello world" data-description="the description"></hero>
	<p>content</p>
	</sidebar>
	</div>`
	htmlContent := input

	r := strings.NewReader(htmlContent)
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	bodyNode := findBody(doc)
	if bodyNode == nil {
		log.Fatal("Couldn't find the body node")
	}
	root := processNode(bodyNode)

	//fmt.Println("root node is ", root.String())
	//fmt.Println("generated node is ", root.Generate())
	//"github.com/yosssi/gohtml"
	//"github.com/tdewolff/minify"
	//fmt.Println(HtmlMinify(root.String()))
	if false {
		generated := root.Generate()
		fmt.Println("printing output generate")
		fmt.Println(gohtml.Format(app.HTMLString(generated)))
	} else {
		generated := root.String()
		fmt.Println("printing output string")
		fmt.Println(gohtml.Format(generated))
	}
}

func HtmlMinify(h string) string {
	result, _ := Minify([]byte(h), nil)
	return string(result)
}
