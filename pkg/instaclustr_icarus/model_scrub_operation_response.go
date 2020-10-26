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

type ScrubOperationResponse struct {
	Type_ string `json:"type"`
	// keyspace to scrub 
	Keyspace string `json:"keyspace"`
	// tables to scrub, empty or not provided will scrub all tables in respective keyspace 
	Tables []string `json:"tables,omitempty"`
	// number of sstables to scrub simultanously, set to 0 to use all available compaction threads 
	Jobs int32 `json:"jobs,omitempty"`
	// scrubbed CFs will be snapshotted first, defaults to false 
	DisableSnapshot bool `json:"disableSnapshot,omitempty"`
	// skip corrupted partitions even when scrubbing counter tables, defaults to false 
	SkipCorrupted bool `json:"skipCorrupted,omitempty"`
	// do not validate columns using column validator, defaults to false 
	NoValidate bool `json:"noValidate,omitempty"`
	// Rewrites rows with overflowed expiration date affected by CASSANDRA-14092 with the maximum supported expiration date of 2038-01-19T03:14:06+00:00. The rows are rewritten with the original timestamp incremented by one millisecond to override/supersede any potential tombstone that may have been generated during compaction of the affected rows. 
	ReinsertOverflowedTTL bool `json:"reinsertOverflowedTTL,omitempty"`
	// unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller's perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint 
	Id string `json:"id"`
	// state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously 
	State string `json:"state"`
	// float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors. 
	Progress float64 `json:"progress"`
	// timestamp telling when this operation was created on Sidecar's side 
	CreationTime string `json:"creationTime"`
	// timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that 
	StartTime string `json:"startTime,omitempty"`
	// timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated. 
	CompletionTime string `json:"completionTime,omitempty"`
	Errors []ErrorObject `json:"errors,omitempty"`
}
