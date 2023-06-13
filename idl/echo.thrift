namespace go echo

struct Request {
	1: string message
}

struct Response {
	1: string message
}

service Echo {
    Response echo(1: Request req)( api.post = '/echo/echo', api.param = 'true')
}
