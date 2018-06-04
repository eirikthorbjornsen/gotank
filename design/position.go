package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("position", func() { // Resources group related API endpoints
	Description("Get the position of the GOTANK")
	BasePath("/api/v1/position") // together. They map to REST resources for REST

	Action("Get", func() { // Actions define a single API endpoint together
		Description("Get the current global position of the gotank") // with its path, parameters (both path
		Routing(GET("/latlng"))                                      // parameters and querystring values) and payload

		Response(OK, PositionLatLngMedia) // Responses define the shape and status code
		Response(InternalServerError)     // of HTTP responses.
	})
})

// BottleMedia defines the media type used to render bottles.
var PositionLatLngMedia = MediaType("application/vnd.gotank.positionlatlng+json", func() {
	Description("A global position of the GOTANK")
	Attributes(func() {
		Attribute("lat", Number, "Current Latitude", func() {
			Minimum(-90)
			Maximum(90)
			Example(63.421)
		})
		Attribute("lon", Number, "Current Longitude", func() {
			Minimum(-180)
			Maximum(180)
			Example(10.423)
		})
		Attribute("orientation", Number, "Current orientation/compass heading (degrees). -1 means no data.", func() {
			Minimum(-1)
			Maximum(359)
			Example(42)
		})
		Attribute("cog", Number, "Course over ground (degrees). -1 means no data.", func() {
			Minimum(-1)
			Maximum(359)
			Example(42)
		})
		Attribute("sog", Number, "Speed over ground (km/h). -1 means no data.", func() {
			Minimum(-1)
			Example(0.5)
		})
		Attribute("hdop", Number, "Horizontal dilution of precision. -1 means no data.", func() {
			Minimum(-1)
			Example(1.9)
		})
		Attribute("numsats", Number, "Number of satellites. -1 means no data.", func() {
			Minimum(-1)
			Example(11)
		})

		Required("lat", "lon", "orientation", "cog", "sog", "hdop", "numsats")
	})
	View("default", func() {
		Attribute("lat")
		Attribute("lon")
		Attribute("orientation")
		Attribute("cog")
		Attribute("sog")
		Attribute("hdop")
		Attribute("numsats")
	})
})
