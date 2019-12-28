package server

type HelloService struct {
}

func (*HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}