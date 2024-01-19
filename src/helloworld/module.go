package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"flamingo.me/helloworld/src/helloworld/interfaces/controller"
)

type (
	// Module defines the dingo module.
	Module struct{}

	routes struct {
		greetingController *controller.GreetingController
	}
)

// Configure defines module configuration.
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

func (r *routes) Inject(gc *controller.GreetingController) {
	r.greetingController = gc
}

// Routes configures the routes handled by the module.
func (r *routes) Routes(registry *web.RouterRegistry) {
	_ = registry.MustRoute("/greet", `helloworld.greet`)
	_ = registry.MustRoute("/greet-person", `helloworld.greetperson`)

	registry.HandleGet("helloworld.greet", r.greetingController.Greet)
	registry.HandleGet("helloworld.greetperson", r.greetingController.GreetPerson)
}
