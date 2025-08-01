// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
)

type InputAppscopeType string

const (
	InputAppscopeTypeAppscope InputAppscopeType = "appscope"
)

func (e InputAppscopeType) ToPointer() *InputAppscopeType {
	return &e
}
func (e *InputAppscopeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "appscope":
		*e = InputAppscopeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeType: %v", v)
	}
}

type InputAppscopeConnection struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputAppscopeConnection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputAppscopeConnection) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputAppscopeMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputAppscopeMode string

const (
	InputAppscopeModeSmart  InputAppscopeMode = "smart"
	InputAppscopeModeAlways InputAppscopeMode = "always"
)

func (e InputAppscopeMode) ToPointer() *InputAppscopeMode {
	return &e
}
func (e *InputAppscopeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputAppscopeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeMode: %v", v)
	}
}

// InputAppscopeCompression - Codec to use to compress the persisted data
type InputAppscopeCompression string

const (
	InputAppscopeCompressionNone InputAppscopeCompression = "none"
	InputAppscopeCompressionGzip InputAppscopeCompression = "gzip"
)

func (e InputAppscopeCompression) ToPointer() *InputAppscopeCompression {
	return &e
}
func (e *InputAppscopeCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputAppscopeCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeCompression: %v", v)
	}
}

type InputAppscopePq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputAppscopeMode `default:"always" json:"mode"`
	// The maximum number of events to hold in memory before writing the events to disk
	MaxBufferSize *float64 `default:"1000" json:"maxBufferSize"`
	// The number of events to send downstream before committing that Stream has read them
	CommitFrequency *float64 `default:"42" json:"commitFrequency"`
	// The maximum size to store in each queue file before closing and optionally compressing. Enter a numeral with units of KB, MB, etc.
	MaxFileSize *string `default:"1 MB" json:"maxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `default:"5GB" json:"maxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>
	Path *string `default:"\\$CRIBL_HOME/state/queues" json:"path"`
	// Codec to use to compress the persisted data
	Compress *InputAppscopeCompression `default:"none" json:"compress"`
}

func (i InputAppscopePq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputAppscopePq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputAppscopePq) GetMode() *InputAppscopeMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputAppscopePq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputAppscopePq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputAppscopePq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputAppscopePq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputAppscopePq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputAppscopePq) GetCompress() *InputAppscopeCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type InputAppscopeMetadatum struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputAppscopeMetadatum) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputAppscopeMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type Allow struct {
	// Specify the name of a process or family of processes.
	Procname string `json:"procname"`
	// Specify a string to substring-match against process command-line.
	Arg *string `json:"arg,omitempty"`
	// Choose a config to apply to processes that match the process name and/or argument.
	Config string `json:"config"`
}

func (o *Allow) GetProcname() string {
	if o == nil {
		return ""
	}
	return o.Procname
}

func (o *Allow) GetArg() *string {
	if o == nil {
		return nil
	}
	return o.Arg
}

func (o *Allow) GetConfig() string {
	if o == nil {
		return ""
	}
	return o.Config
}

type InputAppscopeFilter struct {
	// Specify processes that AppScope should be loaded into, and the config to use.
	Allow []Allow `json:"allow,omitempty"`
	// To override the UNIX domain socket or address/port specified in General Settings (while leaving Authentication settings as is), enter a URL.
	TransportURL *string `json:"transportURL,omitempty"`
}

func (o *InputAppscopeFilter) GetAllow() []Allow {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *InputAppscopeFilter) GetTransportURL() *string {
	if o == nil {
		return nil
	}
	return o.TransportURL
}

type InputAppscopeDataCompressionFormat string

const (
	InputAppscopeDataCompressionFormatNone InputAppscopeDataCompressionFormat = "none"
	InputAppscopeDataCompressionFormatGzip InputAppscopeDataCompressionFormat = "gzip"
)

func (e InputAppscopeDataCompressionFormat) ToPointer() *InputAppscopeDataCompressionFormat {
	return &e
}
func (e *InputAppscopeDataCompressionFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputAppscopeDataCompressionFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeDataCompressionFormat: %v", v)
	}
}

type InputAppscopePersistence struct {
	// Spool events and metrics on disk for Cribl Edge and Search
	Enable *bool `default:"false" json:"enable"`
	// Time span for each file bucket
	TimeWindow *string `default:"10m" json:"timeWindow"`
	// Maximum disk space allowed to be consumed (examples: 420MB, 4GB). When limit is reached, older data will be deleted.
	MaxDataSize *string `default:"1GB" json:"maxDataSize"`
	// Maximum amount of time to retain data (examples: 2h, 4d). When limit is reached, older data will be deleted.
	MaxDataTime *string                             `default:"24h" json:"maxDataTime"`
	Compress    *InputAppscopeDataCompressionFormat `default:"gzip" json:"compress"`
	// Path to use to write metrics. Defaults to $CRIBL_HOME/state/appscope
	DestPath *string `default:"\\$CRIBL_HOME/state/appscope" json:"destPath"`
}

