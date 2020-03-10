package todo

// Todo : struct for Todo
type Todo struct {
	Task     string `json:"task" bson:"task"`
	Duedate  string `json:"duedate" bson:"duedate"`
	Labels   string `json:"labels" bson:"labels"`
	Comments string `json:"comments" bson:"comments"`
}