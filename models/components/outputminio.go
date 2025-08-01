// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
)

type OutputMinioType string

const (
	OutputMinioTypeMinio OutputMinioType = "minio"
)

func (e OutputMinioType) ToPointer() *OutputMinioType {
	return &e
}
func (e *OutputMinioType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "minio":
		*e = OutputMinioType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioType: %v", v)
	}
}

// OutputMinioAuthenticationMethod - AWS authentication method. Choose Auto to use IAM roles.
type OutputMinioAuthenticationMethod string

const (
	OutputMinioAuthenticationMethodAuto   OutputMinioAuthenticationMethod = "auto"
	OutputMinioAuthenticationMethodManual OutputMinioAuthenticationMethod = "manual"
	OutputMinioAuthenticationMethodSecret OutputMinioAuthenticationMethod = "secret"
)

func (e OutputMinioAuthenticationMethod) ToPointer() *OutputMinioAuthenticationMethod {
	return &e
}
func (e *OutputMinioAuthenticationMethod) UnmarshalJSON(data []byte) error {
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
		*e = OutputMinioAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioAuthenticationMethod: %v", v)
	}
}

// OutputMinioSignatureVersion - Signature version to use for signing MinIO requests
type OutputMinioSignatureVersion string

const (
	OutputMinioSignatureVersionV2 OutputMinioSignatureVersion = "v2"
	OutputMinioSignatureVersionV4 OutputMinioSignatureVersion = "v4"
)

func (e OutputMinioSignatureVersion) ToPointer() *OutputMinioSignatureVersion {
	return &e
}
func (e *OutputMinioSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = OutputMinioSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioSignatureVersion: %v", v)
	}
}

// OutputMinioObjectACL - Object ACL to assign to uploaded objects
type OutputMinioObjectACL string

const (
	OutputMinioObjectACLPrivate                OutputMinioObjectACL = "private"
	OutputMinioObjectACLPublicRead             OutputMinioObjectACL = "public-read"
	OutputMinioObjectACLPublicReadWrite        OutputMinioObjectACL = "public-read-write"
	OutputMinioObjectACLAuthenticatedRead      OutputMinioObjectACL = "authenticated-read"
	OutputMinioObjectACLAwsExecRead            OutputMinioObjectACL = "aws-exec-read"
	OutputMinioObjectACLBucketOwnerRead        OutputMinioObjectACL = "bucket-owner-read"
	OutputMinioObjectACLBucketOwnerFullControl OutputMinioObjectACL = "bucket-owner-full-control"
)

func (e OutputMinioObjectACL) ToPointer() *OutputMinioObjectACL {
	return &e
}
func (e *OutputMinioObjectACL) UnmarshalJSON(data []byte) error {
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
		*e = OutputMinioObjectACL(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioObjectACL: %v", v)
	}
}

// OutputMinioStorageClass - Storage class to select for uploaded objects
type OutputMinioStorageClass string

const (
	OutputMinioStorageClassStandard          OutputMinioStorageClass = "STANDARD"
	OutputMinioStorageClassReducedRedundancy OutputMinioStorageClass = "REDUCED_REDUNDANCY"
)

func (e OutputMinioStorageClass) ToPointer() *OutputMinioStorageClass {
	return &e
}
func (e *OutputMinioStorageClass) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		fallthrough
	case "REDUCED_REDUNDANCY":
		*e = OutputMinioStorageClass(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioStorageClass: %v", v)
	}
}

// ServerSideEncryption - Server-side encryption for uploaded objects
type ServerSideEncryption string

const (
	ServerSideEncryptionAes256 ServerSideEncryption = "AES256"
)

func (e ServerSideEncryption) ToPointer() *ServerSideEncryption {
	return &e
}
func (e *ServerSideEncryption) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AES256":
		*e = ServerSideEncryption(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerSideEncryption: %v", v)
	}
}

// OutputMinioDataFormat - Format of the output data
type OutputMinioDataFormat string

const (
	OutputMinioDataFormatJSON    OutputMinioDataFormat = "json"
	OutputMinioDataFormatRaw     OutputMinioDataFormat = "raw"
	OutputMinioDataFormatParquet OutputMinioDataFormat = "parquet"
)

func (e OutputMinioDataFormat) ToPointer() *OutputMinioDataFormat {
	return &e
}
func (e *OutputMinioDataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "raw":
		fallthrough
	case "parquet":
		*e = OutputMinioDataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioDataFormat: %v", v)
	}
}

