# Go API client for instaclustr-icarus

REST API for Instaclustr Icarus - a sidecar for Cassandra

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./instaclustr-icarus"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4567*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigApi* | [**ConfigGet**](docs/ConfigApi.md#configget) | **Get** /config | returns configuration of a Cassandra node as in its cassandra.yaml file
*OperationsApi* | [**OperationsGet**](docs/OperationsApi.md#operationsget) | **Get** /operations | All operations of Icarus
*OperationsApi* | [**OperationsOperationIdGet**](docs/OperationsApi.md#operationsoperationidget) | **Get** /operations/{operationId} | gets operation by its ID
*OperationsApi* | [**OperationsPost**](docs/OperationsApi.md#operationspost) | **Post** /operations | Submits an operation to this Sidecar
*StatusApi* | [**StatusGet**](docs/StatusApi.md#statusget) | **Get** /status | returns a state of a Cassandra node
*TopologyApi* | [**TopologyDcGet**](docs/TopologyApi.md#topologydcget) | **Get** /topology/{dc} | returns topology of a datacenter of a cluster
*TopologyApi* | [**TopologyGet**](docs/TopologyApi.md#topologyget) | **Get** /topology | returns topology of a cluster as seen from this node
*VersionApi* | [**VersionCassandraGet**](docs/VersionApi.md#versioncassandraget) | **Get** /version/cassandra | returns version of Cassandra node
*VersionApi* | [**VersionGet**](docs/VersionApi.md#versionget) | **Get** /version | returns version of Cassandra Sidecar itself
*VersionApi* | [**VersionSchemaGet**](docs/VersionApi.md#versionschemaget) | **Get** /version/schema | returns schema version this Cassandra node is on, same as calling StorageServiceMBean#getSchemaVersion
*VersionApi* | [**VersionSidecarGet**](docs/VersionApi.md#versionsidecarget) | **Get** /version/sidecar | alias for /version endpoint, returns version of Cassandra Sidecar itself

## Documentation For Models

 - [AllOfRestoreOperationRequestImport_](docs/AllOfRestoreOperationRequestImport_.md)
 - [BackupOperationRequest](docs/BackupOperationRequest.md)
 - [BackupOperationResponse](docs/BackupOperationResponse.md)
 - [BaseOperation](docs/BaseOperation.md)
 - [Body](docs/Body.md)
 - [CassandraSchemaVersion](docs/CassandraSchemaVersion.md)
 - [CassandraSchemaVersionException](docs/CassandraSchemaVersionException.md)
 - [CassandraStatus](docs/CassandraStatus.md)
 - [CassandraStatusException](docs/CassandraStatusException.md)
 - [CassandraVersion](docs/CassandraVersion.md)
 - [CleanupOperationRequest](docs/CleanupOperationRequest.md)
 - [CleanupOperationResponse](docs/CleanupOperationResponse.md)
 - [ClusterTopology](docs/ClusterTopology.md)
 - [DataRate](docs/DataRate.md)
 - [DecommissionOperationRequest](docs/DecommissionOperationRequest.md)
 - [DecommissionOperationResponse](docs/DecommissionOperationResponse.md)
 - [DrainOperationRequest](docs/DrainOperationRequest.md)
 - [DrainOperationResponse](docs/DrainOperationResponse.md)
 - [FlushOperationRequest](docs/FlushOperationRequest.md)
 - [FlushOperationResponse](docs/FlushOperationResponse.md)
 - [ImportOperationRequest](docs/ImportOperationRequest.md)
 - [ImportOperationResponse](docs/ImportOperationResponse.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [NodeTopology](docs/NodeTopology.md)
 - [OneOfbody](docs/OneOfbody.md)
 - [OneOfinlineResponse200](docs/OneOfinlineResponse200.md)
 - [RebuildOperationRequest](docs/RebuildOperationRequest.md)
 - [RebuildOperationResponse](docs/RebuildOperationResponse.md)
 - [RefreshOperationRequest](docs/RefreshOperationRequest.md)
 - [RefreshOperationResponse](docs/RefreshOperationResponse.md)
 - [RestartOperationRequest](docs/RestartOperationRequest.md)
 - [RestartOperationResponse](docs/RestartOperationResponse.md)
 - [RestoreOperationRequest](docs/RestoreOperationRequest.md)
 - [RestoreOperationResponse](docs/RestoreOperationResponse.md)
 - [ScrubOperationRequest](docs/ScrubOperationRequest.md)
 - [ScrubOperationResponse](docs/ScrubOperationResponse.md)
 - [SidecarVersion](docs/SidecarVersion.md)
 - [TokenRange](docs/TokenRange.md)
 - [TruncateOperationRequest](docs/TruncateOperationRequest.md)
 - [TruncateOperationResponse](docs/TruncateOperationResponse.md)
 - [UpgradeSsTablesOperationRequest](docs/UpgradeSsTablesOperationRequest.md)
 - [UpgradeSsTablesOperationResponse](docs/UpgradeSsTablesOperationResponse.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

support@instaclustr.com
