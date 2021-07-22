package goTeam

type signInReq struct {
	Name string `example:"Mark"`
	Sig  string `example:"3835c2448a04d7bd74b6a99c3a6dc1147"`
}

type signInResp struct {
	Token string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type signUpReq struct {
	Name string `example:"Mark"`
}

type signUpResp struct {
	Id    int    `example:"1"`
	Token string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type StatusOk struct {
	Status string `example:"OK"`
}
