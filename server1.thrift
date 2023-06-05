namespace go api

struct AddRequest {
	1: double first
	2: double second
}

struct SubtractRequest {
	1: double first
	2: double second
}

struct MultiplyRequest {
	1: double first
	2: double second
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

service Server1 {
	AddResponse add(1: AddRequest req)
	SubtractResponse subtract(1: SubtractRequest req)
	MultiplyResponse multiply(1: MultiplyRequest req)
}
