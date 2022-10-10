package components

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/lithammer/shortuuid/v4"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"strings"
)
func isFocus(path string) bool {
	fmt.Println("according menu ", path)
	if strings.Contains(app.Window().URL().Path, path) {
		return true
	}
	return false
}

func urlLink(str string) string {
	return strcase.ToLowerCamel(str)
}
// Modal is a generic modal with a close handler
type Modal struct {
	app.Compo
	Size string //Size is the size of the modal
	ID           string   // HTML ID of the modal; must be unique across the page
	Icon         string   // Class of the icon to use to the left of the title; may be empty
	Title        string   // Title of the modal
	Class        string   // Class to be applied to the modal's outmost component
	Body         []app.UI // Body of the modal
	Footer       []app.UI // Footer of the modal
	//DisableFocus bool     // Disable auto-focusing the modal; useful if a child component, i.e. an input should be focused instead
	OnClose func() // Handler to call when closing/cancelling the modal
}
/*
<div class="modal active" id="modal-id">
  <a href="#close" class="modal-overlay" aria-label="Close"></a>
  <div class="modal-container">
    <div class="modal-header">
      <a href="#close" class="btn btn-clear float-right" aria-label="Close"></a>
      <div class="modal-title h5">Modal title</div>
    </div>
    <div class="modal-body">
      <div class="content">
        <!-- content here -->
      </div>
    </div>
    <div class="modal-footer">
      ...
    </div>
  </div>
</div>
 */

func (m *Modal) Render() app.UI {
	modal := app.Div().Class("modal active " + m.Size + " " + m.Class).OnClick(func(ctx app.Context, e app.Event) {
		m.OnClose()
	})
	modal.ID(m.ID).Body(
			app.A().Class("modal-overlay").Aria("label", "Close"),
		app.Div().Class("modal-container").Body(
			app.Div().Class("modal-header").Body(
				app.A().Class("btn btn-clear float-right").Aria("label", "Close"),
				app.Div().Class("modal-title h5").Text(m.Title),
			),
			app.Div().Class("modal-body").Body(
				m.Body...,
			),
			app.Div().Class("modal-footer").Body(
				m.Footer...,
			),
		),
	)
	return modal
}
//Avatar
/*
<!-- Show initals when avatar image is unavailable or not fully loaded -->
<figure class="avatar avatar-xl" data-initial="YZ" style="background-color: #5755d9;">
  <img src="img/avatar-1.png" alt="...">
</figure>
 */
func Avatar(size, src, initials, color string) app.HTMLFigure {
	return app.Figure().Class("avatar").Class(size).Attr("data-initial", initials).Style("background-color", color).Body(
		app.If(src != "",
			app.Img().Src(src).Alt("..."),
		),
	)
}

func Button(text, class string, onClick func(ctx app.Context, e app.Event)) app.HTMLButton {
	return app.Button().Class("btn").Class(class).Text(text).OnClick(onClick)
}

func ImageFigure(src string) app.HTMLDiv {
	return app.Div().Class("avatar").Body(
		app.Img().Src(src))
}
func NotificationBadge(text string, dataBadge int) app.HTMLSpan {
	return app.Span().Class("badge").Attr("data-badge", dataBadge).Text(text)
}
func ButtonBadge(text string, dataBadge int, onClick func(ctx app.Context, e app.Event)) app.HTMLSpan {
	return app.Span().Class("btn badge").Attr("data-badge", dataBadge).Text(text).OnClick(onClick)
}
func FigureBadge(figureSrc string, dataBadge int, onClick func(ctx app.Context, e app.Event)) app.HTMLFigure {
	figure := app.Figure().Class("avatar badge").Attr("data-badge", dataBadge).Body(
		app.Img().Src(figureSrc),
	).OnClick(onClick)
	if onClick != nil {
		figure.Class("c-hand")
	}
	return figure
}


// Card
/*
<div class="card">
  <div class="card-image">
    <img src="img/osx-el-capitan.jpg" class="img-responsive">
  </div>
  <div class="card-header">
    <div class="card-title h5">...</div>
    <div class="card-subtitle text-gray">...</div>
  </div>
  <div class="card-body">
    ...
  </div>
  <div class="card-footer">
    <button class="btn btn-primary">...</button>
  </div>
</div>
 */
func Card(title, body app.UI, subTitle []app.UI, src string, footer app.UI) app.HTMLDiv {
	return app.Div().Class("card").Body(
		app.Div().Class("card-image").Body(
			app.Img().Src(src).Class("img-responsive"),
		),
		app.Div().Class("card-header").Body(
			app.Div().Class("card-title h5").Body(title),
			app.Range(subTitle).Slice(func(i int) app.UI{
				return app.Div().Class("card-subtitle text-gray").Body(subTitle[i])
			}),
		),
		app.Div().Class("card-body").Body(
			body,
		),
		app.Div().Class("card-footer").Body(
			footer,
		),
	)
}

