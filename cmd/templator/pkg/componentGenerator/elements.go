package componentGenerator

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Element interface
type Element interface {
	String() string
	Generate() app.UI
}

// Base element holding common attributes
type BaseElement struct {
	Class string
	ID    string
}
