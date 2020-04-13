package commands

import (
	"flag"

	arguments "git-cleanup/utils/arguments"

	prompt "github.com/c-bata/go-prompt"
)

// Command ...
type Command interface {
	MinArgs() int
	Initialize(flagSet *flag.FlagSet)
	Execute(args []string)
	Help()
	Suggestions(in prompt.Document) []prompt.Suggest
}

func getCommandByName(commandName string) Command {
	return GetCommand([]string{"cleanup", commandName}...)
}

// GetCommand ...
func GetCommand(args ...string) Command {
	commandName := args[1]
	var instance Command = nil
	switch commandName {
	case LOCAL:
		instance = &localCommand{}
		break
	case PROTECT:
		instance = &protectCommand{}
		break
	case UNPROTECT:
		instance = &unprotectCommand{}
		break
	case RESET:
		instance = &resetCommand{}
	default:
		return nil
	}

	flagSet := flag.NewFlagSet(commandName, flag.ExitOnError)

	instance.Initialize(flagSet)

	flagArgs := arguments.OnlyFlags(args)

	flagSet.Parse(flagArgs)

	return instance
}
