// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type OutputSecurityLakeType string

const (
	OutputSecurityLakeTypeSecurityLake OutputSecurityLakeType = "security_lake"
)

func (e OutputSecurityLakeType) ToPointer() *OutputSecurityLakeType {
	return &e
}
func (e *OutputSecurityLakeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "security_lake":
		*e = OutputSecurityLakeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeType: %v", v)
	}
}

// OutputSecurityLakeAuthenticationMethod - AWS authentication method. Choose Auto to use IAM roles.
type OutputSecurityLakeAuthenticationMethod string

const (
	OutputSecurityLakeAuthenticationMethodAuto   OutputSecurityLakeAuthenticationMethod = "auto"
	OutputSecurityLakeAuthenticationMethodManual OutputSecurityLakeAuthenticationMethod = "manual"
	OutputSecurityLakeAuthenticationMethodSecret OutputSecurityLakeAuthenticationMethod = "secret"
)

func (e OutputSecurityLakeAuthenticationMethod) ToPointer() *OutputSecurityLakeAuthenticationMethod {
	return &e
}
func (e *OutputSecurityLakeAuthenticationMethod) UnmarshalJSON(data []byte) error {
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
		*e = OutputSecurityLakeAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeAuthenticationMethod: %v", v)
	}
}

// OutputSecurityLakeSignatureVersion - Signature version to use for signing Amazon Security Lake requests
type OutputSecurityLakeSignatureVersion string

const (
	OutputSecurityLakeSignatureVersionV2 OutputSecurityLakeSignatureVersion = "v2"
	OutputSecurityLakeSignatureVersionV4 OutputSecurityLakeSignatureVersion = "v4"
)

func (e OutputSecurityLakeSignatureVersion) ToPointer() *OutputSecurityLakeSignatureVersion {
	return &e
}
func (e *OutputSecurityLakeSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = OutputSecurityLakeSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeSignatureVersion: %v", v)
	}
}

// OutputSecurityLakeObjectACL - Object ACL to assign to uploaded objects
type OutputSecurityLakeObjectACL string

const (
	OutputSecurityLakeObjectACLPrivate                OutputSecurityLakeObjectACL = "private"
	OutputSecurityLakeObjectACLPublicRead             OutputSecurityLakeObjectACL = "public-read"
	OutputSecurityLakeObjectACLPublicReadWrite        OutputSecurityLakeObjectACL = "public-read-write"
	OutputSecurityLakeObjectACLAuthenticatedRead      OutputSecurityLakeObjectACL = "authenticated-read"
	OutputSecurityLakeObjectACLAwsExecRead            OutputSecurityLakeObjectACL = "aws-exec-read"
	OutputSecurityLakeObjectACLBucketOwnerRead        OutputSecurityLakeObjectACL = "bucket-owner-read"
	OutputSecurityLakeObjectACLBucketOwnerFullControl OutputSecurityLakeObjectACL = "bucket-owner-full-control"
)

func (e OutputSecurityLakeObjectACL) ToPointer() *OutputSecurityLakeObjectACL {
	return &e
}
func (e *OutputSecurityLakeObjectACL) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "private":
		fallthrough
	case "public-read":
		fallthrough
	case "public-read-write":
		fallthrough
	case "authenticated-read":
		fallthrough
	case "aws-exec-read":
		fallthrough
	case "bucket-owner-read":
		fallthrough
	case "bucket-owner-full-control":
		*e = OutputSecurityLakeObjectACL(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeObjectACL: %v", v)
	}
}

// OutputSecurityLakeStorageClass - Storage class to select for uploaded objects
type OutputSecurityLakeStorageClass string

