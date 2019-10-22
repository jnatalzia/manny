package types

type Location struct {
	ID        string
	Neighbors []string
}

type Map map[string]Location
