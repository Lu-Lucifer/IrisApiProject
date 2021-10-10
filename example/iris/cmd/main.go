/*
 Cmd is a tool for iris-admin.
 You can use it easy, like this: `iris-admin init`.
 It has some helpful commands for build your program. Example  init , migrate , rollback and so on.

 COMMAND
 Init is a command for initialize your program.
 Migrate is a command for migrate your database.
 Rollback is a command for rollback your migrations which you executed .
*/

package main

import (
	"github.com/snowlyg/iris-admin/modules/migration"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
	"github.com/spf13/cobra"
)

func main() {
	// cmdInit 初始化项目
	var cmdInit = &cobra.Command{
		Use:   "init",
		Short: "initialize program and set config",
		Long:  `initialize this program by set his config param before migrate migration and seed data`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return web_iris.InitConfig()
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "exec run migration",
		Long:  `exec run  migrations which are you writed in  migrate.go file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return migration.Gormigrate().Migrate()
		},
	}

	var cmdRollback = &cobra.Command{
		Use:   "rollback",
		Short: "exec rollback",
		Long:  `exec rollback migrate command which are you execed`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return migration.Gormigrate().RollbackLast()
		},
	}

	var rootCmd = &cobra.Command{Use: "iris-admin"}
	rootCmd.AddCommand(cmdInit, cmdRun, cmdRollback)
	rootCmd.Execute()
}
