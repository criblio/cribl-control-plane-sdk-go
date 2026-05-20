# InputResponseInputSyslogUnion


## Supported Types

### InputResponseInputSyslogSyslog1

```go
inputResponseInputSyslogUnion := components.CreateInputResponseInputSyslogUnionInputResponseInputSyslogSyslog1(components.InputResponseInputSyslogSyslog1{/* values here */})
```

### InputResponseInputSyslogSyslog2

```go
inputResponseInputSyslogUnion := components.CreateInputResponseInputSyslogUnionInputResponseInputSyslogSyslog2(components.InputResponseInputSyslogSyslog2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputResponseInputSyslogUnion.Type {
	case components.InputResponseInputSyslogUnionTypeInputResponseInputSyslogSyslog1:
		// inputResponseInputSyslogUnion.InputResponseInputSyslogSyslog1 is populated
	case components.InputResponseInputSyslogUnionTypeInputResponseInputSyslogSyslog2:
		// inputResponseInputSyslogUnion.InputResponseInputSyslogSyslog2 is populated
}
```
