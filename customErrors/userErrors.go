package customErrors

type UserNotFoundError struct {
	username string
}
func (unfe *UserNotFoundError) Error() string {
	return "User with username '" + unfe.username + "' wasn't found."
}
