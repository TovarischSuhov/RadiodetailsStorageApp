package domain

type Storage struct {
	Height       int64
	Width        int64
	Name         string
	StorageUnits [][]StorageUnit
}
