package impl

type EchoStruct struct{}

func (EchoStruct) Echo(input string) string {
	return input
}
