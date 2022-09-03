package media

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type youtubePlayer struct {
	app.Compo
	isrc   string
}

func NewYoutubePlayer() *youtubePlayer {
	return &youtubePlayer{}
}

func (y *youtubePlayer) Src(src string) *youtubePlayer {
	y.isrc = src
	return y
}
func (y *youtubePlayer) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("yt-contianer").Body(
			//app.Script().
			//	Src("//www.youtube.com/iframe_api").
			//	Async(true),
			app.IFrame().
				//Width(560).
				//Height(600).
				ID("yt-container").
				Allow("autoplay").
				Allow("accelerometer").
				Allow("encrypted-media").
				Allow("picture-in-picture").
				Sandbox("allow-presentation allow-same-origin allow-scripts allow-popups").
				Src("https://www.youtube.com/embed/"+y.isrc).Class("yt-video"),
		),
	)
}