func (i InputAppscopePersistence) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputAppscopePersistence) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputAppscopePersistence) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

func (o *InputAppscopePersistence) GetTimeWindow() *string {
	if o == nil {
		return nil
	}
	return o.TimeWindow
}

func (o *InputAppscopePersistence) GetMaxDataSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataSize
}

func (o *InputAppscopePersistence) GetMaxDataTime() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataTime
}

func (o *InputAppscopePersistence) GetCompress() *InputAppscopeDataCompressionFormat {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputAppscopePersistence) GetDestPath() *string {
	if o == nil {
		return nil
	}
	return o.DestPath
}

// InputAppscopeAuthenticationMethod - Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
type InputAppscopeAuthenticationMethod string

const (
	InputAppscopeAuthenticationMethodManual InputAppscopeAuthenticationMethod = "manual"
	InputAppscopeAuthenticationMethodSecret InputAppscopeAuthenticationMethod = "secret"
)

func (e InputAppscopeAuthenticationMethod) ToPointer() *InputAppscopeAuthenticationMethod {
	return &e
}
func (e *InputAppscopeAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "secret":
		*e = InputAppscopeAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeAuthenticationMethod: %v", v)
	}
}

type InputAppscopeMinimumTLSVersion string

const (
	InputAppscopeMinimumTLSVersionTlSv1  InputAppscopeMinimumTLSVersion = "TLSv1"
	InputAppscopeMinimumTLSVersionTlSv11 InputAppscopeMinimumTLSVersion = "TLSv1.1"
	InputAppscopeMinimumTLSVersionTlSv12 InputAppscopeMinimumTLSVersion = "TLSv1.2"
	InputAppscopeMinimumTLSVersionTlSv13 InputAppscopeMinimumTLSVersion = "TLSv1.3"
)

func (e InputAppscopeMinimumTLSVersion) ToPointer() *InputAppscopeMinimumTLSVersion {
	return &e
}
func (e *InputAppscopeMinimumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputAppscopeMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeMinimumTLSVersion: %v", v)
	}
}

type InputAppscopeMaximumTLSVersion string

const (
	InputAppscopeMaximumTLSVersionTlSv1  InputAppscopeMaximumTLSVersion = "TLSv1"
	InputAppscopeMaximumTLSVersionTlSv11 InputAppscopeMaximumTLSVersion = "TLSv1.1"
	InputAppscopeMaximumTLSVersionTlSv12 InputAppscopeMaximumTLSVersion = "TLSv1.2"
	InputAppscopeMaximumTLSVersionTlSv13 InputAppscopeMaximumTLSVersion = "TLSv1.3"
)

func (e InputAppscopeMaximumTLSVersion) ToPointer() *InputAppscopeMaximumTLSVersion {
	return &e
}
func (e *InputAppscopeMaximumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputAppscopeMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputAppscopeMaximumTLSVersion: %v", v)
	}
}

type InputAppscopeTLSSettingsServerSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// The name of the predefined certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Passphrase to use to decrypt private key
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert        *bool                           `default:"false" json:"requestCert"`
	RejectUnauthorized any                             `json:"rejectUnauthorized,omitempty"`
	CommonNameRegex    any                             `json:"commonNameRegex,omitempty"`
	MinVersion         *InputAppscopeMinimumTLSVersion `json:"minVersion,omitempty"`
	MaxVersion         *InputAppscopeMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputAppscopeTLSSettingsServerSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputAppscopeTLSSettingsServerSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputAppscopeTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputAppscopeTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputAppscopeTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputAppscopeTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputAppscopeTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputAppscopeTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputAppscopeTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

