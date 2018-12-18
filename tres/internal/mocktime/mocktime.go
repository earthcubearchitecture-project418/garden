package mocktime

import (
	restful "github.com/emicklei/go-restful"
)

// New test service
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/api/v1/time").
		Doc("Organic free text search services").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

		// add in start point and length cursors
	service.Route(service.GET("/range").To(SearchCall).
		Doc("Query string search call").
		Param(service.QueryParameter("s", "Start time").DataType("string")).
		Param(service.QueryParameter("e", "Ending time").DataType("string")).
		// Param(service.QueryParameter("e", "Ending time").DataType("string")).
		// Writes([]OrganicResults{}).
		Operation("SearchCall"))

	return service
}

// SearchCall First test function..   opens each time..  not what we want..
// need to open indexes and maintain state
func SearchCall(request *restful.Request, response *restful.Response) {
	s := request.QueryParameter("s")
	e := request.QueryParameter("e")
	r := SimpleRange(s, e)

	response.Write([]byte(r))
}

// GeologicCall
// Take a geologic stage name and convert it to numeric range and
// a full stack of stage names through the hierarchy
// Leverage Simon's work and SPARQL calls.
func GeologicCall(request *restful.Request, response *restful.Response) {
	response.Write([]byte("hello world"))
}
