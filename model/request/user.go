package request

type UserLocalLoginReqStruct struct {
	LoginName			string
	Password 			string
}
type UserRegReqStruct struct {
	UserLocalLoginReqStruct
	NickName 			string
	HeadImg 			string
}

