# cribl-control-plane-sdk-go
<!-- Start Summary [summary] -->
## Summary

Cribl API Reference: This API Reference lists available REST endpoints, along with their supported operations for accessing, creating, updating, or deleting resources. See our complementary product documentation at [docs.cribl.io](http://docs.cribl.io).
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [cribl-control-plane-sdk-go](#cribl-control-plane-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Json Streaming](#json-streaming)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/criblio/cribl-control-plane-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
	)

	res, err := s.Projects.GetSystemProjectsSubscriptionsByGroupIDByAndProjectID(ctx, "<id>", "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      | Environment Variable            |
| ------------ | ---- | ----------- | ------------------------------- |
| `BearerAuth` | http | HTTP Bearer | `CRIBLCONTROLPLANE_BEARER_AUTH` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
	)

	res, err := s.Projects.GetSystemProjectsSubscriptionsByGroupIDByAndProjectID(ctx, "<id>", "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AiSettings](docs/sdks/aisettings/README.md)

* [GetAiSettingsFeaturesByFeature](docs/sdks/aisettings/README.md#getaisettingsfeaturesbyfeature) - Returns setting for a specific feature
* [CreateAiSettingsFeaturesByFeature](docs/sdks/aisettings/README.md#createaisettingsfeaturesbyfeature) - Updates setting for a specific feature
* [GetAiSettingsFeatures](docs/sdks/aisettings/README.md#getaisettingsfeatures) - Returns all feature settings
* [CreateAiSettingsFeatures](docs/sdks/aisettings/README.md#createaisettingsfeatures) - Updates all AI settings at once

### [AppscopeConfigs](docs/sdks/appscopeconfigs/README.md)

* [ListAppscopeLibEntry](docs/sdks/appscopeconfigs/README.md#listappscopelibentry) - Get a list of AppscopeLibEntry objects
* [CreateAppscopeLibEntry](docs/sdks/appscopeconfigs/README.md#createappscopelibentry) - Create AppscopeLibEntry
* [GetAppscopeLibEntryByID](docs/sdks/appscopeconfigs/README.md#getappscopelibentrybyid) - Get AppscopeLibEntry by ID
* [UpdateAppscopeLibEntryByID](docs/sdks/appscopeconfigs/README.md#updateappscopelibentrybyid) - Update AppscopeLibEntry
* [DeleteAppscopeLibEntryByID](docs/sdks/appscopeconfigs/README.md#deleteappscopelibentrybyid) - Delete AppscopeLibEntry

### [Auth](docs/sdks/auth/README.md)

* [GetAuthAuthorizationCodeCallback](docs/sdks/auth/README.md#getauthauthorizationcodecallback) - API call that the IDP should use for an authorization code callback
* [GetAuthMetadata](docs/sdks/auth/README.md#getauthmetadata) - Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider
* [Login](docs/sdks/auth/README.md#login) - Log in and obtain Auth token
* [Logout](docs/sdks/auth/README.md#logout) - Log current user out
* [CreateAuthSloCallback](docs/sdks/auth/README.md#createauthslocallback) - API call that the IDP should use for a logout request
* [GetAuthSloCallback](docs/sdks/auth/README.md#getauthslocallback) - Accepts a logout request from an IDP and logs out the user
* [GetAuthSlo](docs/sdks/auth/README.md#getauthslo) - Redirect user to IDP with logout request
* [CreateAuthSsoCallback](docs/sdks/auth/README.md#createauthssocallback) - API call that the IDP should use for an authentication request
* [GetAuthSsoCallback](docs/sdks/auth/README.md#getauthssocallback) - Accepts an authentication request from an IDP and authenticates the user
* [GetAuthSso](docs/sdks/auth/README.md#getauthsso) - Obtain redirect information
* [GetAuthGroups](docs/sdks/auth/README.md#getauthgroups) - List the external authentication system's groups
* [GetAuthMultiFactor](docs/sdks/auth/README.md#getauthmultifactor) - Get PIV configuration
* [DeleteAuthUsersTokenByID](docs/sdks/auth/README.md#deleteauthuserstokenbyid) - Invalidate token(s) for user *id*
* [DeleteAuthUsersInvalidationByID](docs/sdks/auth/README.md#deleteauthusersinvalidationbyid) - Remove invalidation for user ID

### [Authorize](docs/sdks/authorize/README.md)

* [GetAuthorizePolicy](docs/sdks/authorize/README.md#getauthorizepolicy) - get the client's authorization policy
* [GetAuthorizeRoles](docs/sdks/authorize/README.md#getauthorizeroles) - get the client's roles

### [Banners](docs/sdks/banners/README.md)

* [ListBannerMessage](docs/sdks/banners/README.md#listbannermessage) - Get a list of BannerMessage objects
* [CreateBannerMessage](docs/sdks/banners/README.md#createbannermessage) - Create BannerMessage
* [GetBannerMessageByID](docs/sdks/banners/README.md#getbannermessagebyid) - Get BannerMessage by ID
* [UpdateBannerMessageByID](docs/sdks/banners/README.md#updatebannermessagebyid) - Update BannerMessage
* [DeleteBannerMessageByID](docs/sdks/banners/README.md#deletebannermessagebyid) - Delete BannerMessage

### [Certificates](docs/sdks/certificates/README.md)

* [ListCertificate](docs/sdks/certificates/README.md#listcertificate) - Get a list of Certificate objects
* [CreateCertificate](docs/sdks/certificates/README.md#createcertificate) - Create Certificate
* [GetCertificateByID](docs/sdks/certificates/README.md#getcertificatebyid) - Get Certificate by ID
* [UpdateCertificateByID](docs/sdks/certificates/README.md#updatecertificatebyid) - Update Certificate
* [DeleteCertificateByID](docs/sdks/certificates/README.md#deletecertificatebyid) - Delete Certificate

### [Changelog](docs/sdks/changelog/README.md)

* [GetChangelogViewed](docs/sdks/changelog/README.md#getchangelogviewed) - Get changelog viewed state

### [ClickHouse](docs/sdks/clickhouse/README.md)

* [CreateOutputClickHouseDescribeTable](docs/sdks/clickhouse/README.md#createoutputclickhousedescribetable) - Retrieve the description of the configured ClickHouse table

### [Clui](docs/sdks/clui/README.md)

* [GetClui](docs/sdks/clui/README.md#getclui) - Get CLUI search results

### [Collectors](docs/sdks/collectors/README.md)

* [ListCollector](docs/sdks/collectors/README.md#listcollector) - Get a list of Collector objects
* [GetCollectorByID](docs/sdks/collectors/README.md#getcollectorbyid) - Get Collector by ID

### [Conditions](docs/sdks/conditions/README.md)

* [ListCondition](docs/sdks/conditions/README.md#listcondition) - Get a list of Condition objects
* [GetConditionByID](docs/sdks/conditions/README.md#getconditionbyid) - Get Condition by ID

### [Consent](docs/sdks/consent/README.md)

* [GetAiConsent](docs/sdks/consent/README.md#getaiconsent) - Fetches the AI consent information, specifically the org GUID and accepted boolean.
* [CreateAiConsent](docs/sdks/consent/README.md#createaiconsent) - Stores the AI consent information, specifically the org GUID and accepted boolean.


### [DashboardCategories](docs/sdks/dashboardcategories/README.md)

* [ListDashboardCategory](docs/sdks/dashboardcategories/README.md#listdashboardcategory) - Get a list of DashboardCategory objects
* [CreateDashboardCategory](docs/sdks/dashboardcategories/README.md#createdashboardcategory) - Create DashboardCategory
* [GetDashboardCategoryByID](docs/sdks/dashboardcategories/README.md#getdashboardcategorybyid) - Get DashboardCategory by ID
* [UpdateDashboardCategoryByID](docs/sdks/dashboardcategories/README.md#updatedashboardcategorybyid) - Update DashboardCategory
* [DeleteDashboardCategoryByID](docs/sdks/dashboardcategories/README.md#deletedashboardcategorybyid) - Delete DashboardCategory

### [Dashboards](docs/sdks/dashboards/README.md)

* [ListSearchDashboard](docs/sdks/dashboards/README.md#listsearchdashboard) - Get a list of SearchDashboard objects
* [CreateSearchDashboard](docs/sdks/dashboards/README.md#createsearchdashboard) - Create SearchDashboard
* [GetSearchDashboardByID](docs/sdks/dashboards/README.md#getsearchdashboardbyid) - Get SearchDashboard by ID
* [UpdateSearchDashboardByID](docs/sdks/dashboards/README.md#updatesearchdashboardbyid) - Update SearchDashboard
* [DeleteSearchDashboardByID](docs/sdks/dashboards/README.md#deletesearchdashboardbyid) - Delete SearchDashboard
* [GetSearchDashboardACLByID](docs/sdks/dashboards/README.md#getsearchdashboardaclbyid) - Get SearchDashboard ACL
* [CreateSearchDashboardACLApplyByID](docs/sdks/dashboards/README.md#createsearchdashboardaclapplybyid) - Modify SearchDashboard ACL
* [GetSearchDashboardACLTeamsByID](docs/sdks/dashboards/README.md#getsearchdashboardaclteamsbyid) - Get SearchDashboard Teams
* [CreateSearchDashboardACLTeamsApplyByID](docs/sdks/dashboards/README.md#createsearchdashboardaclteamsapplybyid) - Modify SearchDashboard Teams ACL

### [DatabaseConnections](docs/sdks/databaseconnections/README.md)

* [GetDatabaseConnectionConfig](docs/sdks/databaseconnections/README.md#getdatabaseconnectionconfig) - Get a list of DatabaseConnection objects
* [CreateDatabaseConnectionConfig](docs/sdks/databaseconnections/README.md#createdatabaseconnectionconfig) - Create DatabaseConnectionConfig
* [GetDatabaseConnectionConfigByID](docs/sdks/databaseconnections/README.md#getdatabaseconnectionconfigbyid) - Get DatabaseConnectionConfig by ID
* [UpdateDatabaseConnectionConfigByID](docs/sdks/databaseconnections/README.md#updatedatabaseconnectionconfigbyid) - Update DatabaseConnectionConfig
* [DeleteDatabaseConnectionConfigByID](docs/sdks/databaseconnections/README.md#deletedatabaseconnectionconfigbyid) - Delete DatabaseConnectionConfig
* [CreateLibDatabaseConnectionsTest](docs/sdks/databaseconnections/README.md#createlibdatabaseconnectionstest) - Test a database connection given a type and connectionString

### [Datasets](docs/sdks/datasets/README.md)

* [ListDatasetProviderType](docs/sdks/datasets/README.md#listdatasetprovidertype) - Get a list of DatasetProviderType objects
* [CreateDatasetProviderType](docs/sdks/datasets/README.md#createdatasetprovidertype) - Create DatasetProviderType
* [GetDatasetProviderTypeByID](docs/sdks/datasets/README.md#getdatasetprovidertypebyid) - Get DatasetProviderType by ID
* [UpdateDatasetProviderTypeByID](docs/sdks/datasets/README.md#updatedatasetprovidertypebyid) - Update DatasetProviderType
* [DeleteDatasetProviderTypeByID](docs/sdks/datasets/README.md#deletedatasetprovidertypebyid) - Delete DatasetProviderType
* [ListDatasetProvider](docs/sdks/datasets/README.md#listdatasetprovider) - Get a list of DatasetProvider objects
* [CreateDatasetProvider](docs/sdks/datasets/README.md#createdatasetprovider) - Create DatasetProvider
* [GetDatasetProviderByID](docs/sdks/datasets/README.md#getdatasetproviderbyid) - Get DatasetProvider by ID
* [UpdateDatasetProviderByID](docs/sdks/datasets/README.md#updatedatasetproviderbyid) - Update DatasetProvider
* [DeleteDatasetProviderByID](docs/sdks/datasets/README.md#deletedatasetproviderbyid) - Delete DatasetProvider
* [ListDataset](docs/sdks/datasets/README.md#listdataset) - Get a list of Dataset objects
* [CreateDataset](docs/sdks/datasets/README.md#createdataset) - Create Dataset
* [GetDatasetByID](docs/sdks/datasets/README.md#getdatasetbyid) - Get Dataset by ID
* [UpdateDatasetByID](docs/sdks/datasets/README.md#updatedatasetbyid) - Update Dataset
* [DeleteDatasetByID](docs/sdks/datasets/README.md#deletedatasetbyid) - Delete Dataset
* [GetDatasetACLByID](docs/sdks/datasets/README.md#getdatasetaclbyid) - Get Dataset ACL
* [CreateDatasetACLApplyByID](docs/sdks/datasets/README.md#createdatasetaclapplybyid) - Modify Dataset ACL
* [GetDatasetACLTeamsByID](docs/sdks/datasets/README.md#getdatasetaclteamsbyid) - Get Dataset Teams
* [CreateDatasetACLTeamsApplyByID](docs/sdks/datasets/README.md#createdatasetaclteamsapplybyid) - Modify Dataset Teams ACL

### [Diag](docs/sdks/diag/README.md)

* [GetHealthInfo](docs/sdks/diag/README.md#gethealthinfo) - Provides health info for REST server
* [GetDiagBundle](docs/sdks/diag/README.md#getdiagbundle) - Returns a diag bundle as a tar.gz archive
* [GetSystemInfo](docs/sdks/diag/README.md#getsysteminfo) - Get basic system information
* [GetSystemDiag](docs/sdks/diag/README.md#getsystemdiag) - Get list of existing diag bundles
* [DeleteSystemDiag](docs/sdks/diag/README.md#deletesystemdiag) - Remove diag bundle
* [CreateSystemDiagSend](docs/sdks/diag/README.md#createsystemdiagsend) - Send a diag bundle (tar.gz archive) to Cribl

### [Distributed](docs/sdks/distributed/README.md)

* [GetSummaryWorkers](docs/sdks/distributed/README.md#getsummaryworkers) - get worker and edge nodes count
* [GetWorkers](docs/sdks/distributed/README.md#getworkers) - get worker and edge nodes
* [UpdateWorkersRestart](docs/sdks/distributed/README.md#updateworkersrestart) - restarts worker nodes
* [GetConfigBundlesByGroupAndVersion](docs/sdks/distributed/README.md#getconfigbundlesbygroupandversion) - Download config bundle (used by remote nodes)
* [GetSummary](docs/sdks/distributed/README.md#getsummary) - Get summary of Distributed deployment

### [Edge](docs/sdks/edge/README.md)

* [CreateEdgeAppscopeProcesses](docs/sdks/edge/README.md#createedgeappscopeprocesses) - Attach AppScope to a process running on the edge host
* [GetEdgeAppscopeProcesses](docs/sdks/edge/README.md#getedgeappscopeprocesses) - Get a detailed list of scoped processes running on the edge host
* [DeleteEdgeAppscopeProcessesByPid](docs/sdks/edge/README.md#deleteedgeappscopeprocessesbypid) - Detach AppScope from a process running on the edge host
* [GetEdgeAppscopeProcessesByPid](docs/sdks/edge/README.md#getedgeappscopeprocessesbypid) - Get details of a scoped process running on the edge host
* [UpdateEdgeAppscopeProcessesByPid](docs/sdks/edge/README.md#updateedgeappscopeprocessesbypid) - Update AppScope configuration for a process running on the edge host
* [GetEdgeEventsCollectors](docs/sdks/edge/README.md#getedgeeventscollectors) - Get list of configured collectors
* [GetEdgeFileinspect](docs/sdks/edge/README.md#getedgefileinspect) - Get details about a file on the edge host.
* [GetEdgeLsByPath](docs/sdks/edge/README.md#getedgelsbypath) - Get a directory listing of the given path
* [CreateProductsEdgeMapQuery](docs/sdks/edge/README.md#createproductsedgemapquery) - Data for the Map View for Edge Fleets (Leader only)
* [GetEdgeMetadata](docs/sdks/edge/README.md#getedgemetadata) - Get the host's metadata structure
* [CreateEdgeFileIngest](docs/sdks/edge/README.md#createedgefileingest) - Ingest a specified file through a specified pipeline to a specified destination or send to routes.
* [GetEdgeFileSample](docs/sdks/edge/README.md#getedgefilesample) - Get some number of bytes from the file at the given path
* [CreateEdgeKubeLogs](docs/sdks/edge/README.md#createedgekubelogs) - Make a request to the K8s API logs endpoint
* [GetEdgeKubeProxy](docs/sdks/edge/README.md#getedgekubeproxy) - Make a GET request to the K8s API
* [GetEdgeContainersByID](docs/sdks/edge/README.md#getedgecontainersbyid) - Get details for a single container on the edge host. Add stream=true to get a stream instead.
* [GetEdgeContainers](docs/sdks/edge/README.md#getedgecontainers) - Get a detailed list of containers running on the edge host.
* [GetEdgeLogs](docs/sdks/edge/README.md#getedgelogs) - list log files
* [GetEdgeProcessesByPid](docs/sdks/edge/README.md#getedgeprocessesbypid) - Get details of a process running on the edge host
* [GetEdgeProcesses](docs/sdks/edge/README.md#getedgeprocesses) - Get a detailed list of processes running on the edge host

### [EdgeAppScopeProcesses](docs/sdks/edgeappscopeprocesses/README.md)

* [CreateEdgeAppscopeProcesses](docs/sdks/edgeappscopeprocesses/README.md#createedgeappscopeprocesses) - Attach AppScope to a process running on the edge host
* [GetEdgeAppscopeProcesses](docs/sdks/edgeappscopeprocesses/README.md#getedgeappscopeprocesses) - Get a detailed list of scoped processes running on the edge host
* [DeleteEdgeAppscopeProcessesByPid](docs/sdks/edgeappscopeprocesses/README.md#deleteedgeappscopeprocessesbypid) - Detach AppScope from a process running on the edge host
* [GetEdgeAppscopeProcessesByPid](docs/sdks/edgeappscopeprocesses/README.md#getedgeappscopeprocessesbypid) - Get details of a scoped process running on the edge host
* [UpdateEdgeAppscopeProcessesByPid](docs/sdks/edgeappscopeprocesses/README.md#updateedgeappscopeprocessesbypid) - Update AppScope configuration for a process running on the edge host

### [EdgeContainers](docs/sdks/edgecontainers/README.md)

* [GetEdgeContainersByID](docs/sdks/edgecontainers/README.md#getedgecontainersbyid) - Get details for a single container on the edge host. Add stream=true to get a stream instead.
* [GetEdgeContainers](docs/sdks/edgecontainers/README.md#getedgecontainers) - Get a detailed list of containers running on the edge host.

### [EdgeEvents](docs/sdks/edgeevents/README.md)

* [GetEdgeEventsCollectors](docs/sdks/edgeevents/README.md#getedgeeventscollectors) - Get list of configured collectors

### [EdgeFiles](docs/sdks/edgefiles/README.md)

* [GetEdgeFileinspect](docs/sdks/edgefiles/README.md#getedgefileinspect) - Get details about a file on the edge host.

### [EdgeLs](docs/sdks/edgels/README.md)

* [GetEdgeLsByPath](docs/sdks/edgels/README.md#getedgelsbypath) - Get a directory listing of the given path

### [EdgeProcesses](docs/sdks/edgeprocesses/README.md)

* [GetEdgeProcessesByPid](docs/sdks/edgeprocesses/README.md#getedgeprocessesbypid) - Get details of a process running on the edge host
* [GetEdgeProcesses](docs/sdks/edgeprocesses/README.md#getedgeprocesses) - Get a detailed list of processes running on the edge host

### [EventBreakerRules](docs/sdks/eventbreakerrules/README.md)

* [ListEventBreakerRuleset](docs/sdks/eventbreakerrules/README.md#listeventbreakerruleset) - Get a list of Event Breaker Ruleset objects
* [CreateEventBreakerRuleset](docs/sdks/eventbreakerrules/README.md#createeventbreakerruleset) - Create Event Breaker Ruleset
* [GetEventBreakerRulesetByID](docs/sdks/eventbreakerrules/README.md#geteventbreakerrulesetbyid) - Get Event Breaker Ruleset by ID
* [UpdateEventBreakerRulesetByID](docs/sdks/eventbreakerrules/README.md#updateeventbreakerrulesetbyid) - Update Event Breaker Ruleset
* [DeleteEventBreakerRulesetByID](docs/sdks/eventbreakerrules/README.md#deleteeventbreakerrulesetbyid) - Delete Event Breaker Ruleset

### [Events](docs/sdks/events/README.md)

* [GetEdgeEventsQuery](docs/sdks/events/README.md#getedgeeventsquery) - Get events generated by a specified source

### [Executors](docs/sdks/executors/README.md)

* [ListExecutor](docs/sdks/executors/README.md#listexecutor) - Get a list of Executor objects
* [GetExecutorByID](docs/sdks/executors/README.md#getexecutorbyid) - Get Executor by ID

### [Expressions](docs/sdks/expressions/README.md)

* [CreateLibExpression](docs/sdks/expressions/README.md#createlibexpression) - Evaluate JavaScript expression

### [Features](docs/sdks/features/README.md)

* [GetFeaturesEntryByID](docs/sdks/features/README.md#getfeaturesentrybyid) - Get feature by id (i.e. 'type-name`)
* [GetFeaturesEntry](docs/sdks/features/README.md#getfeaturesentry) - List all features

### [File](docs/sdks/file/README.md)

* [CreateEdgeFileIngest](docs/sdks/file/README.md#createedgefileingest) - Ingest a specified file through a specified pipeline to a specified destination or send to routes.

### [FileSampler](docs/sdks/filesampler/README.md)

* [GetEdgeFileSample](docs/sdks/filesampler/README.md#getedgefilesample) - Get some number of bytes from the file at the given path

### [Functions](docs/sdks/functions/README.md)

* [ListFunction](docs/sdks/functions/README.md#listfunction) - Get a list of Function objects
* [GetFunctionByID](docs/sdks/functions/README.md#getfunctionbyid) - Get Function by ID
* [GetFunctionByPack](docs/sdks/functions/README.md#getfunctionbypack) - Get a list of Function objects within a Pack
* [GetFunctionByPackAndID](docs/sdks/functions/README.md#getfunctionbypackandid) - Get Function by ID within a Pack

### [Git](docs/sdks/git/README.md)

* [GetSystemProjectsVersionCountByGroupIDAndProjectID](docs/sdks/git/README.md#getsystemprojectsversioncountbygroupidandprojectid) - Get the count of files of changed
* [GetSystemProjectsVersionDiffByGroupIDAndProjectID](docs/sdks/git/README.md#getsystemprojectsversiondiffbygroupidandprojectid) - Get the textual diff for given commit
* [GetSystemProjectsVersionFilesByGroupIDAndProjectID](docs/sdks/git/README.md#getsystemprojectsversionfilesbygroupidandprojectid) - Get the files changed
* [CreateSystemProjectsVersionCommitByGroupIDAndProjectID](docs/sdks/git/README.md#createsystemprojectsversioncommitbygroupidandprojectid) - Commit project changes
* [CreateSystemProjectsVersionCommitByProjectID](docs/sdks/git/README.md#createsystemprojectsversioncommitbyprojectid) - Commit project changes.
* [GetSystemProjectsVersionCountByProjectID](docs/sdks/git/README.md#getsystemprojectsversioncountbyprojectid) - get the count of files of changed
* [GetSystemProjectsVersionDiffByProjectID](docs/sdks/git/README.md#getsystemprojectsversiondiffbyprojectid) - get the textual diff for given commit
* [GetSystemProjectsVersionFilesByProjectID](docs/sdks/git/README.md#getsystemprojectsversionfilesbyprojectid) - get the files changed
* [CreateSystemProjectsVersionRevertByProjectID](docs/sdks/git/README.md#createsystemprojectsversionrevertbyprojectid) - Revert project changes.
* [GetSystemProjectsVersionShowByProjectID](docs/sdks/git/README.md#getsystemprojectsversionshowbyprojectid) - Show project changes.
* [GetSystemSettingsGitSettings](docs/sdks/git/README.md#getsystemsettingsgitsettings) - Get git settings
* [UpdateSystemSettingsGitSettings](docs/sdks/git/README.md#updatesystemsettingsgitsettings) - Update git settings
* [UndoLastCommit](docs/sdks/git/README.md#undolastcommit) - undo the last commit
* [GetVersionBranch](docs/sdks/git/README.md#getversionbranch) - get the list of branches
* [CreateVersionCommit](docs/sdks/git/README.md#createversioncommit) - create a new commit containing the current configs the given log message describing the changes.
* [GetVersionCount](docs/sdks/git/README.md#getversioncount) - get the count of files of changed
* [GetVersionCurrentBranch](docs/sdks/git/README.md#getversioncurrentbranch) - returns git branch that the config is checked out to, if any
* [GetVersionDiff](docs/sdks/git/README.md#getversiondiff) - get the textual diff for given commit
* [GetVersionFiles](docs/sdks/git/README.md#getversionfiles) - get the files changed
* [GetVersionInfo](docs/sdks/git/README.md#getversioninfo) - Get info about versioning availability
* [CreateVersionPush](docs/sdks/git/README.md#createversionpush) - push the current configs to the remote repository.
* [CreateVersionRevert](docs/sdks/git/README.md#createversionrevert) - revert a commit
* [GetVersionShow](docs/sdks/git/README.md#getversionshow) - get the log message and textual diff for given commit
* [GetVersionStatus](docs/sdks/git/README.md#getversionstatus) - get the the working tree status
* [CreateVersionSync](docs/sdks/git/README.md#createversionsync) - syncs with remote repo via POST requests

### [GlobalVariables](docs/sdks/globalvariables/README.md)

* [GetGlobalVariable](docs/sdks/globalvariables/README.md#getglobalvariable) - List all GlobalVars with references
* [CreateGlobalVariable](docs/sdks/globalvariables/README.md#createglobalvariable) - Create Global Variable
* [GetGlobalVariableByID](docs/sdks/globalvariables/README.md#getglobalvariablebyid) - Get Global Variable by ID
* [UpdateGlobalVariableByID](docs/sdks/globalvariables/README.md#updateglobalvariablebyid) - Update Global Variable
* [DeleteGlobalVariableByID](docs/sdks/globalvariables/README.md#deleteglobalvariablebyid) - Delete Global Variable
* [GetGlobalVariableLibVarsByPack](docs/sdks/globalvariables/README.md#getglobalvariablelibvarsbypack) - List all GlobalVars with references within a Pack
* [CreateGlobalVariableLibVarsByPack](docs/sdks/globalvariables/README.md#createglobalvariablelibvarsbypack) - Create Global Variable within a Pack
* [GetGlobalVariableLibVarsByPackAndID](docs/sdks/globalvariables/README.md#getglobalvariablelibvarsbypackandid) - Get Global Variable by ID within a Pack
* [UpdateGlobalVariableLibVarsByPackAndID](docs/sdks/globalvariables/README.md#updateglobalvariablelibvarsbypackandid) - Update Global Variable within a Pack
* [DeleteGlobalVariableLibVarsByPackAndID](docs/sdks/globalvariables/README.md#deleteglobalvariablelibvarsbypackandid) - Delete Global Variable within a Pack

### [Grokfiles](docs/sdks/grokfiles/README.md)

* [ListGrokFile](docs/sdks/grokfiles/README.md#listgrokfile) - Get a list of GrokFile objects
* [CreateGrokFile](docs/sdks/grokfiles/README.md#creategrokfile) - Create GrokFile
* [GetGrokFileByID](docs/sdks/grokfiles/README.md#getgrokfilebyid) - Get GrokFile by ID
* [UpdateGrokFileByID](docs/sdks/grokfiles/README.md#updategrokfilebyid) - Update GrokFile
* [DeleteGrokFileByID](docs/sdks/grokfiles/README.md#deletegrokfilebyid) - Delete GrokFile

### [Groups](docs/sdks/groups/README.md)

* [GetGroupsConfigVersionByID](docs/sdks/groups/README.md#getgroupsconfigversionbyid) - Get effective bundle version for given Group
* [CreateProductsGroupsByProduct](docs/sdks/groups/README.md#createproductsgroupsbyproduct) - Create a Fleet or Worker Group
* [GetProductsGroupsByProduct](docs/sdks/groups/README.md#getproductsgroupsbyproduct) - Get a list of ConfigGroup objects
* [UpdateGroupsDeployByID](docs/sdks/groups/README.md#updategroupsdeploybyid) - Deploy commits for a Fleet or Worker Group
* [GetGroupsByID](docs/sdks/groups/README.md#getgroupsbyid) - Get a specific ConfigGroup object
* [GetGroupsACLByID](docs/sdks/groups/README.md#getgroupsaclbyid) - ACL of members with permissions for resources in this Group

### [Health](docs/sdks/health/README.md)

* [GetHealthInfo](docs/sdks/health/README.md#gethealthinfo) - Provides health info for REST server

### [HmacFunctions](docs/sdks/hmacfunctions/README.md)

* [ListHmacFunction](docs/sdks/hmacfunctions/README.md#listhmacfunction) - Get a list of HmacFunction objects
* [CreateHmacFunction](docs/sdks/hmacfunctions/README.md#createhmacfunction) - Create HmacFunction
* [GetHmacFunctionByID](docs/sdks/hmacfunctions/README.md#gethmacfunctionbyid) - Get HmacFunction by ID
* [UpdateHmacFunctionByID](docs/sdks/hmacfunctions/README.md#updatehmacfunctionbyid) - Update HmacFunction
* [DeleteHmacFunctionByID](docs/sdks/hmacfunctions/README.md#deletehmacfunctionbyid) - Delete HmacFunction

### [Ingest](docs/sdks/ingest/README.md)

* [CreateEdgeFileIngest](docs/sdks/ingest/README.md#createedgefileingest) - Ingest a specified file through a specified pipeline to a specified destination or send to routes.

### [Inputs](docs/sdks/inputs/README.md)

* [ListInput](docs/sdks/inputs/README.md#listinput) - Get a list of Input objects
* [CreateInput](docs/sdks/inputs/README.md#createinput) - Create Input
* [GetInputByID](docs/sdks/inputs/README.md#getinputbyid) - Get Input by ID
* [UpdateInputByID](docs/sdks/inputs/README.md#updateinputbyid) - Update Input
* [DeleteInputByID](docs/sdks/inputs/README.md#deleteinputbyid) - Delete Input
* [CreateInputHecTokenByID](docs/sdks/inputs/README.md#createinputhectokenbyid) - Add token and optional metadata to an existing hec input
* [UpdateInputHecTokenByIDAndToken](docs/sdks/inputs/README.md#updateinputhectokenbyidandtoken) - Update token metadata on existing hec input
* [ListInputStatus](docs/sdks/inputs/README.md#listinputstatus) - Get a list of InputStatus objects
* [GetInputStatusByID](docs/sdks/inputs/README.md#getinputstatusbyid) - Get InputStatus by ID

### [Iometrics](docs/sdks/iometrics/README.md)

* [ListIoMetricsEntry](docs/sdks/iometrics/README.md#listiometricsentry) - Get a list of IoMetricsEntry objects
* [GetIoMetricsEntryByID](docs/sdks/iometrics/README.md#getiometricsentrybyid) - Get IoMetricsEntry by ID
* [UpdateIoMetricsEntryByID](docs/sdks/iometrics/README.md#updateiometricsentrybyid) - Update IoMetricsEntry
* [DeleteIoMetricsEntryByID](docs/sdks/iometrics/README.md#deleteiometricsentrybyid) - Delete IoMetricsEntry

### [Jobs](docs/sdks/jobs/README.md)

* [GetJobResultsByID](docs/sdks/jobs/README.md#getjobresultsbyid) - Get results for a discover job by instance id
* [UpdateJobsCancelByID](docs/sdks/jobs/README.md#updatejobscancelbyid) - Cancel a job by instance id
* [CreateJobs](docs/sdks/jobs/README.md#createjobs) - Run or schedule a job
* [GetJobs](docs/sdks/jobs/README.md#getjobs) - Get info on jobs
* [DeleteJobsByID](docs/sdks/jobs/README.md#deletejobsbyid) - Remove job from job inspector by instance id
* [GetJobsByID](docs/sdks/jobs/README.md#getjobsbyid) - Get job info by instance id
* [UpdateJobsKeepByID](docs/sdks/jobs/README.md#updatejobskeepbyid) - prevent job from being deleted automatically
* [UpdateJobsPauseByID](docs/sdks/jobs/README.md#updatejobspausebyid) - Pause a job by instance id
* [UpdateJobsResumeByID](docs/sdks/jobs/README.md#updatejobsresumebyid) - Resume a job by instance id
* [GetJobsResultsByIDAndGroup](docs/sdks/jobs/README.md#getjobsresultsbyidandgroup) - Get results for a discover job by instance id
* [GetJobsErrorsByIDAndGroup](docs/sdks/jobs/README.md#getjobserrorsbyidandgroup) - Get Task errors for a job by id
* [GetJobsErrorsByID](docs/sdks/jobs/README.md#getjobserrorsbyid) - Get Task errors for a job by id

### [Keys](docs/sdks/keys/README.md)

* [ListKeyMetadataEntity](docs/sdks/keys/README.md#listkeymetadataentity) - Get a list of KeyMetadataEntity objects
* [CreateKeyMetadataEntity](docs/sdks/keys/README.md#createkeymetadataentity) - Create KeyMetadataEntity
* [GetKeyMetadataEntityByID](docs/sdks/keys/README.md#getkeymetadataentitybyid) - Get KeyMetadataEntity by ID
* [UpdateKeyMetadataEntityByID](docs/sdks/keys/README.md#updatekeymetadataentitybyid) - Update KeyMetadataEntity
* [DeleteKeyMetadataEntityByID](docs/sdks/keys/README.md#deletekeymetadataentitybyid) - Delete KeyMetadataEntity

### [KubeLogs](docs/sdks/kubelogs/README.md)

* [CreateEdgeKubeLogs](docs/sdks/kubelogs/README.md#createedgekubelogs) - Make a request to the K8s API logs endpoint

### [KubeProxy](docs/sdks/kubeproxy/README.md)

* [GetEdgeKubeProxy](docs/sdks/kubeproxy/README.md#getedgekubeproxy) - Make a GET request to the K8s API

### [Lake](docs/sdks/lake/README.md)

* [CreateCriblLakeDatasetByLakeID](docs/sdks/lake/README.md#createcribllakedatasetbylakeid) - Create a Dataset in the specified Lake
* [GetCriblLakeDatasetByLakeID](docs/sdks/lake/README.md#getcribllakedatasetbylakeid) - Get the list of Dataset contained in the specified Lake
* [DeleteCriblLakeDatasetByLakeIDAndID](docs/sdks/lake/README.md#deletecribllakedatasetbylakeidandid) - Delete a Dataset in the specified Lake
* [GetCriblLakeDatasetByLakeIDAndID](docs/sdks/lake/README.md#getcribllakedatasetbylakeidandid) - Get a Dataset in the specified Lake
* [UpdateCriblLakeDatasetByLakeIDAndID](docs/sdks/lake/README.md#updatecribllakedatasetbylakeidandid) - Update a Dataset in the specified Lake
* [CreateCriblLakeStorageLocationByLakeID](docs/sdks/lake/README.md#createcribllakestoragelocationbylakeid) - Create a Storage Location in the specified Lake
* [GetCriblLakeStorageLocationByLakeID](docs/sdks/lake/README.md#getcribllakestoragelocationbylakeid) - Get the list of Storage Locations contained in the specified Lake
* [DeleteCriblLakeStorageLocationByLakeIDAndID](docs/sdks/lake/README.md#deletecribllakestoragelocationbylakeidandid) - Delete a Storage Location in the specified Lake
* [GetCriblLakeStorageLocationByLakeIDAndID](docs/sdks/lake/README.md#getcribllakestoragelocationbylakeidandid) - Get a Storage Location in the specified Lake
* [UpdateCriblLakeStorageLocationByLakeIDAndID](docs/sdks/lake/README.md#updatecribllakestoragelocationbylakeidandid) - Update a Storage Location in the specified Lake

### [Licenses](docs/sdks/licenses/README.md)

* [ListLicense](docs/sdks/licenses/README.md#listlicense) - Get a list of License objects
* [CreateLicense](docs/sdks/licenses/README.md#createlicense) - Add a license to your deployment
* [GetLicenseByID](docs/sdks/licenses/README.md#getlicensebyid) - Get License by ID
* [DeleteLicenseByID](docs/sdks/licenses/README.md#deletelicensebyid) - Delete License
* [GetLicense](docs/sdks/licenses/README.md#getlicense) - Get license usage metrics, aggregated by day, up to last 90 days

### [Logger](docs/sdks/logger/README.md)

* [ListLoggerConfig](docs/sdks/logger/README.md#listloggerconfig) - Get a list of LoggerConfig objects
* [GetLoggerConfigByID](docs/sdks/logger/README.md#getloggerconfigbyid) - Get LoggerConfig by ID
* [UpdateLoggerConfigByID](docs/sdks/logger/README.md#updateloggerconfigbyid) - Update LoggerConfig
* [DeleteLoggerConfigByID](docs/sdks/logger/README.md#deleteloggerconfigbyid) - Delete LoggerConfig

### [Logging](docs/sdks/logging/README.md)

* [GetSystemJobsLogsByIDAndGroupID](docs/sdks/logging/README.md#getsystemjobslogsbyidandgroupid) - Get contents of the log file
* [GetSystemLogsByID](docs/sdks/logging/README.md#getsystemlogsbyid) - Get contents of the log file
* [GetSystemLogs](docs/sdks/logging/README.md#getsystemlogs) - Get a list of log files
* [GetSystemLogsSearch](docs/sdks/logging/README.md#getsystemlogssearch) - Get contents of the log file

### [Lookups](docs/sdks/lookups/README.md)

* [ListLookupFile](docs/sdks/lookups/README.md#listlookupfile) - Get a list of LookupFile objects
* [CreateLookupFile](docs/sdks/lookups/README.md#createlookupfile) - Create LookupFile
* [UpdateLookupFile](docs/sdks/lookups/README.md#updatelookupfile) - Upload LookupFile
* [GetLookupFileByID](docs/sdks/lookups/README.md#getlookupfilebyid) - Get LookupFile by ID
* [UpdateLookupFileByID](docs/sdks/lookups/README.md#updatelookupfilebyid) - Update LookupFile
* [DeleteLookupFileByID](docs/sdks/lookups/README.md#deletelookupfilebyid) - Delete LookupFile
* [UpdateLookupFileCloneByID](docs/sdks/lookups/README.md#updatelookupfileclonebyid) - Clone a lookup from one context to another

### [Macros](docs/sdks/macros/README.md)

* [ListSearchMacro](docs/sdks/macros/README.md#listsearchmacro) - Get a list of SearchMacro objects
* [CreateSearchMacro](docs/sdks/macros/README.md#createsearchmacro) - Create SearchMacro
* [GetSearchMacroByID](docs/sdks/macros/README.md#getsearchmacrobyid) - Get SearchMacro by ID
* [UpdateSearchMacroByID](docs/sdks/macros/README.md#updatesearchmacrobyid) - Update SearchMacro
* [DeleteSearchMacroByID](docs/sdks/macros/README.md#deletesearchmacrobyid) - Delete SearchMacro

### [Messages](docs/sdks/messages/README.md)

* [ListBulletinMessage](docs/sdks/messages/README.md#listbulletinmessage) - Get a list of BulletinMessage objects
* [CreateBulletinMessage](docs/sdks/messages/README.md#createbulletinmessage) - Create BulletinMessage
* [GetBulletinMessageByID](docs/sdks/messages/README.md#getbulletinmessagebyid) - Get BulletinMessage by ID
* [DeleteBulletinMessageByID](docs/sdks/messages/README.md#deletebulletinmessagebyid) - Delete BulletinMessage

### [Metrics](docs/sdks/metrics/README.md)

* [CreateSystemMetricsEnum](docs/sdks/metrics/README.md#createsystemmetricsenum) - Enumerate all internal system metrics
* [GetSystemMetrics](docs/sdks/metrics/README.md#getsystemmetrics) - Query raw internal system metrics
* [CreateSystemMetricsQuery](docs/sdks/metrics/README.md#createsystemmetricsquery) - Aggregate raw internal system metrics

### [Notifications](docs/sdks/notifications/README.md)

* [ListNotification](docs/sdks/notifications/README.md#listnotification) - Get a list of Notification objects
* [CreateNotification](docs/sdks/notifications/README.md#createnotification) - Create Notification
* [GetNotificationByID](docs/sdks/notifications/README.md#getnotificationbyid) - Get Notification by ID
* [UpdateNotificationByID](docs/sdks/notifications/README.md#updatenotificationbyid) - Update Notification
* [DeleteNotificationByID](docs/sdks/notifications/README.md#deletenotificationbyid) - Delete Notification

### [NotificationTargets](docs/sdks/notificationtargets/README.md)

* [ListNotificationTarget](docs/sdks/notificationtargets/README.md#listnotificationtarget) - Get a list of NotificationTarget objects
* [CreateNotificationTarget](docs/sdks/notificationtargets/README.md#createnotificationtarget) - Create NotificationTarget
* [GetNotificationTargetByID](docs/sdks/notificationtargets/README.md#getnotificationtargetbyid) - Get NotificationTarget by ID
* [UpdateNotificationTargetByID](docs/sdks/notificationtargets/README.md#updatenotificationtargetbyid) - Update NotificationTarget
* [DeleteNotificationTargetByID](docs/sdks/notificationtargets/README.md#deletenotificationtargetbyid) - Delete NotificationTarget

### [OnlyOnPrem](docs/sdks/onlyonprem/README.md)

* [ListUser](docs/sdks/onlyonprem/README.md#listuser) - Get a list of User objects
* [CreateUser](docs/sdks/onlyonprem/README.md#createuser) - Create User
* [GetUserByID](docs/sdks/onlyonprem/README.md#getuserbyid) - Get User by ID
* [UpdateUserByID](docs/sdks/onlyonprem/README.md#updateuserbyid) - Update User properties – admin use only
* [DeleteUserByID](docs/sdks/onlyonprem/README.md#deleteuserbyid) - Delete User
* [UpdateUserInfoByID](docs/sdks/onlyonprem/README.md#updateuserinfobyid) - Update User except for their roles

### [Outputs](docs/sdks/outputs/README.md)

* [ListOutput](docs/sdks/outputs/README.md#listoutput) - Get a list of Output objects
* [CreateOutput](docs/sdks/outputs/README.md#createoutput) - Create Output
* [GetOutputByID](docs/sdks/outputs/README.md#getoutputbyid) - Get Output by ID
* [UpdateOutputByID](docs/sdks/outputs/README.md#updateoutputbyid) - Update Output
* [DeleteOutputByID](docs/sdks/outputs/README.md#deleteoutputbyid) - Delete Output
* [DeleteOutputPqByID](docs/sdks/outputs/README.md#deleteoutputpqbyid) - Clears destination persistent queue
* [GetOutputPqByID](docs/sdks/outputs/README.md#getoutputpqbyid) - Retrieves status of latest clear PQ job for an output
* [GetOutputSamplesByID](docs/sdks/outputs/README.md#getoutputsamplesbyid) - Retrieve samples data for the specified output. Used to get sample data for the test action.
* [CreateOutputTestByID](docs/sdks/outputs/README.md#createoutputtestbyid) - Send sample data to an output to validate configuration or test connectivity
* [ListOutputStatus](docs/sdks/outputs/README.md#listoutputstatus) - Get a list of OutputStatus objects
* [GetOutputStatusByID](docs/sdks/outputs/README.md#getoutputstatusbyid) - Get OutputStatus by ID

### [Packs](docs/sdks/packs/README.md)

* [CreatePacks](docs/sdks/packs/README.md#createpacks) - Install Pack
* [GetPacks](docs/sdks/packs/README.md#getpacks) - Get info on packs
* [UpdatePacks](docs/sdks/packs/README.md#updatepacks) - Upload Pack
* [DeletePacksByID](docs/sdks/packs/README.md#deletepacksbyid) - Uninstall Pack from the system
* [UpdatePacksByID](docs/sdks/packs/README.md#updatepacksbyid) - Upgrade Pack
* [CreatePacksClone](docs/sdks/packs/README.md#createpacksclone) - Clone Pack
* [GetPacksRefsByID](docs/sdks/packs/README.md#getpacksrefsbyid) - Returns all entities that reference objects from outside of the pack
* [GetPacksExportByID](docs/sdks/packs/README.md#getpacksexportbyid) - Export Pack
* [CreatePacksPublishByID](docs/sdks/packs/README.md#createpackspublishbyid) - Publish Pack back to its repository

### [Parquetschemas](docs/sdks/parquetschemas/README.md)

* [ListSchema](docs/sdks/parquetschemas/README.md#listschema) - Get a list of Schema objects
* [CreateSchema](docs/sdks/parquetschemas/README.md#createschema) - Create Schema
* [GetSchemaByID](docs/sdks/parquetschemas/README.md#getschemabyid) - Get Schema by ID
* [UpdateSchemaByID](docs/sdks/parquetschemas/README.md#updateschemabyid) - Update Schema
* [DeleteSchemaByID](docs/sdks/parquetschemas/README.md#deleteschemabyid) - Delete Schema

### [Parsers](docs/sdks/parsers/README.md)

* [ListParser](docs/sdks/parsers/README.md#listparser) - Get a list of Parser objects
* [CreateParser](docs/sdks/parsers/README.md#createparser) - Create Parser
* [GetParserByID](docs/sdks/parsers/README.md#getparserbyid) - Get Parser by ID
* [UpdateParserByID](docs/sdks/parsers/README.md#updateparserbyid) - Update Parser
* [DeleteParserByID](docs/sdks/parsers/README.md#deleteparserbyid) - Delete Parser

### [Pipelines](docs/sdks/pipelines/README.md)

* [GetSystemProjectsPipelinesByGroupIDAndProjectID](docs/sdks/pipelines/README.md#getsystemprojectspipelinesbygroupidandprojectid) - Get A list of Pipeline objects for specified Project
* [CreateSystemProjectsPipelinesByGroupIDAndProjectID](docs/sdks/pipelines/README.md#createsystemprojectspipelinesbygroupidandprojectid) - Create Pipeline
* [GetSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineID](docs/sdks/pipelines/README.md#getsystemprojectspipelinesbygroupidandprojectidandpipelineid) - Get Pipeline by ID in specified Project
* [UpdateSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineID](docs/sdks/pipelines/README.md#updatesystemprojectspipelinesbygroupidandprojectidandpipelineid) - Update Pipeline in specified Project
* [DeleteSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineID](docs/sdks/pipelines/README.md#deletesystemprojectspipelinesbygroupidandprojectidandpipelineid) - Delete Pipeline in specified Project
* [ListPipeline](docs/sdks/pipelines/README.md#listpipeline) - Get a list of Pipeline objects
* [CreatePipeline](docs/sdks/pipelines/README.md#createpipeline) - Create Pipeline
* [GetPipelineByID](docs/sdks/pipelines/README.md#getpipelinebyid) - Get Pipeline by ID
* [UpdatePipelineByID](docs/sdks/pipelines/README.md#updatepipelinebyid) - Update Pipeline
* [DeletePipelineByID](docs/sdks/pipelines/README.md#deletepipelinebyid) - Delete Pipeline
* [GetPipelineByPack](docs/sdks/pipelines/README.md#getpipelinebypack) - Get a list of Pipeline objects within a Pack
* [CreatePipelineByPack](docs/sdks/pipelines/README.md#createpipelinebypack) - Create Pipeline within a Pack
* [GetPipelineByPackAndID](docs/sdks/pipelines/README.md#getpipelinebypackandid) - Get Pipeline by ID within a Pack
* [UpdatePipelineByPackAndID](docs/sdks/pipelines/README.md#updatepipelinebypackandid) - Update Pipeline within a Pack
* [DeletePipelineByPackAndID](docs/sdks/pipelines/README.md#deletepipelinebypackandid) - Delete Pipeline within a Pack

### [Policies](docs/sdks/policies/README.md)

* [ListPolicyRule](docs/sdks/policies/README.md#listpolicyrule) - Get a list of PolicyRule objects
* [CreatePolicyRule](docs/sdks/policies/README.md#createpolicyrule) - Create PolicyRule
* [GetPolicyRuleByID](docs/sdks/policies/README.md#getpolicyrulebyid) - Get PolicyRule by ID
* [UpdatePolicyRuleByID](docs/sdks/policies/README.md#updatepolicyrulebyid) - Update PolicyRule
* [DeletePolicyRuleByID](docs/sdks/policies/README.md#deletepolicyrulebyid) - Delete PolicyRule

### [Preview](docs/sdks/preview/README.md)

* [CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionID](docs/sdks/preview/README.md#createsystemprojectssubscriptionscapturebygroupidandsubscriptionid) - Capture live incoming data from a particular Project and Subscription at the Subscription
* [CreateSystemProjectsCaptureByGroupIDAndProjectID](docs/sdks/preview/README.md#createsystemprojectscapturebygroupidandprojectid) - Capture live incoming data from a particular Project and Subscription at the Destination
* [CreateSystemProjectsPreviewByGroupIDAndProjectID](docs/sdks/preview/README.md#createsystemprojectspreviewbygroupidandprojectid) - Sends sample events through a Pipeline  for specified Project and returns the results
* [CreateSystemCapture](docs/sdks/preview/README.md#createsystemcapture) - Capture live incoming data
* [CreateSystemProjectsCaptureByProjectID](docs/sdks/preview/README.md#createsystemprojectscapturebyprojectid) - Capture live incoming data from a particular project and subscription at the destination
* [CreateSystemProjectsSubscriptionsCaptureByProjectIDAndSubscriptionID](docs/sdks/preview/README.md#createsystemprojectssubscriptionscapturebyprojectidandsubscriptionid) - Capture live incoming data from a particular project and subscription at the subscription
* [CreatePreview](docs/sdks/preview/README.md#createpreview) - Sends sample events through a pipeline and returns the results

### [Processes](docs/sdks/processes/README.md)

* [GetSystemProcesses](docs/sdks/processes/README.md#getsystemprocesses) - Get a list of processes under management

### [Profiler](docs/sdks/profiler/README.md)

* [ListProfilerItem](docs/sdks/profiler/README.md#listprofileritem) - Get a list of ProfilerItem objects
* [CreateProfilerItem](docs/sdks/profiler/README.md#createprofileritem) - Create ProfilerItem
* [GetProfilerItemByID](docs/sdks/profiler/README.md#getprofileritembyid) - Get ProfilerItem by ID
* [UpdateProfilerItemByID](docs/sdks/profiler/README.md#updateprofileritembyid) - Update ProfilerItem
* [DeleteProfilerItemByID](docs/sdks/profiler/README.md#deleteprofileritembyid) - Delete ProfilerItem

### [Projects](docs/sdks/projects/README.md)

* [GetSystemProjectsSubscriptionsByGroupIDByAndProjectID](docs/sdks/projects/README.md#getsystemprojectssubscriptionsbygroupidbyandprojectid) - Get the Subscriptions associated with the Project
* [GetSystemProjectsVersionCountByGroupIDAndProjectID](docs/sdks/projects/README.md#getsystemprojectsversioncountbygroupidandprojectid) - Get the count of files of changed
* [GetSystemProjectsVersionDiffByGroupIDAndProjectID](docs/sdks/projects/README.md#getsystemprojectsversiondiffbygroupidandprojectid) - Get the textual diff for given commit
* [GetSystemProjectsVersionFilesByGroupIDAndProjectID](docs/sdks/projects/README.md#getsystemprojectsversionfilesbygroupidandprojectid) - Get the files changed
* [CreateSystemProjectsVersionCommitByGroupIDAndProjectID](docs/sdks/projects/README.md#createsystemprojectsversioncommitbygroupidandprojectid) - Commit project changes
* [CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionID](docs/sdks/projects/README.md#createsystemprojectssubscriptionscapturebygroupidandsubscriptionid) - Capture live incoming data from a particular Project and Subscription at the Subscription
* [CreateSystemProjectsCaptureByGroupIDAndProjectID](docs/sdks/projects/README.md#createsystemprojectscapturebygroupidandprojectid) - Capture live incoming data from a particular Project and Subscription at the Destination
* [CreateSystemProjectsPreviewByGroupIDAndProjectID](docs/sdks/projects/README.md#createsystemprojectspreviewbygroupidandprojectid) - Sends sample events through a Pipeline  for specified Project and returns the results
* [GetSystemProjectsPipelinesByGroupIDAndProjectID](docs/sdks/projects/README.md#getsystemprojectspipelinesbygroupidandprojectid) - Get A list of Pipeline objects for specified Project
* [CreateSystemProjectsPipelinesByGroupIDAndProjectID](docs/sdks/projects/README.md#createsystemprojectspipelinesbygroupidandprojectid) - Create Pipeline
* [GetSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineID](docs/sdks/projects/README.md#getsystemprojectspipelinesbygroupidandprojectidandpipelineid) - Get Pipeline by ID in specified Project
* [UpdateSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineID](docs/sdks/projects/README.md#updatesystemprojectspipelinesbygroupidandprojectidandpipelineid) - Update Pipeline in specified Project
* [DeleteSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineID](docs/sdks/projects/README.md#deletesystemprojectspipelinesbygroupidandprojectidandpipelineid) - Delete Pipeline in specified Project
* [CreateSystemProjectsVersionCommitByProjectID](docs/sdks/projects/README.md#createsystemprojectsversioncommitbyprojectid) - Commit project changes.
* [GetSystemProjectsVersionCountByProjectID](docs/sdks/projects/README.md#getsystemprojectsversioncountbyprojectid) - get the count of files of changed
* [GetSystemProjectsVersionDiffByProjectID](docs/sdks/projects/README.md#getsystemprojectsversiondiffbyprojectid) - get the textual diff for given commit
* [GetSystemProjectsVersionFilesByProjectID](docs/sdks/projects/README.md#getsystemprojectsversionfilesbyprojectid) - get the files changed
* [CreateSystemProjectsVersionRevertByProjectID](docs/sdks/projects/README.md#createsystemprojectsversionrevertbyprojectid) - Revert project changes.
* [GetSystemProjectsVersionShowByProjectID](docs/sdks/projects/README.md#getsystemprojectsversionshowbyprojectid) - Show project changes.
* [ListProject](docs/sdks/projects/README.md#listproject) - Get a list of Project objects
* [CreateProject](docs/sdks/projects/README.md#createproject) - Create Project
* [GetProjectByID](docs/sdks/projects/README.md#getprojectbyid) - Get Project by ID
* [UpdateProjectByID](docs/sdks/projects/README.md#updateprojectbyid) - Update Project
* [DeleteProjectByID](docs/sdks/projects/README.md#deleteprojectbyid) - Delete Project
* [GetProjectACLByID](docs/sdks/projects/README.md#getprojectaclbyid) - Get Project ACL
* [CreateProjectACLApplyByID](docs/sdks/projects/README.md#createprojectaclapplybyid) - Modify Project ACL
* [GetProjectACLTeamsByID](docs/sdks/projects/README.md#getprojectaclteamsbyid) - Get Project Teams
* [CreateProjectACLTeamsApplyByID](docs/sdks/projects/README.md#createprojectaclteamsapplybyid) - Modify Project Teams ACL
* [GetProjectDestinationsByProjectID](docs/sdks/projects/README.md#getprojectdestinationsbyprojectid) - Lists destinations associated with this project.
* [GetSubscriptionByProjectID](docs/sdks/projects/README.md#getsubscriptionbyprojectid) - Get the subscriptions associated with the project

### [Protobuflibraries](docs/sdks/protobuflibraries/README.md)

* [GetProtobufLibraryConfig](docs/sdks/protobuflibraries/README.md#getprotobuflibraryconfig) - Show list of Protobuf encodings for a given Library
* [GetProtobufLibraryConfigByID](docs/sdks/protobuflibraries/README.md#getprotobuflibraryconfigbyid) - Show Library by Id
* [GetProtobufLibraryConfigEncodingsByIDAndEncID](docs/sdks/protobuflibraries/README.md#getprotobuflibraryconfigencodingsbyidandencid) - Show Protobuf library encodings by encoding id
* [GetProtobufLibraryConfigEncodingsByID](docs/sdks/protobuflibraries/README.md#getprotobuflibraryconfigencodingsbyid) - Show list of Protobuf encodings for a given Library

### [Query](docs/sdks/query/README.md)

* [GetSearchQuery](docs/sdks/query/README.md#getsearchquery) - Runs the query and returns the results

### [Regexes](docs/sdks/regexes/README.md)

* [ListRegexLibEntry](docs/sdks/regexes/README.md#listregexlibentry) - Get a list of RegexLibEntry objects
* [CreateRegexLibEntry](docs/sdks/regexes/README.md#createregexlibentry) - Create RegexLibEntry
* [GetRegexLibEntryByID](docs/sdks/regexes/README.md#getregexlibentrybyid) - Get RegexLibEntry by ID
* [UpdateRegexLibEntryByID](docs/sdks/regexes/README.md#updateregexlibentrybyid) - Update RegexLibEntry
* [DeleteRegexLibEntryByID](docs/sdks/regexes/README.md#deleteregexlibentrybyid) - Delete RegexLibEntry

### [Roles](docs/sdks/roles/README.md)

* [ListRole](docs/sdks/roles/README.md#listrole) - Get a list of Role objects
* [CreateRole](docs/sdks/roles/README.md#createrole) - Create Role
* [GetRoleByID](docs/sdks/roles/README.md#getrolebyid) - Get Role by ID
* [UpdateRoleByID](docs/sdks/roles/README.md#updaterolebyid) - Update Role
* [DeleteRoleByID](docs/sdks/roles/README.md#deleterolebyid) - Delete Role

### [Routes](docs/sdks/routes/README.md)

* [ListRoutes](docs/sdks/routes/README.md#listroutes) - Get a list of Routes objects
* [GetRoutesByID](docs/sdks/routes/README.md#getroutesbyid) - Get Routes by ID
* [UpdateRoutesByID](docs/sdks/routes/README.md#updateroutesbyid) - Update Routes
* [CreateRoutesAppendByID](docs/sdks/routes/README.md#createroutesappendbyid) - Appends routes to the end of the routing table
* [GetRoutesByPack](docs/sdks/routes/README.md#getroutesbypack) - Get a list of Routes objects within a Pack
* [GetRoutesByPackAndID](docs/sdks/routes/README.md#getroutesbypackandid) - Get Routes by ID within a Pack
* [UpdateRoutesByPackAndID](docs/sdks/routes/README.md#updateroutesbypackandid) - Update Routes within a Pack
* [CreateRoutesAppendByPackAndID](docs/sdks/routes/README.md#createroutesappendbypackandid) - Appends routes to the end of the routing table within a Pack

### [Samples](docs/sdks/samples/README.md)

* [CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionID](docs/sdks/samples/README.md#createsystemprojectssubscriptionscapturebygroupidandsubscriptionid) - Capture live incoming data from a particular Project and Subscription at the Subscription
* [CreateSystemProjectsCaptureByGroupIDAndProjectID](docs/sdks/samples/README.md#createsystemprojectscapturebygroupidandprojectid) - Capture live incoming data from a particular Project and Subscription at the Destination
* [CreateSystemProjectsPreviewByGroupIDAndProjectID](docs/sdks/samples/README.md#createsystemprojectspreviewbygroupidandprojectid) - Sends sample events through a Pipeline  for specified Project and returns the results
* [ListDataSample](docs/sdks/samples/README.md#listdatasample) - Get a list of DataSample objects
* [CreateDataSample](docs/sdks/samples/README.md#createdatasample) - Create DataSample
* [GetDataSampleByID](docs/sdks/samples/README.md#getdatasamplebyid) - Get DataSample by ID
* [UpdateDataSampleByID](docs/sdks/samples/README.md#updatedatasamplebyid) - Update DataSample
* [DeleteDataSampleByID](docs/sdks/samples/README.md#deletedatasamplebyid) - Delete DataSample
* [GetDataSampleContentByID](docs/sdks/samples/README.md#getdatasamplecontentbyid) - Get sample content by ID
* [CreateSystemCapture](docs/sdks/samples/README.md#createsystemcapture) - Capture live incoming data
* [CreateSystemProjectsCaptureByProjectID](docs/sdks/samples/README.md#createsystemprojectscapturebyprojectid) - Capture live incoming data from a particular project and subscription at the destination
* [CreateSystemProjectsSubscriptionsCaptureByProjectIDAndSubscriptionID](docs/sdks/samples/README.md#createsystemprojectssubscriptionscapturebyprojectidandsubscriptionid) - Capture live incoming data from a particular project and subscription at the subscription
* [CreatePreview](docs/sdks/samples/README.md#createpreview) - Sends sample events through a pipeline and returns the results

### [SavedJobs](docs/sdks/savedjobs/README.md)

* [ListSavedJob](docs/sdks/savedjobs/README.md#listsavedjob) - Get a list of SavedJob objects
* [CreateSavedJob](docs/sdks/savedjobs/README.md#createsavedjob) - Create SavedJob
* [GetSavedJobByID](docs/sdks/savedjobs/README.md#getsavedjobbyid) - Get SavedJob by ID
* [UpdateSavedJobByID](docs/sdks/savedjobs/README.md#updatesavedjobbyid) - Update SavedJob
* [DeleteSavedJobByID](docs/sdks/savedjobs/README.md#deletesavedjobbyid) - Delete SavedJob

### [SavedQueries](docs/sdks/savedqueries/README.md)

* [ListSavedQuery](docs/sdks/savedqueries/README.md#listsavedquery) - Get a list of SavedQuery objects
* [CreateSavedQuery](docs/sdks/savedqueries/README.md#createsavedquery) - Create SavedQuery
* [GetSavedQueryByID](docs/sdks/savedqueries/README.md#getsavedquerybyid) - Get SavedQuery by ID
* [UpdateSavedQueryByID](docs/sdks/savedqueries/README.md#updatesavedquerybyid) - Update SavedQuery
* [DeleteSavedQueryByID](docs/sdks/savedqueries/README.md#deletesavedquerybyid) - Delete SavedQuery

### [Schemas](docs/sdks/schemas/README.md)

* [ListLibSchemas](docs/sdks/schemas/README.md#listlibschemas) - Get a list of Schema objects
* [CreateLibSchemas](docs/sdks/schemas/README.md#createlibschemas) - Create Schema
* [GetLibSchemasByID](docs/sdks/schemas/README.md#getlibschemasbyid) - Get Schema by ID
* [UpdateLibSchemasByID](docs/sdks/schemas/README.md#updatelibschemasbyid) - Update Schema
* [DeleteLibSchemasByID](docs/sdks/schemas/README.md#deletelibschemasbyid) - Delete Schema
* [GetSchemaLibByPack](docs/sdks/schemas/README.md#getschemalibbypack) - Get a list of Schema objects within a Pack
* [CreateSchemaLibByPack](docs/sdks/schemas/README.md#createschemalibbypack) - Create Schema within a Pack
* [GetSchemaLibByPackAndID](docs/sdks/schemas/README.md#getschemalibbypackandid) - Get Schema by ID within a Pack
* [UpdateSchemaLibByPackAndID](docs/sdks/schemas/README.md#updateschemalibbypackandid) - Update Schema within a Pack
* [DeleteSchemaLibByPackAndID](docs/sdks/schemas/README.md#deleteschemalibbypackandid) - Delete Schema within a Pack

### [Scripts](docs/sdks/scripts/README.md)

* [ListScript](docs/sdks/scripts/README.md#listscript) - Get a list of Script objects
* [CreateScript](docs/sdks/scripts/README.md#createscript) - Create Script
* [GetScriptByID](docs/sdks/scripts/README.md#getscriptbyid) - Get Script by ID
* [UpdateScriptByID](docs/sdks/scripts/README.md#updatescriptbyid) - Update Script
* [DeleteScriptByID](docs/sdks/scripts/README.md#deletescriptbyid) - Delete Script

### [SdsRules](docs/sdks/sdsrules/README.md)

* [ListSensitiveDataRule](docs/sdks/sdsrules/README.md#listsensitivedatarule) - Get a list of SensitiveDataRule objects
* [CreateSensitiveDataRule](docs/sdks/sdsrules/README.md#createsensitivedatarule) - Create SensitiveDataRule
* [GetSensitiveDataRuleByID](docs/sdks/sdsrules/README.md#getsensitivedatarulebyid) - Get SensitiveDataRule by ID
* [UpdateSensitiveDataRuleByID](docs/sdks/sdsrules/README.md#updatesensitivedatarulebyid) - Update SensitiveDataRule
* [DeleteSensitiveDataRuleByID](docs/sdks/sdsrules/README.md#deletesensitivedatarulebyid) - Delete SensitiveDataRule

### [Search](docs/sdks/search/README.md)

* [ListSearchJob](docs/sdks/search/README.md#listsearchjob) - Get a list of SearchJob objects
* [CreateSearchJob](docs/sdks/search/README.md#createsearchjob) - Create SearchJob
* [GetSearchJobByID](docs/sdks/search/README.md#getsearchjobbyid) - Get SearchJob by ID
* [UpdateSearchJobByID](docs/sdks/search/README.md#updatesearchjobbyid) - Update SearchJob
* [DeleteSearchJobByID](docs/sdks/search/README.md#deletesearchjobbyid) - Delete SearchJob
* [CreateSearchJobDispatchExecutorsByID](docs/sdks/search/README.md#createsearchjobdispatchexecutorsbyid) - internal endpoint, dispatch search *id* to worker nodes filtered by worker node filter using RPC
* [GetSearchJobSettingsByID](docs/sdks/search/README.md#getsearchjobsettingsbyid) - Get search job settings
* [CreateSearchJobSettingsByID](docs/sdks/search/README.md#createsearchjobsettingsbyid) - Save search job settings
* [GetSearchJobStatusByID](docs/sdks/search/README.md#getsearchjobstatusbyid) - Get job status
* [GetSearchJobTimelineByID](docs/sdks/search/README.md#getsearchjobtimelinebyid) - Get search timeline
* [GetSearchJobFieldSummariesByID](docs/sdks/search/README.md#getsearchjobfieldsummariesbyid) - List field summaries
* [GetSearchJobLogsByID](docs/sdks/search/README.md#getsearchjoblogsbyid) - Get search logs
* [CreateSearchEventBreakerPreview](docs/sdks/search/README.md#createsearcheventbreakerpreview) - Runs an event breaker rule on the specified data
* [GetSearchDocs](docs/sdks/search/README.md#getsearchdocs) - Get Search documentation
* [GetSearchHealthcheck](docs/sdks/search/README.md#getsearchhealthcheck) - Get health check metric for search
* [GetSearchJobsMetricsByID](docs/sdks/search/README.md#getsearchjobsmetricsbyid) - Get search job metrics
* [GetSearchJobMetrics](docs/sdks/search/README.md#getsearchjobmetrics) - List metrics for all search jobs
* [CreateSearchPreview](docs/sdks/search/README.md#createsearchpreview) - Applies a query snippet on a set of input events for preview
* [CreateSearchUIMetrics](docs/sdks/search/README.md#createsearchuimetrics) - Send metric events from the UI endpoint
* [GetSearchJobsDiagByID](docs/sdks/search/README.md#getsearchjobsdiagbyid) - Get full search job data
* [GetSearchJobsLogsByIDAndFilename](docs/sdks/search/README.md#getsearchjobslogsbyidandfilename) - Get the contents of a log file
* [GetSearchJobsStagesResultsByIDAndStageID](docs/sdks/search/README.md#getsearchjobsstagesresultsbyidandstageid) - List search results for a given stage. Note that this cannot be the root stage!
* [GetSearchJobsResultsByID](docs/sdks/search/README.md#getsearchjobsresultsbyid) - List search results, when lower/upper bound is provided, offset is relative to the time range.
* [GetSearchJobsResultsPollByID](docs/sdks/search/README.md#getsearchjobsresultspollbyid) - List search results

### [Secrets](docs/sdks/secrets/README.md)

* [ListRestSecret](docs/sdks/secrets/README.md#listrestsecret) - Get a list of RestSecret objects
* [CreateRestSecret](docs/sdks/secrets/README.md#createrestsecret) - Create RestSecret
* [GetRestSecretByID](docs/sdks/secrets/README.md#getrestsecretbyid) - Get RestSecret by ID
* [UpdateRestSecretByID](docs/sdks/secrets/README.md#updaterestsecretbyid) - Update RestSecret
* [DeleteRestSecretByID](docs/sdks/secrets/README.md#deleterestsecretbyid) - Delete RestSecret

### [Security](docs/sdks/security/README.md)

* [GetSecurityKmsConfig](docs/sdks/security/README.md#getsecuritykmsconfig) - Get Cribl KMS config
* [UpdateSecurityKmsConfig](docs/sdks/security/README.md#updatesecuritykmsconfig) - Update Cribl KMS config
* [GetSecurityKmsHealth](docs/sdks/security/README.md#getsecuritykmshealth) - Get Cribl KMS health

### [Subscriptions](docs/sdks/subscriptions/README.md)

* [GetSystemProjectsSubscriptionsByGroupIDByAndProjectID](docs/sdks/subscriptions/README.md#getsystemprojectssubscriptionsbygroupidbyandprojectid) - Get the Subscriptions associated with the Project
* [ListSubscription](docs/sdks/subscriptions/README.md#listsubscription) - Get a list of Subscription objects
* [CreateSubscription](docs/sdks/subscriptions/README.md#createsubscription) - Create subscription
* [GetSubscriptionByID](docs/sdks/subscriptions/README.md#getsubscriptionbyid) - Get Subscription by ID
* [UpdateSubscriptionByID](docs/sdks/subscriptions/README.md#updatesubscriptionbyid) - Update subscription
* [DeleteSubscriptionByID](docs/sdks/subscriptions/README.md#deletesubscriptionbyid) - Delete subscription
* [GetSubscriptionByProjectID](docs/sdks/subscriptions/README.md#getsubscriptionbyprojectid) - Get the subscriptions associated with the project

### [System](docs/sdks/system/README.md)

* [UpdateChangelogViewed](docs/sdks/system/README.md#updatechangelogviewed) - Update changelog viewed state
* [CreateSystemDistributedUpgradeCancelByGroup](docs/sdks/system/README.md#createsystemdistributedupgradecancelbygroup) - Cancel a running group upgrade
* [CreateSystemDistributedUpgradeStageByGroup](docs/sdks/system/README.md#createsystemdistributedupgradestagebygroup) - Stage distributed group upgrade
* [CreateSystemDistributedUpgradeByGroup](docs/sdks/system/README.md#createsystemdistributedupgradebygroup) - Execute distributed group upgrade
* [GetSystemDistributedUpgradeDownloadByFile](docs/sdks/system/README.md#getsystemdistributedupgradedownloadbyfile) - Get the previously downloaded Cribl package
* [ReloadSettings](docs/sdks/system/README.md#reloadsettings) - Reload Cribl settings from the filesystem
* [TriggerRestart](docs/sdks/system/README.md#triggerrestart) - Restart Cribl server
* [GetSystemSettingsAuth](docs/sdks/system/README.md#getsystemsettingsauth) - Get authentication settings
* [UpdateSystemSettingsAuth](docs/sdks/system/README.md#updatesystemsettingsauth) - Update authentication settings
* [GetSystemSettingsConf](docs/sdks/system/README.md#getsystemsettingsconf) - Get Cribl system settings
* [UpdateSystemSettingsConf](docs/sdks/system/README.md#updatesystemsettingsconf) - Update Cribl system settings
* [GetSystemSettingsGitSettings](docs/sdks/system/README.md#getsystemsettingsgitsettings) - Get git settings
* [UpdateSystemSettingsGitSettings](docs/sdks/system/README.md#updatesystemsettingsgitsettings) - Update git settings
* [GetSystemSettingsSearchLimits](docs/sdks/system/README.md#getsystemsettingssearchlimits) - Get search limits
* [~~GetSystemSettings~~](docs/sdks/system/README.md#getsystemsettings) - Get Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/redis-cache-limits, /system/services-limits, /system/settings/git-settings, and /system/settings/conf respectively :warning: **Deprecated**
* [~~UpdateSystemSettings~~](docs/sdks/system/README.md#updatesystemsettings) - Update Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/settings/git-settings, /system/settings/auth and /system/settings/conf respectively :warning: **Deprecated**
* [GetSystemSettingsCribl](docs/sdks/system/README.md#getsystemsettingscribl) - Get public settings visible to any logged user
* [GetSystemSettingsUpgrade](docs/sdks/system/README.md#getsystemsettingsupgrade) - Get a list of Cribl versions available for upgrade
* [CreateSystemSettingsUpgradeByVersion](docs/sdks/system/README.md#createsystemsettingsupgradebyversion) - Upgrade Cribl to a given version
* [CreateSystemSettingsUpgradeFromPackage](docs/sdks/system/README.md#createsystemsettingsupgradefrompackage) - Upgrade master node with the provided package

### [Teams](docs/sdks/teams/README.md)

* [CreateTeam](docs/sdks/teams/README.md#createteam) - Create Team
* [GetTeam](docs/sdks/teams/README.md#getteam) - Get a list of Team objects
* [GetTeamByID](docs/sdks/teams/README.md#getteambyid) - Get Team by ID
* [UpdateTeamByID](docs/sdks/teams/README.md#updateteambyid) - Update Team
* [DeleteTeamByID](docs/sdks/teams/README.md#deleteteambyid) - Delete Team
* [GetTeamACLByID](docs/sdks/teams/README.md#getteamaclbyid) - Get Teams's Access Control List
* [GetTeamRolesByID](docs/sdks/teams/README.md#getteamrolesbyid) - Get user's product roles
* [GetTeamUsersByID](docs/sdks/teams/README.md#getteamusersbyid) - Get users on a team
* [CreateTeamUsersByID](docs/sdks/teams/README.md#createteamusersbyid) - Update existing users on a team – admin use only
* [GetProductsGroupsACLTeamsByProductAndID](docs/sdks/teams/README.md#getproductsgroupsaclteamsbyproductandid) - ACL of team with permissions for resources in this Group

### [TeamsACL](docs/sdks/teamsacl/README.md)

* [GetDatasetACLTeamsByID](docs/sdks/teamsacl/README.md#getdatasetaclteamsbyid) - Get Dataset Teams
* [CreateDatasetACLTeamsApplyByID](docs/sdks/teamsacl/README.md#createdatasetaclteamsapplybyid) - Modify Dataset Teams ACL
* [GetSearchDashboardACLTeamsByID](docs/sdks/teamsacl/README.md#getsearchdashboardaclteamsbyid) - Get SearchDashboard Teams
* [CreateSearchDashboardACLTeamsApplyByID](docs/sdks/teamsacl/README.md#createsearchdashboardaclteamsapplybyid) - Modify SearchDashboard Teams ACL
* [GetProjectACLTeamsByID](docs/sdks/teamsacl/README.md#getprojectaclteamsbyid) - Get Project Teams
* [CreateProjectACLTeamsApplyByID](docs/sdks/teamsacl/README.md#createprojectaclteamsapplybyid) - Modify Project Teams ACL

### [TrustPolicies](docs/sdks/trustpolicies/README.md)

* [ListTrustPolicy](docs/sdks/trustpolicies/README.md#listtrustpolicy) - Get a list of TrustPolicy objects

### [UIState](docs/sdks/uistate/README.md)

* [GetUIByKey](docs/sdks/uistate/README.md#getuibykey) - Get UI state by key
* [UpdateUIByKey](docs/sdks/uistate/README.md#updateuibykey) - Update UI state by key

### [UsageGroups](docs/sdks/usagegroups/README.md)

* [ListUsageGroup](docs/sdks/usagegroups/README.md#listusagegroup) - Get a list of UsageGroup objects
* [CreateUsageGroup](docs/sdks/usagegroups/README.md#createusagegroup) - Create UsageGroup
* [GetUsageGroupByID](docs/sdks/usagegroups/README.md#getusagegroupbyid) - Get UsageGroup by ID
* [UpdateUsageGroupByID](docs/sdks/usagegroups/README.md#updateusagegroupbyid) - Update UsageGroup
* [DeleteUsageGroupByID](docs/sdks/usagegroups/README.md#deleteusagegroupbyid) - Delete UsageGroup

### [Users](docs/sdks/users/README.md)

* [ListUser](docs/sdks/users/README.md#listuser) - Get a list of User objects
* [CreateUser](docs/sdks/users/README.md#createuser) - Create User
* [GetUserByID](docs/sdks/users/README.md#getuserbyid) - Get User by ID
* [UpdateUserByID](docs/sdks/users/README.md#updateuserbyid) - Update User properties – admin use only
* [DeleteUserByID](docs/sdks/users/README.md#deleteuserbyid) - Delete User
* [UpdateUserInfoByID](docs/sdks/users/README.md#updateuserinfobyid) - Update User except for their roles
* [GetProductsUsersACLByProductAndID](docs/sdks/users/README.md#getproductsusersaclbyproductandid) - Get user's Access Control List
* [GetProductsUsersByProduct](docs/sdks/users/README.md#getproductsusersbyproduct) - Get Users belonging to a product
* [DeleteProductsUsersCacheByProduct](docs/sdks/users/README.md#deleteproductsuserscachebyproduct) - Invalidate the members cache for a given product in SaaS deployment.

### [UsersACL](docs/sdks/usersacl/README.md)

* [GetDatasetACLByID](docs/sdks/usersacl/README.md#getdatasetaclbyid) - Get Dataset ACL
* [CreateDatasetACLApplyByID](docs/sdks/usersacl/README.md#createdatasetaclapplybyid) - Modify Dataset ACL
* [GetSearchDashboardACLByID](docs/sdks/usersacl/README.md#getsearchdashboardaclbyid) - Get SearchDashboard ACL
* [CreateSearchDashboardACLApplyByID](docs/sdks/usersacl/README.md#createsearchdashboardaclapplybyid) - Modify SearchDashboard ACL
* [GetProjectACLByID](docs/sdks/usersacl/README.md#getprojectaclbyid) - Get Project ACL
* [CreateProjectACLApplyByID](docs/sdks/usersacl/README.md#createprojectaclapplybyid) - Modify Project ACL
* [GetGroupsACLByID](docs/sdks/usersacl/README.md#getgroupsaclbyid) - ACL of members with permissions for resources in this Group

### [Versioning](docs/sdks/versioning/README.md)

* [GetSystemProjectsVersionCountByGroupIDAndProjectID](docs/sdks/versioning/README.md#getsystemprojectsversioncountbygroupidandprojectid) - Get the count of files of changed
* [GetSystemProjectsVersionDiffByGroupIDAndProjectID](docs/sdks/versioning/README.md#getsystemprojectsversiondiffbygroupidandprojectid) - Get the textual diff for given commit
* [GetSystemProjectsVersionFilesByGroupIDAndProjectID](docs/sdks/versioning/README.md#getsystemprojectsversionfilesbygroupidandprojectid) - Get the files changed
* [CreateSystemProjectsVersionCommitByGroupIDAndProjectID](docs/sdks/versioning/README.md#createsystemprojectsversioncommitbygroupidandprojectid) - Commit project changes
* [CreateSystemProjectsVersionCommitByProjectID](docs/sdks/versioning/README.md#createsystemprojectsversioncommitbyprojectid) - Commit project changes.
* [GetSystemProjectsVersionCountByProjectID](docs/sdks/versioning/README.md#getsystemprojectsversioncountbyprojectid) - get the count of files of changed
* [GetSystemProjectsVersionDiffByProjectID](docs/sdks/versioning/README.md#getsystemprojectsversiondiffbyprojectid) - get the textual diff for given commit
* [GetSystemProjectsVersionFilesByProjectID](docs/sdks/versioning/README.md#getsystemprojectsversionfilesbyprojectid) - get the files changed
* [CreateSystemProjectsVersionRevertByProjectID](docs/sdks/versioning/README.md#createsystemprojectsversionrevertbyprojectid) - Revert project changes.
* [GetSystemProjectsVersionShowByProjectID](docs/sdks/versioning/README.md#getsystemprojectsversionshowbyprojectid) - Show project changes.
* [UndoLastCommit](docs/sdks/versioning/README.md#undolastcommit) - undo the last commit
* [GetVersionBranch](docs/sdks/versioning/README.md#getversionbranch) - get the list of branches
* [CreateVersionCommit](docs/sdks/versioning/README.md#createversioncommit) - create a new commit containing the current configs the given log message describing the changes.
* [GetVersionCount](docs/sdks/versioning/README.md#getversioncount) - get the count of files of changed
* [GetVersionCurrentBranch](docs/sdks/versioning/README.md#getversioncurrentbranch) - returns git branch that the config is checked out to, if any
* [GetVersionDiff](docs/sdks/versioning/README.md#getversiondiff) - get the textual diff for given commit
* [GetVersionFiles](docs/sdks/versioning/README.md#getversionfiles) - get the files changed
* [GetVersionInfo](docs/sdks/versioning/README.md#getversioninfo) - Get info about versioning availability
* [CreateVersionPush](docs/sdks/versioning/README.md#createversionpush) - push the current configs to the remote repository.
* [CreateVersionRevert](docs/sdks/versioning/README.md#createversionrevert) - revert a commit
* [GetVersionShow](docs/sdks/versioning/README.md#getversionshow) - get the log message and textual diff for given commit
* [GetVersionStatus](docs/sdks/versioning/README.md#getversionstatus) - get the the working tree status
* [CreateVersionSync](docs/sdks/versioning/README.md#createversionsync) - syncs with remote repo via POST requests

### [Workers](docs/sdks/workers/README.md)

* [GetSummaryWorkers](docs/sdks/workers/README.md#getsummaryworkers) - get worker and edge nodes count
* [GetWorkers](docs/sdks/workers/README.md#getworkers) - get worker and edge nodes
* [UpdateWorkersRestart](docs/sdks/workers/README.md#updateworkersrestart) - restarts worker nodes

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Json Streaming [jsonl] -->
## Json Streaming

Json Streaming ([jsonl][jsonl-format] / [x-ndjson][x-ndjson]) content type can be used to stream content from certain operations. These operations expose the stream that can be consumed using a `for` loop in Go. The loop will terminate when the server no longer has any events to send and closes the underlying connection.

Here's an example of consuming a JSONL stream:

```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
	)

	res, err := s.Search.GetSearchJobsResultsByID(ctx, operations.GetSearchJobsResultsByIDRequest{
		ID: "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.SearchJobResults != nil {
		for res.SearchJobResults.Next() {
			event, _ := res.SearchJobResults.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

[jsonl-format]: https://jsonlines.org/
[x-ndjson]: https://github.com/ndjson/ndjson-spec
<!-- End Json Streaming [jsonl] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
	)

	res, err := s.Projects.GetSystemProjectsSubscriptionsByGroupIDByAndProjectID(ctx, "<id>", "<id>", operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
	)

	res, err := s.Projects.GetSystemProjectsSubscriptionsByGroupIDByAndProjectID(ctx, "<id>", "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetSystemProjectsSubscriptionsByGroupIDByAndProjectID` function may return the following errors:

| Error Type         | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| apierrors.Error    | 500         | application/json |
| apierrors.APIError | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
	)

	res, err := s.Projects.GetSystemProjectsSubscriptionsByGroupIDByAndProjectID(ctx, "<id>", "<id>")
	if err != nil {

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
