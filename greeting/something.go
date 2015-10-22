// Package greeting stores greeting settings in context.
package greeting

import (
    "net/http"
    "golang.org/x/net/context"
    "golang.org/x/text/language"
)

// For more information about context and why we're doing this,
// see https://blog.golang.org/context
type ctxkey int

var key ctxkey = 0

var greetings = map[language.Tag]string{
    language.AmericanEnglish: "Yo",
    language.Japanese:        "こんにちは",
}

// Guess is kami middleware that examines Accept-Language and sets
// the greeting to a better one if possible.
func Guess(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
    if tag, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language")); err == nil {
        for _, t := range tag {
            if g, ok := greetings[t]; ok {
                ctx = WithContext(ctx, g)
                return ctx
            }
        }
    }
    return ctx
}

// WithContext returns a new context with the given greeting.
func WithContext(ctx context.Context, greeting string) context.Context {
    return context.WithValue(ctx, key, greeting)
}

// FromContext retrieves the greeting from this context,
// or returns an empty string if missing.
func FromContext(ctx context.Context) string {
    hello, _ := ctx.Value(key).(string)
    return hello
}
