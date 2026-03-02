package service

import "context"

// Health contains the liveness payload served to clients.
type Health struct {
	Status string
}

// BuildInfo contains build metadata injected at build time.
type BuildInfo struct {
	GitSHA    string
	BuildTime string
}

// Service provides the MVP status information.
type Service struct {
	gitSHA    string
	buildTime string
}

// New constructs the status service.
func New(gitSHA string, buildTime string) *Service {
	return &Service{
		gitSHA:    gitSHA,
		buildTime: buildTime,
	}
}

// Health returns the basic liveness response.
func (s *Service) Health(_ context.Context) Health {
	return Health{Status: "ok"}
}

// BuildInfo returns the backend build metadata.
func (s *Service) BuildInfo(_ context.Context) BuildInfo {
	return BuildInfo{
		GitSHA:    s.gitSHA,
		BuildTime: s.buildTime,
	}
}
