package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("config", func() {
	BasePath("/api/v1/config")
	Description("Configuration")

	Action("Get", func() {
		Description("Get generic configuration")
		Routing(GET("/generic"))
		Payload(ConfigurationMedia)
		Response(InternalServerError)
	})

	Action("Modify", func() {
		Description("Modify generic configuration")
		Routing(PUT("/generic"))
		Payload(ConfigurationMedia)
		Response(OK, "plain/text")
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})
})

// ConfigurationMedia describes the configuration data
var ConfigurationMedia = MediaType("application/vnd.gotank.configuration+json", func() {
	Description("Configuration parameters")
	Attributes(func() {
		Attribute("mode", Integer, "Mode to use", func() {
			Minimum(1)
			Maximum(5)
			Example(3)
		})
		Attribute("gps", String, "GPS provider setting", func() {
			Enum("onboard", "static", "external")
		})
		Attribute("compass", String, "Compass provider setting", func() {
			Enum("onboard", "static", "external")
		})

		Required("mode", "gps", "compass")
	})
	View("default", func() {
		Attribute("mode")
		Attribute("gps")
		Attribute("compass")
	})
})
