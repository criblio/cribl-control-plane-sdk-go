// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
)

type InputPrometheusType string

const (
	InputPrometheusTypePrometheus InputPrometheusType = "prometheus"
)

func (e InputPrometheusType) ToPointer() *InputPrometheusType {
	return &e
}
func (e *InputPrometheusType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prometheus":
		*e = InputPrometheusType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusType: %v", v)
	}
}

type InputPrometheusConnection struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputPrometheusConnection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputPrometheusConnection) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputPrometheusMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputPrometheusMode string

const (
	InputPrometheusModeSmart  InputPrometheusMode = "smart"
	InputPrometheusModeAlways InputPrometheusMode = "always"
)

func (e InputPrometheusMode) ToPointer() *InputPrometheusMode {
	return &e
}
func (e *InputPrometheusMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputPrometheusMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusMode: %v", v)
	}
}

// InputPrometheusCompression - Codec to use to compress the persisted data
type InputPrometheusCompression string

const (
	InputPrometheusCompressionNone InputPrometheusCompression = "none"
	InputPrometheusCompressionGzip InputPrometheusCompression = "gzip"
)

func (e InputPrometheusCompression) ToPointer() *InputPrometheusCompression {
	return &e
}
func (e *InputPrometheusCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputPrometheusCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusCompression: %v", v)
	}
}

type InputPrometheusPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputPrometheusMode `default:"always" json:"mode"`
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
	Compress *InputPrometheusCompression `default:"none" json:"compress"`
}

func (i InputPrometheusPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputPrometheusPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputPrometheusPq) GetMode() *InputPrometheusMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputPrometheusPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputPrometheusPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputPrometheusPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputPrometheusPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputPrometheusPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputPrometheusPq) GetCompress() *InputPrometheusCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputPrometheusDiscoveryType - Target discovery mechanism. Use static to manually enter a list of targets.
type InputPrometheusDiscoveryType string

const (
	InputPrometheusDiscoveryTypeStatic InputPrometheusDiscoveryType = "static"
	InputPrometheusDiscoveryTypeDNS    InputPrometheusDiscoveryType = "dns"
	InputPrometheusDiscoveryTypeEc2    InputPrometheusDiscoveryType = "ec2"
)

func (e InputPrometheusDiscoveryType) ToPointer() *InputPrometheusDiscoveryType {
	return &e
}
func (e *InputPrometheusDiscoveryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		fallthrough
	case "dns":
		fallthrough
	case "ec2":
		*e = InputPrometheusDiscoveryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusDiscoveryType: %v", v)
	}
}

// InputPrometheusLogLevel - Collector runtime Log Level
type InputPrometheusLogLevel string

const (
	InputPrometheusLogLevelError InputPrometheusLogLevel = "error"
	InputPrometheusLogLevelWarn  InputPrometheusLogLevel = "warn"
	InputPrometheusLogLevelInfo  InputPrometheusLogLevel = "info"
	InputPrometheusLogLevelDebug InputPrometheusLogLevel = "debug"
)

func (e InputPrometheusLogLevel) ToPointer() *InputPrometheusLogLevel {
	return &e
}
func (e *InputPrometheusLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "warn":
		fallthrough
	case "info":
		fallthrough
	case "debug":
		*e = InputPrometheusLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusLogLevel: %v", v)
	}
}

type InputPrometheusMetadatum struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputPrometheusMetadatum) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputPrometheusMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// InputPrometheusAuthTypeAuthenticationMethod - Enter credentials directly, or select a stored secret
type InputPrometheusAuthTypeAuthenticationMethod string

const (
	InputPrometheusAuthTypeAuthenticationMethodManual InputPrometheusAuthTypeAuthenticationMethod = "manual"
	InputPrometheusAuthTypeAuthenticationMethodSecret InputPrometheusAuthTypeAuthenticationMethod = "secret"
)

