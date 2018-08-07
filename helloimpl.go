package main

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}
func (h *HelloHandler) HelloString(para string) (string, error) {
	return para, nil
}
func (h *HelloHandler) HelloBoolean(para bool) (r bool, e error) {
	return para, nil
}
func (h *HelloHandler) HelloInt(para int32) (r int32, e error) {
	return para, nil
}
func (h *HelloHandler) HelloVoid() (e error) {
	return nil
}
func (h *HelloHandler) HelloNull() (r string, e error) {
	return "hello null", nil
}
