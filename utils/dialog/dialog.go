package dialog

import (
	prompt "github.com/c-bata/go-prompt"
)

func displayQuestion(writer prompt.ConsoleWriter, text string) {

	writer.SetColor(prompt.Yellow, prompt.Black, true)
	writer.WriteStr(text)
}

func displayNewline(writer prompt.ConsoleWriter) {
	writer.WriteStr("\n")
}
