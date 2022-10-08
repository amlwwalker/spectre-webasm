package page_layouts

import (
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	pages "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// InstallationPage
type HeroPage struct {
	app.Compo
}
// gettingStarted render method
func (h *HeroPage) Render() app.UI {
	return pages.Page("Hero", "Heros are large title blocks", "/web/documents/heros.md",
		layouts.Hero("An example title", "an example description"),
	)
}
//
//func (h *HeroPage) Render() app.UI {
//	return pages.Page("Tiles", "Tiles are repeatable or embeddable information blocks.", "/web/documents/tiles.md",
//		app.Div().Class("columns").Body(
//			app.Div().Class("column col-9 col-sm-12").Body(
//				components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
//				components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
//			),
//		),
//	)
//}
