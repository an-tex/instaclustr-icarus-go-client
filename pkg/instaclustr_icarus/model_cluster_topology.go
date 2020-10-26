/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.2
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

type ClusterTopology struct {
	// Timestamp this topology was created at
	Timestamp float64 `json:"timestamp,omitempty"`
	// name of a cluster
	ClusterName string `json:"clusterName,omitempty"`
	// Database schema version of given cluster
	SchemaVersion string `json:"schemaVersion,omitempty"`
	Topology []NodeTopology `json:"topology,omitempty"`
}
