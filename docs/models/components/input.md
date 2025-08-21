# Input


## Supported Types

### InputCollection

```go
input := components.CreateInputInputCollection(components.InputCollection{/* values here */})
```

### InputKafka

```go
input := components.CreateInputInputKafka(components.InputKafka{/* values here */})
```

### InputMsk

```go
input := components.CreateInputInputMsk(components.InputMsk{/* values here */})
```

### InputHTTP

```go
input := components.CreateInputInputHTTP(components.InputHTTP{/* values here */})
```

### InputSplunk

```go
input := components.CreateInputInputSplunk(components.InputSplunk{/* values here */})
```

### InputSplunkSearch

```go
input := components.CreateInputInputSplunkSearch(components.InputSplunkSearch{/* values here */})
```

### InputSplunkHec

```go
input := components.CreateInputInputSplunkHec(components.InputSplunkHec{/* values here */})
```

### InputAzureBlob

```go
input := components.CreateInputInputAzureBlob(components.InputAzureBlob{/* values here */})
```

### InputElastic

```go
input := components.CreateInputInputElastic(components.InputElastic{/* values here */})
```

### InputConfluentCloud

```go
input := components.CreateInputInputConfluentCloud(components.InputConfluentCloud{/* values here */})
```

### InputGrafana

```go
input := components.CreateInputInputGrafana(components.InputGrafana{/* values here */})
```

### InputLoki

```go
input := components.CreateInputInputLoki(components.InputLoki{/* values here */})
```

### InputPrometheusRw

```go
input := components.CreateInputInputPrometheusRw(components.InputPrometheusRw{/* values here */})
```

### InputPrometheus

```go
input := components.CreateInputInputPrometheus(components.InputPrometheus{/* values here */})
```

### InputEdgePrometheus

```go
input := components.CreateInputInputEdgePrometheus(components.InputEdgePrometheus{/* values here */})
```

### InputOffice365Mgmt

```go
input := components.CreateInputInputOffice365Mgmt(components.InputOffice365Mgmt{/* values here */})
```

### InputOffice365Service

```go
input := components.CreateInputInputOffice365Service(components.InputOffice365Service{/* values here */})
```

### InputOffice365MsgTrace

```go
input := components.CreateInputInputOffice365MsgTrace(components.InputOffice365MsgTrace{/* values here */})
```

### InputEventhub

```go
input := components.CreateInputInputEventhub(components.InputEventhub{/* values here */})
```

### InputExec

```go
input := components.CreateInputInputExec(components.InputExec{/* values here */})
```

### InputFirehose

```go
input := components.CreateInputInputFirehose(components.InputFirehose{/* values here */})
```

### InputGooglePubsub

```go
input := components.CreateInputInputGooglePubsub(components.InputGooglePubsub{/* values here */})
```

### InputCribl

```go
input := components.CreateInputInputCribl(components.InputCribl{/* values here */})
```

### InputCriblTCP

```go
input := components.CreateInputInputCriblTCP(components.InputCriblTCP{/* values here */})
```

### InputCriblHTTP

```go
input := components.CreateInputInputCriblHTTP(components.InputCriblHTTP{/* values here */})
```

### InputCriblLakeHTTP

```go
input := components.CreateInputInputCriblLakeHTTP(components.InputCriblLakeHTTP{/* values here */})
```

### InputTcpjson

```go
input := components.CreateInputInputTcpjson(components.InputTcpjson{/* values here */})
```

### InputSystemMetrics

```go
input := components.CreateInputInputSystemMetrics(components.InputSystemMetrics{/* values here */})
```

### InputSystemState

```go
input := components.CreateInputInputSystemState(components.InputSystemState{/* values here */})
```

### InputKubeMetrics

```go
input := components.CreateInputInputKubeMetrics(components.InputKubeMetrics{/* values here */})
```

### InputKubeLogs

```go
input := components.CreateInputInputKubeLogs(components.InputKubeLogs{/* values here */})
```

### InputKubeEvents

```go
input := components.CreateInputInputKubeEvents(components.InputKubeEvents{/* values here */})
```

### InputWindowsMetrics

```go
input := components.CreateInputInputWindowsMetrics(components.InputWindowsMetrics{/* values here */})
```

### InputCrowdstrike

```go
input := components.CreateInputInputCrowdstrike(components.InputCrowdstrike{/* values here */})
```

### InputDatadogAgent

```go
input := components.CreateInputInputDatadogAgent(components.InputDatadogAgent{/* values here */})
```

### InputDatagen

```go
input := components.CreateInputInputDatagen(components.InputDatagen{/* values here */})
```

### InputHTTPRaw

```go
input := components.CreateInputInputHTTPRaw(components.InputHTTPRaw{/* values here */})
```

### InputKinesis

```go
input := components.CreateInputInputKinesis(components.InputKinesis{/* values here */})
```

### InputCriblmetrics

```go
input := components.CreateInputInputCriblmetrics(components.InputCriblmetrics{/* values here */})
```

### InputMetrics

```go
input := components.CreateInputInputMetrics(components.InputMetrics{/* values here */})
```

### InputS3

```go
input := components.CreateInputInputS3(components.InputS3{/* values here */})
```

### InputS3Inventory

```go
input := components.CreateInputInputS3Inventory(components.InputS3Inventory{/* values here */})
```

### InputSnmp

```go
input := components.CreateInputInputSnmp(components.InputSnmp{/* values here */})
```

### InputOpenTelemetry

```go
input := components.CreateInputInputOpenTelemetry(components.InputOpenTelemetry{/* values here */})
```

### InputModelDrivenTelemetry

```go
input := components.CreateInputInputModelDrivenTelemetry(components.InputModelDrivenTelemetry{/* values here */})
```

### InputSqs

```go
input := components.CreateInputInputSqs(components.InputSqs{/* values here */})
```

### InputSyslog

```go
input := components.CreateInputInputSyslog(components.InputSyslog{/* values here */})
```

### InputFile

```go
input := components.CreateInputInputFile(components.InputFile{/* values here */})
```

### InputTCP

```go
input := components.CreateInputInputTCP(components.InputTCP{/* values here */})
```

### InputAppscope

```go
input := components.CreateInputInputAppscope(components.InputAppscope{/* values here */})
```

### InputWef

```go
input := components.CreateInputInputWef(components.InputWef{/* values here */})
```

### InputWinEventLogs

```go
input := components.CreateInputInputWinEventLogs(components.InputWinEventLogs{/* values here */})
```

### InputRawUDP

```go
input := components.CreateInputInputRawUDP(components.InputRawUDP{/* values here */})
```

### InputJournalFiles

```go
input := components.CreateInputInputJournalFiles(components.InputJournalFiles{/* values here */})
```

### InputWiz

```go
input := components.CreateInputInputWiz(components.InputWiz{/* values here */})
```

### InputWizWebhook

```go
input := components.CreateInputInputWizWebhook(components.InputWizWebhook{/* values here */})
```

### InputNetflow

```go
input := components.CreateInputInputNetflow(components.InputNetflow{/* values here */})
```

### InputSecurityLake

```go
input := components.CreateInputInputSecurityLake(components.InputSecurityLake{/* values here */})
```

### InputZscalerHec

```go
input := components.CreateInputInputZscalerHec(components.InputZscalerHec{/* values here */})
```

