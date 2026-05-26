# InputSyslogInputUnion


## Supported Types

### InputSyslogSyslogInput1

```go
inputSyslogInputUnion := components.CreateInputSyslogInputUnionInputSyslogSyslogInput1(components.InputSyslogSyslogInput1{/* values here */})
```

### InputSyslogSyslogInput2

```go
inputSyslogInputUnion := components.CreateInputSyslogInputUnionInputSyslogSyslogInput2(components.InputSyslogSyslogInput2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputSyslogInputUnion.Type {
	case components.InputSyslogInputUnionTypeInputSyslogSyslogInput1:
		// inputSyslogInputUnion.InputSyslogSyslogInput1 is populated
	case components.InputSyslogInputUnionTypeInputSyslogSyslogInput2:
		// inputSyslogInputUnion.InputSyslogSyslogInput2 is populated
}
```
