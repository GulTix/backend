package forms

func NewHandler() Handler {
	return &handlerImpl{}
}
