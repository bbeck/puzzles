package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const (
	TargetMin = 17
	TargetMax = 61
)

func main() {
	bots, actions := InputToBotsAndActions()
	values := make(map[int][]int)

	var bot Bot
	for len(actions) > 0 {
		var action Action
		action, actions = actions[0], actions[1:]

		bot = bots[action.Bot]
		if len(values[bot.ID]) == 0 {
			values[bot.ID] = append(values[bot.ID], action.Value)
			continue
		}

		min := Min(values[bot.ID][0], action.Value)
		max := Max(values[bot.ID][0], action.Value)
		values[bot.ID] = nil

		if min == TargetMin && max == TargetMax {
			break
		}

		if bot.LowKind == "bot" {
			actions = append(actions, Action{Value: min, Bot: bot.Low})
		}
		if bot.HighKind == "bot" {
			actions = append(actions, Action{Value: max, Bot: bot.High})
		}
	}

	fmt.Println(bot.ID)
}

type Bot struct {
	ID                int
	LowKind, HighKind string
	Low, High         int
}

type Action struct {
	Value int
	Bot   int
}

func InputToBotsAndActions() (map[int]Bot, []Action) {
	var bots = make(map[int]Bot)
	var actions []Action

	for in.HasNext() {
		switch {
		case in.HasPrefix("bot"):
			var bot Bot
			in.Scanf("bot %d gives low to %s %d and high to %s %d", &bot.ID, &bot.LowKind, &bot.Low, &bot.HighKind, &bot.High)
			bots[bot.ID] = bot

		case in.HasPrefix("value"):
			var action Action
			in.Scanf("value %d goes to bot %d", &action.Value, &action.Bot)
			actions = append(actions, action)
		}
	}

	return bots, actions
}
