package utils

var (
	LoginVerify    = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}}
)
