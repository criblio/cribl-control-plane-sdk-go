// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type InputSystemStateType string

const (
	InputSystemStateTypeSystemState InputSystemStateType = "system_state"
)

func (e InputSystemStateType) ToPointer() *InputSystemStateType {
	return &e
}
func (e *InputSystemStateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system_state":
		*e = InputSystemStateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemStateType: %v", v)
	}
}

type InputSystemStateConnection struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputSystemStateConnection) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSystemStateConnection) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputSystemStateMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputSystemStateMode string

const (
	InputSystemStateModeSmart  InputSystemStateMode = "smart"
	InputSystemStateModeAlways InputSystemStateMode = "always"
)

func (e InputSystemStateMode) ToPointer() *InputSystemStateMode {
	return &e
}
func (e *InputSystemStateMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputSystemStateMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemStateMode: %v", v)
	}
}

// InputSystemStateCompression - Codec to use to compress the persisted data
type InputSystemStateCompression string

const (
	InputSystemStateCompressionNone InputSystemStateCompression = "none"
	InputSystemStateCompressionGzip InputSystemStateCompression = "gzip"
)

func (e InputSystemStateCompression) ToPointer() *InputSystemStateCompression {
	return &e
}
func (e *InputSystemStateCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSystemStateCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemStateCompression: %v", v)
	}
}

type InputSystemStatePq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputSystemStateMode `default:"always" json:"mode"`
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
	Compress *InputSystemStateCompression `default:"none" json:"compress"`
}

func (i InputSystemStatePq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemStatePq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemStatePq) GetMode() *InputSystemStateMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSystemStatePq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputSystemStatePq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputSystemStatePq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputSystemStatePq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputSystemStatePq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputSystemStatePq) GetCompress() *InputSystemStateCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type InputSystemStateMetadatum struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputSystemStateMetadatum) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputSystemStateMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// HostsFile - Creates events based on entries collected from the hosts file
type HostsFile struct {
	Enable *bool `default:"true" json:"enable"`
}

func (h HostsFile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HostsFile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HostsFile) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// Interfaces - Creates events for each of the host’s network interfaces
type Interfaces struct {
	Enable *bool `default:"true" json:"enable"`
}

func (i Interfaces) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Interfaces) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Interfaces) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// DisksAndFileSystems - Creates events for physical disks, partitions, and file systems
type DisksAndFileSystems struct {
	Enable *bool `default:"true" json:"enable"`
}

func (d DisksAndFileSystems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DisksAndFileSystems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DisksAndFileSystems) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// HostInfo - Creates events based on the host system’s current state
type HostInfo struct {
	Enable *bool `default:"true" json:"enable"`
}

func (h HostInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HostInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HostInfo) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// InputSystemStateRoutes - Creates events based on entries collected from the host’s network routes
type InputSystemStateRoutes struct {
	Enable *bool `default:"true" json:"enable"`
}

func (i InputSystemStateRoutes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemStateRoutes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemStateRoutes) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// DNS - Creates events for DNS resolvers and search entries
type DNS struct {
	Enable *bool `default:"true" json:"enable"`
}

func (d DNS) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DNS) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DNS) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// UsersAndGroups - Creates events for local users and groups
type UsersAndGroups struct {
	Enable *bool `default:"true" json:"enable"`
}

func (u UsersAndGroups) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UsersAndGroups) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UsersAndGroups) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// Firewall - Creates events for Firewall rules entries
type Firewall struct {
	Enable *bool `default:"true" json:"enable"`
}

func (f Firewall) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *Firewall) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Firewall) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// Services - Creates events from the list of services
type Services struct {
	Enable *bool `default:"true" json:"enable"`
}

func (s Services) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Services) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Services) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// ListeningPorts - Creates events from list of listening ports
type ListeningPorts struct {
	Enable *bool `default:"true" json:"enable"`
}

func (l ListeningPorts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListeningPorts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListeningPorts) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

// LoggedInUsers - Creates events from list of logged-in users
type LoggedInUsers struct {
	Enable *bool `default:"true" json:"enable"`
}

func (l LoggedInUsers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggedInUsers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggedInUsers) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

type Collectors struct {
	// Creates events based on entries collected from the hosts file
	Hostsfile *HostsFile `json:"hostsfile,omitempty"`
	// Creates events for each of the host’s network interfaces
	Interfaces *Interfaces `json:"interfaces,omitempty"`
	// Creates events for physical disks, partitions, and file systems
	Disk *DisksAndFileSystems `json:"disk,omitempty"`
	// Creates events based on the host system’s current state
	Metadata *HostInfo `json:"metadata,omitempty"`
	// Creates events based on entries collected from the host’s network routes
	Routes *InputSystemStateRoutes `json:"routes,omitempty"`
	// Creates events for DNS resolvers and search entries
	DNS *DNS `json:"dns,omitempty"`
	// Creates events for local users and groups
	User *UsersAndGroups `json:"user,omitempty"`
	// Creates events for Firewall rules entries
	Firewall *Firewall `json:"firewall,omitempty"`
	// Creates events from the list of services
	Services *Services `json:"services,omitempty"`
	// Creates events from list of listening ports
	Ports *ListeningPorts `json:"ports,omitempty"`
	// Creates events from list of logged-in users
	LoginUsers *LoggedInUsers `json:"loginUsers,omitempty"`
}

