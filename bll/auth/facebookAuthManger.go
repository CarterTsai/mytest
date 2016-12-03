package bll

import "mytest/dal"

// FacebookAuthManager interface
type FacebookAuthManager struct {
	Rep dal.IFacebookAuthRepository
}

// Login implement
func (f *FacebookAuthManager) Login(uid string) bool {
	return f.Rep.Login(uid)
}
