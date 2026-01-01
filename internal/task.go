package internal

type Task struct {
	ID       int64    `json:"id"`
	JobName  string   `json:"jobName"`
	StartAt  string   `json:"startAt"`
	StopAt   string   `json:"stopAt"`
	IsClose  bool     `json:"isClose"`
	Priority string   `json:"priority"`
	Tags     []string `json:"tags"`
}
