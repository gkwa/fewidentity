package cmd

import (
	"fmt"
	
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of {{ cookiecutter.project_slug }}",
	Long:  `All software has versions. This is {{ cookiecutter.project_slug }}'s`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println("Version info", "version", buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
