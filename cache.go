package gdl

type Cache interface {
	Get(s string) (interface{}, bool)
}
