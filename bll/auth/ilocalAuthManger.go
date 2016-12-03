package bll

// ILocalAuthManager interface
type ILocalAuthManager interface {
	Login(name string, password string) bool
}
