package response

const (
	CodeSuccess      = 20001 // Success
	CodeParamInvalid = 20003 // Email is invalid
	TokenInvalid     = 30001 // Token is invalid
)

// message
var msg = map[int]string{
	CodeSuccess:      "Success",
	CodeParamInvalid: "Email is invalid",
	TokenInvalid:     "Token is invalid",
}
