package server

type Option func(*Server)

func WithConfig(cfg *Config) Option {
	return func(s *Server) {
		s.cfg = cfg
	}
}
