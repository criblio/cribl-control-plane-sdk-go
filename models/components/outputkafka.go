// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type OutputKafkaType string

const (
	OutputKafkaTypeKafka OutputKafkaType = "kafka"
)

func (e OutputKafkaType) ToPointer() *OutputKafkaType {
	return &e
}
func (e *OutputKafkaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kafka":
		*e = OutputKafkaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaType: %v", v)
	}
}

// OutputKafkaAcknowledgments - Control the number of required acknowledgments.
type OutputKafkaAcknowledgments int64

const (
	OutputKafkaAcknowledgmentsOne    OutputKafkaAcknowledgments = 1
	OutputKafkaAcknowledgmentsZero   OutputKafkaAcknowledgments = 0
	OutputKafkaAcknowledgmentsMinus1 OutputKafkaAcknowledgments = -1
)

func (e OutputKafkaAcknowledgments) ToPointer() *OutputKafkaAcknowledgments {
	return &e
}
func (e *OutputKafkaAcknowledgments) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 0:
		fallthrough
	case -1:
		*e = OutputKafkaAcknowledgments(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaAcknowledgments: %v", v)
	}
}

// OutputKafkaRecordDataFormat - Format to use to serialize events before writing to Kafka.
type OutputKafkaRecordDataFormat string

const (
	OutputKafkaRecordDataFormatJSON     OutputKafkaRecordDataFormat = "json"
	OutputKafkaRecordDataFormatRaw      OutputKafkaRecordDataFormat = "raw"
	OutputKafkaRecordDataFormatProtobuf OutputKafkaRecordDataFormat = "protobuf"
)

func (e OutputKafkaRecordDataFormat) ToPointer() *OutputKafkaRecordDataFormat {
	return &e
}
func (e *OutputKafkaRecordDataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "raw":
		fallthrough
	case "protobuf":
		*e = OutputKafkaRecordDataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaRecordDataFormat: %v", v)
	}
}

// OutputKafkaCompression - Codec to use to compress the data before sending to Kafka
type OutputKafkaCompression string

const (
	OutputKafkaCompressionNone   OutputKafkaCompression = "none"
	OutputKafkaCompressionGzip   OutputKafkaCompression = "gzip"
	OutputKafkaCompressionSnappy OutputKafkaCompression = "snappy"
	OutputKafkaCompressionLz4    OutputKafkaCompression = "lz4"
)

func (e OutputKafkaCompression) ToPointer() *OutputKafkaCompression {
	return &e
}
func (e *OutputKafkaCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		fallthrough
	case "snappy":
		fallthrough
	case "lz4":
		*e = OutputKafkaCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaCompression: %v", v)
	}
}

// OutputKafkaAuth - Credentials to use when authenticating with the schema registry using basic HTTP authentication
type OutputKafkaAuth struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
}

func (o OutputKafkaAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputKafkaAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputKafkaAuth) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputKafkaAuth) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

type OutputKafkaKafkaSchemaRegistryMinimumTLSVersion string

const (
	OutputKafkaKafkaSchemaRegistryMinimumTLSVersionTlSv1  OutputKafkaKafkaSchemaRegistryMinimumTLSVersion = "TLSv1"
	OutputKafkaKafkaSchemaRegistryMinimumTLSVersionTlSv11 OutputKafkaKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.1"
	OutputKafkaKafkaSchemaRegistryMinimumTLSVersionTlSv12 OutputKafkaKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.2"
	OutputKafkaKafkaSchemaRegistryMinimumTLSVersionTlSv13 OutputKafkaKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.3"
)

func (e OutputKafkaKafkaSchemaRegistryMinimumTLSVersion) ToPointer() *OutputKafkaKafkaSchemaRegistryMinimumTLSVersion {
	return &e
}
func (e *OutputKafkaKafkaSchemaRegistryMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputKafkaKafkaSchemaRegistryMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaKafkaSchemaRegistryMinimumTLSVersion: %v", v)
	}
}

