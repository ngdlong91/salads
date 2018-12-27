package authentication_usecase

import "github.com/sirupsen/logrus"

type ChangePassword struct {
	req    struct{}
	resp   struct{}
	logger *logrus.Entry
	goNext bool
}

func NewChangePasswordWorker(logger *logrus.Entry) *ChangePassword {
	uc := &ChangePassword{
		logger: logger,
	}
	return uc
}
