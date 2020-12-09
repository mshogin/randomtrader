package storage

type testGCEClient struct{}

// GetGCETestClient ...
var GetGCETestClient = func() (Storage, error) {
	return &testGCEClient{}, nil
}

// SaveObject ...
func (m *testGCEClient) SaveObject(prefix, fpath string) error {
	return nil
}
