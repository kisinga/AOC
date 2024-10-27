package types

type Challenge interface {
	Solve() error
}
