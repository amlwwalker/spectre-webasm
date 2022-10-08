package markdown

import (
	"github.com/amlwwalker/spectre-webasm/pkg/http"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

type status int

const (
	neverLoaded status = iota
	loading
	loadingErr
	loaded
)

const (
	GetMarkdown = "/markdown/get"
)

func HandleGetMarkdown(ctx app.Context, a app.Action) {
	path := a.Tags.Get("path")
	if path == "" {
		app.Log(errors.New("getting markdown failed").
			Tag("reason", "empty path"))
		return
	}
	state := markdownState(path)

	var md markdownContent
	ctx.GetState(state, &md)
	switch md.Status {
	case loading, loaded:
		return
	}

	md.Status = loading
	md.Err = nil
	ctx.SetState(state, md)

	res, err := http.Get(ctx, path)
	if err != nil {
		md.Status = loadingErr
		md.Err = errors.New("getting markdown failed").Wrap(err)
		ctx.SetState(state, md)
		return
	}

	md.Status = loaded
	md.Data = string(res)
	ctx.SetState(state, md)
}

func markdownState(src string) string {
	return src
}

type markdownContent struct {
	Status status
	Err    error
	Data   string
}
