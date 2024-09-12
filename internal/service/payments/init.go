package payments

func NewService() Service {
	return &ServiceImpl{}
}