func (e InputPrometheusAuthTypeAuthenticationMethod) ToPointer() *InputPrometheusAuthTypeAuthenticationMethod {
	return &e
}
func (e *InputPrometheusAuthTypeAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "secret":
		*e = InputPrometheusAuthTypeAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusAuthTypeAuthenticationMethod: %v", v)
	}
}

// InputPrometheusRecordType - DNS Record type to resolve
type InputPrometheusRecordType string

const (
	InputPrometheusRecordTypeSrv  InputPrometheusRecordType = "SRV"
	InputPrometheusRecordTypeA    InputPrometheusRecordType = "A"
	InputPrometheusRecordTypeAaaa InputPrometheusRecordType = "AAAA"
)

func (e InputPrometheusRecordType) ToPointer() *InputPrometheusRecordType {
	return &e
}
func (e *InputPrometheusRecordType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SRV":
		fallthrough
	case "A":
		fallthrough
	case "AAAA":
		*e = InputPrometheusRecordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusRecordType: %v", v)
	}
}

// MetricsProtocol - Protocol to use when collecting metrics
type MetricsProtocol string

const (
	MetricsProtocolHTTP  MetricsProtocol = "http"
	MetricsProtocolHTTPS MetricsProtocol = "https"
)

func (e MetricsProtocol) ToPointer() *MetricsProtocol {
	return &e
}
func (e *MetricsProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = MetricsProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MetricsProtocol: %v", v)
	}
}

type InputPrometheusSearchFilter struct {
	// Search filter attribute name, see: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html for more information. Attributes can be manually entered if not present in the drop down list
	Name string `json:"Name"`
	// Search Filter Values, if empty only "running" EC2 instances will be returned
	Values []string `json:"Values"`
}

func (o *InputPrometheusSearchFilter) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputPrometheusSearchFilter) GetValues() []string {
	if o == nil {
		return []string{}
	}
	return o.Values
}

// InputPrometheusAwsAuthenticationMethodAuthenticationMethod - AWS authentication method. Choose Auto to use IAM roles.
type InputPrometheusAwsAuthenticationMethodAuthenticationMethod string

const (
	InputPrometheusAwsAuthenticationMethodAuthenticationMethodAuto   InputPrometheusAwsAuthenticationMethodAuthenticationMethod = "auto"
	InputPrometheusAwsAuthenticationMethodAuthenticationMethodManual InputPrometheusAwsAuthenticationMethodAuthenticationMethod = "manual"
	InputPrometheusAwsAuthenticationMethodAuthenticationMethodSecret InputPrometheusAwsAuthenticationMethodAuthenticationMethod = "secret"
)

func (e InputPrometheusAwsAuthenticationMethodAuthenticationMethod) ToPointer() *InputPrometheusAwsAuthenticationMethodAuthenticationMethod {
	return &e
}
func (e *InputPrometheusAwsAuthenticationMethodAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "manual":
		fallthrough
	case "secret":
		*e = InputPrometheusAwsAuthenticationMethodAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusAwsAuthenticationMethodAuthenticationMethod: %v", v)
	}
}

// InputPrometheusSignatureVersion - Signature version to use for signing EC2 requests
type InputPrometheusSignatureVersion string

const (
	InputPrometheusSignatureVersionV2 InputPrometheusSignatureVersion = "v2"
	InputPrometheusSignatureVersionV4 InputPrometheusSignatureVersion = "v4"
)

func (e InputPrometheusSignatureVersion) ToPointer() *InputPrometheusSignatureVersion {
	return &e
}
func (e *InputPrometheusSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = InputPrometheusSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputPrometheusSignatureVersion: %v", v)
	}
}

