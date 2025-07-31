package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Путь к JSON-файлу можно менять через флаг --storage
	storagePath string

	RootCmd = &cobra.Command{
		Use:   "taskmanager",
		Short: "taskmanager — пример CLI для работы с задачами",
	}
)

func init() {
	// глобальный флаг для всего CLI
	RootCmd.PersistentFlags().StringVar(&storagePath, "storage", "storage/tasks.json", "путь к файлу с задачами")
	// регистрируем subcommand list
	RootCmd.AddCommand(ListCmd)
}
