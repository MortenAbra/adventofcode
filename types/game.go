package types

type Games struct {
	Game map[string]Cubes
}

type Cubes struct {
	Red   int
	Green int
	Blue  int
}
