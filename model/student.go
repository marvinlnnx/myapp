package model

type Student struct {
	Name  string `redis:"name"`
	Score int    `redis:"score"`
	Age   int    `redis:"age"`
	Job   string `redis:"job"`
}
