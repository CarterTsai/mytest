package dal

// IFacebookAuthRepository interface
type IFacebookAuthRepository interface {
	Login(uid string) bool
}