type OutputKafkaKafkaSchemaRegistryMaximumTLSVersion string

const (
	OutputKafkaKafkaSchemaRegistryMaximumTLSVersionTlSv1  OutputKafkaKafkaSchemaRegistryMaximumTLSVersion = "TLSv1"
	OutputKafkaKafkaSchemaRegistryMaximumTLSVersionTlSv11 OutputKafkaKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.1"
	OutputKafkaKafkaSchemaRegistryMaximumTLSVersionTlSv12 OutputKafkaKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.2"
	OutputKafkaKafkaSchemaRegistryMaximumTLSVersionTlSv13 OutputKafkaKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.3"
)

func (e OutputKafkaKafkaSchemaRegistryMaximumTLSVersion) ToPointer() *OutputKafkaKafkaSchemaRegistryMaximumTLSVersion {
	return &e
}
func (e *OutputKafkaKafkaSchemaRegistryMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputKafkaKafkaSchemaRegistryMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaKafkaSchemaRegistryMaximumTLSVersion: %v", v)
	}
}

type OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Reject certificates that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (such as the system's). Defaults to Enabled. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
	// The name of the predefined certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Passphrase to use to decrypt private key
	Passphrase *string                                          `json:"passphrase,omitempty"`
	MinVersion *OutputKafkaKafkaSchemaRegistryMinimumTLSVersion `json:"minVersion,omitempty"`
	MaxVersion *OutputKafkaKafkaSchemaRegistryMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (o OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetMinVersion() *OutputKafkaKafkaSchemaRegistryMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide) GetMaxVersion() *OutputKafkaKafkaSchemaRegistryMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type OutputKafkaKafkaSchemaRegistryAuthentication struct {
	Disabled *bool `default:"true" json:"disabled"`
	// URL for accessing the Confluent Schema Registry. Example: http://localhost:8081. To connect over TLS, use https instead of http.
	SchemaRegistryURL *string `default:"http://localhost:8081" json:"schemaRegistryURL"`
	// Maximum time to wait for a Schema Registry connection to complete successfully
	ConnectionTimeout *float64 `default:"30000" json:"connectionTimeout"`
	// Maximum time to wait for the Schema Registry to respond to a request
	RequestTimeout *float64 `default:"30000" json:"requestTimeout"`
	// Maximum number of times to try fetching schemas from the Schema Registry
	MaxRetries *float64 `default:"1" json:"maxRetries"`
	// Credentials to use when authenticating with the schema registry using basic HTTP authentication
	Auth *OutputKafkaAuth                                     `json:"auth,omitempty"`
	TLS  *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide `json:"tls,omitempty"`
	// Used when __keySchemaIdOut is not present, to transform key values, leave blank if key transformation is not required by default.
	DefaultKeySchemaID *float64 `json:"defaultKeySchemaId,omitempty"`
	// Used when __valueSchemaIdOut is not present, to transform _raw, leave blank if value transformation is not required by default.
	DefaultValueSchemaID *float64 `json:"defaultValueSchemaId,omitempty"`
}

func (o OutputKafkaKafkaSchemaRegistryAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetSchemaRegistryURL() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryURL
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetMaxRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetries
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetAuth() *OutputKafkaAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetTLS() *OutputKafkaKafkaSchemaRegistryTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetDefaultKeySchemaID() *float64 {
	if o == nil {
		return nil
	}
	return o.DefaultKeySchemaID
}

func (o *OutputKafkaKafkaSchemaRegistryAuthentication) GetDefaultValueSchemaID() *float64 {
	if o == nil {
		return nil
	}
	return o.DefaultValueSchemaID
}

type OutputKafkaSASLMechanism string

const (
	OutputKafkaSASLMechanismPlain       OutputKafkaSASLMechanism = "plain"
	OutputKafkaSASLMechanismScramSha256 OutputKafkaSASLMechanism = "scram-sha-256"
	OutputKafkaSASLMechanismScramSha512 OutputKafkaSASLMechanism = "scram-sha-512"
	OutputKafkaSASLMechanismKerberos    OutputKafkaSASLMechanism = "kerberos"
)

func (e OutputKafkaSASLMechanism) ToPointer() *OutputKafkaSASLMechanism {
	return &e
}
func (e *OutputKafkaSASLMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plain":
		fallthrough
	case "scram-sha-256":
		fallthrough
	case "scram-sha-512":
		fallthrough
	case "kerberos":
		*e = OutputKafkaSASLMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaSASLMechanism: %v", v)
	}
}

// OutputKafkaAuthentication - Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
type OutputKafkaAuthentication struct {
	Disabled  *bool                     `default:"true" json:"disabled"`
	Mechanism *OutputKafkaSASLMechanism `default:"plain" json:"mechanism"`
}

func (o OutputKafkaAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputKafkaAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputKafkaAuthentication) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputKafkaAuthentication) GetMechanism() *OutputKafkaSASLMechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}

