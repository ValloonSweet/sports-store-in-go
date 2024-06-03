package admin

type OrdersHandler struct{}

func (handler OrdersHandler) GetData() string {
	return "This is orders handler"
}
