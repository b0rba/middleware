package impl

type EchoRPC struct{}

func (t *EchoRPC) Echo(req int, reply *int) error {
	*reply = req * 2
	return nil
}
