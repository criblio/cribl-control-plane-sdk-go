# SnmpTrapSerializeV3UserAuthProtocolNotNone


## Supported Types

### SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNone

```go
snmpTrapSerializeV3UserAuthProtocolNotNone := components.CreateSnmpTrapSerializeV3UserAuthProtocolNotNoneNone(components.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone

```go
snmpTrapSerializeV3UserAuthProtocolNotNone := components.CreateSnmpTrapSerializeV3UserAuthProtocolNotNoneDes(components.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone

```go
snmpTrapSerializeV3UserAuthProtocolNotNone := components.CreateSnmpTrapSerializeV3UserAuthProtocolNotNoneAes(components.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone

```go
snmpTrapSerializeV3UserAuthProtocolNotNone := components.CreateSnmpTrapSerializeV3UserAuthProtocolNotNoneAes256b(components.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone{/* values here */})
```

### SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone

```go
snmpTrapSerializeV3UserAuthProtocolNotNone := components.CreateSnmpTrapSerializeV3UserAuthProtocolNotNoneAes256r(components.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch snmpTrapSerializeV3UserAuthProtocolNotNone.Type {
	case components.SnmpTrapSerializeV3UserAuthProtocolNotNoneTypeNone:
		// snmpTrapSerializeV3UserAuthProtocolNotNone.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNone is populated
	case components.SnmpTrapSerializeV3UserAuthProtocolNotNoneTypeDes:
		// snmpTrapSerializeV3UserAuthProtocolNotNone.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone is populated
	case components.SnmpTrapSerializeV3UserAuthProtocolNotNoneTypeAes:
		// snmpTrapSerializeV3UserAuthProtocolNotNone.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone is populated
	case components.SnmpTrapSerializeV3UserAuthProtocolNotNoneTypeAes256b:
		// snmpTrapSerializeV3UserAuthProtocolNotNone.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone is populated
	case components.SnmpTrapSerializeV3UserAuthProtocolNotNoneTypeAes256r:
		// snmpTrapSerializeV3UserAuthProtocolNotNone.SnmpTrapSerializeV3UserAuthProtocolNotNonePrivProtocolNotNone is populated
	default:
		// Unknown type - use snmpTrapSerializeV3UserAuthProtocolNotNone.GetUnknownRaw() for raw JSON
}
```
