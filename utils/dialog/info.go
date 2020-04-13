package dialog

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func displayInfo(writer prompt.ConsoleWriter, text string) {

	writer.SetColor(prompt.Yellow, prompt.Black, true)
	writer.WriteStr(text)
}

// Info ...
func Info(format string, params ...interface{}) {
	writer := prompt.NewStandardOutputWriter()
	displayInfo(writer, fmt.Sprintf(format, params...))
	displayNewline(writer)
	writer.Flush()
}
