# ServerSideEncryptionForUploadedObjectsOptions

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ServerSideEncryptionForUploadedObjectsOptionsAes256

// Open enum: custom values can be created with a direct type cast
custom := components.ServerSideEncryptionForUploadedObjectsOptions("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `ServerSideEncryptionForUploadedObjectsOptionsAes256` | AES256                                                |
| `ServerSideEncryptionForUploadedObjectsOptionsAwsKms` | aws:kms                                               |