package cmd

import (
	"github.com/spf13/cobra"
)

var (
	storagePath     string
	descriptionTask string
	statusTask      string

	RootCmd = &cobra.Command{
		Use:   "taskmanager",
		Short: "taskmanager — пример CLI для работы с задачами",
	}
)

func init() {
	// глобальный флаг для всего CLI
	RootCmd.PersistentFlags().StringVar(&storagePath, "storage", "storage/tasks.json", "путь к файлу с задачами")
	RootCmd.PersistentFlags().StringVar(&descriptionTask, "d", "", "Описание задачи")
	RootCmd.PersistentFlags().StringVar(&statusTask, "s", "", "Статус задачи")
	// регистрируем subcommand list
	RootCmd.AddCommand(ListCmd, AddCmd)
}
