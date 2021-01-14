// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/crowdstrike/gofalcon/falcon/client/cloud_connect_aws"
	"github.com/crowdstrike/gofalcon/falcon/client/cspm_registration"
	"github.com/crowdstrike/gofalcon/falcon/client/custom_ioa"
	"github.com/crowdstrike/gofalcon/falcon/client/d4c_registration"
	"github.com/crowdstrike/gofalcon/falcon/client/detects"
	"github.com/crowdstrike/gofalcon/falcon/client/device_control_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/event_streams"
	"github.com/crowdstrike/gofalcon/falcon/client/falconx_sandbox"
	"github.com/crowdstrike/gofalcon/falcon/client/firewall_management"
	"github.com/crowdstrike/gofalcon/falcon/client/firewall_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/host_group"
	"github.com/crowdstrike/gofalcon/falcon/client/hosts"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
	"github.com/crowdstrike/gofalcon/falcon/client/installation_tokens"
	"github.com/crowdstrike/gofalcon/falcon/client/intel"
	"github.com/crowdstrike/gofalcon/falcon/client/ioa_exclusions"
	"github.com/crowdstrike/gofalcon/falcon/client/iocs"
	"github.com/crowdstrike/gofalcon/falcon/client/malquery"
	"github.com/crowdstrike/gofalcon/falcon/client/ml_exclusions"
	"github.com/crowdstrike/gofalcon/falcon/client/oauth2"
	"github.com/crowdstrike/gofalcon/falcon/client/prevention_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/quick_scan"
	"github.com/crowdstrike/gofalcon/falcon/client/real_time_response"
	"github.com/crowdstrike/gofalcon/falcon/client/real_time_response_admin"
	"github.com/crowdstrike/gofalcon/falcon/client/sample_uploads"
	"github.com/crowdstrike/gofalcon/falcon/client/sensor_download"
	"github.com/crowdstrike/gofalcon/falcon/client/sensor_update_policies"
	"github.com/crowdstrike/gofalcon/falcon/client/sensor_visibility_exclusions"
	"github.com/crowdstrike/gofalcon/falcon/client/spotlight_vulnerabilities"
	"github.com/crowdstrike/gofalcon/falcon/client/user_management"
)

// Default crowd strike API specification HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.crowdstrike.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new crowd strike API specification HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CrowdStrikeAPISpecification {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new crowd strike API specification HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CrowdStrikeAPISpecification {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new crowd strike API specification client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CrowdStrikeAPISpecification {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CrowdStrikeAPISpecification)
	cli.Transport = transport
	cli.CloudConnectAws = cloud_connect_aws.New(transport, formats)
	cli.CspmRegistration = cspm_registration.New(transport, formats)
	cli.CustomIoa = custom_ioa.New(transport, formats)
	cli.D4cRegistration = d4c_registration.New(transport, formats)
	cli.Detects = detects.New(transport, formats)
	cli.DeviceControlPolicies = device_control_policies.New(transport, formats)
	cli.EventStreams = event_streams.New(transport, formats)
	cli.FalconxSandbox = falconx_sandbox.New(transport, formats)
	cli.FirewallManagement = firewall_management.New(transport, formats)
	cli.FirewallPolicies = firewall_policies.New(transport, formats)
	cli.HostGroup = host_group.New(transport, formats)
	cli.Hosts = hosts.New(transport, formats)
	cli.Incidents = incidents.New(transport, formats)
	cli.InstallationTokens = installation_tokens.New(transport, formats)
	cli.Intel = intel.New(transport, formats)
	cli.IoaExclusions = ioa_exclusions.New(transport, formats)
	cli.Iocs = iocs.New(transport, formats)
	cli.Malquery = malquery.New(transport, formats)
	cli.MlExclusions = ml_exclusions.New(transport, formats)
	cli.Oauth2 = oauth2.New(transport, formats)
	cli.PreventionPolicies = prevention_policies.New(transport, formats)
	cli.QuickScan = quick_scan.New(transport, formats)
	cli.RealTimeResponse = real_time_response.New(transport, formats)
	cli.RealTimeResponseAdmin = real_time_response_admin.New(transport, formats)
	cli.SampleUploads = sample_uploads.New(transport, formats)
	cli.SensorDownload = sensor_download.New(transport, formats)
	cli.SensorUpdatePolicies = sensor_update_policies.New(transport, formats)
	cli.SensorVisibilityExclusions = sensor_visibility_exclusions.New(transport, formats)
	cli.SpotlightVulnerabilities = spotlight_vulnerabilities.New(transport, formats)
	cli.UserManagement = user_management.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CrowdStrikeAPISpecification is a client for crowd strike API specification