const (
	OutputSecurityLakeStorageClassStandard           OutputSecurityLakeStorageClass = "STANDARD"
	OutputSecurityLakeStorageClassReducedRedundancy  OutputSecurityLakeStorageClass = "REDUCED_REDUNDANCY"
	OutputSecurityLakeStorageClassStandardIa         OutputSecurityLakeStorageClass = "STANDARD_IA"
	OutputSecurityLakeStorageClassOnezoneIa          OutputSecurityLakeStorageClass = "ONEZONE_IA"
	OutputSecurityLakeStorageClassIntelligentTiering OutputSecurityLakeStorageClass = "INTELLIGENT_TIERING"
	OutputSecurityLakeStorageClassGlacier            OutputSecurityLakeStorageClass = "GLACIER"
	OutputSecurityLakeStorageClassGlacierIr          OutputSecurityLakeStorageClass = "GLACIER_IR"
	OutputSecurityLakeStorageClassDeepArchive        OutputSecurityLakeStorageClass = "DEEP_ARCHIVE"
)

func (e OutputSecurityLakeStorageClass) ToPointer() *OutputSecurityLakeStorageClass {
	return &e
}
func (e *OutputSecurityLakeStorageClass) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		fallthrough
	case "REDUCED_REDUNDANCY":
		fallthrough
	case "STANDARD_IA":
		fallthrough
	case "ONEZONE_IA":
		fallthrough
	case "INTELLIGENT_TIERING":
		fallthrough
	case "GLACIER":
		fallthrough
	case "GLACIER_IR":
		fallthrough
	case "DEEP_ARCHIVE":
		*e = OutputSecurityLakeStorageClass(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeStorageClass: %v", v)
	}
}

type OutputSecurityLakeServerSideEncryptionForUploadedObjects string

const (
	OutputSecurityLakeServerSideEncryptionForUploadedObjectsAes256 OutputSecurityLakeServerSideEncryptionForUploadedObjects = "AES256"
	OutputSecurityLakeServerSideEncryptionForUploadedObjectsAwsKms OutputSecurityLakeServerSideEncryptionForUploadedObjects = "aws:kms"
)

func (e OutputSecurityLakeServerSideEncryptionForUploadedObjects) ToPointer() *OutputSecurityLakeServerSideEncryptionForUploadedObjects {
	return &e
}
func (e *OutputSecurityLakeServerSideEncryptionForUploadedObjects) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AES256":
		fallthrough
	case "aws:kms":
		*e = OutputSecurityLakeServerSideEncryptionForUploadedObjects(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeServerSideEncryptionForUploadedObjects: %v", v)
	}
}

// OutputSecurityLakeBackpressureBehavior - How to handle events when all receivers are exerting backpressure
type OutputSecurityLakeBackpressureBehavior string

const (
	OutputSecurityLakeBackpressureBehaviorBlock OutputSecurityLakeBackpressureBehavior = "block"
	OutputSecurityLakeBackpressureBehaviorDrop  OutputSecurityLakeBackpressureBehavior = "drop"
)

func (e OutputSecurityLakeBackpressureBehavior) ToPointer() *OutputSecurityLakeBackpressureBehavior {
	return &e
}
func (e *OutputSecurityLakeBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputSecurityLakeBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeBackpressureBehavior: %v", v)
	}
}

// OutputSecurityLakeDiskSpaceProtection - How to handle events when disk space is below the global 'Min free disk space' limit
type OutputSecurityLakeDiskSpaceProtection string

const (
	OutputSecurityLakeDiskSpaceProtectionBlock OutputSecurityLakeDiskSpaceProtection = "block"
	OutputSecurityLakeDiskSpaceProtectionDrop  OutputSecurityLakeDiskSpaceProtection = "drop"
)

func (e OutputSecurityLakeDiskSpaceProtection) ToPointer() *OutputSecurityLakeDiskSpaceProtection {
	return &e
}
func (e *OutputSecurityLakeDiskSpaceProtection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputSecurityLakeDiskSpaceProtection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeDiskSpaceProtection: %v", v)
	}
}

// OutputSecurityLakeParquetVersion - Determines which data types are supported and how they are represented
type OutputSecurityLakeParquetVersion string

const (
	OutputSecurityLakeParquetVersionParquet10 OutputSecurityLakeParquetVersion = "PARQUET_1_0"
	OutputSecurityLakeParquetVersionParquet24 OutputSecurityLakeParquetVersion = "PARQUET_2_4"
	OutputSecurityLakeParquetVersionParquet26 OutputSecurityLakeParquetVersion = "PARQUET_2_6"
)

