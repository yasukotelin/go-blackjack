package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/yasukotelin/scrlib"
)

type Result int

const (
	Win Result = iota
	Loose
	Even
)

const (
	app = "go-blackjack"
)

var (
	dealer *Dealer
	player *Player
)

func wait() {
	time.Sleep(500 * time.Millisecond)
}

func printLogo() {
	var sb strings.Builder
	sb.WriteString("                   _     _            _     _            _\n")
	sb.WriteString("  __ _  ___       | |__ | | __ _  ___| | __(_) __ _  ___| | __\n")
	sb.WriteString(" / _` |/ _ \\ _____| '_ \\| |/ _` |/ __| |/ /| |/ _` |/ __| |/ /\n")
	sb.WriteString("| (_| | (_) |_____| |_) | | (_| | (__|   < | | (_| | (__|   <\n")
	sb.WriteString(" \\__, |\\___/      |_.__/|_|\\__,_|\\___|_|\\_\\/ |\\__,_|\\___|_|\\_\\\n")
	sb.WriteString(" |___/                                   |__/\n")

	fmt.Println(sb.String())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	dealer = NewDealer("Dealer")
	player = NewPlayer("You")

	scrlib.Clear()
	fmt.Println()
	printLogo()
	fmt.Println()

	dealer.Sayln(fmt.Sprintf("Hi. I'm %s.", dealer.Name()))
	dealer.Sayln(fmt.Sprintf("Welcome to %s", app))
	wait()
	fmt.Println()
	dealer.Sayln("Get 3 wins earlier than me!")
	fmt.Println()
	dealer.Ask("Are you ready? (Please press any key.) ")

	scrlib.Clear()

	result := playGameTimes(3)

	switch result {
	case Win:
		scrlib.Clear()
		fmt.Println()
		dealer.Sayln("You totally win!!")
		dealer.Sayln("Congratulations!")
	case Loose:
		scrlib.Clear()
		fmt.Println()
		dealer.Sayln("You totally loose.")
	}

	dealer.Ask("(Please press any key)")
	scrlib.Clear()
}

func playGameTimes(times int) Result {
	var result Result
	for {
		if player.WinCount >= times {
			result = Win
			break
		}
		if dealer.WinCount >= times {
			result = Loose
			break
		}
		dealer.Setup()
		player.Setup()

		r := playGame()

		switch r {
		case Win:
			player.WinCount++
			fmt.Println()
			dealer.Sayln("You win!!")
		case Loose:
			dealer.WinCount++
			fmt.Println()
			dealer.Sayln("You loose")
		case Even:
			fmt.Println()
			dealer.Sayln("just even")
		}
		dealer.Ask("(Please press any key)")
		scrlib.Clear()
	}
	return result
}

func playGame() Result {
	dealer.SetFirst(player)
	showHideHand(dealer)
	wait()

	// プレイヤーターン
	for {
		showHand(player)
		if isBust(player) {
			return Loose
		}
		var in string
		for {
			in = dealer.Ask("HIT [h] or STAND [s]: ")
			if in == "h" || in == "s" {
				break
			}
		}

		if in == "h" {
			dealer.Pass(player)
		} else {
			break
		}
	}

	// ディーラーターン
	for {
		wait()
		showHand(dealer)
		if isBust(dealer) {
			return Win
		}
		if dealer.Total() > 16 {
			break
		}
		dealer.Pass(dealer)
	}

	// 勝敗判定
	if player.Total() == dealer.Total() {
		return Even
	} else if player.Total() > dealer.Total() {
		return Win
	} else {
		return Loose
	}
}

func showHand(p IPlayer) {
	fmt.Println()
	fmt.Printf("=== %s ======\n", p.Name())
	p.ShowHand()
	fmt.Println()
}

func showHideHand(d *Dealer) {
	fmt.Println()
	fmt.Printf("=== %s ======\n", d.Name())
	d.showHideHand()
	fmt.Println()
}

func isBust(p IPlayer) bool {
	if p.Total() > 21 {
		return true
	}
	return false
}
