# InputSyslog


## Supported Types

### InputSyslogSyslog1

```go
inputSyslog := operations.CreateInputSyslogInputSyslogSyslog1(operations.InputSyslogSyslog1{/* values here */})
```

### InputSyslogSyslog2

```go
inputSyslog := operations.CreateInputSyslogInputSyslogSyslog2(operations.InputSyslogSyslog2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputSyslog.Type {
	case operations.InputSyslogTypeInputSyslogSyslog1:
		// inputSyslog.InputSyslogSyslog1 is populated
	case operations.InputSyslogTypeInputSyslogSyslog2:
		// inputSyslog.InputSyslogSyslog2 is populated
}
```
