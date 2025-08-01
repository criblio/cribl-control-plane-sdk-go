// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type InputPrometheusRwType string

const (
	InputPrometheusRwTypePrometheusRw InputPrometheusRwType = "prometheus_rw"
)

func (e InputPrometheusRwType) ToPointer() *InputPrometheusRwType {
	return &e
}
func (e *InputPrometheusRwType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prometheus_rw":
		*e = InputPrometheusRwType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRwType: %v", v)
	}
}

type InputPrometheusRwConnection struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputPrometheusRwConnection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputPrometheusRwConnection) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputPrometheusRwMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputPrometheusRwMode string

const (
	InputPrometheusRwModeSmart  InputPrometheusRwMode = "smart"
	InputPrometheusRwModeAlways InputPrometheusRwMode = "always"
)

func (e InputPrometheusRwMode) ToPointer() *InputPrometheusRwMode {
	return &e
}
func (e *InputPrometheusRwMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputPrometheusRwMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRwMode: %v", v)
	}
}

// InputPrometheusRwCompression - Codec to use to compress the persisted data
type InputPrometheusRwCompression string

const (
	InputPrometheusRwCompressionNone InputPrometheusRwCompression = "none"
	InputPrometheusRwCompressionGzip InputPrometheusRwCompression = "gzip"
)

func (e InputPrometheusRwCompression) ToPointer() *InputPrometheusRwCompression {
	return &e
}
func (e *InputPrometheusRwCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputPrometheusRwCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRwCompression: %v", v)
	}
}

type InputPrometheusRwPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputPrometheusRwMode `default:"always" json:"mode"`
	// The maximum number of events to hold in memory before writing the events to disk
	MaxBufferSize *float64 `default:"1000" json:"maxBufferSize"`
	// The number of events to send downstream before committing that Stream has read them
	CommitFrequency *float64 `default:"42" json:"commitFrequency"`
	// The maximum size to store in each queue file before closing and optionally compressing. Enter a numeral with units of KB, MB, etc.
	MaxFileSize *string `default:"1 MB" json:"maxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `default:"5GB" json:"maxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>
	Path *string `default:"$CRIBL_HOME/state/queues" json:"path"`
	// Codec to use to compress the persisted data
	Compress *InputPrometheusRwCompression `default:"none" json:"compress"`
}

func (i InputPrometheusRwPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputPrometheusRwPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputPrometheusRwPq) GetMode() *InputPrometheusRwMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputPrometheusRwPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputPrometheusRwPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputPrometheusRwPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputPrometheusRwPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputPrometheusRwPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputPrometheusRwPq) GetCompress() *InputPrometheusRwCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type InputPrometheusRwMinimumTLSVersion string

const (
	InputPrometheusRwMinimumTLSVersionTlSv1  InputPrometheusRwMinimumTLSVersion = "TLSv1"
	InputPrometheusRwMinimumTLSVersionTlSv11 InputPrometheusRwMinimumTLSVersion = "TLSv1.1"
	InputPrometheusRwMinimumTLSVersionTlSv12 InputPrometheusRwMinimumTLSVersion = "TLSv1.2"
	InputPrometheusRwMinimumTLSVersionTlSv13 InputPrometheusRwMinimumTLSVersion = "TLSv1.3"
)

func (e InputPrometheusRwMinimumTLSVersion) ToPointer() *InputPrometheusRwMinimumTLSVersion {
	return &e
}
func (e *InputPrometheusRwMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputPrometheusRwMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRwMinimumTLSVersion: %v", v)
	}
}

type InputPrometheusRwMaximumTLSVersion string

const (
	InputPrometheusRwMaximumTLSVersionTlSv1  InputPrometheusRwMaximumTLSVersion = "TLSv1"
	InputPrometheusRwMaximumTLSVersionTlSv11 InputPrometheusRwMaximumTLSVersion = "TLSv1.1"
	InputPrometheusRwMaximumTLSVersionTlSv12 InputPrometheusRwMaximumTLSVersion = "TLSv1.2"
	InputPrometheusRwMaximumTLSVersionTlSv13 InputPrometheusRwMaximumTLSVersion = "TLSv1.3"
)

func (e InputPrometheusRwMaximumTLSVersion) ToPointer() *InputPrometheusRwMaximumTLSVersion {
	return &e
}
func (e *InputPrometheusRwMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputPrometheusRwMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRwMaximumTLSVersion: %v", v)
	}
}

