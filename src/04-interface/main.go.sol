package main
func main() {}

type Speaker interface {
	Hello() string
}

type EnglishSpeaker struct{}
type GermanSpeaker struct {
	greeting string
}

func (*EnglishSpeaker) Hello() string {
	return "Hello"
}
func (gs *GermanSpeaker) Hello() string {
	return "Guten " + gs.greeting
}
