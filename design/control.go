package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("control", func() {
	Description("Control the GOTANK")
	BasePath("/api/v1/control")
	Action("ControlManualy", func() {
		Description("Control the GOTANK manualy")
		Routing(PUT("/manualy"))
		Payload(ControlManualyMedia)
		Response(OK)
		Response(InternalServerError)
	})
})

var ControlManualyMedia = MediaType("application/vnd.gotank.control+json", func() {
	Description("Manualy control input")
	Attributes(func() {
		Attribute("forward", Number, "Percentage forwards control input", func() {
			Minimum(0)
			Maximum(100)
			Example(10)
		})
		Attribute("backward", Number, "Percentage backwards control input", func() {
			Minimum(0)
			Maximum(100)
			Example(20)
		})
		Attribute("left", Number, "Percentage left control input", func() {
			Minimum(0)
			Maximum(100)
			Example(30)
		})
		Attribute("right", Number, "Percentage right control input", func() {
			Minimum(0)
			Maximum(100)
			Example(40)
		})
		Required("forward", "backward", "left", "right")
	})
	View("default", func() {
		Attribute("forward")
		Attribute("backward")
		Attribute("left")
		Attribute("right")
	})
})
