package model

type Todo struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Item      string `json:"item,omitempty" bson:"item,omitempty"`
	Completed bool   `json:"completed" bson:"completed"`
}

type TodoError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (todoError TodoError) Error() string {
	return todoError.Message
}
