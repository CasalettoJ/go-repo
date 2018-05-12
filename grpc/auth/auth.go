package auth

import (
	"golang.org/x/net/context"
)

// Auth contains data for authenticating a user
// Implements https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials interface
type Auth struct {
	User string
}

// GetRequestMetadata returns a mapping of the Auth struct
func (a *Auth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"user": a.User,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Auth) RequireTransportSecurity() bool {
	return true
}
