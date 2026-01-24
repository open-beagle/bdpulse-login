package beagle

import (
	"net/http"

	"github.com/open-beagle/go-login/login"
	"github.com/open-beagle/go-login/login/internal/oauth2"
	"github.com/open-beagle/go-login/login/logger"
)

var _ login.Middleware = (*Config)(nil)

// Config configures the GitLab auth provider.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string
	Scope        []string
	Client       *http.Client
	Logger       logger.Logger
	Dumper       logger.Dumper
}

// Handler returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   c.Server + "/awecloud/dex/oauth/token",
		AuthorizationURL: c.Server + "/awecloud/dex/oauth/authorize",
		Scope:            c.Scope,
	})

}
