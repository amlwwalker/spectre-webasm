package layouts

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// FlexBox auto fits content
func FlexBox(width string, body ...app.UI) app.HTMLDiv {
	return app.Div().Class("container").Style("max-width", width).Body(
		app.Div().Class("columns").Body(body...),
	)
}
// ThreeColumn is a helper function forces two column content
func ThreeColumn(body ...app.UI) app.HTMLDiv {
	var wrappers []app.UI
	for _, v := range body {
		wrappers = append(wrappers, app.Div().Class("column col-4 col-xs-12").Body(v))
	}
	return FlexBox("800px",
		wrappers...,
	)
}
// FourColumn is a helper function forces two column content
func FourColumn(body ...app.UI) app.HTMLDiv {
	var wrappers []app.UI
	for _, v := range body {
		wrappers = append(wrappers, app.Div().Class("column col-3 col-xs-12").Body(v))
	}
	return FlexBox("800px",
		wrappers...,
	)
}
// TwoColumn is a helper function forces two column content
func TwoColumn(body ...app.UI) app.HTMLDiv {
	var wrappers []app.UI
	for _, v := range body {
		wrappers = append(wrappers, app.Div().Class("column col-6 col-xs-12").Body(v))
	}
	return FlexBox("800px",
		wrappers...,
	)
}

func GaplessFlexBox(width string, body ...app.UI) app.HTMLDiv {
	return FlexBox(width, body...).Class("col-gapless")
}
