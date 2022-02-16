/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"embed"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

//go:embed boilerplate/*
var boilerplateFiles embed.FS

var pluginInitQuestions = []promptui.Prompt{
	{
		Label:   "Plugin name",
		Default: "My Plugin",
	},
	{
		Label:   "Plugin directory name",
		Default: "my-plugin",
	},
	{
		Label: "Plugin author",
	},
	{
		Label:   "Plugin description",
		Default: "A plugin I made myself!",
	},
	{
		Label: "Plugin URL",
	},
	{
		Label: "Plugin author URL",
	},
	{
		Label:   "Plugin version",
		Default: "1.0",
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

	for _, prompt := range pluginInitQuestions {
		prompt.Run()
	}

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
