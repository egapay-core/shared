package utils

type ContextKey string

const (
	// ContextKeyClientPlatform is the key for the client platform in the context (MOBILE, WEB etc.)
	ContextKeyClientPlatform ContextKey = "x-client-platform"

	// ContextKeyLanguageID is the key for the language id in the context
	ContextKeyLanguageID ContextKey = "x-language-id"

	// ContextKeyCountryCode is the key for the country code in the context
	ContextKeyCountryCode ContextKey = "x-country-code"

	// ContextKeyUserAccessToken is the key for the user access token in the context
	ContextKeyUserAccessToken ContextKey = "x-ega-user-access-token"
)