// Toast
/*
<div class="toast toast-primary">
  <button class="btn btn-clear float-right"></button>
  Lorem ipsum dolor sit amet, consectetur adipiscing elit.
</div>
*/
func Toast(message, toastType string, active, close bool) app.HTMLDiv {
	toast := app.Div().Class("toast toast-"+toastType)
	if !active {
		toast.Class("d-none")
	}
	var body []app.UI
	if close {
			body = append(body, app.Button().Class("btn btn-clear float-right"))
	}
	body = append(body, app.Span().Text(message))
	toast.Body(body...)
	//toast.Text(message)
	return toast
}
// Steps
/*
<ul class="step">ds
  <li class="step-item">
    <a href="#" class="tooltip" data-tooltip="Step 1">Step 1</a>
  </li>
  <li class="step-item active">
    <a href="#" class="tooltip" data-tooltip="Step 2">Step 2</a>
  </li>
  <li class="step-item">
    <a href="#" class="tooltip" data-tooltip="Step 3">Step 3</a>
  </li>
  <li class="step-item">
    <a href="#" class="tooltip" data-tooltip="Step 4">Step 4</a>
  </li>
</ul>
 */
func Steps() app.HTMLDiv {
	fmt.Println("app.Window().URL().Fragment", app.Window().URL().Fragment)
	return app.Div().Class("docs-demo columns").Body(
		app.Div().Class("column col-12").Body(
			app.Ul().Class("step").Body(
				app.If( app.Window().URL().Fragment == "step1",
					app.Li().Class("step-item active").Body(
						app.A().Class("tooltip").Href("#step1").Attr("data-tooltip", "Step 1 Tooltip"),
					),
				).Else(
					app.Li().Class("step-item").Body(
						app.A().Class("tooltip").Href("#step1").Attr("data-tooltip", "Step 1 Tooltip"),
					),
				),
				app.If( app.Window().URL().Fragment == "step2",
					app.Li().Class("step-item active").Body(
						app.A().Class("tooltip").Href("#step2").Attr("data-tooltip", "Step 2 Tooltip"),
					),
				).Else(
					app.Li().Class("step-item").Body(
						app.A().Class("tooltip").Href("#step2").Attr("data-tooltip", "Step 2 Tooltip"),
					),
				),
				app.If( app.Window().URL().Fragment == "step3",
					app.Li().Class("step-item active").Body(
						app.A().Class("tooltip").Href("#step3").Attr("data-tooltip", "Step 3 Tooltip"),
					),
				).Else(
					app.Li().Class("step-item").Body(
						app.A().Class("tooltip").Href("#step3").Attr("data-tooltip", "Step 3 Tooltip"),
					),
				),
			),
		),
	)
}
//Nav
/*
<ul class="nav">
  <li class="nav-item">
    <a href="#">Elements</a>
  </li>
  <li class="nav-item active">
    <a href="#">Layout</a>
    <ul class="nav">
      <li class="nav-item">
        <a href="#">Flexbox grid</a>
      </li>
      <li class="nav-item">
        <a href="#">Responsive</a>
      </li>
      <li class="nav-item">
        <a href="#">Navbar</a>
      </li>
      <li class="nav-item">
        <a href="#">Empty states</a>
      </li>
    </ul>
  </li>
  <li class="nav-item">
    <a href="#">Components</a>
  </li>
  <li class="nav-item">
    <a href="#">Utilities</a>
  </li>
</ul>
 */
func Nav() app.HTMLUl {
	return app.Ul().Class("docs-nav").Class("nav").Body(
		app.Li().Class("nav-item").Body(
			app.A().Href("#").Text("Getting Started"),
		),
		app.Li().Class("nav-item").Body(
			app.A().Href("#").Text("Elements"),
		),
		app.Li().Class("nav-item active").Body(
			app.A().Href("#").Text("Layout"),
			app.Ul().Class("nav").Body(
				app.Li().Class("nav-item").Body(
					app.A().Href("#").Text("Flexbox grid"),
				),
				app.Li().Class("nav-item").Body(
					app.A().Href("#").Text("Responsive"),
				),
				app.Li().Class("nav-item").Body(
					app.A().Href("#").Text("Hero"),
				),
				app.Li().Class("nav-item").Body(
					app.A().Href("#").Text("Navbar"),
				),
			),
		),
		app.Li().Class("nav-item").Body(
			app.A().Href("#").Text("Components"),
		),
		app.Li().Class("nav-item").Body(
			app.A().Href("#").Text("Utilities"),
		),
		app.Li().Class("nav-item").Body(
			app.A().Href("#").Text("Experimentals"),
		),
	)
}

func Accordion(menus ...app.UI) app.HTMLDiv {
	return app.Div().Class("accordion-container").Body(
		menus...
	)
}

