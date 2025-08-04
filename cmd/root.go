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
	RootCmd.PersistentFlags().StringVar(&storagePath, "storage", "storage/tasks.json", "путь к файлу с задачами, дефолтно используется (storage/tasks.json)")
	RootCmd.PersistentFlags().StringVar(&descriptionTask, "d", "", "Описание задачи")
	RootCmd.PersistentFlags().StringVar(&statusTask, "s", "", "Статус задачи, доступные варианты (new,in_progress,done)")
	ListCmd.Flags().StringVarP(&statuses, "status", "f", "", "фильтр по статусам (new, in_progress, done, разделенные запятой)")
	RootCmd.AddCommand(ListCmd, AddCmd)
}