// OutputMinioBackpressureBehavior - How to handle events when all receivers are exerting backpressure
type OutputMinioBackpressureBehavior string

const (
	OutputMinioBackpressureBehaviorBlock OutputMinioBackpressureBehavior = "block"
	OutputMinioBackpressureBehaviorDrop  OutputMinioBackpressureBehavior = "drop"
)

func (e OutputMinioBackpressureBehavior) ToPointer() *OutputMinioBackpressureBehavior {
	return &e
}
func (e *OutputMinioBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputMinioBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioBackpressureBehavior: %v", v)
	}
}

// OutputMinioDiskSpaceProtection - How to handle events when disk space is below the global 'Min free disk space' limit
type OutputMinioDiskSpaceProtection string

const (
	OutputMinioDiskSpaceProtectionBlock OutputMinioDiskSpaceProtection = "block"
	OutputMinioDiskSpaceProtectionDrop  OutputMinioDiskSpaceProtection = "drop"
)

func (e OutputMinioDiskSpaceProtection) ToPointer() *OutputMinioDiskSpaceProtection {
	return &e
}
func (e *OutputMinioDiskSpaceProtection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputMinioDiskSpaceProtection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioDiskSpaceProtection: %v", v)
	}
}

// OutputMinioCompression - Data compression format to apply to HTTP content before it is delivered
type OutputMinioCompression string

const (
	OutputMinioCompressionNone OutputMinioCompression = "none"
	OutputMinioCompressionGzip OutputMinioCompression = "gzip"
)

func (e OutputMinioCompression) ToPointer() *OutputMinioCompression {
	return &e
}
func (e *OutputMinioCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputMinioCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioCompression: %v", v)
	}
}

// OutputMinioCompressionLevel - Compression level to apply before moving files to final destination
type OutputMinioCompressionLevel string

const (
	OutputMinioCompressionLevelBestSpeed       OutputMinioCompressionLevel = "best_speed"
	OutputMinioCompressionLevelNormal          OutputMinioCompressionLevel = "normal"
	OutputMinioCompressionLevelBestCompression OutputMinioCompressionLevel = "best_compression"
)

func (e OutputMinioCompressionLevel) ToPointer() *OutputMinioCompressionLevel {
	return &e
}
func (e *OutputMinioCompressionLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "best_speed":
		fallthrough
	case "normal":
		fallthrough
	case "best_compression":
		*e = OutputMinioCompressionLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioCompressionLevel: %v", v)
	}
}

// OutputMinioParquetVersion - Determines which data types are supported and how they are represented
type OutputMinioParquetVersion string

const (
	OutputMinioParquetVersionParquet10 OutputMinioParquetVersion = "PARQUET_1_0"
	OutputMinioParquetVersionParquet24 OutputMinioParquetVersion = "PARQUET_2_4"
	OutputMinioParquetVersionParquet26 OutputMinioParquetVersion = "PARQUET_2_6"
)

func (e OutputMinioParquetVersion) ToPointer() *OutputMinioParquetVersion {
	return &e
}
func (e *OutputMinioParquetVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputMinioParquetVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioParquetVersion: %v", v)
	}
}

// OutputMinioDataPageVersion - Serialization format of data pages. Note that some reader implementations use Data page V2's attributes to work more efficiently, while others ignore it.
type OutputMinioDataPageVersion string

const (
	OutputMinioDataPageVersionDataPageV1 OutputMinioDataPageVersion = "DATA_PAGE_V1"
	OutputMinioDataPageVersionDataPageV2 OutputMinioDataPageVersion = "DATA_PAGE_V2"
)

func (e OutputMinioDataPageVersion) ToPointer() *OutputMinioDataPageVersion {
	return &e
}
func (e *OutputMinioDataPageVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DATA_PAGE_V1":
		fallthrough
	case "DATA_PAGE_V2":
		*e = OutputMinioDataPageVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioDataPageVersion: %v", v)
	}
}

type OutputMinioKeyValueMetadatum struct {
	Key   *string `default:"" json:"key"`
	Value string  `json:"value"`
}

