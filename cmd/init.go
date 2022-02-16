package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

//go:embed boilerplate/*
var boilerplateFiles embed.FS

// InitAnswers holds the answers entered by the user
type InitAnswers struct {
	Name        string
	Directory   string
	Author      string
	Description string
	URL         string
	AuthorURL   string
	Version     string
	Prefix      string
	Support     []string
}

var initQuestions = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "Enter a name for your plugin:",
			Default: "My Plugin",
		},
	},
	{
		Name: "directory",
		Prompt: &survey.Input{
			Message: "Enter the plugin's directory name:",
			Default: "my-plugin",
		},
	},
	{
		Name: "prefix",
		Prompt: &survey.Input{
			Message: "Enter a prefix for the plugin's global functions:",
			Default: "my-plugin",
		},
	},
	{
		Name: "description",
		Prompt: &survey.Input{
			Message: "Enter a description for your plugin:",
		},
	},
	{
		Name: "URL",
		Prompt: &survey.Input{
			Message: "Enter the plugin's website:",
		},
	},
	{
		Name: "author",
		Prompt: &survey.Input{
			Message: "Enter the plugin's author:",
		},
	},
	{
		Name: "authorURL",
		Prompt: &survey.Input{
			Message: "Enter the author's website:",
		},
	},
	{
		Name: "version",
		Prompt: &survey.Input{
			Message: "Enter the plugin's initial version:",
			Default: "1.0",
		},
	},
	{
		Name: "version",
		Prompt: &survey.Input{
			Message: "Enter the plugin's initial version:",
			Default: "1.0",
		},
	},
	{
		Name: "support",
		Prompt: &survey.MultiSelect{
			Message: "Choose what to include in your plugin:",
			Options: []string{
				"Custom page templates",
				"Extending the REST API",
				"Custom post types",
			},
		},
	},
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new WordPress plugin",
	Long: `
	You can use wpdev plugin init to initialize a new plugin.

	You will have the option to include many features, such as custom post types and extending the REST API.

	At the end, a new folder will be created containing all the required plugin files. 
	
	Please note that the plugin generated follows a structure that requires Composer to work correctly.`,
	Run: Run,
}

func Run(cmd *cobra.Command, args []string) {

	var answers InitAnswers

	err := survey.Ask(initQuestions, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Answers: %+v", answers)

	err = os.Mkdir(answers.Directory, 0755)

}

func init() {
	pluginCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
