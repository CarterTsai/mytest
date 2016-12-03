package dal

// FacebookAuthRepository Implement
type FacebookAuthRepository struct{}

// Login Implement
func (f *FacebookAuthRepository) Login(uid string) bool {
	return true
}
