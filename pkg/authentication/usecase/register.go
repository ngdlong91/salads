package authentication_usecase

import "github.com/sirupsen/logrus"

type Register struct {
	req    struct{}
	resp   struct{}
	logger *logrus.Entry
	goNext bool
}

func NewRegisterWorker(logger *logrus.Entry) *Register {
	uc := &Register{
		logger: logger,
	}
	return uc
}