type CrowdStrikeAPISpecification struct {
	CloudConnectAws cloud_connect_aws.ClientService

	CspmRegistration cspm_registration.ClientService

	CustomIoa custom_ioa.ClientService

	D4cRegistration d4c_registration.ClientService

	Detects detects.ClientService

	DeviceControlPolicies device_control_policies.ClientService

	EventStreams event_streams.ClientService

	FalconxSandbox falconx_sandbox.ClientService

	FirewallManagement firewall_management.ClientService

	FirewallPolicies firewall_policies.ClientService

	HostGroup host_group.ClientService

	Hosts hosts.ClientService

	Incidents incidents.ClientService

	InstallationTokens installation_tokens.ClientService

	Intel intel.ClientService

	IoaExclusions ioa_exclusions.ClientService

	Iocs iocs.ClientService

	Malquery malquery.ClientService

	MlExclusions ml_exclusions.ClientService

	Oauth2 oauth2.ClientService

	PreventionPolicies prevention_policies.ClientService

	QuickScan quick_scan.ClientService

	RealTimeResponse real_time_response.ClientService

	RealTimeResponseAdmin real_time_response_admin.ClientService

	SampleUploads sample_uploads.ClientService

	SensorDownload sensor_download.ClientService

	SensorUpdatePolicies sensor_update_policies.ClientService

	SensorVisibilityExclusions sensor_visibility_exclusions.ClientService

	SpotlightVulnerabilities spotlight_vulnerabilities.ClientService

	UserManagement user_management.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CrowdStrikeAPISpecification) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.CloudConnectAws.SetTransport(transport)
	c.CspmRegistration.SetTransport(transport)
	c.CustomIoa.SetTransport(transport)
	c.D4cRegistration.SetTransport(transport)
	c.Detects.SetTransport(transport)
	c.DeviceControlPolicies.SetTransport(transport)
	c.EventStreams.SetTransport(transport)
	c.FalconxSandbox.SetTransport(transport)
	c.FirewallManagement.SetTransport(transport)
	c.FirewallPolicies.SetTransport(transport)
	c.HostGroup.SetTransport(transport)
	c.Hosts.SetTransport(transport)
	c.Incidents.SetTransport(transport)
	c.InstallationTokens.SetTransport(transport)
	c.Intel.SetTransport(transport)
	c.IoaExclusions.SetTransport(transport)
	c.Iocs.SetTransport(transport)
	c.Malquery.SetTransport(transport)
	c.MlExclusions.SetTransport(transport)
	c.Oauth2.SetTransport(transport)
	c.PreventionPolicies.SetTransport(transport)
	c.QuickScan.SetTransport(transport)
	c.RealTimeResponse.SetTransport(transport)
	c.RealTimeResponseAdmin.SetTransport(transport)
	c.SampleUploads.SetTransport(transport)
	c.SensorDownload.SetTransport(transport)
	c.SensorUpdatePolicies.SetTransport(transport)
	c.SensorVisibilityExclusions.SetTransport(transport)
	c.SpotlightVulnerabilities.SetTransport(transport)
	c.UserManagement.SetTransport(transport)
}
