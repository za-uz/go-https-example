package types

import (
	"time"
)

type Entity struct {
	Password string
	Balance int
}

type Transaction struct {
	From string
	To string
	Amount int
}

type Authentication struct {
	Transaction
	Time time.Time
	Digest string
}

type Information struct {
	Name string
	Balance int
}

