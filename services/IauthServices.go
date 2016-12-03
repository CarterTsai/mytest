package services

import "mytest/bll/auth"

// IAuthServices interface
type IAuthServices interface {
	FacebookAuthManager() bll.IFacebookAuthManager
	LocalAuthManager() bll.ILocalAuthManager
}
