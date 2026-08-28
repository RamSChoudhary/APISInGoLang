package types

type Student struct {
	id    int
	name  string `validate:"required"`
	email string `validate:"required"`
	age   int    `validate:"required"`
}