func (e OutputSecurityLakeParquetVersion) ToPointer() *OutputSecurityLakeParquetVersion {
	return &e
}
func (e *OutputSecurityLakeParquetVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARQUET_1_0":
		fallthrough
	case "PARQUET_2_4":
		fallthrough
	case "PARQUET_2_6":
		*e = OutputSecurityLakeParquetVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeParquetVersion: %v", v)
	}
}

// OutputSecurityLakeDataPageVersion - Serialization format of data pages. Note that some reader implementations use Data page V2's attributes to work more efficiently, while others ignore it.
type OutputSecurityLakeDataPageVersion string

const (
	OutputSecurityLakeDataPageVersionDataPageV1 OutputSecurityLakeDataPageVersion = "DATA_PAGE_V1"
	OutputSecurityLakeDataPageVersionDataPageV2 OutputSecurityLakeDataPageVersion = "DATA_PAGE_V2"
)

func (e OutputSecurityLakeDataPageVersion) ToPointer() *OutputSecurityLakeDataPageVersion {
	return &e
}
func (e *OutputSecurityLakeDataPageVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DATA_PAGE_V1":
		fallthrough
	case "DATA_PAGE_V2":
		*e = OutputSecurityLakeDataPageVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSecurityLakeDataPageVersion: %v", v)
	}
}

type OutputSecurityLakeKeyValueMetadatum struct {
	Key   *string `default:"" json:"key"`
	Value string  `json:"value"`
}

