package main

func main() {}

type Speaker interface {
	Hello() string
}

type EnglishSpeaker struct{}
type GermanSpeaker struct {
	greeting string
}

// TODO: Implémenter l'interface Speaker pour les deux types
//
// func (type) nomFonction ... { ... }
//
