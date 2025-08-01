// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type OutputAzureLogsType string

const (
	OutputAzureLogsTypeAzureLogs OutputAzureLogsType = "azure_logs"
)

func (e OutputAzureLogsType) ToPointer() *OutputAzureLogsType {
	return &e
}
func (e *OutputAzureLogsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure_logs":
		*e = OutputAzureLogsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsType: %v", v)
	}
}

type OutputAzureLogsExtraHTTPHeader struct {
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value"`
}

func (o *OutputAzureLogsExtraHTTPHeader) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputAzureLogsExtraHTTPHeader) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputAzureLogsFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputAzureLogsFailedRequestLoggingMode string

const (
	OutputAzureLogsFailedRequestLoggingModePayload           OutputAzureLogsFailedRequestLoggingMode = "payload"
	OutputAzureLogsFailedRequestLoggingModePayloadAndHeaders OutputAzureLogsFailedRequestLoggingMode = "payloadAndHeaders"
	OutputAzureLogsFailedRequestLoggingModeNone              OutputAzureLogsFailedRequestLoggingMode = "none"
)

func (e OutputAzureLogsFailedRequestLoggingMode) ToPointer() *OutputAzureLogsFailedRequestLoggingMode {
	return &e
}
func (e *OutputAzureLogsFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payload":
		fallthrough
	case "payloadAndHeaders":
		fallthrough
	case "none":
		*e = OutputAzureLogsFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsFailedRequestLoggingMode: %v", v)
	}
}

type OutputAzureLogsResponseRetrySetting struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputAzureLogsResponseRetrySetting) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputAzureLogsResponseRetrySetting) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputAzureLogsResponseRetrySetting) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputAzureLogsResponseRetrySetting) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputAzureLogsResponseRetrySetting) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputAzureLogsResponseRetrySetting) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputAzureLogsTimeoutRetrySettings struct {
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputAzureLogsTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputAzureLogsTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputAzureLogsTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputAzureLogsTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputAzureLogsTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputAzureLogsTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputAzureLogsBackpressureBehavior - How to handle events when all receivers are exerting backpressure
type OutputAzureLogsBackpressureBehavior string

const (
	OutputAzureLogsBackpressureBehaviorBlock OutputAzureLogsBackpressureBehavior = "block"
	OutputAzureLogsBackpressureBehaviorDrop  OutputAzureLogsBackpressureBehavior = "drop"
	OutputAzureLogsBackpressureBehaviorQueue OutputAzureLogsBackpressureBehavior = "queue"
)

func (e OutputAzureLogsBackpressureBehavior) ToPointer() *OutputAzureLogsBackpressureBehavior {
	return &e
}
func (e *OutputAzureLogsBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		fallthrough
	case "queue":
		*e = OutputAzureLogsBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsBackpressureBehavior: %v", v)
	}
}

// OutputAzureLogsAuthenticationMethod - Enter workspace ID and workspace key directly, or select a stored secret
type OutputAzureLogsAuthenticationMethod string

const (
	OutputAzureLogsAuthenticationMethodManual OutputAzureLogsAuthenticationMethod = "manual"
	OutputAzureLogsAuthenticationMethodSecret OutputAzureLogsAuthenticationMethod = "secret"
)

func (e OutputAzureLogsAuthenticationMethod) ToPointer() *OutputAzureLogsAuthenticationMethod {
	return &e
}
func (e *OutputAzureLogsAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "secret":
		*e = OutputAzureLogsAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsAuthenticationMethod: %v", v)
	}
}

// OutputAzureLogsCompression - Codec to use to compress the persisted data
type OutputAzureLogsCompression string

const (
	OutputAzureLogsCompressionNone OutputAzureLogsCompression = "none"
	OutputAzureLogsCompressionGzip OutputAzureLogsCompression = "gzip"
)

func (e OutputAzureLogsCompression) ToPointer() *OutputAzureLogsCompression {
	return &e
}
func (e *OutputAzureLogsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputAzureLogsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsCompression: %v", v)
	}
}

// OutputAzureLogsQueueFullBehavior - How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputAzureLogsQueueFullBehavior string

const (
	OutputAzureLogsQueueFullBehaviorBlock OutputAzureLogsQueueFullBehavior = "block"
	OutputAzureLogsQueueFullBehaviorDrop  OutputAzureLogsQueueFullBehavior = "drop"
)

func (e OutputAzureLogsQueueFullBehavior) ToPointer() *OutputAzureLogsQueueFullBehavior {
	return &e
}
func (e *OutputAzureLogsQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputAzureLogsQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsQueueFullBehavior: %v", v)
	}
}

// OutputAzureLogsMode - In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.
type OutputAzureLogsMode string

const (
	OutputAzureLogsModeError        OutputAzureLogsMode = "error"
	OutputAzureLogsModeBackpressure OutputAzureLogsMode = "backpressure"
	OutputAzureLogsModeAlways       OutputAzureLogsMode = "always"
)

func (e OutputAzureLogsMode) ToPointer() *OutputAzureLogsMode {
	return &e
}
func (e *OutputAzureLogsMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "backpressure":
		fallthrough
	case "always":
		*e = OutputAzureLogsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputAzureLogsMode: %v", v)
	}
}

type OutputAzureLogsPqControls struct {
}