type InputPrometheusRwTLSSettingsServerSide struct {
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
	RequestCert        *bool                               `default:"false" json:"requestCert"`
	RejectUnauthorized any                                 `json:"rejectUnauthorized,omitempty"`
	CommonNameRegex    any                                 `json:"commonNameRegex,omitempty"`
	MinVersion         *InputPrometheusRwMinimumTLSVersion `json:"minVersion,omitempty"`
	MaxVersion         *InputPrometheusRwMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputPrometheusRwTLSSettingsServerSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputPrometheusRwTLSSettingsServerSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetRejectUnauthorized() any {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetCommonNameRegex() any {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetMinVersion() *InputPrometheusRwMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputPrometheusRwTLSSettingsServerSide) GetMaxVersion() *InputPrometheusRwMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

// InputPrometheusRwAuthenticationType - Remote Write authentication type
type InputPrometheusRwAuthenticationType string

const (
	InputPrometheusRwAuthenticationTypeNone              InputPrometheusRwAuthenticationType = "none"
	InputPrometheusRwAuthenticationTypeBasic             InputPrometheusRwAuthenticationType = "basic"
	InputPrometheusRwAuthenticationTypeCredentialsSecret InputPrometheusRwAuthenticationType = "credentialsSecret"
	InputPrometheusRwAuthenticationTypeToken             InputPrometheusRwAuthenticationType = "token"
	InputPrometheusRwAuthenticationTypeTextSecret        InputPrometheusRwAuthenticationType = "textSecret"
	InputPrometheusRwAuthenticationTypeOauth             InputPrometheusRwAuthenticationType = "oauth"
)

func (e InputPrometheusRwAuthenticationType) ToPointer() *InputPrometheusRwAuthenticationType {
	return &e
}
func (e *InputPrometheusRwAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "basic":
		fallthrough
	case "credentialsSecret":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		fallthrough
	case "oauth":
		*e = InputPrometheusRwAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRwAuthenticationType: %v", v)
	}
}

type InputPrometheusRwMetadatum struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputPrometheusRwMetadatum) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputPrometheusRwMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputPrometheusRwOauthParam struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

func (o *InputPrometheusRwOauthParam) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputPrometheusRwOauthParam) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputPrometheusRwOauthHeader struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

func (o *InputPrometheusRwOauthHeader) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputPrometheusRwOauthHeader) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputPrometheusRw struct {
	// Unique ID for this input
	ID       *string                `json:"id,omitempty"`
	Type     *InputPrometheusRwType `json:"type,omitempty"`
	Disabled *bool                  `default:"false" json:"disabled"`
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
	Connections []InputPrometheusRwConnection `json:"connections,omitempty"`
	Pq          *InputPrometheusRwPq          `json:"pq,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host *string `default:"0.0.0.0" json:"host"`
	// Port to listen on
	Port float64                                 `json:"port"`
	TLS  *InputPrometheusRwTLSSettingsServerSide `json:"tls,omitempty"`
	// Maximum number of active requests allowed per Worker Process. Set to 0 for unlimited. Caution: Increasing the limit above the default value, or setting it to unlimited, may degrade performance and reduce throughput.
	MaxActiveReq *float64 `default:"256" json:"maxActiveReq"`
	// Maximum number of requests per socket before @{product} instructs the client to close the connection. Default is 0 (unlimited).
	MaxRequestsPerSocket *int64 `default:"0" json:"maxRequestsPerSocket"`
	// Extract the client IP and port from PROXY protocol v1/v2. When enabled, the X-Forwarded-For header is ignored. Disable to use the X-Forwarded-For header for client IP extraction.
	EnableProxyHeader *bool `default:"false" json:"enableProxyHeader"`
	// Add request headers to events, in the __headers field
	CaptureHeaders *bool `default:"false" json:"captureHeaders"`
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *float64 `default:"100" json:"activityLogSampleRate"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *float64 `default:"0" json:"requestTimeout"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *float64 `default:"0" json:"socketTimeout"`
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 second, maximum 600 seconds (10 minutes).
	KeepAliveTimeout *float64 `default:"5" json:"keepAliveTimeout"`
	// Expose the /cribl_health endpoint, which returns 200 OK when this Source is healthy
	EnableHealthCheck *bool `default:"false" json:"enableHealthCheck"`
	// Messages from matched IP addresses will be processed, unless also matched by the denylist
	IPAllowlistRegex *string `default:"/.*/" json:"ipAllowlistRegex"`
	// Messages from matched IP addresses will be ignored. This takes precedence over the allowlist.
	IPDenylistRegex *string `default:"/^\\$/" json:"ipDenylistRegex"`
	// Absolute path on which to listen for Prometheus requests. Defaults to /write, which will expand as: http://<your‑upstream‑URL>:<your‑port>/write.
	PrometheusAPI *string `default:"/write" json:"prometheusAPI"`
	// Remote Write authentication type
	AuthType *InputPrometheusRwAuthenticationType `default:"none" json:"authType"`
	// Fields to add to events from this input
	Metadata    []InputPrometheusRwMetadatum `json:"metadata,omitempty"`
	Description *string                      `json:"description,omitempty"`
	Username    *string                      `json:"username,omitempty"`
	Password    *string                      `json:"password,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Select or create a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `default:"Bearer \\${token}" json:"authHeaderExpr"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *float64 `default:"3600" json:"tokenTimeoutSecs"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []InputPrometheusRwOauthParam `json:"oauthParams,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []InputPrometheusRwOauthHeader `json:"oauthHeaders,omitempty"`
}

