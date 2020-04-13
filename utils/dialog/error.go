package dialog

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func displayError(writer prompt.ConsoleWriter, text string) {

	writer.SetColor(prompt.Red, prompt.Black, true)
	writer.WriteStr(text)
}

// Error ...
func Error(format string, params ...interface{}) {
	writer := prompt.NewStandardOutputWriter()
	displayError(writer, fmt.Sprintf(format, params...))
	displayNewline(writer)
	writer.Flush()
}
