package authentication_usecase

import "github.com/sirupsen/logrus"

type Logout struct {
	req    struct{}
	resp   struct{}
	logger *logrus.Entry
	goNext bool
}

func NewLogoutWorker(logger *logrus.Entry) *Logout {
	uc := &Logout{
		logger: logger,
	}
	return uc
}
