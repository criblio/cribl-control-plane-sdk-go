# InputSyslog


## Supported Types

### InputSyslogSyslog1

```go
inputSyslog := components.CreateInputSyslogInputSyslogSyslog1(components.InputSyslogSyslog1{/* values here */})
```

### InputSyslogSyslog2

```go
inputSyslog := components.CreateInputSyslogInputSyslogSyslog2(components.InputSyslogSyslog2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputSyslog.Type {
	case components.InputSyslogTypeInputSyslogSyslog1:
		// inputSyslog.InputSyslogSyslog1 is populated
	case components.InputSyslogTypeInputSyslogSyslog2:
		// inputSyslog.InputSyslogSyslog2 is populated
}
```
