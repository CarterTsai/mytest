package bll

import "mytest/dal"

// LocalAuthManager interface
type LocalAuthManager struct {
	Rep dal.ILocalAuthRepository
}

// Login implement
func (f *LocalAuthManager) Login(name string, password string) bool {
	return f.Rep.Login(name, password)
}
