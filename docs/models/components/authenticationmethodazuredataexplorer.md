# AuthenticationMethodAzureDataExplorer

The type of OAuth 2.0 client credentials grant flow to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodAzureDataExplorerClientSecret

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodAzureDataExplorer("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `AuthenticationMethodAzureDataExplorerClientSecret`     | clientSecret                                            |
| `AuthenticationMethodAzureDataExplorerClientTextSecret` | clientTextSecret                                        |
| `AuthenticationMethodAzureDataExplorerCertificate`      | certificate                                             |