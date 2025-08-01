// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
)

type OutputDynatraceOtlpType string

const (
	OutputDynatraceOtlpTypeDynatraceOtlp OutputDynatraceOtlpType = "dynatrace_otlp"
)

func (e OutputDynatraceOtlpType) ToPointer() *OutputDynatraceOtlpType {
	return &e
}
func (e *OutputDynatraceOtlpType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dynatrace_otlp":
		*e = OutputDynatraceOtlpType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpType: %v", v)
	}
}

// OutputDynatraceOtlpProtocol - Select a transport option for Dynatrace
type OutputDynatraceOtlpProtocol string

const (
	OutputDynatraceOtlpProtocolHTTP OutputDynatraceOtlpProtocol = "http"
)

func (e OutputDynatraceOtlpProtocol) ToPointer() *OutputDynatraceOtlpProtocol {
	return &e
}
func (e *OutputDynatraceOtlpProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		*e = OutputDynatraceOtlpProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpProtocol: %v", v)
	}
}

// OutputDynatraceOTLPOTLPVersion - The version of OTLP Protobuf definitions to use when structuring data to send
type OutputDynatraceOTLPOTLPVersion string

const (
	OutputDynatraceOTLPOTLPVersionOneDot3Dot1 OutputDynatraceOTLPOTLPVersion = "1.3.1"
)

func (e OutputDynatraceOTLPOTLPVersion) ToPointer() *OutputDynatraceOTLPOTLPVersion {
	return &e
}
func (e *OutputDynatraceOTLPOTLPVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1.3.1":
		*e = OutputDynatraceOTLPOTLPVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOTLPOTLPVersion: %v", v)
	}
}

// OutputDynatraceOtlpCompressCompression - Type of compression to apply to messages sent to the OpenTelemetry endpoint
type OutputDynatraceOtlpCompressCompression string

const (
	OutputDynatraceOtlpCompressCompressionNone    OutputDynatraceOtlpCompressCompression = "none"
	OutputDynatraceOtlpCompressCompressionDeflate OutputDynatraceOtlpCompressCompression = "deflate"
	OutputDynatraceOtlpCompressCompressionGzip    OutputDynatraceOtlpCompressCompression = "gzip"
)

func (e OutputDynatraceOtlpCompressCompression) ToPointer() *OutputDynatraceOtlpCompressCompression {
	return &e
}
func (e *OutputDynatraceOtlpCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "deflate":
		fallthrough
	case "gzip":
		*e = OutputDynatraceOtlpCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpCompressCompression: %v", v)
	}
}

// OutputDynatraceOtlpHTTPCompressCompression - Type of compression to apply to messages sent to the OpenTelemetry endpoint
type OutputDynatraceOtlpHTTPCompressCompression string

const (
	OutputDynatraceOtlpHTTPCompressCompressionNone OutputDynatraceOtlpHTTPCompressCompression = "none"
	OutputDynatraceOtlpHTTPCompressCompressionGzip OutputDynatraceOtlpHTTPCompressCompression = "gzip"
)

func (e OutputDynatraceOtlpHTTPCompressCompression) ToPointer() *OutputDynatraceOtlpHTTPCompressCompression {
	return &e
}
func (e *OutputDynatraceOtlpHTTPCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputDynatraceOtlpHTTPCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpHTTPCompressCompression: %v", v)
	}
}

type OutputDynatraceOtlpMetadatum struct {
	Key   *string `default:"" json:"key"`
	Value string  `json:"value"`
}

func (o OutputDynatraceOtlpMetadatum) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputDynatraceOtlpMetadatum) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputDynatraceOtlpMetadatum) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *OutputDynatraceOtlpMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputDynatraceOtlpFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputDynatraceOtlpFailedRequestLoggingMode string

const (
	OutputDynatraceOtlpFailedRequestLoggingModePayload           OutputDynatraceOtlpFailedRequestLoggingMode = "payload"
	OutputDynatraceOtlpFailedRequestLoggingModePayloadAndHeaders OutputDynatraceOtlpFailedRequestLoggingMode = "payloadAndHeaders"
	OutputDynatraceOtlpFailedRequestLoggingModeNone              OutputDynatraceOtlpFailedRequestLoggingMode = "none"
)

