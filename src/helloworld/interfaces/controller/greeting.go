package controller

import (
	"context"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

// GreetingController defines handlers to greeting requests.
type GreetingController struct {
	responder *web.Responder
}

// Inject injects controller dependencies.
func (c *GreetingController) Inject(r *web.Responder) {
	c.responder = r
}

// Greet greets the user.
func (c *GreetingController) Greet(_ context.Context, _ *web.Request) web.Result {
	return c.responder.Render("greeting/greet", nil)
}

// GreetPerson greets the user by name.
func (c *GreetingController) GreetPerson(_ context.Context, r *web.Request) web.Result {
	name, _ := r.Query1("name")
	return c.responder.Render("greeting/greet-person", struct {
		Name  string
		Today time.Time
	}{
		Name:  name,
		Today: time.Now(),
	})
}
