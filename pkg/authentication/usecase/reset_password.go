package authentication_usecase

import "github.com/sirupsen/logrus"

type ResetPassword struct {
	req    struct{}
	resp   struct{}
	logger *logrus.Entry
	goNext bool
}

func NewResetPasswordWorker(logger *logrus.Entry) *ResetPassword {
	uc := &ResetPassword{
		logger: logger,
	}
	return uc
}