func (e OutputDynatraceOtlpFailedRequestLoggingMode) ToPointer() *OutputDynatraceOtlpFailedRequestLoggingMode {
	return &e
}
func (e *OutputDynatraceOtlpFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputDynatraceOtlpFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpFailedRequestLoggingMode: %v", v)
	}
}

// EndpointType - Select the type of Dynatrace endpoint configured
type EndpointType string

const (
	EndpointTypeSaas EndpointType = "saas"
	EndpointTypeAg   EndpointType = "ag"
)

func (e EndpointType) ToPointer() *EndpointType {
	return &e
}
func (e *EndpointType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saas":
		fallthrough
	case "ag":
		*e = EndpointType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EndpointType: %v", v)
	}
}

// OutputDynatraceOtlpBackpressureBehavior - How to handle events when all receivers are exerting backpressure
type OutputDynatraceOtlpBackpressureBehavior string

const (
	OutputDynatraceOtlpBackpressureBehaviorBlock OutputDynatraceOtlpBackpressureBehavior = "block"
	OutputDynatraceOtlpBackpressureBehaviorDrop  OutputDynatraceOtlpBackpressureBehavior = "drop"
	OutputDynatraceOtlpBackpressureBehaviorQueue OutputDynatraceOtlpBackpressureBehavior = "queue"
)

func (e OutputDynatraceOtlpBackpressureBehavior) ToPointer() *OutputDynatraceOtlpBackpressureBehavior {
	return &e
}
func (e *OutputDynatraceOtlpBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputDynatraceOtlpBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpBackpressureBehavior: %v", v)
	}
}

type OutputDynatraceOtlpExtraHTTPHeader struct {
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value"`
}

func (o *OutputDynatraceOtlpExtraHTTPHeader) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputDynatraceOtlpExtraHTTPHeader) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputDynatraceOtlpResponseRetrySetting struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputDynatraceOtlpResponseRetrySetting) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputDynatraceOtlpResponseRetrySetting) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputDynatraceOtlpResponseRetrySetting) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputDynatraceOtlpResponseRetrySetting) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputDynatraceOtlpResponseRetrySetting) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputDynatraceOtlpResponseRetrySetting) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputDynatraceOtlpTimeoutRetrySettings struct {
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputDynatraceOtlpTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputDynatraceOtlpTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputDynatraceOtlpTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputDynatraceOtlpTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputDynatraceOtlpTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputDynatraceOtlpTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputDynatraceOtlpPqCompressCompression - Codec to use to compress the persisted data
type OutputDynatraceOtlpPqCompressCompression string

const (
	OutputDynatraceOtlpPqCompressCompressionNone OutputDynatraceOtlpPqCompressCompression = "none"
	OutputDynatraceOtlpPqCompressCompressionGzip OutputDynatraceOtlpPqCompressCompression = "gzip"
)

func (e OutputDynatraceOtlpPqCompressCompression) ToPointer() *OutputDynatraceOtlpPqCompressCompression {
	return &e
}
func (e *OutputDynatraceOtlpPqCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputDynatraceOtlpPqCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpPqCompressCompression: %v", v)
	}
}

// OutputDynatraceOtlpQueueFullBehavior - How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputDynatraceOtlpQueueFullBehavior string

const (
	OutputDynatraceOtlpQueueFullBehaviorBlock OutputDynatraceOtlpQueueFullBehavior = "block"
	OutputDynatraceOtlpQueueFullBehaviorDrop  OutputDynatraceOtlpQueueFullBehavior = "drop"
)

func (e OutputDynatraceOtlpQueueFullBehavior) ToPointer() *OutputDynatraceOtlpQueueFullBehavior {
	return &e
}
func (e *OutputDynatraceOtlpQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputDynatraceOtlpQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpQueueFullBehavior: %v", v)
	}
}

// OutputDynatraceOtlpMode - In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.
type OutputDynatraceOtlpMode string

const (
	OutputDynatraceOtlpModeError        OutputDynatraceOtlpMode = "error"
	OutputDynatraceOtlpModeBackpressure OutputDynatraceOtlpMode = "backpressure"
	OutputDynatraceOtlpModeAlways       OutputDynatraceOtlpMode = "always"
)

