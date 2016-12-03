package dal

// LocalAuthRepository Implement
type LocalAuthRepository struct{}

// Login Implement
func (f *LocalAuthRepository) Login(name string, password string) bool {
	return true
}
