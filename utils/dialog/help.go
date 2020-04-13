package dialog

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func displayHelp(writer prompt.ConsoleWriter, text string) {

	writer.SetColor(prompt.LightGray, prompt.Black, true)
	writer.WriteStr(text)
}

// Help ...
func Help(format string, params ...interface{}) {
	writer := prompt.NewStandardOutputWriter()
	displayHelp(writer, fmt.Sprintf(format, params...))
	displayNewline(writer)
	writer.Flush()
}
