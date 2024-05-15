package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Config struct {
	ConnType     string `yaml:"conn_type"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ReadTimeout  int64  `yaml:"read_timeout"`
	WriteTimeout int64  `yaml:"write_timeout"`
	IdleTimeout  int64  `yaml:"idle_timeout"`
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func New(option ...Option) *Server {
	srv := new(Server)

	for _, opt := range option {
		opt(srv)
	}

	return srv
}

func (s *Server) Run(ctx context.Context) error {
	ln, err := net.Listen(s.cfg.ConnType, fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))

	if err != nil {
		return err
	}

	s.srv = &http.Server{
		Addr:         ln.Addr().String(),
		Handler:      http.NewServeMux(), // TODO create another mx
		ReadTimeout:  s.setReadTimeout(),
		WriteTimeout: s.setWriteTimeout(),
		IdleTimeout:  s.setIdleTimeout(),
	}

	return s.srv.Serve(ln)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) setReadTimeout() time.Duration {
	return time.Duration(s.cfg.ReadTimeout) * time.Second
}

func (s *Server) setWriteTimeout() time.Duration {
	return time.Duration(s.cfg.WriteTimeout) * time.Second
}

func (s *Server) setIdleTimeout() time.Duration {
	return time.Duration(s.cfg.IdleTimeout) * time.Second
}
