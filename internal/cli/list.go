package cli

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"strings"
	"taskmanager/internal/storage"
)

var statuses string

var ListCmd = &cobra.Command{
	Use:     "list",
	Example: "list",
	Short:   "вывести все созданные задачи",
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

		filteredTasks := tasks
		if statuses != "" {
			statusList := strings.Split(statuses, ",")
			filteredTasks = filterTasks(tasks, statusList)
		}

		// вывод результата в таблицу
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.Header([]string{"ID", "Описание задачи", "Status"})
		for _, t := range filteredTasks {
			err := table.Append([]string{
				t.ID,
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

func filterTasks(tasks []storage.Task, statuses []string) []storage.Task {
	var filtered []storage.Task
	for _, task := range tasks {
		for _, status := range statuses {
			if strings.TrimSpace(status) == task.Status {
				filtered = append(filtered, task)
				break
			}
		}
	}
	return filtered
}

var AddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Добавить новую задачу",
	Example: "add --d \"Текст задачи\" --s \"in_progress\"",
	RunE: func(cmd *cobra.Command, args []string) error {
		// считываем данные из json
		tasks, err := storage.LoadTasks(storagePath)
		if err != nil {
			return err
		}
		// ToDo Надо бы создавать файл если его почему-то нет
		if len(tasks) == 0 {
			fmt.Println("задач нет")
			return nil
		}

		err = storage.CreateTask(descriptionTask, statusTask, tasks)
		if err != nil {
			return err
		}

		return nil
	},
}