func (i InputPrometheusRw) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputPrometheusRw) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputPrometheusRw) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputPrometheusRw) GetType() *InputPrometheusRwType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputPrometheusRw) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputPrometheusRw) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputPrometheusRw) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputPrometheusRw) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputPrometheusRw) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputPrometheusRw) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputPrometheusRw) GetConnections() []InputPrometheusRwConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputPrometheusRw) GetPq() *InputPrometheusRwPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputPrometheusRw) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputPrometheusRw) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *InputPrometheusRw) GetTLS() *InputPrometheusRwTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputPrometheusRw) GetMaxActiveReq() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveReq
}

func (o *InputPrometheusRw) GetMaxRequestsPerSocket() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestsPerSocket
}

func (o *InputPrometheusRw) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputPrometheusRw) GetCaptureHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.CaptureHeaders
}

func (o *InputPrometheusRw) GetActivityLogSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.ActivityLogSampleRate
}

func (o *InputPrometheusRw) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputPrometheusRw) GetSocketTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketTimeout
}

func (o *InputPrometheusRw) GetKeepAliveTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTimeout
}

func (o *InputPrometheusRw) GetEnableHealthCheck() *bool {
	if o == nil {
		return nil
	}
	return o.EnableHealthCheck
}

func (o *InputPrometheusRw) GetIPAllowlistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPAllowlistRegex
}

func (o *InputPrometheusRw) GetIPDenylistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPDenylistRegex
}

func (o *InputPrometheusRw) GetPrometheusAPI() *string {
	if o == nil {
		return nil
	}
	return o.PrometheusAPI
}

func (o *InputPrometheusRw) GetAuthType() *InputPrometheusRwAuthenticationType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *InputPrometheusRw) GetMetadata() []InputPrometheusRwMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputPrometheusRw) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputPrometheusRw) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *InputPrometheusRw) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *InputPrometheusRw) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *InputPrometheusRw) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

func (o *InputPrometheusRw) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *InputPrometheusRw) GetLoginURL() *string {
	if o == nil {
		return nil
	}
	return o.LoginURL
}

func (o *InputPrometheusRw) GetSecretParamName() *string {
	if o == nil {
		return nil
	}
	return o.SecretParamName
}

func (o *InputPrometheusRw) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *InputPrometheusRw) GetTokenAttributeName() *string {
	if o == nil {
		return nil
	}
	return o.TokenAttributeName
}

func (o *InputPrometheusRw) GetAuthHeaderExpr() *string {
	if o == nil {
		return nil
	}
	return o.AuthHeaderExpr
}

func (o *InputPrometheusRw) GetTokenTimeoutSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TokenTimeoutSecs
}

func (o *InputPrometheusRw) GetOauthParams() []InputPrometheusRwOauthParam {
	if o == nil {
		return nil
	}
	return o.OauthParams
}

func (o *InputPrometheusRw) GetOauthHeaders() []InputPrometheusRwOauthHeader {
	if o == nil {
		return nil
	}
	return o.OauthHeaders
}