type InputPrometheus struct {
	// Unique ID for this input
	ID       *string              `json:"id,omitempty"`
	Type     *InputPrometheusType `json:"type,omitempty"`
	Disabled *bool                `default:"false" json:"disabled"`
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
	Connections []InputPrometheusConnection `json:"connections,omitempty"`
	Pq          *InputPrometheusPq          `json:"pq,omitempty"`
	// Other dimensions to include in events
	DimensionList []string `json:"dimensionList,omitempty"`
	// Target discovery mechanism. Use static to manually enter a list of targets.
	DiscoveryType *InputPrometheusDiscoveryType `default:"static" json:"discoveryType"`
	// How often in minutes to scrape targets for metrics, 60 must be evenly divisible by the value or save will fail.
	Interval *float64 `default:"15" json:"interval"`
	// Collector runtime Log Level
	LogLevel *InputPrometheusLogLevel `default:"info" json:"logLevel"`
	// Reject certificates that cannot be verified against a valid CA, such as self-signed certificates
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// How often workers should check in with the scheduler to keep job subscription alive
	KeepAliveTime *float64 `default:"30" json:"keepAliveTime"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `default:"0" json:"jobTimeout"`
	// The number of Keep Alive Time periods before an inactive worker will have its job subscription revoked.
	MaxMissedKeepAlives *float64 `default:"3" json:"maxMissedKeepAlives"`
	// Time to keep the job's artifacts on disk after job completion. This also affects how long a job is listed in the Job Inspector.
	TTL *string `default:"4h" json:"ttl"`
	// When enabled, this job's artifacts are not counted toward the Worker Group's finished job artifacts limit. Artifacts will be removed only after the Collector's configured time to live.
	IgnoreGroupJobsLimit *bool `default:"false" json:"ignoreGroupJobsLimit"`
	// Fields to add to events from this input
	Metadata []InputPrometheusMetadatum `json:"metadata,omitempty"`
	// Enter credentials directly, or select a stored secret
	AuthType    *InputPrometheusAuthTypeAuthenticationMethod `default:"manual" json:"authType"`
	Description *string                                      `json:"description,omitempty"`
	// List of Prometheus targets to pull metrics from. Values can be in URL or host[:port] format. For example: http://localhost:9090/metrics, localhost:9090, or localhost. In cases where just host[:port] is specified, the endpoint will resolve to 'http://host[:port]/metrics'.
	TargetList []string `json:"targetList,omitempty"`
	// List of DNS names to resolve
	NameList []string `json:"nameList,omitempty"`
	// DNS Record type to resolve
	RecordType *InputPrometheusRecordType `default:"SRV" json:"recordType"`
	// Protocol to use when collecting metrics
	ScrapeProtocol *MetricsProtocol `default:"http" json:"scrapeProtocol"`
	// Path to use when collecting metrics from discovered targets
	ScrapePath *string `default:"/metrics" json:"scrapePath"`
	// Use public IP address for discovered targets. Set to false if the private IP address should be used.
	UsePublicIP *bool `default:"true" json:"usePublicIp"`
	// The port number in the metrics URL for discovered targets.
	ScrapePort *float64 `default:"9090" json:"scrapePort"`
	// EC2 Instance Search Filter
	SearchFilter []InputPrometheusSearchFilter `json:"searchFilter,omitempty"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *InputPrometheusAwsAuthenticationMethodAuthenticationMethod `default:"auto" json:"awsAuthenticationMethod"`
	AwsSecretKey            *string                                                     `json:"awsSecretKey,omitempty"`
	// Region where the EC2 is located
	Region *string `json:"region,omitempty"`
	// EC2 service endpoint. If empty, defaults to the AWS Region-specific endpoint. Otherwise, it must point to EC2-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Signature version to use for signing EC2 requests
	SignatureVersion *InputPrometheusSignatureVersion `default:"v4" json:"signatureVersion"`
	// Reuse connections between requests, which can improve performance
	ReuseConnections *bool `default:"true" json:"reuseConnections"`
	// Use Assume Role credentials to access EC2
	EnableAssumeRole *bool `default:"false" json:"enableAssumeRole"`
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Duration of the assumed role's session, in seconds. Minimum is 900 (15 minutes), default is 3600 (1 hour), and maximum is 43200 (12 hours).
	DurationSeconds *float64 `default:"3600" json:"durationSeconds"`
	// Username for Prometheus Basic authentication
	Username *string `json:"username,omitempty"`
	// Password for Prometheus Basic authentication
	Password *string `json:"password,omitempty"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
}

func (i InputPrometheus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputPrometheus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputPrometheus) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputPrometheus) GetType() *InputPrometheusType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputPrometheus) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputPrometheus) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputPrometheus) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputPrometheus) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputPrometheus) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputPrometheus) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputPrometheus) GetConnections() []InputPrometheusConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputPrometheus) GetPq() *InputPrometheusPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputPrometheus) GetDimensionList() []string {
	if o == nil {
		return nil
	}
	return o.DimensionList
}

func (o *InputPrometheus) GetDiscoveryType() *InputPrometheusDiscoveryType {
	if o == nil {
		return nil
	}
	return o.DiscoveryType
}

func (o *InputPrometheus) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputPrometheus) GetLogLevel() *InputPrometheusLogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *InputPrometheus) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputPrometheus) GetKeepAliveTime() *float64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTime
}

func (o *InputPrometheus) GetJobTimeout() *string {
	if o == nil {
		return nil
	}
	return o.JobTimeout
}

func (o *InputPrometheus) GetMaxMissedKeepAlives() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxMissedKeepAlives
}

func (o *InputPrometheus) GetTTL() *string {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *InputPrometheus) GetIgnoreGroupJobsLimit() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreGroupJobsLimit
}

func (o *InputPrometheus) GetMetadata() []InputPrometheusMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputPrometheus) GetAuthType() *InputPrometheusAuthTypeAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *InputPrometheus) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputPrometheus) GetTargetList() []string {
	if o == nil {
		return nil
	}
	return o.TargetList
}

func (o *InputPrometheus) GetNameList() []string {
	if o == nil {
		return nil
	}
	return o.NameList
}

func (o *InputPrometheus) GetRecordType() *InputPrometheusRecordType {
	if o == nil {
		return nil
	}
	return o.RecordType
}

func (o *InputPrometheus) GetScrapeProtocol() *MetricsProtocol {
	if o == nil {
		return nil
	}
	return o.ScrapeProtocol
}

func (o *InputPrometheus) GetScrapePath() *string {
	if o == nil {
		return nil
	}
	return o.ScrapePath
}

func (o *InputPrometheus) GetUsePublicIP() *bool {
	if o == nil {
		return nil
	}
	return o.UsePublicIP
}

func (o *InputPrometheus) GetScrapePort() *float64 {
	if o == nil {
		return nil
	}
	return o.ScrapePort
}

func (o *InputPrometheus) GetSearchFilter() []InputPrometheusSearchFilter {
	if o == nil {
		return nil
	}
	return o.SearchFilter
}

func (o *InputPrometheus) GetAwsAuthenticationMethod() *InputPrometheusAwsAuthenticationMethodAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AwsAuthenticationMethod
}

func (o *InputPrometheus) GetAwsSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretKey
}

func (o *InputPrometheus) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *InputPrometheus) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *InputPrometheus) GetSignatureVersion() *InputPrometheusSignatureVersion {
	if o == nil {
		return nil
	}
	return o.SignatureVersion
}

func (o *InputPrometheus) GetReuseConnections() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnections
}

func (o *InputPrometheus) GetEnableAssumeRole() *bool {
	if o == nil {
		return nil
	}
	return o.EnableAssumeRole
}

func (o *InputPrometheus) GetAssumeRoleArn() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleArn
}

func (o *InputPrometheus) GetAssumeRoleExternalID() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleExternalID
}

func (o *InputPrometheus) GetDurationSeconds() *float64 {
	if o == nil {
		return nil
	}
	return o.DurationSeconds
}

func (o *InputPrometheus) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *InputPrometheus) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *InputPrometheus) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}
