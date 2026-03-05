# S3AwsAuthenticationMethodSecretPartitioningScheme

Partitioning scheme used for this dataset. Using a known scheme like DDSS enables more efficient data reading and retrieval.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.S3AwsAuthenticationMethodSecretPartitioningSchemeNone

// Open enum: custom values can be created with a direct type cast
custom := components.S3AwsAuthenticationMethodSecretPartitioningScheme("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `S3AwsAuthenticationMethodSecretPartitioningSchemeNone` | none                                                    |
| `S3AwsAuthenticationMethodSecretPartitioningSchemeDdss` | ddss                                                    |