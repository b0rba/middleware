package echoer

// Echo empty struct
type Echo struct{}

// Ech makes a echo
func (Echo) Ech(input string) string {
	return input
}