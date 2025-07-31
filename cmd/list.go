package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"strconv"
	"taskmanager/storage"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "вывести все созданные задачи",
	RunE: func(cmd *cobra.Command, args []string) error {
		// считываем данные из json
		tasks, err := storage.LoadTasks(storagePath)
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("задач нет")
			return nil
		}

		// вывод результата в таблице
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.Header([]string{"ID", "Описание задачи", "Status"})
		for _, t := range tasks {
			err := table.Append([]string{
				strconv.Itoa(t.ID),
				t.Description,
				t.Status,
			})
			if err != nil {
				return err
			}
		}

		err = table.Render()
		if err != nil {
			return err
		}

		return nil
	},
}
