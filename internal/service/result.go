package service

type ScanResult struct {
	FileCount int
	Manifest  string
}

type SyncResult struct {
	Downloaded int
	Deleted    int
	Unchanged  int
}
