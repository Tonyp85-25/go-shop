package stubs

type StubIdProvider struct {
}

func (s StubIdProvider) GetId() (string, error) {
	return "123", nil
}
