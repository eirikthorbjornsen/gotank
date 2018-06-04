package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("gotank", func() {
	Title("The GOTANK API")
	Description(`This is the API to the gotank!

		It is the interface to control and get information about the GOTANK.
		It is also used to set different configurations and setting for the GOTANK `)
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

var _ = Resource("about", func() {
	BasePath("/api")
	Action("ApiVersion", func() {
		Routing(GET("/"))
		Response(OK, APIVersionMedia)
	})
})

var APIVersionMedia = MediaType("application/vnd.gotank.apiVersion", func() {
	Description("Returns supported api versions")
	Attributes(func() {
		Attribute("versions", ArrayOf(String), "Supported api versions", func() {
			Example([]string{"v1", "v2"})
		})
		Required("versions")
	})
	View("default", func() {
		Attribute("versions")
	})
})
