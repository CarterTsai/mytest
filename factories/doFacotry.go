package factories

import "mytest/dal"

// GetFacebookAuthRepository get auth repository
func GetFacebookAuthRepository() dal.IFacebookAuthRepository {
	var i dal.IFacebookAuthRepository
	d := dal.FacebookAuthRepository{}
	i = &d
	return i
}

// GetLocalAuthRepository get auth repository
func GetLocalAuthRepository() dal.ILocalAuthRepository {
	var i dal.ILocalAuthRepository
	d := dal.LocalAuthRepository{}
	i = &d
	return i
}
