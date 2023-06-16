namespace go common

enum Err
{
	// gateway 10001- 19999
    BadRequest            = 10001,
	Unauthorized          = 10002,
	ServerNotFound        = 10003,
	ServerMethodNotFound  = 10004,
	MissingParameters     = 10005,
	RequestServerFail     = 10006,
	ServerHandleFail      = 10007,
	ResponseUnableParse   = 10008,

}
