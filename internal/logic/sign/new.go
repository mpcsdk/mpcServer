package sign

import "li17server/internal/service"

type sGenerator struct{}

func New() *sGenerator {
	return &sGenerator{}
}

func init() {
	service.RegisterGenerator(New())
}
