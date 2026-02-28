# V3User


## Supported Types

### SnmpTrapSerializeV3UserAuthProtocolNone

```go
v3User := components.CreateV3UserNone(components.SnmpTrapSerializeV3UserAuthProtocolNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNone

```go
v3User := components.CreateV3UserMd5(components.SnmpTrapSerializeV3UserAuthProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNone

```go
v3User := components.CreateV3UserSha(components.SnmpTrapSerializeV3UserAuthProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNone

```go
v3User := components.CreateV3UserSha224(components.SnmpTrapSerializeV3UserAuthProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNone

```go
v3User := components.CreateV3UserSha256(components.SnmpTrapSerializeV3UserAuthProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNone

```go
v3User := components.CreateV3UserSha384(components.SnmpTrapSerializeV3UserAuthProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNone

```go
v3User := components.CreateV3UserSha512(components.SnmpTrapSerializeV3UserAuthProtocolNotNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3User.Type {
	case components.V3UserTypeNone:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNone is populated
	case components.V3UserTypeMd5:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNotNone is populated
	case components.V3UserTypeSha:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNotNone is populated
	case components.V3UserTypeSha224:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNotNone is populated
	case components.V3UserTypeSha256:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNotNone is populated
	case components.V3UserTypeSha384:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNotNone is populated
	case components.V3UserTypeSha512:
		// v3User.SnmpTrapSerializeV3UserAuthProtocolNotNone is populated
	default:
		// Unknown type - use v3User.GetUnknownRaw() for raw JSON
}
```