type OutputAzureLogs struct {
	// Unique ID for this output
	ID   *string             `json:"id,omitempty"`
	Type OutputAzureLogsType `json:"type"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// The Log Type of events sent to this LogAnalytics workspace. Defaults to `Cribl`. Use only letters, numbers, and `_` characters, and can't exceed 100 characters. Can be overwritten by event field __logType.
	LogType *string `default:"Cribl" json:"logType"`
	// Optional Resource ID of the Azure resource to associate the data with. Can be overridden by the __resourceId event field. This ID populates the _ResourceId property, allowing the data to be included in resource-centric queries. If the ID is neither specified nor overridden, resource-centric queries will omit the data.
	ResourceID *string `json:"resourceId,omitempty"`
	// Maximum number of ongoing requests before blocking
	Concurrency *float64 `default:"5" json:"concurrency"`
	// Maximum size, in KB, of the request body
	MaxPayloadSizeKB *float64 `default:"1024" json:"maxPayloadSizeKB"`
	// Maximum number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *float64 `default:"0" json:"maxPayloadEvents"`
	Compress         *bool    `json:"compress,omitempty"`
	// Reject certificates not authorized by a CA in the CA certificate path or by another trusted CA (such as the system's).
	//         Enabled by default. When this setting is also present in TLS Settings (Client Side),
	//         that value will take precedence.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Amount of time, in seconds, to wait for a request to complete before canceling it
	TimeoutSec *float64 `default:"30" json:"timeoutSec"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Body size limit.
	FlushPeriodSec *float64 `default:"1" json:"flushPeriodSec"`
	// Headers to add to all events
	ExtraHTTPHeaders []OutputAzureLogsExtraHTTPHeader `json:"extraHttpHeaders,omitempty"`
	// Enable round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputAzureLogsFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// The DNS name of the Log API endpoint that sends log data to a Log Analytics workspace in Azure Monitor. Defaults to .ods.opinsights.azure.com. @{product} will add a prefix and suffix to construct a URI in this format: <https://<Workspace_ID><your_DNS_name>/api/logs?api-version=<API version>.
	APIURL *string `default:".ods.opinsights.azure.com" json:"apiUrl"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable)
	ResponseRetrySettings []OutputAzureLogsResponseRetrySetting `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputAzureLogsTimeoutRetrySettings  `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// How to handle events when all receivers are exerting backpressure
	OnBackpressure *OutputAzureLogsBackpressureBehavior `default:"block" json:"onBackpressure"`
	// Enter workspace ID and workspace key directly, or select a stored secret
	AuthType    *OutputAzureLogsAuthenticationMethod `default:"manual" json:"authType"`
	Description *string                              `json:"description,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.)
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data
	PqCompress *OutputAzureLogsCompression `default:"none" json:"pqCompress"`
	// How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputAzureLogsQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputAzureLogsMode       `default:"error" json:"pqMode"`
	PqControls *OutputAzureLogsPqControls `json:"pqControls,omitempty"`
	// Azure Log Analytics Workspace ID. See Azure Dashboard Workspace > Advanced settings.
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// Azure Log Analytics Workspace Primary or Secondary Shared Key. See Azure Dashboard Workspace > Advanced settings.
	WorkspaceKey *string `json:"workspaceKey,omitempty"`
	// Select or create a stored secret that references your access key and secret key
	KeypairSecret *string `json:"keypairSecret,omitempty"`
}

func (o OutputAzureLogs) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputAzureLogs) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputAzureLogs) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputAzureLogs) GetType() OutputAzureLogsType {
	if o == nil {
		return OutputAzureLogsType("")
	}
	return o.Type
}

func (o *OutputAzureLogs) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputAzureLogs) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputAzureLogs) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputAzureLogs) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputAzureLogs) GetLogType() *string {
	if o == nil {
		return nil
	}
	return o.LogType
}

func (o *OutputAzureLogs) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *OutputAzureLogs) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputAzureLogs) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputAzureLogs) GetMaxPayloadEvents() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputAzureLogs) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputAzureLogs) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputAzureLogs) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputAzureLogs) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputAzureLogs) GetExtraHTTPHeaders() []OutputAzureLogsExtraHTTPHeader {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputAzureLogs) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputAzureLogs) GetFailedRequestLoggingMode() *OutputAzureLogsFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputAzureLogs) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputAzureLogs) GetAPIURL() *string {
	if o == nil {
		return nil
	}
	return o.APIURL
}

func (o *OutputAzureLogs) GetResponseRetrySettings() []OutputAzureLogsResponseRetrySetting {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputAzureLogs) GetTimeoutRetrySettings() *OutputAzureLogsTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputAzureLogs) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputAzureLogs) GetOnBackpressure() *OutputAzureLogsBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputAzureLogs) GetAuthType() *OutputAzureLogsAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputAzureLogs) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputAzureLogs) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputAzureLogs) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputAzureLogs) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputAzureLogs) GetPqCompress() *OutputAzureLogsCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputAzureLogs) GetPqOnBackpressure() *OutputAzureLogsQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputAzureLogs) GetPqMode() *OutputAzureLogsMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputAzureLogs) GetPqControls() *OutputAzureLogsPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}

func (o *OutputAzureLogs) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *OutputAzureLogs) GetWorkspaceKey() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceKey
}

func (o *OutputAzureLogs) GetKeypairSecret() *string {
	if o == nil {
		return nil
	}
	return o.KeypairSecret
}
