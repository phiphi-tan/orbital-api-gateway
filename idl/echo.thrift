namespace go echo

struct Request {
	1: string message (api.body = "message,required")
}

struct Response {
	1: string message
}

service Echo {
    Response echo(1: Request req)
}
