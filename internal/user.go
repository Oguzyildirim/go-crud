// Package internal defines the types used to create Users and their corresponding attributes.
package internal

// User is an activity that needs to be completed within a period of time.
type User struct {
	ID       string
	Name     string
	Lastname string
	Username string
	Country  string
}

// Validate ...
func (u User) Validate() error {
	if u.Name == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "Name is required")
	}
	if u.Lastname == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "Lastname is required")
	}
	if u.Username == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "Username is required")
	}
	if u.Country == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "Country is required")
	}
	return nil
}
