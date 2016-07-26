package mockstore

//go:generate go get github.com/golang/mock/mockgen
//go:generate sh -c "mockgen -package=mockstore github.com/citwild/wfe/store AccountsStore > accounts.go"
//go:generate sh -c "mockgen -package=mockstore github.com/citwild/wfe/store PasswordStore > password.go"
