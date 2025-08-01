// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type InputExecType string

const (
	InputExecTypeExec InputExecType = "exec"
)

func (e InputExecType) ToPointer() *InputExecType {
	return &e
}
func (e *InputExecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "exec":
		*e = InputExecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputExecType: %v", v)
	}
}

type InputExecConnection struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputExecConnection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputExecConnection) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputExecMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputExecMode string

const (
	InputExecModeSmart  InputExecMode = "smart"
	InputExecModeAlways InputExecMode = "always"
)

func (e InputExecMode) ToPointer() *InputExecMode {
	return &e
}
func (e *InputExecMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputExecMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputExecMode: %v", v)
	}
}

// InputExecCompression - Codec to use to compress the persisted data
type InputExecCompression string

const (
	InputExecCompressionNone InputExecCompression = "none"
	InputExecCompressionGzip InputExecCompression = "gzip"
)

func (e InputExecCompression) ToPointer() *InputExecCompression {
	return &e
}
func (e *InputExecCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputExecCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputExecCompression: %v", v)
	}
}

type InputExecPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputExecMode `default:"always" json:"mode"`
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
	Compress *InputExecCompression `default:"none" json:"compress"`
}

func (i InputExecPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputExecPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputExecPq) GetMode() *InputExecMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputExecPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputExecPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputExecPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputExecPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputExecPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputExecPq) GetCompress() *InputExecCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// ScheduleType - Select a schedule type; either an interval (in seconds) or a cron-style schedule.
type ScheduleType string

const (
	ScheduleTypeInterval     ScheduleType = "interval"
	ScheduleTypeCronSchedule ScheduleType = "cronSchedule"
)

func (e ScheduleType) ToPointer() *ScheduleType {
	return &e
}
func (e *ScheduleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "interval":
		fallthrough
	case "cronSchedule":
		*e = ScheduleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleType: %v", v)
	}
}

type InputExecMetadatum struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputExecMetadatum) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputExecMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputExec struct {
	// Unique ID for this input
	ID       *string       `json:"id,omitempty"`
	Type     InputExecType `json:"type"`
	Disabled *bool         `default:"false" json:"disabled"`
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
	Connections []InputExecConnection `json:"connections,omitempty"`
	Pq          *InputExecPq          `json:"pq,omitempty"`
	// Command to execute; supports Bourne shell (or CMD on Windows) syntax
	Command string `json:"command"`
	// Maximum number of retry attempts in the event that the command fails
	Retries *float64 `default:"10" json:"retries"`
	// Select a schedule type; either an interval (in seconds) or a cron-style schedule.
	ScheduleType *ScheduleType `default:"interval" json:"scheduleType"`
	// A list of event-breaking rulesets that will be applied, in order, to the input data stream
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// How long (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel before flushing the data stream out, as is, to the Pipelines
	StaleChannelFlushMs *float64 `default:"10000" json:"staleChannelFlushMs"`
	// Fields to add to events from this input
	Metadata    []InputExecMetadatum `json:"metadata,omitempty"`
	Description *string              `json:"description,omitempty"`
	// Interval between command executions in seconds.
	Interval *float64 `default:"60" json:"interval"`
	// Cron schedule to execute the command on.
	CronSchedule *string `default:"* * * * *" json:"cronSchedule"`
}

func (i InputExec) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputExec) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputExec) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputExec) GetType() InputExecType {
	if o == nil {
		return InputExecType("")
	}
	return o.Type
}

func (o *InputExec) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputExec) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputExec) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputExec) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputExec) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputExec) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputExec) GetConnections() []InputExecConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputExec) GetPq() *InputExecPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputExec) GetCommand() string {
	if o == nil {
		return ""
	}
	return o.Command
}

func (o *InputExec) GetRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

func (o *InputExec) GetScheduleType() *ScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *InputExec) GetBreakerRulesets() []string {
	if o == nil {
		return nil
	}
	return o.BreakerRulesets
}

func (o *InputExec) GetStaleChannelFlushMs() *float64 {
	if o == nil {
		return nil
	}
	return o.StaleChannelFlushMs
}

func (o *InputExec) GetMetadata() []InputExecMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputExec) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputExec) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputExec) GetCronSchedule() *string {
	if o == nil {
		return nil
	}
	return o.CronSchedule
}
