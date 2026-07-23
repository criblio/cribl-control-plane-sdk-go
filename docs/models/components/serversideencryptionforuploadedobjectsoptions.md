# ServerSideEncryptionForUploadedObjectsOptions

Server-side encryption to use for uploaded objects

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.ServerSideEncryptionForUploadedObjectsOptionsAes256Value

// Open enum: custom values can be created with a direct type cast
custom := components.ServerSideEncryptionForUploadedObjectsOptions("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ServerSideEncryptionForUploadedObjectsOptionsAes256Value` | AES256                                                     |
| `ServerSideEncryptionForUploadedObjectsOptionsAwsKms`      | aws:kms                                                    |