func (o *InputAppscopeTLSSettingsServerSide) GetRejectUnauthorized() any {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputAppscopeTLSSettingsServerSide) GetCommonNameRegex() any {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputAppscopeTLSSettingsServerSide) GetMinVersion() *InputAppscopeMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputAppscopeTLSSettingsServerSide) GetMaxVersion() *InputAppscopeMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type InputAppscope struct {
	// Unique ID for this input
	ID       string            `json:"id"`
	Type     InputAppscopeType `json:"type"`
	Disabled *bool             `default:"false" json:"disabled"`
	// Pipeline to process data from this Source before sending it through the Routes
	Pipeline *string `json:"pipeline,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `default:"true" json:"sendToRoutes"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Use a disk queue to minimize data loss when connected services block. See [Cribl Docs](https://docs.cribl.io/stream/persistent-queues) for PQ defaults (Cribl-managed Cloud Workers) and configuration options (on-prem and hybrid Workers).
	PqEnabled *bool `default:"false" json:"pqEnabled"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Direct connections to Destinations, and optionally via a Pipeline or a Pack
	Connections []InputAppscopeConnection `json:"connections,omitempty"`
	Pq          *InputAppscopePq          `json:"pq,omitempty"`
	// Regex matching IP addresses that are allowed to establish a connection
	IPWhitelistRegex *string `default:"/.*/" json:"ipWhitelistRegex"`
	// Maximum number of active connections allowed per Worker Process. Use 0 for unlimited.
	MaxActiveCxn *float64 `default:"1000" json:"maxActiveCxn"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. After this time, the connection will be closed. Leave at 0 for no inactive socket monitoring.
	SocketIdleTimeout *float64 `default:"0" json:"socketIdleTimeout"`
	// How long the server will wait after initiating a closure for a client to close its end of the connection. If the client doesn't close the connection within this time, the server will forcefully terminate the socket to prevent resource leaks and ensure efficient connection cleanup and system stability. Leave at 0 for no inactive socket monitoring.
	SocketEndingMaxWait *float64 `default:"30" json:"socketEndingMaxWait"`
	// The maximum duration a socket can remain open, even if active. This helps manage resources and mitigate issues caused by TCP pinning. Set to 0 to disable.
	SocketMaxLifespan *float64 `default:"0" json:"socketMaxLifespan"`
	// Enable if the connection is proxied by a device that supports proxy protocol v1 or v2
	EnableProxyHeader *bool `default:"false" json:"enableProxyHeader"`
	// Fields to add to events from this input
	Metadata []InputAppscopeMetadatum `json:"metadata,omitempty"`
	// A list of event-breaking rulesets that will be applied, in order, to the input data stream
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// How long (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel before flushing the data stream out, as is, to the Pipelines
	StaleChannelFlushMs *float64 `default:"10000" json:"staleChannelFlushMs"`
	// Toggle to Yes to specify a file-backed UNIX domain socket connection, instead of a network host and port.
	EnableUnixPath *bool                     `default:"false" json:"enableUnixPath"`
	Filter         *InputAppscopeFilter      `json:"filter,omitempty"`
	Persistence    *InputAppscopePersistence `json:"persistence,omitempty"`
	// Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
	AuthType    *InputAppscopeAuthenticationMethod `default:"manual" json:"authType"`
	Description *string                            `json:"description,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host *string `json:"host,omitempty"`
	// Port to listen on
	Port *float64                            `json:"port,omitempty"`
	TLS  *InputAppscopeTLSSettingsServerSide `json:"tls,omitempty"`
	// Path to the UNIX domain socket to listen on.
	UnixSocketPath *string `default:"\\$CRIBL_HOME/state/appscope.sock" json:"unixSocketPath"`
	// Permissions to set for socket e.g., 777. If empty, falls back to the runtime user's default permissions.
	UnixSocketPerms *string `json:"unixSocketPerms,omitempty"`
	// Shared secret to be provided by any client (in authToken header field). If empty, unauthorized access is permitted.
	AuthToken *string `default:"" json:"authToken"`
	// Select or create a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
}

func (i InputAppscope) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputAppscope) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputAppscope) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputAppscope) GetType() InputAppscopeType {
	if o == nil {
		return InputAppscopeType("")
	}
	return o.Type
}

func (o *InputAppscope) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputAppscope) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputAppscope) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputAppscope) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputAppscope) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputAppscope) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputAppscope) GetConnections() []InputAppscopeConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputAppscope) GetPq() *InputAppscopePq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputAppscope) GetIPWhitelistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPWhitelistRegex
}

func (o *InputAppscope) GetMaxActiveCxn() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveCxn
}

func (o *InputAppscope) GetSocketIdleTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketIdleTimeout
}

func (o *InputAppscope) GetSocketEndingMaxWait() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketEndingMaxWait
}

func (o *InputAppscope) GetSocketMaxLifespan() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketMaxLifespan
}

func (o *InputAppscope) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputAppscope) GetMetadata() []InputAppscopeMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputAppscope) GetBreakerRulesets() []string {
	if o == nil {
		return nil
	}
	return o.BreakerRulesets
}

func (o *InputAppscope) GetStaleChannelFlushMs() *float64 {
	if o == nil {
		return nil
	}
	return o.StaleChannelFlushMs
}

func (o *InputAppscope) GetEnableUnixPath() *bool {
	if o == nil {
		return nil
	}
	return o.EnableUnixPath
}

func (o *InputAppscope) GetFilter() *InputAppscopeFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *InputAppscope) GetPersistence() *InputAppscopePersistence {
	if o == nil {
		return nil
	}
	return o.Persistence
}

func (o *InputAppscope) GetAuthType() *InputAppscopeAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *InputAppscope) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputAppscope) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputAppscope) GetPort() *float64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *InputAppscope) GetTLS() *InputAppscopeTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputAppscope) GetUnixSocketPath() *string {
	if o == nil {
		return nil
	}
	return o.UnixSocketPath
}

func (o *InputAppscope) GetUnixSocketPerms() *string {
	if o == nil {
		return nil
	}
	return o.UnixSocketPerms
}

func (o *InputAppscope) GetAuthToken() *string {
	if o == nil {
		return nil
	}
	return o.AuthToken
}

func (o *InputAppscope) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}
