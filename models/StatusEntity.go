package models

const (
	ToDo       string = "To Do"
	Reopened   string = "Reopened"
	InProgress string = "In Progress"
	CodeReview string = "Code Review"
	Approved   string = "Approved"
	Done       string = "Done"
	Feedback   string = "Feedback"
	Rejected   string = "Rejected"
	Blocked    string = "Blocked"

	ToDoID       string = "10006"
	ReopenedID   string = "4"
	InProgressID string = "3"
	CodeReviewID string = "10002"
	ApprovedID   string = "10005"
	DoneID       string = "10007"
	FeedbackID   string = "10003"
	RejectedID   string = "10001"
	BlockedID    string = "5"
)

type StatusEntity struct {
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

// IsToDo означает, является ли статус задачи "новым", то есть может ли человек взять ее в работу.
func (status StatusEntity) IsToDo() bool {
	return status.ID == ToDoID || status.ID == ReopenedID
}
