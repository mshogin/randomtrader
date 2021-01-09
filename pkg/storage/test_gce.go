package storage

type testGCEClient struct{}

// GetGCETestClient ...
func GetGCETestClient() Storage {
	return &testGCEClient{}
}

// SaveObject ...
func (m *testGCEClient) SaveObject(prefix, fpath string) error {
	return nil
}

func (m *testGCEClient) DownloadObjects(prefix, fpath string) error {
	return nil
}