func AccordionMenu(pathID, menuTitle string, menuElements []string) app.UI {
	uuid := shortuuid.New()
	return app.Div().Class("accordion").Body(
		app.Input().ID("accordion-"+ pathID + "-" + uuid).Type("checkbox").Name("docs-accordion-checkbox").Hidden(true).Checked(isFocus(pathID)),
		app.Label().Class("accordion header c-hand").For("accordion-"+ pathID + "-" + uuid).Text(menuTitle),
		app.Div().Class("accordion-body").Body(
			app.Ul().Class("menu menu-nav").Body(
				app.Range(menuElements).Slice(func(i int) app.UI {
					return app.Li().Class("menu-item").Body(
						app.A().Href("/"+ pathID +"/"+ urlLink(menuElements[i])).Text(menuElements[i]),
					)
				}),
			),
		),
	)
}

func Panel(titleContent, subtitleContent, navContent, bodyContent, footerContent app.UI) app.HTMLDiv {
	return app.Div().Class("panel").Body(
		app.Div().Class("panel-header").Body(
			app.Div().Class("panel-title").Body(
				titleContent,
			).Class("text-center"),
			app.Div().Class("panel-subtitle").Body(
				subtitleContent,
			).Class("text-center"),
		),
		app.Div().Class("panel-nav").Body(
			navContent,
		),
		app.Div().Class("panel-body").Body(
			bodyContent,
		),
		app.Div().Class("panel-footer").Body(
			footerContent,
		),
	)
}

// Tile
/*
<div class="tile">
  <div class="tile-icon">
    <div class="example-tile-icon">
      <i class="icon icon-file centered"></i>
    </div>
  </div>
  <div class="tile-content">
    <p class="tile-title">The Avengers</p>
    <p class="tile-subtitle">Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...</p>
  </div>
  <div class="tile-action">
    <button class="btn btn-primary">Join</button>
  </div>
</div>
 */
func Tile(title, subTitle, buttonText, imageSrc string) app.HTMLDiv {
	return app.Div().Class("tile").Body(
		app.If(imageSrc != "",
			app.Div().Class("tile-icon").Body(
			ImageFigure(imageSrc),
		)),
		app.Div().Class("tile-content").Body(
			app.P().Class("tile-title").Class("h4").Text(title),
			app.P().Class("tile-subtitle").Text(subTitle),
		),
		app.Div().Class("tile-action").Body(
			app.Button().Class("btn btn-primary").Text(buttonText),
		),
	)
}

// Tabs
/*
<ul class="tab tab-block">
  <li class="tab-item active">
    <a href="#">Music</a>
  </li>
  <li class="tab-item">
    <a href="#" class="active">Playlists</a>
  </li>
  <li class="tab-item">
    <a href="#">Radio</a>
  </li>
  <li class="tab-item">
    <a href="#">Connect</a>
  </li>
</ul>
 */
func Tabs(tabs ...app.UI) app.HTMLUl {
	return app.Ul().Class("tab tab-block").Body(
		tabs...
	)
}

func Tab(text string, active bool, onClick func(ctx app.Context, e app.Event)) app.HTMLLi {
	tab := app.Li().Class("tab-item").Body(
		app.A().Href("#").Text(text),
	).OnClick(onClick)
	if active {
		tab.Class("active")
	}
	return tab
}

func Breadcrumbs(content []app.UI) app.HTMLUl {
	return app.Ul().Class("breadcrumb").Body(
		app.Range(content).Slice(func(i int) app.UI {
			return app.Li().Class("breadcrumb-item").Body(
				content[i],
			)
		}),
	)
}

// Chip
/*
<span class="chip">Crime</span>

<span class="chip">
  Biography
  <a href="#" class="btn btn-clear" aria-label="Close" role="button"></a>
</span>

<div class="chip">
  <img src="img/avatar-1.png" class="avatar avatar-sm">
  Yan Zhu
  <a href="#" class="btn btn-clear" aria-label="Close" role="button"></a>
</div>
 */
func Chip(text, size, src string, close bool) app.HTMLSpan {
	return app.Span().Class("chip").Body(
		app.If(src != "", app.Img().Src(src).Class("avatar").Class(size)),
		app.Span().Text(text),
		app.If(close, app.A().Href("#").Class("btn btn-clear").Aria("label", "Close").Attr("role", "button")),
	)
}


// EmptyState
/*
<div class="empty">
  <div class="empty-icon">
    <i class="icon icon-people"></i>
  </div>
  <p class="empty-title h5">You have no new messages</p>
  <p class="empty-subtitle">Click the button to start a conversation.</p>
  <div class="empty-action">
    <button class="btn btn-primary">Send a message</button>
  </div>
</div>
 */
func EmptyState(icon, title, subTitle, buttonClass, buttonText string, onClick func(ctx app.Context, e app.Event)) app.HTMLDiv {
	return app.Div().Class("empty").Body(
		app.Div().Class("empty-icon").Body(
			app.I().Class("icon").Class(icon),
		),
		app.P().Class("empty-title h5").Text(title),
		app.P().Class("empty-subtitle").Text(subTitle),
		app.Div().Class("empty-action").Body(
			app.Button().Class("btn").Class(buttonClass).Text(
				buttonText,
			).OnClick(onClick),
		),
	)
}
