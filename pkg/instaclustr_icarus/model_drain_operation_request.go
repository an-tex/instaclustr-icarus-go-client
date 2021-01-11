/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.3
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

// drain a node, this operation will be successful only in case a node is in state NORMAL, when a node was already drained or it is in the middle of draining, this operation returns immediately. 
type DrainOperationRequest struct {
	Type_ string `json:"type"`
}
