package factories

import (
	"mytest/bll/auth"
	"mytest/services"
)

// IGetAuthServiceInterface interface
type IGetAuthServiceInterface interface {
	GetFacebookAuthManager() bll.IFacebookAuthManager
	GetLocalAuthManager() bll.ILocalAuthManager
}

// GetAuthServiceInterface get auth service
type GetAuthServiceInterface struct {
	Auth *services.AuthServices
}

// GetFacebookAuthManager get facebookAuth manager
func (i *GetAuthServiceInterface) GetFacebookAuthManager() bll.IFacebookAuthManager {
	return i.Auth.FacebookAuthManager(GetFacebookAuthRepository())
}

// GetLocalAuthManager get localAuth manager
func (i *GetAuthServiceInterface) GetLocalAuthManager() bll.ILocalAuthManager {
	return i.Auth.LocalAuthManager(GetLocalAuthRepository())
}

// CreateAuthServiceInterface get auth service
func CreateAuthServiceInterface() IGetAuthServiceInterface {
	var x IGetAuthServiceInterface
	a := GetAuthServiceInterface{}
	t := services.AuthServices{}
	a.Auth = &t
	x = &a
	return x
}
