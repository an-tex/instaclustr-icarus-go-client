/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.5
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

type OneOfinlineResponse200 struct {
    CleanupOperationResponse
    DecommissionOperationResponse
    DrainOperationResponse
    FlushOperationResponse
    RebuildOperationResponse
    RefreshOperationResponse
    RestartOperationResponse
    ScrubOperationResponse
    UpgradeSsTablesOperationResponse
    ImportOperationResponse
    TruncateOperationResponse
    BackupOperationResponse
    RestoreOperationResponse
}
