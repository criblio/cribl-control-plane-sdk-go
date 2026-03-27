# CreateInputSystemByPackInputSyslogUnion


## Supported Types

### CreateInputSystemByPackInputSyslogSyslog1

```go
createInputSystemByPackInputSyslogUnion := operations.CreateCreateInputSystemByPackInputSyslogUnionCreateInputSystemByPackInputSyslogSyslog1(operations.CreateInputSystemByPackInputSyslogSyslog1{/* values here */})
```

### CreateInputSystemByPackInputSyslogSyslog2

```go
createInputSystemByPackInputSyslogUnion := operations.CreateCreateInputSystemByPackInputSyslogUnionCreateInputSystemByPackInputSyslogSyslog2(operations.CreateInputSystemByPackInputSyslogSyslog2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createInputSystemByPackInputSyslogUnion.Type {
	case operations.CreateInputSystemByPackInputSyslogUnionTypeCreateInputSystemByPackInputSyslogSyslog1:
		// createInputSystemByPackInputSyslogUnion.CreateInputSystemByPackInputSyslogSyslog1 is populated
	case operations.CreateInputSystemByPackInputSyslogUnionTypeCreateInputSystemByPackInputSyslogSyslog2:
		// createInputSystemByPackInputSyslogUnion.CreateInputSystemByPackInputSyslogSyslog2 is populated
}
```
