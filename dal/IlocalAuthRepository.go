package dal

// ILocalAuthRepository interface
type ILocalAuthRepository interface {
	Login(name string, password string) bool
}
