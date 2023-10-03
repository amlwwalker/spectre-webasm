package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Extend the app.UI interface with a Body method.
type UIBody interface {
	app.UI
	Body(children ...app.UI) interface{}
}

var htmlToGoAppMap = map[string]func() app.UI{
	"a":          func() app.UI { return app.A() },
	"abbr":       func() app.UI { return app.Abbr() },
	"address":    func() app.UI { return app.Address() },
	"article":    func() app.UI { return app.Article() },
	"aside":      func() app.UI { return app.Aside() },
	"audio":      func() app.UI { return app.Audio() },
	"b":          func() app.UI { return app.B() },
	"blockquote": func() app.UI { return app.Blockquote() },
	"body":       func() app.UI { return app.Body() },
	"br":         func() app.UI { return app.Br() },
	"button":     func() app.UI { return app.Button() },
	"canvas":     func() app.UI { return app.Canvas() },
	"cite":       func() app.UI { return app.Cite() },
	"code":       func() app.UI { return app.Code() },
	"div":        func() app.UI { return app.Div() },
	"em":         func() app.UI { return app.Em() },
	"footer":     func() app.UI { return app.Footer() },
	"form":       func() app.UI { return app.Form() },
	"h1":         func() app.UI { return app.H1() },
	"h2":         func() app.UI { return app.H2() },
	"h3":         func() app.UI { return app.H3() },
	"h4":         func() app.UI { return app.H4() },
	"h5":         func() app.UI { return app.H5() },
	"h6":         func() app.UI { return app.H6() },
	"head":       func() app.UI { return app.Head() },
	"header":     func() app.UI { return app.Header() },
	"hr":         func() app.UI { return app.Hr() },
	"i":          func() app.UI { return app.I() },
	"iframe":     func() app.UI { return app.IFrame() },
	"img":        func() app.UI { return app.Img() },
	"input":      func() app.UI { return app.Input() },
	"label":      func() app.UI { return app.Label() },
}
