package cmd

import (
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of {{ cookiecutter.project_slug }}",
	Long:  `All software has versions. This is {{ cookiecutter.project_slug }}'s`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		buildInfo := version.GetBuildInfo()
		logger.Info("Version info", "version", buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
