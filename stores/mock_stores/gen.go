package mock_stores

//go:generate go get github.com/golang/mock/mockgen
//go:generate sh -c "mockgen github.com/citwild/wfe/stores Accounts > accounts.go"
//go:generate sh -c "mockgen github.com/citwild/wfe/stores Password > password.go"
