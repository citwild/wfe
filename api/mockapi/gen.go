package mockapi

//go:generate go get github.com/golang/mock/mockgen
//go:generate sh -c "mockgen -package=mockapi github.com/citwild/wfe/api AccountsServer > accounts.go"
