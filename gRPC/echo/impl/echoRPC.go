package impl

type EchoRPC struct{}

func (t *EchoRPC) Echo(req string, reply *string) error {
	*reply = req
	return nil
}