func (o *Collectors) GetHostsfile() *HostsFile {
	if o == nil {
		return nil
	}
	return o.Hostsfile
}

func (o *Collectors) GetInterfaces() *Interfaces {
	if o == nil {
		return nil
	}
	return o.Interfaces
}

func (o *Collectors) GetDisk() *DisksAndFileSystems {
	if o == nil {
		return nil
	}
	return o.Disk
}

func (o *Collectors) GetMetadata() *HostInfo {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Collectors) GetRoutes() *InputSystemStateRoutes {
	if o == nil {
		return nil
	}
	return o.Routes
}

func (o *Collectors) GetDNS() *DNS {
	if o == nil {
		return nil
	}
	return o.DNS
}

func (o *Collectors) GetUser() *UsersAndGroups {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Collectors) GetFirewall() *Firewall {
	if o == nil {
		return nil
	}
	return o.Firewall
}

func (o *Collectors) GetServices() *Services {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *Collectors) GetPorts() *ListeningPorts {
	if o == nil {
		return nil
	}
	return o.Ports
}

func (o *Collectors) GetLoginUsers() *LoggedInUsers {
	if o == nil {
		return nil
	}
	return o.LoginUsers
}

type InputSystemStateDataCompressionFormat string

const (
	InputSystemStateDataCompressionFormatNone InputSystemStateDataCompressionFormat = "none"
	InputSystemStateDataCompressionFormatGzip InputSystemStateDataCompressionFormat = "gzip"
)

func (e InputSystemStateDataCompressionFormat) ToPointer() *InputSystemStateDataCompressionFormat {
	return &e
}
func (e *InputSystemStateDataCompressionFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSystemStateDataCompressionFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSystemStateDataCompressionFormat: %v", v)
	}
}

type InputSystemStatePersistence struct {
	// Spool metrics to disk for Cribl Edge and Search
	Enable *bool `default:"false" json:"enable"`
	// Time span for each file bucket
	TimeWindow *string `default:"10m" json:"timeWindow"`
	// Maximum disk space allowed to be consumed (examples: 420MB, 4GB). When limit is reached, older data will be deleted.
	MaxDataSize *string `default:"1GB" json:"maxDataSize"`
	// Maximum amount of time to retain data (examples: 2h, 4d). When limit is reached, older data will be deleted.
	MaxDataTime *string                                `default:"24h" json:"maxDataTime"`
	Compress    *InputSystemStateDataCompressionFormat `default:"none" json:"compress"`
	// Path to use to write metrics. Defaults to $CRIBL_HOME/state/system_state
	DestPath *string `default:"$CRIBL_HOME/state/system_state" json:"destPath"`
}

func (i InputSystemStatePersistence) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemStatePersistence) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemStatePersistence) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

func (o *InputSystemStatePersistence) GetTimeWindow() *string {
	if o == nil {
		return nil
	}
	return o.TimeWindow
}

func (o *InputSystemStatePersistence) GetMaxDataSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataSize
}

func (o *InputSystemStatePersistence) GetMaxDataTime() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataTime
}

func (o *InputSystemStatePersistence) GetCompress() *InputSystemStateDataCompressionFormat {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputSystemStatePersistence) GetDestPath() *string {
	if o == nil {
		return nil
	}
	return o.DestPath
}

type InputSystemState struct {
	// Unique ID for this input
	ID       string               `json:"id"`
	Type     InputSystemStateType `json:"type"`
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
	Connections []InputSystemStateConnection `json:"connections,omitempty"`
	Pq          *InputSystemStatePq          `json:"pq,omitempty"`
	// Time, in seconds, between consecutive state collections. Default is 300 seconds (5 minutes).
	Interval *float64 `default:"300" json:"interval"`
	// Fields to add to events from this input
	Metadata    []InputSystemStateMetadatum  `json:"metadata,omitempty"`
	Collectors  *Collectors                  `json:"collectors,omitempty"`
	Persistence *InputSystemStatePersistence `json:"persistence,omitempty"`
	// Enable to use built-in tools (PowerShell) to collect events instead of native API (default) [Learn more](https://docs.cribl.io/edge/sources-system-state/#advanced-tab)
	DisableNativeModule *bool   `default:"false" json:"disableNativeModule"`
	Description         *string `json:"description,omitempty"`
}

func (i InputSystemState) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSystemState) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSystemState) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputSystemState) GetType() InputSystemStateType {
	if o == nil {
		return InputSystemStateType("")
	}
	return o.Type
}

func (o *InputSystemState) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSystemState) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSystemState) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputSystemState) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputSystemState) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputSystemState) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputSystemState) GetConnections() []InputSystemStateConnection {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputSystemState) GetPq() *InputSystemStatePq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputSystemState) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputSystemState) GetMetadata() []InputSystemStateMetadatum {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputSystemState) GetCollectors() *Collectors {
	if o == nil {
		return nil
	}
	return o.Collectors
}

func (o *InputSystemState) GetPersistence() *InputSystemStatePersistence {
	if o == nil {
		return nil
	}
	return o.Persistence
}

func (o *InputSystemState) GetDisableNativeModule() *bool {
	if o == nil {
		return nil
	}
	return o.DisableNativeModule
}

func (o *InputSystemState) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
