package dialog

import (
	prompt "github.com/c-bata/go-prompt"
)

const (
	PromptHeader = "> "
)

func Input(suggester prompt.Completer) string {
	return prompt.Input(PromptHeader, suggester, prompt.OptionInputTextColor(prompt.Green))
}

func defaultSuggester(d prompt.Document, s []prompt.Suggest) []prompt.Suggest {
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func defaultCompleter(s []prompt.Suggest) prompt.Completer {
	return func(d prompt.Document) []prompt.Suggest {
		return defaultSuggester(d, s)
	}
}
