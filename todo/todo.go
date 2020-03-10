package todo

type Todo struct {
	//	TID			primitive.ObjectID 	`json:"_id" 			bson:"_id"`
	Task     string `json:"task" 			bson:"task"`
	Duedate  string `json:"duedate" 		bson:"duedate"`
	Labels   string `json:"labels"			bson:"labels"`
	Comments string `json:"comments"  		bson:"comments"`
}