func (e OutputDynatraceOtlpMode) ToPointer() *OutputDynatraceOtlpMode {
	return &e
}
func (e *OutputDynatraceOtlpMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputDynatraceOtlpMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDynatraceOtlpMode: %v", v)
	}
}

type OutputDynatraceOtlpPqControls struct {
}

type OutputDynatraceOtlp struct {
	// Unique ID for this output
	ID   *string                  `json:"id,omitempty"`
	Type *OutputDynatraceOtlpType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Select a transport option for Dynatrace
	Protocol *OutputDynatraceOtlpProtocol `default:"http" json:"protocol"`
	// The endpoint where Dynatrace events will be sent. Enter any valid URL or an IP address (IPv4 or IPv6; enclose IPv6 addresses in square brackets)
	Endpoint *string `default:"https://{your-environment-id}.live.dynatrace.com/api/v2/otlp" json:"endpoint"`
	// The version of OTLP Protobuf definitions to use when structuring data to send
	OtlpVersion *OutputDynatraceOTLPOTLPVersion `default:"1.3.1" json:"otlpVersion"`
	// Type of compression to apply to messages sent to the OpenTelemetry endpoint
	Compress *OutputDynatraceOtlpCompressCompression `default:"gzip" json:"compress"`
	// Type of compression to apply to messages sent to the OpenTelemetry endpoint
	HTTPCompress *OutputDynatraceOtlpHTTPCompressCompression `default:"gzip" json:"httpCompress"`
	// If you want to send traces to the default `{endpoint}/v1/traces` endpoint, leave this field empty; otherwise, specify the desired endpoint
	HTTPTracesEndpointOverride *string `json:"httpTracesEndpointOverride,omitempty"`
	// If you want to send metrics to the default `{endpoint}/v1/metrics` endpoint, leave this field empty; otherwise, specify the desired endpoint
	HTTPMetricsEndpointOverride *string `json:"httpMetricsEndpointOverride,omitempty"`
	// If you want to send logs to the default `{endpoint}/v1/logs` endpoint, leave this field empty; otherwise, specify the desired endpoint
	HTTPLogsEndpointOverride *string `json:"httpLogsEndpointOverride,omitempty"`
	// List of key-value pairs to send with each gRPC request. Value supports JavaScript expressions that are evaluated just once, when the destination gets started. To pass credentials as metadata, use 'C.Secret'.
	Metadata []OutputDynatraceOtlpMetadatum `json:"metadata,omitempty"`
	// Maximum number of ongoing requests before blocking
	Concurrency *float64 `default:"5" json:"concurrency"`
	// Maximum size (in KB) of the request body. The maximum payload size is 4 MB. If this limit is exceeded, the entire OTLP message is dropped
	MaxPayloadSizeKB *float64 `default:"2048" json:"maxPayloadSizeKB"`
	// Amount of time, in seconds, to wait for a request to complete before canceling it
	TimeoutSec *float64 `default:"30" json:"timeoutSec"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Body size limit.
	FlushPeriodSec *float64 `default:"1" json:"flushPeriodSec"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputDynatraceOtlpFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// Amount of time (milliseconds) to wait for the connection to establish before retrying
	ConnectionTimeout *float64 `default:"10000" json:"connectionTimeout"`
	// How often the sender should ping the peer to keep the connection open
	KeepAliveTime *float64 `default:"30" json:"keepAliveTime"`
	// Disable to close the connection immediately after sending the outgoing request
	KeepAlive *bool `default:"true" json:"keepAlive"`
	// Select the type of Dynatrace endpoint configured
	EndpointType *EndpointType `default:"saas" json:"endpointType"`
	// Select or create a stored text secret
	TokenSecret   string  `json:"tokenSecret"`
	AuthTokenName *string `default:"Authorization" json:"authTokenName"`
	// How to handle events when all receivers are exerting backpressure
	OnBackpressure *OutputDynatraceOtlpBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                                  `json:"description,omitempty"`
	// Reject certificates not authorized by a CA in the CA certificate path or by another trusted CA (such as the system's).
	//         Enabled by default. When this setting is also present in TLS Settings (Client Side),
	//         that value will take precedence.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Enable round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Headers to add to all events
	ExtraHTTPHeaders []OutputDynatraceOtlpExtraHTTPHeader `json:"extraHttpHeaders,omitempty"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable)
	ResponseRetrySettings []OutputDynatraceOtlpResponseRetrySetting `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputDynatraceOtlpTimeoutRetrySettings  `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.)
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data
	PqCompress *OutputDynatraceOtlpPqCompressCompression `default:"none" json:"pqCompress"`
	// How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputDynatraceOtlpQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputDynatraceOtlpMode       `default:"error" json:"pqMode"`
	PqControls *OutputDynatraceOtlpPqControls `json:"pqControls,omitempty"`
}

func (o OutputDynatraceOtlp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputDynatraceOtlp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputDynatraceOtlp) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputDynatraceOtlp) GetType() *OutputDynatraceOtlpType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputDynatraceOtlp) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputDynatraceOtlp) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputDynatraceOtlp) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputDynatraceOtlp) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputDynatraceOtlp) GetProtocol() *OutputDynatraceOtlpProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *OutputDynatraceOtlp) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *OutputDynatraceOtlp) GetOtlpVersion() *OutputDynatraceOTLPOTLPVersion {
	if o == nil {
		return nil
	}
	return o.OtlpVersion
}

func (o *OutputDynatraceOtlp) GetCompress() *OutputDynatraceOtlpCompressCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputDynatraceOtlp) GetHTTPCompress() *OutputDynatraceOtlpHTTPCompressCompression {
	if o == nil {
		return nil
	}
	return o.HTTPCompress
}

func (o *OutputDynatraceOtlp) GetHTTPTracesEndpointOverride() *string {
	if o == nil {
		return nil
	}
	return o.HTTPTracesEndpointOverride
}

func (o *OutputDynatraceOtlp) GetHTTPMetricsEndpointOverride() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMetricsEndpointOverride
}

func (o *OutputDynatraceOtlp) GetHTTPLogsEndpointOverride() *string {
	if o == nil {
		return nil
	}
	return o.HTTPLogsEndpointOverride
}

func (o *OutputDynatraceOtlp) GetMetadata() []OutputDynatraceOtlpMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *OutputDynatraceOtlp) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputDynatraceOtlp) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputDynatraceOtlp) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputDynatraceOtlp) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputDynatraceOtlp) GetFailedRequestLoggingMode() *OutputDynatraceOtlpFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputDynatraceOtlp) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *OutputDynatraceOtlp) GetKeepAliveTime() *float64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTime
}

func (o *OutputDynatraceOtlp) GetKeepAlive() *bool {
	if o == nil {
		return nil
	}
	return o.KeepAlive
}

func (o *OutputDynatraceOtlp) GetEndpointType() *EndpointType {
	if o == nil {
		return nil
	}
	return o.EndpointType
}

func (o *OutputDynatraceOtlp) GetTokenSecret() string {
	if o == nil {
		return ""
	}
	return o.TokenSecret
}

func (o *OutputDynatraceOtlp) GetAuthTokenName() *string {
	if o == nil {
		return nil
	}
	return o.AuthTokenName
}

func (o *OutputDynatraceOtlp) GetOnBackpressure() *OutputDynatraceOtlpBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputDynatraceOtlp) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputDynatraceOtlp) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputDynatraceOtlp) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputDynatraceOtlp) GetExtraHTTPHeaders() []OutputDynatraceOtlpExtraHTTPHeader {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputDynatraceOtlp) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputDynatraceOtlp) GetResponseRetrySettings() []OutputDynatraceOtlpResponseRetrySetting {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputDynatraceOtlp) GetTimeoutRetrySettings() *OutputDynatraceOtlpTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputDynatraceOtlp) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputDynatraceOtlp) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputDynatraceOtlp) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputDynatraceOtlp) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputDynatraceOtlp) GetPqCompress() *OutputDynatraceOtlpPqCompressCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputDynatraceOtlp) GetPqOnBackpressure() *OutputDynatraceOtlpQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputDynatraceOtlp) GetPqMode() *OutputDynatraceOtlpMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputDynatraceOtlp) GetPqControls() *OutputDynatraceOtlpPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
