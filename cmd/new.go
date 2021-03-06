package cmd

import (
	"context"
	"go/build"
	"os"
	"path/filepath"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/new"
)

// TODO: github.com/gobuffalo/envy
func getGopath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		name := ""
		root := ""

		if len(args) == 0 {
			name = filepath.Base(pwd)
			root = pwd
		} else {
			name = args[0]
			root = filepath.Join(pwd, name)
		}

		run := genny.WetRunner(context.Background())
		run.Root = root
		run.Logger = logger.New(logger.DebugLevel)
		run.Logger.Infof("Creating new project in %s", root)

		gopath := getGopath()
		rel, err := filepath.Rel(gopath+"/src", root)
		if err != nil {
			return err
		}

		gg, err := new.New(&new.Options{
			Name:    name,
			Package: rel,
		})
		if err != nil {
			return err
		}

		run.WithGroup(gg)
		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
