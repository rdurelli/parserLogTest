package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLogRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(logRoutes LogRoutes) Routes {
	return Routes{
		logRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
