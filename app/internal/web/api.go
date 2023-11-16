package web

//API interface tell that who ever is a API can be passed to the web server as a API handler

const (
	GET     = "GET"
	LIST    = "LIST"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
)

type API interface {
	GET()     //get the data
	LIST()    //get list of data as paginatiion
	POST()    // craete data
	PUT()     // update data
	DELETE()  //delete data
	HEAD()    //get required headers
	OPTIONS() //get options

	GetRoutMapping() map[string]string
}
