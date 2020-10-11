/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.1
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

type OneOfbody struct {
    CleanupOperationRequest
    DecommissionOperationRequest
    DrainOperationRequest
    FlushOperationRequest
    RebuildOperationRequest
    RefreshOperationRequest
    RestartOperationRequest
    ScrubOperationRequest
    UpgradeSsTablesOperationRequest
    ImportOperationRequest
    TruncateOperationRequest
    BackupOperationRequest
    RestoreOperationRequest
}