func (o OutputMinioKeyValueMetadatum) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputMinioKeyValueMetadatum) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputMinioKeyValueMetadatum) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *OutputMinioKeyValueMetadatum) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputMinio struct {
	// Unique ID for this output
	ID   *string          `json:"id,omitempty"`
	Type *OutputMinioType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// MinIO service url (e.g. http://minioHost:9000)
	Endpoint string `json:"endpoint"`
	// Name of the destination MinIO bucket. This value can be a constant or a JavaScript expression that can only be evaluated at init time. Example referencing a Global Variable: `myBucket-${C.vars.myVar}`
	Bucket string `json:"bucket"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *OutputMinioAuthenticationMethod `default:"auto" json:"awsAuthenticationMethod"`
	// Secret key. This value can be a constant or a JavaScript expression, such as `${C.env.SOME_SECRET}`).
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// Region where the MinIO service/cluster is located
	Region *string `json:"region,omitempty"`
	// Filesystem location in which to buffer files, before compressing and moving to final destination. Use performant stable storage.
	StagePath *string `default:"$CRIBL_HOME/state/outputs/staging" json:"stagePath"`
	// Add the Output ID value to staging location
	AddIDToStagePath *bool `default:"true" json:"addIdToStagePath"`
	// Root directory to prepend to path before uploading. Enter a constant, or a JavaScript expression enclosed in quotes or backticks.
	DestPath *string `json:"destPath,omitempty"`
	// Signature version to use for signing MinIO requests
	SignatureVersion *OutputMinioSignatureVersion `default:"v4" json:"signatureVersion"`
	// Object ACL to assign to uploaded objects
	ObjectACL *OutputMinioObjectACL `default:"private" json:"objectACL"`
	// Storage class to select for uploaded objects
	StorageClass *OutputMinioStorageClass `json:"storageClass,omitempty"`
	// Server-side encryption for uploaded objects
	ServerSideEncryption *ServerSideEncryption `json:"serverSideEncryption,omitempty"`
	// Reuse connections between requests, which can improve performance
	ReuseConnections *bool `default:"true" json:"reuseConnections"`
	// Reject certificates that cannot be verified against a valid CA, such as self-signed certificates)
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Disable if you can access files within the bucket but not the bucket itself
	VerifyPermissions *bool `default:"true" json:"verifyPermissions"`
	// Remove empty staging directories after moving files
	RemoveEmptyDirs *bool `default:"true" json:"removeEmptyDirs"`
	// JavaScript expression defining how files are partitioned and organized. Default is date-based. If blank, Stream will fall back to the event's __partition field value – if present – otherwise to each location's root directory.
	PartitionExpr *string `default:"C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')" json:"partitionExpr"`
	// Format of the output data
	Format *OutputMinioDataFormat `default:"json" json:"format"`
	// JavaScript expression to define the output filename prefix (can be constant)
	BaseFileName *string `default:"CriblOut" json:"baseFileName"`
	// JavaScript expression to define the output filename suffix (can be constant).  The `__format` variable refers to the value of the `Data format` field (`json` or `raw`).  The `__compression` field refers to the kind of compression being used (`none` or `gzip`).
	FileNameSuffix *string `default:".\\${C.env[\"CRIBL_WORKER_ID\"]}.\\${__format}\\${__compression === \"gzip\" ? \".gz\" : \"\"}" json:"fileNameSuffix"`
	// Maximum uncompressed output file size. Files of this size will be closed and moved to final output location.
	MaxFileSizeMB *float64 `default:"32" json:"maxFileSizeMB"`
	// Maximum number of files to keep open concurrently. When exceeded, @{product} will close the oldest open files and move them to the final output location.
	MaxOpenFiles *float64 `default:"100" json:"maxOpenFiles"`
	// If set, this line will be written to the beginning of each output file
	HeaderLine *string `default:"" json:"headerLine"`
	// Buffer size used to write to a file
	WriteHighWaterMark *float64 `default:"64" json:"writeHighWaterMark"`
	// How to handle events when all receivers are exerting backpressure
	OnBackpressure *OutputMinioBackpressureBehavior `default:"block" json:"onBackpressure"`
	// If a file fails to move to its final destination after the maximum number of retries, move it to a designated directory to prevent further errors
	DeadletterEnabled *bool `default:"false" json:"deadletterEnabled"`
	// How to handle events when disk space is below the global 'Min free disk space' limit
	OnDiskFullBackpressure *OutputMinioDiskSpaceProtection `default:"block" json:"onDiskFullBackpressure"`
	// Maximum amount of time to write to a file. Files open for longer than this will be closed and moved to final output location.
	MaxFileOpenTimeSec *float64 `default:"300" json:"maxFileOpenTimeSec"`
	// Maximum amount of time to keep inactive files open. Files open for longer than this will be closed and moved to final output location.
	MaxFileIdleTimeSec *float64 `default:"30" json:"maxFileIdleTimeSec"`
	// Maximum number of parts to upload in parallel per file. Minimum part size is 5MB.
	MaxConcurrentFileParts *float64 `default:"4" json:"maxConcurrentFileParts"`
	Description            *string  `json:"description,omitempty"`
	// This value can be a constant or a JavaScript expression (`${C.env.SOME_ACCESS_KEY}`)
	AwsAPIKey *string `json:"awsApiKey,omitempty"`
	// Select or create a stored secret that references your access key and secret key
	AwsSecret *string `json:"awsSecret,omitempty"`
	// Data compression format to apply to HTTP content before it is delivered
	Compress *OutputMinioCompression `default:"gzip" json:"compress"`
	// Compression level to apply before moving files to final destination
	CompressionLevel *OutputMinioCompressionLevel `default:"best_speed" json:"compressionLevel"`
	// Automatically calculate the schema based on the events of each Parquet file generated
	AutomaticSchema *bool `default:"false" json:"automaticSchema"`
	// Determines which data types are supported and how they are represented
	ParquetVersion *OutputMinioParquetVersion `default:"PARQUET_2_6" json:"parquetVersion"`
	// Serialization format of data pages. Note that some reader implementations use Data page V2's attributes to work more efficiently, while others ignore it.
	ParquetDataPageVersion *OutputMinioDataPageVersion `default:"DATA_PAGE_V2" json:"parquetDataPageVersion"`
	// The number of rows that every group will contain. The final group can contain a smaller number of rows.
	ParquetRowGroupLength *float64 `default:"10000" json:"parquetRowGroupLength"`
	// Target memory size for page segments, such as 1MB or 128MB. Generally, lower values improve reading speed, while higher values improve compression.
	ParquetPageSize *string `default:"1MB" json:"parquetPageSize"`
	// Log up to 3 rows that @{product} skips due to data mismatch
	ShouldLogInvalidRows *bool `json:"shouldLogInvalidRows,omitempty"`
	// The metadata of files the Destination writes will include the properties you add here as key-value pairs. Useful for tagging. Examples: "key":"OCSF Event Class", "value":"9001"
	KeyValueMetadata []OutputMinioKeyValueMetadatum `json:"keyValueMetadata,omitempty"`
	// Statistics profile an entire file in terms of minimum/maximum values within data, numbers of nulls, etc. You can use Parquet tools to view statistics.
	EnableStatistics *bool `default:"true" json:"enableStatistics"`
	// One page index contains statistics for one data page. Parquet readers use statistics to enable page skipping.
	EnableWritePageIndex *bool `default:"true" json:"enableWritePageIndex"`
	// Parquet tools can use the checksum of a Parquet page to verify data integrity
	EnablePageChecksum *bool `default:"false" json:"enablePageChecksum"`
	// How frequently, in seconds, to clean up empty directories
	EmptyDirCleanupSec *float64 `default:"300" json:"emptyDirCleanupSec"`
	// Storage location for files that fail to reach their final destination after maximum retries are exceeded
	DeadletterPath *string `default:"$CRIBL_HOME/state/outputs/dead-letter" json:"deadletterPath"`
	// The maximum number of times a file will attempt to move to its final destination before being dead-lettered
	MaxRetryNum *float64 `default:"20" json:"maxRetryNum"`
}

func (o OutputMinio) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputMinio) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputMinio) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputMinio) GetType() *OutputMinioType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputMinio) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputMinio) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputMinio) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputMinio) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputMinio) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *OutputMinio) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *OutputMinio) GetAwsAuthenticationMethod() *OutputMinioAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AwsAuthenticationMethod
}

func (o *OutputMinio) GetAwsSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretKey
}

func (o *OutputMinio) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *OutputMinio) GetStagePath() *string {
	if o == nil {
		return nil
	}
	return o.StagePath
}

func (o *OutputMinio) GetAddIDToStagePath() *bool {
	if o == nil {
		return nil
	}
	return o.AddIDToStagePath
}

func (o *OutputMinio) GetDestPath() *string {
	if o == nil {
		return nil
	}
	return o.DestPath
}

func (o *OutputMinio) GetSignatureVersion() *OutputMinioSignatureVersion {
	if o == nil {
		return nil
	}
	return o.SignatureVersion
}

func (o *OutputMinio) GetObjectACL() *OutputMinioObjectACL {
	if o == nil {
		return nil
	}
	return o.ObjectACL
}

func (o *OutputMinio) GetStorageClass() *OutputMinioStorageClass {
	if o == nil {
		return nil
	}
	return o.StorageClass
}

func (o *OutputMinio) GetServerSideEncryption() *ServerSideEncryption {
	if o == nil {
		return nil
	}
	return o.ServerSideEncryption
}

func (o *OutputMinio) GetReuseConnections() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnections
}

func (o *OutputMinio) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputMinio) GetVerifyPermissions() *bool {
	if o == nil {
		return nil
	}
	return o.VerifyPermissions
}

func (o *OutputMinio) GetRemoveEmptyDirs() *bool {
	if o == nil {
		return nil
	}
	return o.RemoveEmptyDirs
}

func (o *OutputMinio) GetPartitionExpr() *string {
	if o == nil {
		return nil
	}
	return o.PartitionExpr
}

func (o *OutputMinio) GetFormat() *OutputMinioDataFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputMinio) GetBaseFileName() *string {
	if o == nil {
		return nil
	}
	return o.BaseFileName
}

func (o *OutputMinio) GetFileNameSuffix() *string {
	if o == nil {
		return nil
	}
	return o.FileNameSuffix
}

func (o *OutputMinio) GetMaxFileSizeMB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxFileSizeMB
}

func (o *OutputMinio) GetMaxOpenFiles() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxOpenFiles
}

func (o *OutputMinio) GetHeaderLine() *string {
	if o == nil {
		return nil
	}
	return o.HeaderLine
}

func (o *OutputMinio) GetWriteHighWaterMark() *float64 {
	if o == nil {
		return nil
	}
	return o.WriteHighWaterMark
}

func (o *OutputMinio) GetOnBackpressure() *OutputMinioBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputMinio) GetDeadletterEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.DeadletterEnabled
}

func (o *OutputMinio) GetOnDiskFullBackpressure() *OutputMinioDiskSpaceProtection {
	if o == nil {
		return nil
	}
	return o.OnDiskFullBackpressure
}

func (o *OutputMinio) GetMaxFileOpenTimeSec() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxFileOpenTimeSec
}

func (o *OutputMinio) GetMaxFileIdleTimeSec() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxFileIdleTimeSec
}

func (o *OutputMinio) GetMaxConcurrentFileParts() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxConcurrentFileParts
}

func (o *OutputMinio) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputMinio) GetAwsAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsAPIKey
}

func (o *OutputMinio) GetAwsSecret() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecret
}

func (o *OutputMinio) GetCompress() *OutputMinioCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputMinio) GetCompressionLevel() *OutputMinioCompressionLevel {
	if o == nil {
		return nil
	}
	return o.CompressionLevel
}

func (o *OutputMinio) GetAutomaticSchema() *bool {
	if o == nil {
		return nil
	}
	return o.AutomaticSchema
}

func (o *OutputMinio) GetParquetVersion() *OutputMinioParquetVersion {
	if o == nil {
		return nil
	}
	return o.ParquetVersion
}

func (o *OutputMinio) GetParquetDataPageVersion() *OutputMinioDataPageVersion {
	if o == nil {
		return nil
	}
	return o.ParquetDataPageVersion
}

func (o *OutputMinio) GetParquetRowGroupLength() *float64 {
	if o == nil {
		return nil
	}
	return o.ParquetRowGroupLength
}

func (o *OutputMinio) GetParquetPageSize() *string {
	if o == nil {
		return nil
	}
	return o.ParquetPageSize
}

func (o *OutputMinio) GetShouldLogInvalidRows() *bool {
	if o == nil {
		return nil
	}
	return o.ShouldLogInvalidRows
}

func (o *OutputMinio) GetKeyValueMetadata() []OutputMinioKeyValueMetadatum {
	if o == nil {
		return nil
	}
	return o.KeyValueMetadata
}

func (o *OutputMinio) GetEnableStatistics() *bool {
	if o == nil {
		return nil
	}
	return o.EnableStatistics
}

func (o *OutputMinio) GetEnableWritePageIndex() *bool {
	if o == nil {
		return nil
	}
	return o.EnableWritePageIndex
}

func (o *OutputMinio) GetEnablePageChecksum() *bool {
	if o == nil {
		return nil
	}
	return o.EnablePageChecksum
}

func (o *OutputMinio) GetEmptyDirCleanupSec() *float64 {
	if o == nil {
		return nil
	}
	return o.EmptyDirCleanupSec
}

func (o *OutputMinio) GetDeadletterPath() *string {
	if o == nil {
		return nil
	}
	return o.DeadletterPath
}

func (o *OutputMinio) GetMaxRetryNum() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryNum
}
