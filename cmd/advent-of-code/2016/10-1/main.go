package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

const (
	TargetMin = 17
	TargetMax = 61
)

func main() {
	bots := InputToBots()
	values := make(map[int][]int)

	var bot Bot
	var actions = InputToInitializations()
	for len(actions) > 0 {
		var action Action
		action, actions = actions[0], actions[1:]

		bot = bots[action.Bot]
		if len(values[bot.ID]) == 0 {
			values[bot.ID] = append(values[bot.ID], action.Value)
			continue
		}

		min := lib.Min(values[bot.ID][0], action.Value)
		max := lib.Max(values[bot.ID][0], action.Value)
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

func InputToBots() map[int]Bot {
	bots := make(map[int]Bot)
	for _, line := range lib.InputToLines() {
		if !strings.HasPrefix(line, "bot") {
			continue
		}

		line = strings.ReplaceAll(line, "gives low to ", "")
		line = strings.ReplaceAll(line, "and high to ", "")

		var bot Bot
		fmt.Sscanf(line, "bot %d %s %d %s %d", &bot.ID, &bot.LowKind, &bot.Low, &bot.HighKind, &bot.High)

		bots[bot.ID] = bot
	}

	return bots
}

type Action struct {
	Value int
	Bot   int
}

func InputToInitializations() []Action {
	var initializations []Action
	for _, line := range lib.InputToLines() {
		if !strings.HasPrefix(line, "value") {
			continue
		}

		line = strings.ReplaceAll(line, "goes to bot ", "")

		var initialization Action
		fmt.Sscanf(line, "value %d %d", &initialization.Value, &initialization.Bot)
		initializations = append(initializations, initialization)
	}

	return initializations
}
