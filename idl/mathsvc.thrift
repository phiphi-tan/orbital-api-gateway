namespace go mathsvc

struct AddRequest {
	1: required double first
	2: required double second
}

struct SubtractRequest {
	1: required double first
	2: required double second
}

struct MultiplyRequest {
	1: required double first
	2: required double second
}

struct AddResponse {
	1: double addresult
}

struct SubtractResponse {
	1: double subtractresult
}

struct MultiplyResponse {
	1: double multiplyresult
}

service MathSvc {
	AddResponse add(1: AddRequest req)( api.post = '/mathsvc/add', api.param = 'true')
	SubtractResponse subtract(1: SubtractRequest req)( api.post = '/mathsvc/subtract', api.param = 'true')
	MultiplyResponse multiply(1: MultiplyRequest req)( api.post = '/mathsvc/multiply', api.param = 'true')
}
