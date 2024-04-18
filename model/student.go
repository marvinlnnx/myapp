package model

type Student struct {
	ID    int    `redis:"-"`
	Name  string `redis:"name"`
	Score int    `redis:"score"`
	Age   int    `redis:"age"`
	Job   string `redis:"job"`
}
