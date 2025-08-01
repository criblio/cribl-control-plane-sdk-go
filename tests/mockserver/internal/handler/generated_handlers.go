// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"context"
	"mockserver/internal/logging"
	"mockserver/internal/tracking"
	"net/http"
)

// GeneratedHandlers returns all generated handlers.
func GeneratedHandlers(ctx context.Context, dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) []*GeneratedHandler {
	return []*GeneratedHandler{
		NewGeneratedHandler(ctx, http.MethodDelete, "/master/groups/{id}", pathDeleteMasterGroupsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/packs/{id}", pathDeletePacksID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/pipelines/{id}", pathDeletePipelinesID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/products/lake/lakes/{lakeId}/datasets/{id}", pathDeleteProductsLakeLakesLakeIDDatasetsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/system/inputs/{id}", pathDeleteSystemInputsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/system/outputs/{id}", pathDeleteSystemOutputsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/system/outputs/{id}/pq", pathDeleteSystemOutputsIDPq(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/health", pathGetHealth(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/master/groups/{id}", pathGetMasterGroupsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/master/groups/{id}/acl", pathGetMasterGroupsIDACL(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/master/groups/{id}/configVersion", pathGetMasterGroupsIDConfigVersion(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/master/summary", pathGetMasterSummary(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/master/summary/workers", pathGetMasterSummaryWorkers(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/master/workers", pathGetMasterWorkers(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/packs", pathGetPacks(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/pipelines", pathGetPipelines(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/pipelines/{id}", pathGetPipelinesID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/products/lake/lakes/{lakeId}/datasets", pathGetProductsLakeLakesLakeIDDatasets(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/products/lake/lakes/{lakeId}/datasets/{id}", pathGetProductsLakeLakesLakeIDDatasetsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/products/{product}/groups", pathGetProductsProductGroups(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/products/{product}/groups/{id}/acl/teams", pathGetProductsProductGroupsIDACLTeams(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/routes", pathGetRoutes(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/routes/{id}", pathGetRoutesID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/system/inputs", pathGetSystemInputs(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/system/inputs/{id}", pathGetSystemInputsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/system/outputs", pathGetSystemOutputs(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/system/outputs/{id}", pathGetSystemOutputsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/system/outputs/{id}/pq", pathGetSystemOutputsIDPq(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/system/outputs/{id}/samples", pathGetSystemOutputsIDSamples(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/branch", pathGetVersionBranch(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/count", pathGetVersionCount(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/current-branch", pathGetVersionCurrentBranch(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/diff", pathGetVersionDiff(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/files", pathGetVersionFiles(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/info", pathGetVersionInfo(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/show", pathGetVersionShow(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/version/status", pathGetVersionStatus(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/master/groups/{id}", pathPatchMasterGroupsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/master/groups/{id}/deploy", pathPatchMasterGroupsIDDeploy(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/master/workers/restart", pathPatchMasterWorkersRestart(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/packs/{id}", pathPatchPacksID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/pipelines/{id}", pathPatchPipelinesID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/products/lake/lakes/{lakeId}/datasets/{id}", pathPatchProductsLakeLakesLakeIDDatasetsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/routes/{id}", pathPatchRoutesID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/system/inputs/{id}", pathPatchSystemInputsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/system/inputs/{id}/hectoken/{token}", pathPatchSystemInputsIDHectokenToken(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPatch, "/system/outputs/{id}", pathPatchSystemOutputsID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/auth/login", pathPostAuthLogin(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/packs", pathPostPacks(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/pipelines", pathPostPipelines(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/products/lake/lakes/{lakeId}/datasets", pathPostProductsLakeLakesLakeIDDatasets(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/products/{product}/groups", pathPostProductsProductGroups(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/routes/{id}/append", pathPostRoutesIDAppend(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/system/inputs", pathPostSystemInputs(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/system/inputs/{id}/hectoken", pathPostSystemInputsIDHectoken(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/system/outputs", pathPostSystemOutputs(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/system/outputs/{id}/test", pathPostSystemOutputsIDTest(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/version/commit", pathPostVersionCommit(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/version/push", pathPostVersionPush(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/version/revert", pathPostVersionRevert(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/version/sync", pathPostVersionSync(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/version/undo", pathPostVersionUndo(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPut, "/packs", pathPutPacks(dir, rt)),
	}
}
