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

// restarts a Cassandra node this Sidecar talks to. This operation will be successfully carried out only in case both Cassandra node as well as this Sidecar are running in Kubernetes. There is an assumption that Cassandra node and Sidecar are running in separate Docker containers as part of the same Kubernetes Pod. The restart is done by executing \"/bin/kill 1\" of Cassandra container where pid 1 stands for Cassandra process. The logic behind restart is that Kubernetes detects that container has finished and it will start it again on its own. Before Cassandra node is stopped, it is drained first so there are not any requests comming to this node whatsover hence restart is safe. 
type RestartOperationRequest struct {
	Type_ string `json:"type"`
}
