package server

import (
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	khttp "github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/phpgao/durl_backend/api/url_shorter/v1"
	"github.com/phpgao/durl_backend/internal/conf"
	"github.com/phpgao/durl_backend/internal/service"
	"github.com/phpgao/durl_backend/internal/util"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, s *service.UrlShortenerService, logger log.Logger) *khttp.Server {
	opts := getServerOptions(c)
	opts = append(opts, khttp.Middleware(
		recovery.Recovery(),
	))
	opts = append(opts, khttp.Filter(urlFilter(s)))

	srv := khttp.NewServer(opts...)
	v1.RegisterUrlShortenerServiceHTTPServer(srv, s)
	return srv
}

func getServerOptions(c *conf.Server) []khttp.ServerOption {
	var opts []khttp.ServerOption
	if c.Http.Network != "" {
		opts = append(opts, khttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, khttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, khttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	return opts
}

func urlFilter(s *service.UrlShortenerService) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/go/") {
				subPath := strings.TrimPrefix(r.URL.Path, "/go/")
				if len(subPath) > 0 && util.IsValidBase62String(subPath) {
					handleRedirect(w, r, s, subPath)
					return
				}
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request, s *service.UrlShortenerService, short string) {
	origin, err := s.GetRedirectURL(r.Context(), &v1.JumpRequest{Url: short})
	if err != nil {
		s.Logger.Errorf("get redirect url error: %v", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
	go s.ShortUrlUseCase.UpdateVisitCount(r.Context(), origin.Id)
	http.Redirect(w, r, origin.Url, http.StatusMovedPermanently)
}
