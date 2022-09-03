package layouts

import (
	"amlwwalker/go-app-tuts/pkg/media"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Dashboard represents a control area left and a work area right
/*
<div class="container">
<div class="columns">
<div class="column col-6">col-6</div>
<div class="column col-3">col-3</div>
<div class="column col-2">col-2</div>
<div class="column col-1">col-1</div>
</div>
</div>
 */
func Dashboard(nav app.UI) app.HTMLDiv {
	return FlexBox("800px",
		app.Div().Class("column col-3").Text("text"),
		app.Div().Class("column col-12").Text("hello right"),
	)
}

//Navbar
/*
<header class="navbar">
  <section class="navbar-section">
    <a href="..." class="navbar-brand mr-2">Spectre.css</a>
    <a href="..." class="btn btn-link">Docs</a>
    <a href="..." class="btn btn-link">GitHub</a>
  </section>
  <section class="navbar-center">
    <!-- centered logo or brand -->
  </section>
  <section class="navbar-section">
    <div class="input-group input-inline">
      <input class="form-input" type="text" placeholder="search">
      <button class="btn btn-primary input-group-btn">Search</button>
    </div>
  </section>
</header>
 */
func NavBar() app.HTMLHeader {
	return app.Header().Class("navbar").Body(
		app.Section().Class("navbar-section").Body(
			app.A().Href("#").Class("navbar-brand mr-2").Text("Spectre.css"),
			app.A().Href("#").Class("btn btn-link").Text("Docs"),
			app.A().Href("#").Class("btn btn-link").Text("Github"),
		),
		app.Section().Class("navbar-center").Body(
			app.Img().Src("https://picturepan2.github.io/spectre/img/spectre-logo.svg"),
		),
		app.Section().Class("navbar-section").Body(
			app.Div().Class("input-group input-inline").Body(
				app.Input().Class("form-input").Type("text").Placeholder("search"),
				app.Button().Class("btn btn-primary input-group-btn").Text("Search"),
			),
		),
	)
}

//Sidebar
/*
<div class="off-canvas">
  <!-- off-screen toggle button -->
  <a class="off-canvas-toggle btn btn-primary btn-action" href="#sidebar-id">
    <i class="icon icon-menu"></i>
  </a>

  <div id="sidebar-id" class="off-canvas-sidebar">
    <!-- off-screen sidebar -->
  </div>

  <a class="off-canvas-overlay" href="#close"></a>

  <div class="off-canvas-content">
    <!-- off-screen content -->
  </div>
</div>
 */
func SideBar(sidebarContent, bodyContent app.UI) app.HTMLDiv {
	return FlexBox("100%",
		app.Div().Class("docs-container").Class("off-canvas").Class("off-canvas-sidebar-show").Style("min-height", "100vh;").Body(
			app.A().Class("off-canvas-toggle btn btn-primary btn-action").Href("#sidebar-id").Body(
				app.I().Class("icon icon-menu"),
			),
			app.Div().ID("sidebar-id").Class("docs-sidebar").Class("off-canvas-sidebar").Class("flex-centered").Body(
				app.Div().Class("docs-nav").Body(
					sidebarContent,
				),
			),
			app.A().Class("off-canvas-overlay").Href("#close"),
			app.Div().Class("off-canvas-content").Body(
				bodyContent,
			),
		),
	)
}

// Hero
/*
<div class="hero bg-gray">
  <div class="hero-body">
    <h1>Hero title</h1>
    <p>This is a hero example</p>
  </div>
</div>
 */
func Hero(title, description string) app.HTMLDiv {
	return app.Div().Class("hero bg-gray").Body(
		app.Div().Class("hero-title").Body(
			app.H1().Text(title),
			app.P().Text(description)),
	)
}

func MediaHero(title, description string, src string) app.HTMLDiv {
	return app.Div().Class("hero bg-gray").Body(
		app.Div().Class("hero-title").Body(
			app.H1().Text(title),
			app.P().Text(description),
			media.NewYoutubePlayer().Src(src),
		),
	)
}

