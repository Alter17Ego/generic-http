module example

go 1.15

replace github.com/Alter17Ego/generic-http => ../

require (
	github.com/Alter17Ego/generic-app v1.0.0
	github.com/Alter17Ego/generic-http v1.0.0
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-chi/render v1.0.1
)
