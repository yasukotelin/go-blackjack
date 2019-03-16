package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yasukotelin/cardlib"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

type IPlayer interface {
	Name() string
	Total() int
	Say(string)
	Sayln(string)
	Setup()
	ShowHand()
	AddHand(*cardlib.Card)
}

type Player struct {
	name     string
	Hand     []cardlib.Card
	WinCount int
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Total() int {
	return total(p.Hand)
}

func (p *Player) Say(msg string) {
	fmt.Printf("%s > %s", p.Name(), msg)
}

func (p *Player) Sayln(msg string) {
	fmt.Printf("%s > %s\n", p.Name(), msg)
}

func (p *Player) Setup() {
	p.Hand = make([]cardlib.Card, 0, 10)
}

func (p *Player) ShowHand() {
	var c []string
	for _, h := range p.Hand {
		c = append(c, h.String())
	}
	p.Sayln(strings.Join(c, " "))
	p.Sayln(fmt.Sprintf("[total] %d", p.Total()))
}

func (p *Player) AddHand(c *cardlib.Card) {
	p.Hand = append(p.Hand, *c)
}

type Dealer struct {
	name     string
	Emoticon string
	Deck     *cardlib.Deck
	Hand     []cardlib.Card
	WinCount int
}

func NewDealer(name string) *Dealer {
	return &Dealer{
		name:     name,
		Emoticon: "ʕ◔ϖ◔ʔ",
	}
}

func (d *Dealer) Name() string {
	return d.name
}

func (d *Dealer) Total() int {
	return total(d.Hand)
}

func (d *Dealer) Say(msg string) {
	fmt.Printf("%s > %s", d.Emoticon, msg)
}

func (d *Dealer) Sayln(msg string) {
	fmt.Printf("%s > %s\n", d.Emoticon, msg)
}

func (d *Dealer) Ask(msg string) string {
	d.Say(msg)
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func (d *Dealer) ShowHand() {
	var c []string
	for _, h := range d.Hand {
		c = append(c, h.String())
	}
	d.Sayln(strings.Join(c, " "))
	d.Sayln(fmt.Sprintf("[total] %d", d.Total()))
}

func (d *Dealer) showHideHand() {
	first := d.Hand[0]
	d.Sayln(fmt.Sprintf("%s ??", first.String()))
	d.Sayln(fmt.Sprintf("[total] %d + ??", numconvBj(first)))
}

func (d *Dealer) AddHand(c *cardlib.Card) {
	d.Hand = append(d.Hand, *c)
}

func (d *Dealer) Setup() {
	d.Deck = cardlib.NewDeck()
	d.Deck.RemoveJoker()
	d.Shuffle()
	d.Hand = make([]cardlib.Card, 0, 10)
}

func (d *Dealer) Shuffle() {
	d.Sayln("Shuffle.")
	d.Deck.Shuffle()
}

func (d *Dealer) Pass(p IPlayer) {
	p.AddHand(d.Deck.Draw())
}

func (d *Dealer) SetFirst(p *Player) {
	d.Sayln("Set.")
	for i := 0; i < 2; i++ {
		p.Hand = append(p.Hand, *d.Deck.Draw())
		d.Hand = append(d.Hand, *d.Deck.Draw())
	}
}

func total(cards []cardlib.Card) int {
	var total int
	for _, c := range cards {
		total += numconvBj(c)
	}
	return total
}

func numconvBj(card cardlib.Card) int {
	switch card.Number {
	case 11, 12, 13:
		return 10
	default:
		return card.Number
	}
}