func (o OutputSecurityLakeKeyValueMetadatum) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSecurityLakeKeyValueMetadatum) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputSecurityLakeKeyValueMetadatum) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *OutputSecurityLakeKeyValueMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputSecurityLake struct {
	// Unique ID for this output
	ID   *string                 `json:"id,omitempty"`
	Type *OutputSecurityLakeType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards. These fields are added as dimensions and labels to generated metrics and logs, respectively.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Name of the destination S3 bucket. Must be a JavaScript expression (which can evaluate to a constant value), enclosed in quotes or backticks. Can be evaluated only at initialization time. Example referencing a Global Variable: `myBucket-${C.vars.myVar}`
	Bucket string `json:"bucket"`
	// Region where the Amazon Security Lake is located.
	Region       string  `json:"region"`
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *OutputSecurityLakeAuthenticationMethod `default:"auto" json:"awsAuthenticationMethod"`
	// Amazon Security Lake service endpoint. If empty, defaults to the AWS Region-specific endpoint. Otherwise, it must point to Amazon Security Lake-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Signature version to use for signing Amazon Security Lake requests
	SignatureVersion *OutputSecurityLakeSignatureVersion `default:"v4" json:"signatureVersion"`
	// Reuse connections between requests, which can improve performance
	ReuseConnections *bool `default:"true" json:"reuseConnections"`
	// Reject certificates that cannot be verified against a valid CA, such as self-signed certificates
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Use Assume Role credentials to access S3
	EnableAssumeRole *bool `default:"false" json:"enableAssumeRole"`
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn string `json:"assumeRoleArn"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Duration of the assumed role's session, in seconds. Minimum is 900 (15 minutes), default is 3600 (1 hour), and maximum is 43200 (12 hours).
	DurationSeconds *float64 `default:"3600" json:"durationSeconds"`
	// Filesystem location in which to buffer files, before compressing and moving to final destination. Use performant and stable storage.
	StagePath *string `default:"$CRIBL_HOME/state/outputs/staging" json:"stagePath"`
	// Add the Output ID value to staging location
	AddIDToStagePath *bool `default:"true" json:"addIdToStagePath"`
	// Object ACL to assign to uploaded objects
	ObjectACL *OutputSecurityLakeObjectACL `default:"private" json:"objectACL"`
	// Storage class to select for uploaded objects
	StorageClass         *OutputSecurityLakeStorageClass                           `json:"storageClass,omitempty"`
	ServerSideEncryption *OutputSecurityLakeServerSideEncryptionForUploadedObjects `json:"serverSideEncryption,omitempty"`
	// ID or ARN of the KMS customer-managed key to use for encryption
	KmsKeyID *string `json:"kmsKeyId,omitempty"`
	// Remove empty staging directories after moving files
	RemoveEmptyDirs *bool `default:"true" json:"removeEmptyDirs"`
	// JavaScript expression to define the output filename prefix (can be constant)
	BaseFileName *string `default:"CriblOut" json:"baseFileName"`
	// Maximum uncompressed output file size. Files of this size will be closed and moved to final output location.
	MaxFileSizeMB *float64 `default:"32" json:"maxFileSizeMB"`
	// Maximum number of files to keep open concurrently. When exceeded, @{product} will close the oldest open files and move them to the final output location.
	MaxOpenFiles *float64 `default:"100" json:"maxOpenFiles"`
	// If set, this line will be written to the beginning of each output file
	HeaderLine *string `default:"" json:"headerLine"`
	// Buffer size used to write to a file
	WriteHighWaterMark *float64 `default:"64" json:"writeHighWaterMark"`
	// How to handle events when all receivers are exerting backpressure
	OnBackpressure *OutputSecurityLakeBackpressureBehavior `default:"block" json:"onBackpressure"`
	// If a file fails to move to its final destination after the maximum number of retries, move it to a designated directory to prevent further errors
	DeadletterEnabled *bool `default:"false" json:"deadletterEnabled"`
	// How to handle events when disk space is below the global 'Min free disk space' limit
	OnDiskFullBackpressure *OutputSecurityLakeDiskSpaceProtection `default:"block" json:"onDiskFullBackpressure"`
	// Maximum amount of time to write to a file. Files open for longer than this will be closed and moved to final output location.
	MaxFileOpenTimeSec *float64 `default:"300" json:"maxFileOpenTimeSec"`
	// Maximum amount of time to keep inactive files open. Files open for longer than this will be closed and moved to final output location.
	MaxFileIdleTimeSec *float64 `default:"30" json:"maxFileIdleTimeSec"`
	// Maximum number of parts to upload in parallel per file. Minimum part size is 5MB.
	MaxConcurrentFileParts *float64 `default:"4" json:"maxConcurrentFileParts"`
	// Disable if you can access files within the bucket but not the bucket itself
	VerifyPermissions *bool `default:"true" json:"verifyPermissions"`
	// Maximum number of files that can be waiting for upload before backpressure is applied
	MaxClosingFilesToBackpressure *float64 `default:"100" json:"maxClosingFilesToBackpressure"`
	// ID of the AWS account whose data the Destination will write to Security Lake. This should have been configured when creating the Amazon Security Lake custom source.
	AccountID string `json:"accountId"`
	// Name of the custom source configured in Amazon Security Lake
	CustomSource string `json:"customSource"`
	// Automatically calculate the schema based on the events of each Parquet file generated
	AutomaticSchema *bool `default:"false" json:"automaticSchema"`
	// Determines which data types are supported and how they are represented
	ParquetVersion *OutputSecurityLakeParquetVersion `default:"PARQUET_2_6" json:"parquetVersion"`
	// Serialization format of data pages. Note that some reader implementations use Data page V2's attributes to work more efficiently, while others ignore it.
	ParquetDataPageVersion *OutputSecurityLakeDataPageVersion `default:"DATA_PAGE_V2" json:"parquetDataPageVersion"`
	// The number of rows that every group will contain. The final group can contain a smaller number of rows.
	ParquetRowGroupLength *float64 `default:"10000" json:"parquetRowGroupLength"`
	// Target memory size for page segments, such as 1MB or 128MB. Generally, lower values improve reading speed, while higher values improve compression.
	ParquetPageSize *string `default:"1MB" json:"parquetPageSize"`
	// Log up to 3 rows that @{product} skips due to data mismatch
	ShouldLogInvalidRows *bool `json:"shouldLogInvalidRows,omitempty"`
	// The metadata of files the Destination writes will include the properties you add here as key-value pairs. Useful for tagging. Examples: "key":"OCSF Event Class", "value":"9001"
	KeyValueMetadata []OutputSecurityLakeKeyValueMetadatum `json:"keyValueMetadata,omitempty"`
	// Statistics profile an entire file in terms of minimum/maximum values within data, numbers of nulls, etc. You can use Parquet tools to view statistics.
	EnableStatistics *bool `default:"true" json:"enableStatistics"`
	// One page index contains statistics for one data page. Parquet readers use statistics to enable page skipping.
	EnableWritePageIndex *bool `default:"true" json:"enableWritePageIndex"`
	// Parquet tools can use the checksum of a Parquet page to verify data integrity
	EnablePageChecksum *bool   `default:"false" json:"enablePageChecksum"`
	Description        *string `json:"description,omitempty"`
	// This value can be a constant or a JavaScript expression (`${C.env.SOME_ACCESS_KEY}`)
	AwsAPIKey *string `json:"awsApiKey,omitempty"`
	// Select or create a stored secret that references your access key and secret key
	AwsSecret *string `json:"awsSecret,omitempty"`
	// How frequently, in seconds, to clean up empty directories
	EmptyDirCleanupSec *float64 `default:"300" json:"emptyDirCleanupSec"`
	// To add a new schema, navigate to Processing > Knowledge > Parquet Schemas
	ParquetSchema *string `json:"parquetSchema,omitempty"`
	// Storage location for files that fail to reach their final destination after maximum retries are exceeded
	DeadletterPath *string `default:"$CRIBL_HOME/state/outputs/dead-letter" json:"deadletterPath"`
	// The maximum number of times a file will attempt to move to its final destination before being dead-lettered
	MaxRetryNum *float64 `default:"20" json:"maxRetryNum"`
}

func (o OutputSecurityLake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSecurityLake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputSecurityLake) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputSecurityLake) GetType() *OutputSecurityLakeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputSecurityLake) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputSecurityLake) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputSecurityLake) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputSecurityLake) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputSecurityLake) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *OutputSecurityLake) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}

func (o *OutputSecurityLake) GetAwsSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretKey
}

func (o *OutputSecurityLake) GetAwsAuthenticationMethod() *OutputSecurityLakeAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AwsAuthenticationMethod
}

func (o *OutputSecurityLake) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *OutputSecurityLake) GetSignatureVersion() *OutputSecurityLakeSignatureVersion {
	if o == nil {
		return nil
	}
	return o.SignatureVersion
}

func (o *OutputSecurityLake) GetReuseConnections() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnections
}

func (o *OutputSecurityLake) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputSecurityLake) GetEnableAssumeRole() *bool {
	if o == nil {
		return nil
	}
	return o.EnableAssumeRole
}

func (o *OutputSecurityLake) GetAssumeRoleArn() string {
	if o == nil {
		return ""
	}
	return o.AssumeRoleArn
}

func (o *OutputSecurityLake) GetAssumeRoleExternalID() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleExternalID
}

func (o *OutputSecurityLake) GetDurationSeconds() *float64 {
	if o == nil {
		return nil
	}
	return o.DurationSeconds
}

func (o *OutputSecurityLake) GetStagePath() *string {
	if o == nil {
		return nil
	}
	return o.StagePath
}

func (o *OutputSecurityLake) GetAddIDToStagePath() *bool {
	if o == nil {
		return nil
	}
	return o.AddIDToStagePath
}

func (o *OutputSecurityLake) GetObjectACL() *OutputSecurityLakeObjectACL {
	if o == nil {
		return nil
	}
	return o.ObjectACL
}

func (o *OutputSecurityLake) GetStorageClass() *OutputSecurityLakeStorageClass {
	if o == nil {
		return nil
	}
	return o.StorageClass
}

func (o *OutputSecurityLake) GetServerSideEncryption() *OutputSecurityLakeServerSideEncryptionForUploadedObjects {
	if o == nil {
		return nil
	}
	return o.ServerSideEncryption
}

func (o *OutputSecurityLake) GetKmsKeyID() *string {
	if o == nil {
		return nil
	}
	return o.KmsKeyID
}

func (o *OutputSecurityLake) GetRemoveEmptyDirs() *bool {
	if o == nil {
		return nil
	}
	return o.RemoveEmptyDirs
}

func (o *OutputSecurityLake) GetBaseFileName() *string {
	if o == nil {
		return nil
	}
	return o.BaseFileName
}

func (o *OutputSecurityLake) GetMaxFileSizeMB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxFileSizeMB
}

func (o *OutputSecurityLake) GetMaxOpenFiles() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxOpenFiles
}

func (o *OutputSecurityLake) GetHeaderLine() *string {
	if o == nil {
		return nil
	}
	return o.HeaderLine
}

func (o *OutputSecurityLake) GetWriteHighWaterMark() *float64 {
	if o == nil {
		return nil
	}
	return o.WriteHighWaterMark
}

func (o *OutputSecurityLake) GetOnBackpressure() *OutputSecurityLakeBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputSecurityLake) GetDeadletterEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.DeadletterEnabled
}

func (o *OutputSecurityLake) GetOnDiskFullBackpressure() *OutputSecurityLakeDiskSpaceProtection {
	if o == nil {
		return nil
	}
	return o.OnDiskFullBackpressure
}

func (o *OutputSecurityLake) GetMaxFileOpenTimeSec() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxFileOpenTimeSec
}

func (o *OutputSecurityLake) GetMaxFileIdleTimeSec() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxFileIdleTimeSec
}

func (o *OutputSecurityLake) GetMaxConcurrentFileParts() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxConcurrentFileParts
}

func (o *OutputSecurityLake) GetVerifyPermissions() *bool {
	if o == nil {
		return nil
	}
	return o.VerifyPermissions
}

func (o *OutputSecurityLake) GetMaxClosingFilesToBackpressure() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxClosingFilesToBackpressure
}

func (o *OutputSecurityLake) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *OutputSecurityLake) GetCustomSource() string {
	if o == nil {
		return ""
	}
	return o.CustomSource
}

func (o *OutputSecurityLake) GetAutomaticSchema() *bool {
	if o == nil {
		return nil
	}
	return o.AutomaticSchema
}

func (o *OutputSecurityLake) GetParquetVersion() *OutputSecurityLakeParquetVersion {
	if o == nil {
		return nil
	}
	return o.ParquetVersion
}

func (o *OutputSecurityLake) GetParquetDataPageVersion() *OutputSecurityLakeDataPageVersion {
	if o == nil {
		return nil
	}
	return o.ParquetDataPageVersion
}

func (o *OutputSecurityLake) GetParquetRowGroupLength() *float64 {
	if o == nil {
		return nil
	}
	return o.ParquetRowGroupLength
}

func (o *OutputSecurityLake) GetParquetPageSize() *string {
	if o == nil {
		return nil
	}
	return o.ParquetPageSize
}

func (o *OutputSecurityLake) GetShouldLogInvalidRows() *bool {
	if o == nil {
		return nil
	}
	return o.ShouldLogInvalidRows
}

func (o *OutputSecurityLake) GetKeyValueMetadata() []OutputSecurityLakeKeyValueMetadatum {
	if o == nil {
		return nil
	}
	return o.KeyValueMetadata
}

func (o *OutputSecurityLake) GetEnableStatistics() *bool {
	if o == nil {
		return nil
	}
	return o.EnableStatistics
}

func (o *OutputSecurityLake) GetEnableWritePageIndex() *bool {
	if o == nil {
		return nil
	}
	return o.EnableWritePageIndex
}

func (o *OutputSecurityLake) GetEnablePageChecksum() *bool {
	if o == nil {
		return nil
	}
	return o.EnablePageChecksum
}

func (o *OutputSecurityLake) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputSecurityLake) GetAwsAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsAPIKey
}

func (o *OutputSecurityLake) GetAwsSecret() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecret
}

func (o *OutputSecurityLake) GetEmptyDirCleanupSec() *float64 {
	if o == nil {
		return nil
	}
	return o.EmptyDirCleanupSec
}

func (o *OutputSecurityLake) GetParquetSchema() *string {
	if o == nil {
		return nil
	}
	return o.ParquetSchema
}

func (o *OutputSecurityLake) GetDeadletterPath() *string {
	if o == nil {
		return nil
	}
	return o.DeadletterPath
}

func (o *OutputSecurityLake) GetMaxRetryNum() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryNum
}
