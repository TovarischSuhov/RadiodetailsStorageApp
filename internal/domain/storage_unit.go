package domain

type StorageUnit struct {
	ID     int64
	Height int64
	Width  int64
	Length int64
	Object Object
}
