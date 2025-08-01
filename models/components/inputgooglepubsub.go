// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type InputGooglePubsubType string

const (
	InputGooglePubsubTypeGooglePubsub InputGooglePubsubType = "google_pubsub"
)

func (e InputGooglePubsubType) ToPointer() *InputGooglePubsubType {
	return &e
}
func (e *InputGooglePubsubType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google_pubsub":
		*e = InputGooglePubsubType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputGooglePubsubType: %v", v)
	}
}

type InputGooglePubsubConnection struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputGooglePubsubConnection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputGooglePubsubConnection) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputGooglePubsubMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputGooglePubsubMode string

const (
	InputGooglePubsubModeSmart  InputGooglePubsubMode = "smart"
	InputGooglePubsubModeAlways InputGooglePubsubMode = "always"
)

func (e InputGooglePubsubMode) ToPointer() *InputGooglePubsubMode {
	return &e
}
func (e *InputGooglePubsubMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputGooglePubsubMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputGooglePubsubMode: %v", v)
	}
}

// InputGooglePubsubCompression - Codec to use to compress the persisted data
type InputGooglePubsubCompression string

const (
	InputGooglePubsubCompressionNone InputGooglePubsubCompression = "none"
	InputGooglePubsubCompressionGzip InputGooglePubsubCompression = "gzip"
)

func (e InputGooglePubsubCompression) ToPointer() *InputGooglePubsubCompression {
	return &e
}
func (e *InputGooglePubsubCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputGooglePubsubCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputGooglePubsubCompression: %v", v)
	}
}

type InputGooglePubsubPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputGooglePubsubMode `default:"always" json:"mode"`
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
	Compress *InputGooglePubsubCompression `default:"none" json:"compress"`
}

func (i InputGooglePubsubPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputGooglePubsubPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputGooglePubsubPq) GetMode() *InputGooglePubsubMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputGooglePubsubPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputGooglePubsubPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputGooglePubsubPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputGooglePubsubPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputGooglePubsubPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputGooglePubsubPq) GetCompress() *InputGooglePubsubCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputGooglePubsubGoogleAuthenticationMethod - Choose Auto to use Google Application Default Credentials (ADC), Manual to enter Google service account credentials directly, or Secret to select or create a stored secret that references Google service account credentials.
type InputGooglePubsubGoogleAuthenticationMethod string

const (
	InputGooglePubsubGoogleAuthenticationMethodAuto   InputGooglePubsubGoogleAuthenticationMethod = "auto"
	InputGooglePubsubGoogleAuthenticationMethodManual InputGooglePubsubGoogleAuthenticationMethod = "manual"
	InputGooglePubsubGoogleAuthenticationMethodSecret InputGooglePubsubGoogleAuthenticationMethod = "secret"
)

func (e InputGooglePubsubGoogleAuthenticationMethod) ToPointer() *InputGooglePubsubGoogleAuthenticationMethod {
	return &e
}
func (e *InputGooglePubsubGoogleAuthenticationMethod) UnmarshalJSON(data []byte) error {
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
		*e = InputGooglePubsubGoogleAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputGooglePubsubGoogleAuthenticationMethod: %v", v)
	}
}

type InputGooglePubsubMetadatum struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputGooglePubsubMetadatum) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputGooglePubsubMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputGooglePubsub struct {
	// Unique ID for this input
	ID       *string                `json:"id,omitempty"`
	Type     *InputGooglePubsubType `json:"type,omitempty"`
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
	Connections []InputGooglePubsubConnection `json:"connections,omitempty"`
	Pq          *InputGooglePubsubPq          `json:"pq,omitempty"`
	// ID of the topic to receive events from
	TopicName string `json:"topicName"`
	// ID of the subscription to use when receiving events
	SubscriptionName string `json:"subscriptionName"`
	// Create topic if it does not exist
	CreateTopic *bool `default:"false" json:"createTopic"`
	// Create subscription if it does not exist
	CreateSubscription *bool `default:"true" json:"createSubscription"`
	// Region to retrieve messages from. Select 'default' to allow Google to auto-select the nearest region. When using ordered delivery, the selected region must be allowed by message storage policy.
	Region *string `json:"region,omitempty"`
	// Choose Auto to use Google Application Default Credentials (ADC), Manual to enter Google service account credentials directly, or Secret to select or create a stored secret that references Google service account credentials.
	GoogleAuthMethod *InputGooglePubsubGoogleAuthenticationMethod `default:"manual" json:"googleAuthMethod"`
	// Contents of service account credentials (JSON keys) file downloaded from Google Cloud. To upload a file, click the upload button at this field's upper right.
	ServiceAccountCredentials *string `json:"serviceAccountCredentials,omitempty"`
	// Select or create a stored text secret
	Secret *string `json:"secret,omitempty"`
	// If Destination exerts backpressure, this setting limits how many inbound events Stream will queue for processing before it stops retrieving events
	MaxBacklog *float64 `default:"1000" json:"maxBacklog"`
	// How many streams to pull messages from at one time. Doubling the value doubles the number of messages this Source pulls from the topic (if available), while consuming more CPU and memory. Defaults to 5.
	Concurrency *float64 `default:"5" json:"concurrency"`
	// Pull request timeout, in milliseconds
	RequestTimeout *float64 `default:"60000" json:"requestTimeout"`
	// Fields to add to events from this input
	Metadata    []InputGooglePubsubMetadatum `json:"metadata,omitempty"`
	Description *string                      `json:"description,omitempty"`
	// Receive events in the order they were added to the queue. The process sending events must have ordering enabled.
	OrderedDelivery *bool `default:"false" json:"orderedDelivery"`
}

func (i InputGooglePubsub) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputGooglePubsub) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputGooglePubsub) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputGooglePubsub) GetType() *InputGooglePubsubType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputGooglePubsub) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputGooglePubsub) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputGooglePubsub) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputGooglePubsub) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputGooglePubsub) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputGooglePubsub) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputGooglePubsub) GetConnections() []InputGooglePubsubConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputGooglePubsub) GetPq() *InputGooglePubsubPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputGooglePubsub) GetTopicName() string {
	if o == nil {
		return ""
	}
	return o.TopicName
}

func (o *InputGooglePubsub) GetSubscriptionName() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionName
}

func (o *InputGooglePubsub) GetCreateTopic() *bool {
	if o == nil {
		return nil
	}
	return o.CreateTopic
}

func (o *InputGooglePubsub) GetCreateSubscription() *bool {
	if o == nil {
		return nil
	}
	return o.CreateSubscription
}

func (o *InputGooglePubsub) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *InputGooglePubsub) GetGoogleAuthMethod() *InputGooglePubsubGoogleAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.GoogleAuthMethod
}

func (o *InputGooglePubsub) GetServiceAccountCredentials() *string {
	if o == nil {
		return nil
	}
	return o.ServiceAccountCredentials
}

func (o *InputGooglePubsub) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *InputGooglePubsub) GetMaxBacklog() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBacklog
}

func (o *InputGooglePubsub) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *InputGooglePubsub) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputGooglePubsub) GetMetadata() []InputGooglePubsubMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputGooglePubsub) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputGooglePubsub) GetOrderedDelivery() *bool {
	if o == nil {
		return nil
	}
	return o.OrderedDelivery
}