type OutputKafkaMinimumTLSVersion string

const (
	OutputKafkaMinimumTLSVersionTlSv1  OutputKafkaMinimumTLSVersion = "TLSv1"
	OutputKafkaMinimumTLSVersionTlSv11 OutputKafkaMinimumTLSVersion = "TLSv1.1"
	OutputKafkaMinimumTLSVersionTlSv12 OutputKafkaMinimumTLSVersion = "TLSv1.2"
	OutputKafkaMinimumTLSVersionTlSv13 OutputKafkaMinimumTLSVersion = "TLSv1.3"
)

func (e OutputKafkaMinimumTLSVersion) ToPointer() *OutputKafkaMinimumTLSVersion {
	return &e
}
func (e *OutputKafkaMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputKafkaMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaMinimumTLSVersion: %v", v)
	}
}

type OutputKafkaMaximumTLSVersion string

const (
	OutputKafkaMaximumTLSVersionTlSv1  OutputKafkaMaximumTLSVersion = "TLSv1"
	OutputKafkaMaximumTLSVersionTlSv11 OutputKafkaMaximumTLSVersion = "TLSv1.1"
	OutputKafkaMaximumTLSVersionTlSv12 OutputKafkaMaximumTLSVersion = "TLSv1.2"
	OutputKafkaMaximumTLSVersionTlSv13 OutputKafkaMaximumTLSVersion = "TLSv1.3"
)

func (e OutputKafkaMaximumTLSVersion) ToPointer() *OutputKafkaMaximumTLSVersion {
	return &e
}
func (e *OutputKafkaMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputKafkaMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaMaximumTLSVersion: %v", v)
	}
}

type OutputKafkaTLSSettingsClientSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Reject certificates that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (such as the system's). Defaults to Enabled. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
	// The name of the predefined certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Passphrase to use to decrypt private key
	Passphrase *string                       `json:"passphrase,omitempty"`
	MinVersion *OutputKafkaMinimumTLSVersion `json:"minVersion,omitempty"`
	MaxVersion *OutputKafkaMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (o OutputKafkaTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputKafkaTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputKafkaTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputKafkaTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputKafkaTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *OutputKafkaTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *OutputKafkaTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *OutputKafkaTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *OutputKafkaTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *OutputKafkaTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *OutputKafkaTLSSettingsClientSide) GetMinVersion() *OutputKafkaMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *OutputKafkaTLSSettingsClientSide) GetMaxVersion() *OutputKafkaMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

// OutputKafkaBackpressureBehavior - How to handle events when all receivers are exerting backpressure
type OutputKafkaBackpressureBehavior string

const (
	OutputKafkaBackpressureBehaviorBlock OutputKafkaBackpressureBehavior = "block"
	OutputKafkaBackpressureBehaviorDrop  OutputKafkaBackpressureBehavior = "drop"
	OutputKafkaBackpressureBehaviorQueue OutputKafkaBackpressureBehavior = "queue"
)

