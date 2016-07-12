package mock_api

//go:generate go get github.com/golang/mock/mockgen
//go:generate sh -c "mockgen github.com/citwild/wfe/api AccountsServer > accounts.go"
