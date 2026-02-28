# CreateInputInputSyslogUnion


## Supported Types

### CreateInputInputSyslogSyslog1

```go
createInputInputSyslogUnion := operations.CreateCreateInputInputSyslogUnionCreateInputInputSyslogSyslog1(operations.CreateInputInputSyslogSyslog1{/* values here */})
```

### CreateInputInputSyslogSyslog2

```go
createInputInputSyslogUnion := operations.CreateCreateInputInputSyslogUnionCreateInputInputSyslogSyslog2(operations.CreateInputInputSyslogSyslog2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createInputInputSyslogUnion.Type {
	case operations.CreateInputInputSyslogUnionTypeCreateInputInputSyslogSyslog1:
		// createInputInputSyslogUnion.CreateInputInputSyslogSyslog1 is populated
	case operations.CreateInputInputSyslogUnionTypeCreateInputInputSyslogSyslog2:
		// createInputInputSyslogUnion.CreateInputInputSyslogSyslog2 is populated
}
```
