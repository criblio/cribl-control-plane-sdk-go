# CreateOutputRequest

New Destination object


## Supported Types

### OutputDefault

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDefault(operations.OutputDefault{/* values here */})
```

### OutputWebhook

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputWebhook(components.OutputWebhook{/* values here */})
```

### OutputSentinel

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSentinel(operations.OutputSentinel{/* values here */})
```

### OutputDevnull

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDevnull(components.OutputDevnull{/* values here */})
```

### OutputSyslog

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSyslog(components.OutputSyslog{/* values here */})
```

### OutputSplunk

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSplunk(operations.OutputSplunk{/* values here */})
```

### OutputSplunkLb

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSplunkLb(operations.OutputSplunkLb{/* values here */})
```

### OutputSplunkHec

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSplunkHec(components.OutputSplunkHec{/* values here */})
```

### OutputTcpjson

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputTcpjson(components.OutputTcpjson{/* values here */})
```

### OutputWavefront

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputWavefront(operations.OutputWavefront{/* values here */})
```

### OutputSignalfx

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSignalfx(operations.OutputSignalfx{/* values here */})
```

### OutputFilesystem

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputFilesystem(operations.OutputFilesystem{/* values here */})
```

### OutputS3

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputS3(operations.OutputS3{/* values here */})
```

### OutputAzureBlob

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputAzureBlob(operations.OutputAzureBlob{/* values here */})
```

### OutputAzureDataExplorer

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputAzureDataExplorer(operations.OutputAzureDataExplorer{/* values here */})
```

### OutputAzureLogs

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputAzureLogs(operations.OutputAzureLogs{/* values here */})
```

### OutputKinesis

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputKinesis(operations.OutputKinesis{/* values here */})
```

### OutputHoneycomb

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputHoneycomb(operations.OutputHoneycomb{/* values here */})
```

### OutputAzureEventhub

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputAzureEventhub(operations.OutputAzureEventhub{/* values here */})
```

### OutputGoogleChronicle

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputGoogleChronicle(operations.OutputGoogleChronicle{/* values here */})
```

### OutputGoogleCloudStorage

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputGoogleCloudStorage(operations.OutputGoogleCloudStorage{/* values here */})
```

### OutputGoogleCloudLogging

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputGoogleCloudLogging(operations.OutputGoogleCloudLogging{/* values here */})
```

### OutputGooglePubsub

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputGooglePubsub(operations.OutputGooglePubsub{/* values here */})
```

### OutputExabeam

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputExabeam(operations.OutputExabeam{/* values here */})
```

### OutputKafka

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputKafka(operations.OutputKafka{/* values here */})
```

### OutputConfluentCloud

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputConfluentCloud(operations.OutputConfluentCloud{/* values here */})
```

### OutputMsk

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputMsk(operations.OutputMsk{/* values here */})
```

### OutputElastic

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputElastic(operations.OutputElastic{/* values here */})
```

### OutputElasticCloud

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputElasticCloud(operations.OutputElasticCloud{/* values here */})
```

### OutputNewrelic

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputNewrelic(components.OutputNewrelic{/* values here */})
```

### OutputNewrelicEvents

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputNewrelicEvents(operations.OutputNewrelicEvents{/* values here */})
```

### OutputInfluxdb

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputInfluxdb(operations.OutputInfluxdb{/* values here */})
```

### OutputCloudwatch

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputCloudwatch(operations.OutputCloudwatch{/* values here */})
```

### OutputMinio

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputMinio(operations.OutputMinio{/* values here */})
```

### OutputStatsd

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputStatsd(operations.OutputStatsd{/* values here */})
```

### OutputStatsdExt

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputStatsdExt(operations.OutputStatsdExt{/* values here */})
```

### OutputGraphite

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputGraphite(operations.OutputGraphite{/* values here */})
```

### OutputRouter

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputRouter(operations.OutputRouter{/* values here */})
```

### OutputSns

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSns(operations.OutputSns{/* values here */})
```

### OutputSqs

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSqs(operations.OutputSqs{/* values here */})
```

### OutputSnmp

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSnmp(operations.OutputSnmp{/* values here */})
```

### OutputSumoLogic

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSumoLogic(operations.OutputSumoLogic{/* values here */})
```

### OutputDatadog

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDatadog(components.OutputDatadog{/* values here */})
```

### OutputGrafanaCloudUnion

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputGrafanaCloudUnion(components.OutputGrafanaCloudUnion{/* values here */})
```

### OutputLoki

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputLoki(operations.OutputLoki{/* values here */})
```

### OutputPrometheus

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputPrometheus(operations.OutputPrometheus{/* values here */})
```

### OutputRing

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputRing(components.OutputRing{/* values here */})
```

### OutputOpenTelemetry

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputOpenTelemetry(operations.OutputOpenTelemetry{/* values here */})
```

### OutputServiceNow

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputServiceNow(operations.OutputServiceNow{/* values here */})
```

### OutputDataset

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDataset(components.OutputDataset{/* values here */})
```

### OutputCriblTCP

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputCriblTCP(components.OutputCriblTCP{/* values here */})
```

### OutputCriblHTTP

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputCriblHTTP(components.OutputCriblHTTP{/* values here */})
```

### OutputHumioHec

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputHumioHec(operations.OutputHumioHec{/* values here */})
```

### OutputCrowdstrikeNextGenSiem

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputCrowdstrikeNextGenSiem(operations.OutputCrowdstrikeNextGenSiem{/* values here */})
```

### OutputDlS3

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDlS3(operations.OutputDlS3{/* values here */})
```

### OutputSecurityLake

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSecurityLake(operations.OutputSecurityLake{/* values here */})
```

### OutputCriblLake

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputCriblLake(components.OutputCriblLake{/* values here */})
```

### OutputDiskSpool

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDiskSpool(components.OutputDiskSpool{/* values here */})
```

### OutputClickHouse

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputClickHouse(operations.OutputClickHouse{/* values here */})
```

### OutputXsiam

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputXsiam(components.OutputXsiam{/* values here */})
```

### OutputNetflow

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputNetflow(operations.OutputNetflow{/* values here */})
```

### OutputDynatraceHTTP

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDynatraceHTTP(operations.OutputDynatraceHTTP{/* values here */})
```

### OutputDynatraceOtlp

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputDynatraceOtlp(operations.OutputDynatraceOtlp{/* values here */})
```

### OutputSentinelOneAiSiem

```go
createOutputRequest := operations.CreateCreateOutputRequestOutputSentinelOneAiSiem(operations.OutputSentinelOneAiSiem{/* values here */})
```

