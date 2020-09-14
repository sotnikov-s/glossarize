package security

import "context"

// identityContextKey is the context key of the identity value.
const identityContextKey = "identity"

// setIdentityInContext returns the ctx copy with the identity key
// populated with the passed identity.
func setIdentityInContext(ctx context.Context, ident *Identity) context.Context {
	return context.WithValue(ctx, identityContextKey, ident)
}
