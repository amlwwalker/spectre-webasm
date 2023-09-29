package markdown

import (
	"fmt"
	custom_renderers "github.com/amlwwalker/spectre-webasm/pkg/markdown/custom-renderers"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"path/filepath"
)

type markdownDoc struct {
	app.Compo

	Iid    string
	Iclass string
	Imd    string
}

func newMarkdownDoc() *markdownDoc {
	return &markdownDoc{}
}

func (d *markdownDoc) ID(v string) *markdownDoc {
	d.Iid = v
	return d
}

func (d *markdownDoc) Class(v string) *markdownDoc {
	d.Iclass = app.AppendClass(d.Iclass, v)
	return d
}

func (d *markdownDoc) MD(v string) *markdownDoc {
	d.Imd = fmt.Sprintf(`<div class="markdown">%s</div>`, parseMarkdown([]byte(v)))
	return d
}

func (d *markdownDoc) OnMount(ctx app.Context) {
	ctx.Defer(d.highlightCode)
}

func (d *markdownDoc) OnUpdate(ctx app.Context) {
	ctx.Defer(d.highlightCode)
}

func (d *markdownDoc) Render() app.UI {
	return app.Div().
		ID(d.Iid).
		Class(d.Iclass).
		Body(
			app.Raw(d.Imd),
		)
}

func (d *markdownDoc) highlightCode(ctx app.Context) {
	//app.Window().Get("Prism").Call("highlightAll")
}

func parseMarkdown(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	renderer := custom_renderers.NewCustomRenderer("ComponentA")

	return markdown.ToHTML(md, parser, renderer)
}

type remoteMarkdownDoc struct {
	app.Compo

	Iid    string
	Iclass string
	Isrc   string

	md markdownContent
}

func NewRemoteMarkdownDoc() *remoteMarkdownDoc {
	return &remoteMarkdownDoc{}
}

func (d *remoteMarkdownDoc) ID(v string) *remoteMarkdownDoc {
	d.Iid = v
	return d
}

func (d *remoteMarkdownDoc) Class(v string) *remoteMarkdownDoc {
	d.Iclass = app.AppendClass(d.Iclass, v)
	return d
}

func (d *remoteMarkdownDoc) Src(v string) *remoteMarkdownDoc {
	d.Isrc = v
	return d
}

func (d *remoteMarkdownDoc) OnMount(ctx app.Context) {
	d.load(ctx)
}

func (d *remoteMarkdownDoc) OnUpdate(ctx app.Context) {
	d.load(ctx)
}

func (d *remoteMarkdownDoc) load(ctx app.Context) {
	src := d.Isrc
	ctx.ObserveState(markdownState(src)).
		While(func() bool {
			return src == d.Isrc
		}).
		OnChange(func() {
			//fmt.Println(d.md.Data)
			//ctx.Defer(scrollTo)
			//Window().ScrollToID(id)
			//app.Window().Call("mountReactComponents") // <-- Add this line here.
			// Create a new event
			event := app.Window().Get("Event").New("markdownUpdated")

			// Dispatch the event on the document
			app.Window().Get("document").Call("dispatchEvent", event)

		}).
		Value(&d.md)

	ctx.NewAction(GetMarkdown, app.T("path", d.Isrc))

}

func (d *remoteMarkdownDoc) Render() app.UI {
	return app.Div().
		ID(d.Iid).
		Class(d.Iclass).
		Body(
			ui.Loader().
				Class("heading").
				Class("fill").
				Loading(d.md.Status == loading).
				Err(d.md.Err).
				Label(fmt.Sprintf("Loading %s...", filepath.Base(d.Isrc))),
			app.If(d.md.Status == loaded,
				newMarkdownDoc().
					Class("fill").
					MD(d.md.Data),
			).Else(),
		)
}
