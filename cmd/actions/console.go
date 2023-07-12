package actions

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/urfave/cli/v2"
)

var suggestions = make([]prompt.Suggest, 0, 1000)

func completer(d prompt.Document) []prompt.Suggest {
	// s := []prompt.Suggest{
	// 	{Text: "users", Description: "Store the username and age"},
	// 	{Text: "articles", Description: "Store the article text posted by user"},
	// 	{Text: "comments", Description: "Store the text commented to articles"},
	// }
	return prompt.FilterContains(suggestions, d.GetWordBeforeCursor(), true)
	// return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func ConsoleAction(c *cli.Context) error {
	c2 := ctx.GetContext(c)
	for folder := range c2.Config.GetFolders(true) {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        folder.Path,
			Description: folder.Command,
		})
	}

	fmt.Println("Please select a folder.")
	t := prompt.Input("> ", completer,
		prompt.OptionCompletionOnDown(),
		prompt.OptionTitle("go-dev"))

	if len(t) == 0 {
		return nil
	}
	for folder := range c2.Config.GetFolders(true) {
		if folder.Path == t {
			justCD := c.Bool(consts.FlagJustCD)
			output := c.String(consts.FlagOutput)
			utils.SetOutput(output)

			openFolder := c.Bool(consts.FlagOpen)
			command := parseCommand(folder, openFolder, justCD)
			utils.WriteOutput(command)
			break
		}
	}

	return nil
}
