package data

type Todo struct {
	ID        int    `json:"id"`
	Value     string `json:"value"`
	Completed bool   `json:"completed"`
}
