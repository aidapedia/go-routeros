package routeros

import "time"

type Option struct {
	address  string
	username string
	password string
	timeout  time.Duration
}

func NewOption(address, username, password string) *Option {
	return &Option{
		address:  address,
		username: username,
		password: password,
	}
}

// SetTimeout sets the timeout for the connection.
func (o *Option) SetTimeout(duration time.Duration) {
	o.timeout = duration
}

// GetTimeout returns the timeout for the connection.
func (o *Option) GetTimeout() time.Duration {
	return o.timeout
}

// SetAddress sets the address of the router.
func (o *Option) SetAddress(address string) {
	o.address = address
}

// GetAddress returns the address of the router.
func (o *Option) GetAddress() string {
	return o.address
}

// SetUsername sets the username for the connection.
func (o *Option) SetUsername(username string) {
	o.username = username
}

// GetUsername returns the username for the connection.
func (o *Option) GetUsername() string {
	return o.username
}

// SetPassword sets the password for the connection.
func (o *Option) SetPassword(password string) {
	o.password = password
}

// GetPassword returns the password for the connection.
func (o *Option) GetPassword() string {
	return o.password
}
