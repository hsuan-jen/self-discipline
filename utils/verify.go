package utils

var (
	LoginVerify       = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify    = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}}
	WechatLoginVerify = Rules{"Token": {NotEmpty()}}
	IssueVerify       = Rules{"Content": {NotEmpty()}}
	PageInfoVerify    = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty(), Le("15"), Ge("1")}}
	GiveVerify        = Rules{"ArticleID": {NotEmpty()}}
)
