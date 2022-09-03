package ace

import (
	"amlwwalker/go-app-tuts/pkg/spectre/components"
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var base = `
package main

import (
	"fmt"
)

`

type editor struct {
	app.Compo
	activeTab string
	ace app.Value
	lesson string
}

func NewEditor() *editor {
	return &editor{}
}

func (e *editor) OnMount(ctx app.Context) {
	e.activeTab = "untitled.txt"
	fmt.Println("mounting")
	e.ace = app.Window().Get("ace").Call("edit", "editor")
	e.ace.Call("setTheme", "ace/theme/dracula")
	e.ace.Call("setFontSize", 16)
	e.ace.Get("session").Call("setMode", "ace/mode/golang")
	fmt.Println("mounted")
}

func (e *editor) OnUpdate(ctx app.Context) {
	fmt.Println("updated")
}

func createLesson(lesson string) string {
	l := base + `func main() {
	fmt.Println("lesson ` + lesson + `")
}`
	fmt.Println(lesson, l)
	return l
}
func (e *editor) Render() app.UI {
	return app.Div().Body(
		components.Tabs(
		components.Tab("untitled.txt", e.activeTab == "untitled.txt", func(ctx app.Context, ev app.Event) {
			e.activeTab = "untitled.txt"
			fmt.Println(e.activeTab)
			e.ace.Get("session").Call("setValue", createLesson(e.activeTab))
			e.Update()
		}),
		components.Tab("untitled1.txt", e.activeTab == "untitled1.txt", func(ctx app.Context, ev app.Event) {
			e.activeTab = "untitled1.txt"
			fmt.Println(e.activeTab)
			e.ace.Get("session").Call("setValue", createLesson(e.activeTab))
			e.Update()
		}),
		components.Tab("untitled2.txt", e.activeTab == "untitled2.txt", func(ctx app.Context, ev app.Event) {
			e.activeTab = "untitled2.txt"
			fmt.Println(e.activeTab)
			e.ace.Get("session").Call("setValue", createLesson(e.activeTab))
			e.Update()
		}),
		components.Tab("untitled3.txt", e.activeTab == "untitled3.txt", func(ctx app.Context, ev app.Event) {
			e.activeTab = "untitled3.txt"
			fmt.Println(e.activeTab)
			e.ace.Get("session").Call("setValue", createLesson(e.activeTab))
			e.Update()
		}),
	),
	app.Div().Class("container").ID("editor-parent").Body(
		app.Div().ID("editor").Text(`comment`),
	),
	)
}

//overlay on code words https://codepen.io/oatssss/pen/oYxJQV
/*
$ace_light = [
	'chrome' => 'Chrome',
	'clouds' => 'Clouds',
	'crimson_editor' => 'Crimson Editor',
	'dawn' => 'Dawn',
	'dreamweaver' => 'Dreamweaver',
	'eclipse' => 'Eclipse',
	'github' => 'GitHub',
	'iplastic' => 'IPlastic',
	'katzenmilch' => 'KatzenMilch',
	'kuroir' => 'Kuroir',
	'solarized_light' => 'Solarized Light',
	'sqlserver' => 'SQL Server',
	'textmate' => 'TextMate',
	'tomorrow' => 'Tomorrow',
	'xcode' => 'XCode'
];
$ace_dark = [
	'ambiance' => 'Ambiance',
	'chaos' => 'Chaos',
	'clouds_midnight' => 'Clouds Midnight',
	'cobalt' => 'Cobalt',
	'dracula' => 'Dracula',
	'gob' => 'Greeon on Black',
	'gruvbox' => 'Gruvbox',
	'idle_fingers' => 'idle Fingers',
	'kr_theme' => 'krTheme',
	'merbivore' => 'Merbivore',
	'merbivore_soft' => 'Merbivore Soft',
	'mono_industrial' => 'Mono Industrial',
	'monokai' => 'Monokai',
	'pastel_on_dark' => 'Pastel on Dark',
	'solarized_dark' => 'Solarized Dark',
	'terminal' => 'Terminal',
	'tomorrow_night' => 'Tomorrow Night',
	'tomorrow_night_blue' => 'Tomorrow Night Blue',
	'tomorrow_night_bright' => 'Tomorrow Night Bright',
	'tomorrow_night_eighties' => 'Tomorrow Night 80s',
	'twilight' => 'Twilight',
	'vibrant_ink' => 'Vibrant Ink'
];

 */
