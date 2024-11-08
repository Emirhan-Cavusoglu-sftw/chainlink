package chainlink

import (
	"github.com/smartcontractkit/chainlink/v2/core/config/toml"
	"github.com/smartcontractkit/chainlink/v2/core/static"
)

type telemetryConfig struct {
	s toml.Telemetry
}

func (b *telemetryConfig) Enabled() bool { return *b.s.Enabled }

func (b *telemetryConfig) InsecureConnection() bool {
	if b.s.InsecureConnection == nil {
		return false
	}
	return *b.s.InsecureConnection
}

func (b *telemetryConfig) CACertFile() string {
	if b.s.CACertFile == nil {
		return ""
	}
	return *b.s.CACertFile
}

func (b *telemetryConfig) OtelExporterGRPCEndpoint() string {
	if b.s.Endpoint == nil {
		return ""
	}
	return *b.s.Endpoint
}

// ResourceAttributes returns the resource attributes set in the TOML config
// by the user, but first sets OTEL required attributes:
//
//	service.name
//	service.version
//
// These can be overridden by the TOML if the user so chooses
func (b *telemetryConfig) ResourceAttributes() map[string]string {
	defaults := map[string]string{
		"service.name":    "chainlink",
		"service.version": static.Version,
	}

	for k, v := range b.s.ResourceAttributes {
		defaults[k] = v
	}

	return defaults
}

func (b *telemetryConfig) TraceSampleRatio() float64 {
	if b.s.TraceSampleRatio == nil {
		return 0.0
	}
	return *b.s.TraceSampleRatio
}