func (e OutputKafkaBackpressureBehavior) ToPointer() *OutputKafkaBackpressureBehavior {
	return &e
}
func (e *OutputKafkaBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputKafkaBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaBackpressureBehavior: %v", v)
	}
}

// OutputKafkaPqCompressCompression - Codec to use to compress the persisted data
type OutputKafkaPqCompressCompression string

const (
	OutputKafkaPqCompressCompressionNone OutputKafkaPqCompressCompression = "none"
	OutputKafkaPqCompressCompressionGzip OutputKafkaPqCompressCompression = "gzip"
)

func (e OutputKafkaPqCompressCompression) ToPointer() *OutputKafkaPqCompressCompression {
	return &e
}
func (e *OutputKafkaPqCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputKafkaPqCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaPqCompressCompression: %v", v)
	}
}

// OutputKafkaQueueFullBehavior - How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputKafkaQueueFullBehavior string

const (
	OutputKafkaQueueFullBehaviorBlock OutputKafkaQueueFullBehavior = "block"
	OutputKafkaQueueFullBehaviorDrop  OutputKafkaQueueFullBehavior = "drop"
)

func (e OutputKafkaQueueFullBehavior) ToPointer() *OutputKafkaQueueFullBehavior {
	return &e
}
func (e *OutputKafkaQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputKafkaQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaQueueFullBehavior: %v", v)
	}
}

// OutputKafkaMode - In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.
type OutputKafkaMode string

const (
	OutputKafkaModeError        OutputKafkaMode = "error"
	OutputKafkaModeBackpressure OutputKafkaMode = "backpressure"
	OutputKafkaModeAlways       OutputKafkaMode = "always"
)

func (e OutputKafkaMode) ToPointer() *OutputKafkaMode {
	return &e
}
func (e *OutputKafkaMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputKafkaMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKafkaMode: %v", v)
	}
}

type OutputKafkaPqControls struct {
}

