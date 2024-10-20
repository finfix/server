package sql

type scanner interface {
	SliceScan() ([]any, error)
	MapScan(dest map[string]any) error
	StructScan(dest any) error
}

type closer interface {
	Close() error
}
