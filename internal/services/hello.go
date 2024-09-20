package services

type HelloculateService interface {
	Hello() (string, error)
}

type helloculateService struct {
}

func (c helloculateService) Hello() (string, error) {
	return "hello world", nil
}

func NewCalculateService() HelloculateService {
	return new(helloculateService)
}