type OutputKafka struct {
	// Unique ID for this output
	ID   *string          `json:"id,omitempty"`
	Type *OutputKafkaType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Enter each Kafka bootstrap server you want to use. Specify hostname and port, e.g., mykafkabroker:9092, or just hostname, in which case @{product} will assign port 9092.
	Brokers []string `json:"brokers"`
	// The topic to publish events to. Can be overridden using the __topicOut field.
	Topic string `json:"topic"`
	// Control the number of required acknowledgments.
	Ack *OutputKafkaAcknowledgments `default:"1" json:"ack"`
	// Format to use to serialize events before writing to Kafka.
	Format *OutputKafkaRecordDataFormat `default:"json" json:"format"`
	// Codec to use to compress the data before sending to Kafka
	Compression *OutputKafkaCompression `default:"gzip" json:"compression"`
	// Maximum size of each record batch before compression. The value must not exceed the Kafka brokers' message.max.bytes setting.
	MaxRecordSizeKB *float64 `default:"768" json:"maxRecordSizeKB"`
	// The maximum number of events you want the Destination to allow in a batch before forcing a flush
	FlushEventCount *float64 `default:"1000" json:"flushEventCount"`
	// The maximum amount of time you want the Destination to wait before forcing a flush. Shorter intervals tend to result in smaller batches being sent.
	FlushPeriodSec      *float64                                      `default:"1" json:"flushPeriodSec"`
	KafkaSchemaRegistry *OutputKafkaKafkaSchemaRegistryAuthentication `json:"kafkaSchemaRegistry,omitempty"`
	// Maximum time to wait for a connection to complete successfully
	ConnectionTimeout *float64 `default:"10000" json:"connectionTimeout"`
	// Maximum time to wait for Kafka to respond to a request
	RequestTimeout *float64 `default:"60000" json:"requestTimeout"`
	// If messages are failing, you can set the maximum number of retries as high as 100 to prevent loss of data
	MaxRetries *float64 `default:"5" json:"maxRetries"`
	// The maximum wait time for a retry, in milliseconds. Default (and minimum) is 30,000 ms (30 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackOff *float64 `default:"30000" json:"maxBackOff"`
	// Initial value used to calculate the retry, in milliseconds. Maximum is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"300" json:"initialBackoff"`
	// Set the backoff multiplier (2-20) to control the retry frequency for failed messages. For faster retries, use a lower multiplier. For slower retries with more delay between attempts, use a higher multiplier. The multiplier is used in an exponential backoff formula; see the Kafka [documentation](https://kafka.js.org/docs/retry-detailed) for details.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// Maximum time to wait for Kafka to respond to an authentication request
	AuthenticationTimeout *float64 `default:"10000" json:"authenticationTimeout"`
	// Specifies a time window during which @{product} can reauthenticate if needed. Creates the window measuring backward from the moment when credentials are set to expire.
	ReauthenticationThreshold *float64 `default:"10000" json:"reauthenticationThreshold"`
	// Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
	Sasl *OutputKafkaAuthentication        `json:"sasl,omitempty"`
	TLS  *OutputKafkaTLSSettingsClientSide `json:"tls,omitempty"`
	// How to handle events when all receivers are exerting backpressure
	OnBackpressure *OutputKafkaBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                          `json:"description,omitempty"`
	// Select a set of Protobuf definitions for the events you want to send
	ProtobufLibraryID *string `json:"protobufLibraryId,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.)
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data
	PqCompress *OutputKafkaPqCompressCompression `default:"none" json:"pqCompress"`
	// How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputKafkaQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputKafkaMode       `default:"error" json:"pqMode"`
	PqControls *OutputKafkaPqControls `json:"pqControls,omitempty"`
}

func (o OutputKafka) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputKafka) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputKafka) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputKafka) GetType() *OutputKafkaType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputKafka) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputKafka) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputKafka) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputKafka) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputKafka) GetBrokers() []string {
	if o == nil {
		return []string{}
	}
	return o.Brokers
}

func (o *OutputKafka) GetTopic() string {
	if o == nil {
		return ""
	}
	return o.Topic
}

func (o *OutputKafka) GetAck() *OutputKafkaAcknowledgments {
	if o == nil {
		return nil
	}
	return o.Ack
}

func (o *OutputKafka) GetFormat() *OutputKafkaRecordDataFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputKafka) GetCompression() *OutputKafkaCompression {
	if o == nil {
		return nil
	}
	return o.Compression
}

func (o *OutputKafka) GetMaxRecordSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRecordSizeKB
}

func (o *OutputKafka) GetFlushEventCount() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushEventCount
}

func (o *OutputKafka) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputKafka) GetKafkaSchemaRegistry() *OutputKafkaKafkaSchemaRegistryAuthentication {
	if o == nil {
		return nil
	}
	return o.KafkaSchemaRegistry
}

func (o *OutputKafka) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *OutputKafka) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *OutputKafka) GetMaxRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetries
}

func (o *OutputKafka) GetMaxBackOff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackOff
}

func (o *OutputKafka) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputKafka) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputKafka) GetAuthenticationTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.AuthenticationTimeout
}

func (o *OutputKafka) GetReauthenticationThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.ReauthenticationThreshold
}

func (o *OutputKafka) GetSasl() *OutputKafkaAuthentication {
	if o == nil {
		return nil
	}
	return o.Sasl
}

func (o *OutputKafka) GetTLS() *OutputKafkaTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *OutputKafka) GetOnBackpressure() *OutputKafkaBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputKafka) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputKafka) GetProtobufLibraryID() *string {
	if o == nil {
		return nil
	}
	return o.ProtobufLibraryID
}

func (o *OutputKafka) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputKafka) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputKafka) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputKafka) GetPqCompress() *OutputKafkaPqCompressCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputKafka) GetPqOnBackpressure() *OutputKafkaQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputKafka) GetPqMode() *OutputKafkaMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputKafka) GetPqControls() *OutputKafkaPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
