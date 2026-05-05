# InputResponseInputSyslogUnion


## Supported Types

### InputResponseInputSyslog1

```go
inputResponseInputSyslogUnion := components.CreateInputResponseInputSyslogUnionInputResponseInputSyslog1(components.InputResponseInputSyslog1{/* values here */})
```

### InputResponseInputSyslog2

```go
inputResponseInputSyslogUnion := components.CreateInputResponseInputSyslogUnionInputResponseInputSyslog2(components.InputResponseInputSyslog2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputResponseInputSyslogUnion.Type {
	case components.InputResponseInputSyslogUnionTypeInputResponseInputSyslog1:
		// inputResponseInputSyslogUnion.InputResponseInputSyslog1 is populated
	case components.InputResponseInputSyslogUnionTypeInputResponseInputSyslog2:
		// inputResponseInputSyslogUnion.InputResponseInputSyslog2 is populated
}
```
