package bll

// IFacebookAuthManager interface
type IFacebookAuthManager interface {
	Login(uid string) bool
}
