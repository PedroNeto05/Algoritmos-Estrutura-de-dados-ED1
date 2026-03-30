package list

type IList interface {
	Add(val int)
	AddOnIndex(val, index int) error
	Get(index int) (int, error)
	Set(val, index int) error
	RemoveOnIndex(index int) error
	Size() int
}
