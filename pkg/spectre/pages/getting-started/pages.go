package getting_started

import (
	pages "amlwwalker/go-app-tuts/pkg/spectre/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// InstallationPage
type InstallationPage struct {
	app.Compo
}
// gettingStarted render method
func (h *InstallationPage) Render() app.UI {
	return pages.Page("Getting Started", "Getting Started with Spectre for Web Assembly (in Go)", "/web/documents/installation.md")
}

// elements pages
type CustomVersionPage struct {
	app.Compo
}

// elements render method
func (h CustomVersionPage) Render() app.UI {
	return pages.Page("Custom Versions", "Customising Spectre in Web Assembly", "/web/documents/customVersion.md")

}
