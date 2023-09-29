package markdown

import (
	"fmt"
	"github.com/amlwwalker/spectre-webasm/pkg/http"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"time"
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
	fmt.Println("tags ", a.Tags)
	path := a.Tags.Get("path")
	if path == "" {
		app.Log(errors.New("getting markdown failed").
			WithTag("reason", "empty path"))
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

	fmt.Println("retrieving markdown from", path)
	res, err := http.Get(ctx, "/retrieve"+path)
	if err != nil {
		md.Status = loadingErr
		md.Err = errors.New("getting markdown failed").Wrap(err)
		ctx.SetState(state, md)
		return
	}
	fmt.Println("warning - we are forcing a delay here")
	time.Sleep(300 * time.Millisecond) //fixme remove
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
