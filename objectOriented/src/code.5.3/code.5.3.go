package main

import "fmt"

type Document struct {
	sc   SpellChecker
	text string
}

func (d *Document) SetSpellChecker(sc SpellChecker) { d.sc = sc }
func (d Document) Spellcheck()                      { d.sc.Check(d.text) }

type SpellChecker interface {
	Check(text string)
}

type SCStupid struct{}
type SCClever struct{}

func (sc SCStupid) Check(text string) { fmt.Println("Stupid!") }
func (sc SCClever) Check(text string) { fmt.Println("Clever!") }

func main() {
	d := new(Document)
	d.SetSpellChecker(new(SCStupid))
	d.Spellcheck()
	d.SetSpellChecker(new(SCClever))
	d.Spellcheck()
}
