package authentication_usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/ngdlong91/salads/key/login"
)



type Login struct {
	req    struct{
		email string
		phone string
		password string
		method int
	}
	Resp   struct{
		Tokens string `json:"tokens"`
	}

	logger *logrus.Entry
	goNext bool
}

func (uc *Login) BeforeProcess() {
	uc.SetMoveNext(true)
}

func (uc *Login) Process() {
	switch uc.req.method {
	case login.ByPhone:
		uc.loginByPhone()
	case login.ByEmail:
		uc.loginByEmail()
	default:
		uc.reject()
	}
}

func (uc *Login) reject()  {

}

func (uc *Login) loginByPhone()  {

}

func (uc *Login) loginByEmail()  {

}


func (uc *Login) AfterProcess() {
	panic("implement me")
}

func (uc *Login) SetMoveNext(goNext bool) {
	panic("implement me")
}

func (uc *Login) IsMoveNext() bool {
	panic("implement me")
}

func NewLoginWorker(logger *logrus.Entry) *Login {
	uc := &Login{
		logger: logger,
	}
	return uc
}
