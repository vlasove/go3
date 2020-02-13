package main

type Authable interface {
	Encode(data interface{}) (string, error)
	Decode(token string) (interface{}, error)
}

type TokenService struct {
	repo Repository
}

func (ts *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

func (ts *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}
