package react

import (
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	pages "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// InstallationPage
type CustomComponentPage struct {
	app.Compo
}

// gettingStarted render method
func (h *CustomComponentPage) Render() app.UI {
	return pages.NewPage("React Custom Components", "Custom React Components are possible!", "/documents/custom-components.md",
		components.CustomComponent("ComponentA", map[string]interface{}{
			"field": "data",
			"data":  "Some data for ComponentA",
		}),
	)
}
