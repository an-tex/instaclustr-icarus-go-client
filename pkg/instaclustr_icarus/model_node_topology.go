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

type NodeTopology struct {
	// hostname of Cassandra node
	Hostname string `json:"hostname,omitempty"`
	// name of a cluster this node belongs to
	Cluster string `json:"cluster,omitempty"`
	// name of a data center this node belongs to
	Dc string `json:"dc,omitempty"`
	// name of a rack this node belongs to
	Rack string `json:"rack,omitempty"`
	// Cassandra host id of this node
	NodeId string `json:"nodeId,omitempty"`
	// IP address of a node
	IpAddress string `json:"ipAddress,omitempty"`
}
