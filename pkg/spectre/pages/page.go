package components

import (
	"encoding/json"
	"fmt"
	"github.com/amlwwalker/spectre-webasm/pkg/markdown"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io/ioutil"
	"net/http"
)

type Page struct {
	app.Compo
	Ctx         app.Context
	Title       string
	Description string
	Src         string
	Body        []app.UI
	Sidebar     []app.UI
	Properties  map[string]string
	Switches    map[string]bool
	States      map[string]interface{}
}

func (p *Page) OnMount(ctx app.Context) {
	p.Ctx = ctx
	for k, v := range p.States {
		ctx.ObserveState(k).Value(&v)
	}
	//this possibly won't work now
	defer func() {
		fmt.Println("calling prism highlight all")
		app.Window().Get("Prism").Call("highlightAll")
	}()
}
func (p *Page) Render() app.UI {
	res, err := http.Get("/links")
	if err != nil {
		return app.Div()
	}
	var links []link
	if byts, err := ioutil.ReadAll(res.Body); err != nil {
		links = []link{}
	} else {
		json.Unmarshal(byts, &links)
	}
	var authorized bool
	for _, v := range links {
		if v.URL == p.Src {
			authorized = true
			break
		}
	}
	//fixme - for properly authorized pages, all the content should come from the server (use frontmatter for titles and descriptions)
	return app.Div().Class("container").Body(
		layouts.NavBar(),
		layouts.SideBar(
			components.Accordion(
				createSideBar(links),
			), app.Div().Class("docs-content").Class("content").Body(
				app.If(authorized,
					layouts.Hero(p.Title, p.Description),
					app.Span().Body(
						p.Body...,
					),
					markdown.NewRemoteMarkdownDoc().Src(p.Src)).Else(
					layouts.Hero("Woops!", "The requested page is not available"),
				),
			),
		),
	)
}

func NewPage(title, description, src string, sidebar []app.UI, body ...app.UI) Page {
	return Page{
		Title:       title,
		Description: description,
		Src:         src,
		Body:        body,
		Sidebar:     sidebar,
		Properties:  make(map[string]string),
		Switches:    make(map[string]bool),
		States:      make(map[string]interface{}),
	}
}

// NewPage
func NewPageComposer(p *Page) func() app.Composer {
	return func() app.Composer {
		return p
	}
}

func (p *Page) SetBody(body ...app.UI) {
	p.Body = body
}
