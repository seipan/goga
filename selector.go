package goga

type Selector interface {
	Select(Individuals) (Individuals, error)
}
