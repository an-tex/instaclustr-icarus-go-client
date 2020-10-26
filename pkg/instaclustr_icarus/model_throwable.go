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

type Throwable struct {
	Cause *ThrowableCause `json:"cause,omitempty"`
	Stacktrace *Stacktrace `json:"stacktrace,omitempty"`
	Message string `json:"message,omitempty"`
	LocalizedMessage string `json:"localizedMessage,omitempty"`
}