// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds

import (
	"time"
)

const (
	ClusterRoleStatusActive  = "ACTIVE"
	ClusterRoleStatusDeleted = "DELETED"
	ClusterRoleStatusPending = "PENDING"
)

const (
	ClusterStatusAvailable                  = "available"
	ClusterStatusBackingUp                  = "backing-up"
	ClusterStatusConfiguringIAMDatabaseAuth = "configuring-iam-database-auth"
	ClusterStatusCreating                   = "creating"
	ClusterStatusDeleting                   = "deleting"
	ClusterStatusMigrating                  = "migrating"
	ClusterStatusModifying                  = "modifying"
	ClusterStatusPreparingDataMigration     = "preparing-data-migration"
	ClusterStatusPromoting                  = "promoting"
	ClusterStatusRebooting                  = "rebooting"
	ClusterStatusRenaming                   = "renaming"
	ClusterStatusResettingMasterCredentials = "resetting-master-credentials"
	ClusterStatusScalingCompute             = "scaling-compute"
	ClusterStatusUpgrading                  = "upgrading"
)

const (
	ClusterSnapshotStatusAvailable = "available"
	ClusterSnapshotStatusCreating  = "creating"
)

const (
	storageTypeStandard    = "standard"
	storageTypeGP2         = "gp2"
	storageTypeGP3         = "gp3"
	storageTypeIO1         = "io1"
	storageTypeAuroraIOPT1 = "aurora-iopt1"
)

func StorageType_Values() []string {
	return []string{
		storageTypeStandard,
		storageTypeGP2,
		storageTypeGP3,
		storageTypeIO1,
		storageTypeAuroraIOPT1,
	}
}

const (
	InstanceEngineCustomPrefix        = "custom-"
	InstanceEngineDB2Advanced         = "db2-ae"
	InstanceEngineDB2Standard         = "db2-se"
	InstanceEngineMariaDB             = "mariadb"
	InstanceEngineMySQL               = "mysql"
	InstanceEngineOracleEnterprise    = "oracle-ee"
	InstanceEngineOracleEnterpriseCDB = "oracle-ee-cdb"
	InstanceEngineOracleStandard2     = "oracle-se2"
	InstanceEngineOracleStandard2CDB  = "oracle-se2-cdb"
	InstanceEnginePostgres            = "postgres"
	InstanceEngineSQLServerEnterprise = "sqlserver-ee"
	InstanceEngineSQLServerExpress    = "sqlserver-ex"
	InstanceEngineSQLServerStandard   = "sqlserver-se"
	InstanceEngineSQLServerWeb        = "sqlserver-ewb"
)

// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/accessing-monitoring.html#Overview.DBInstance.Status.
const (
	InstanceStatusAvailable                                    = "available"
	InstanceStatusBackingUp                                    = "backing-up"
	InstanceStatusConfiguringEnhancedMonitoring                = "configuring-enhanced-monitoring"
	InstanceStatusConfiguringIAMDatabaseAuth                   = "configuring-iam-database-auth"
	InstanceStatusConfiguringLogExports                        = "configuring-log-exports"
	InstanceStatusConvertingToVPC                              = "converting-to-vpc"
	InstanceStatusCreating                                     = "creating"
	InstanceStatusDeletePreCheck                               = "delete-precheck"
	InstanceStatusDeleting                                     = "deleting"
	InstanceStatusFailed                                       = "failed"
	InstanceStatusInaccessibleEncryptionCredentials            = "inaccessible-encryption-credentials"
	InstanceStatusInaccessibleEncryptionCredentialsRecoverable = "inaccessible-encryption-credentials-recoverable"
	InstanceStatusIncompatiblCreate                            = "incompatible-create"
	InstanceStatusIncompatibleNetwork                          = "incompatible-network"
	InstanceStatusIncompatibleOptionGroup                      = "incompatible-option-group"
	InstanceStatusIncompatibleParameters                       = "incompatible-parameters"
	InstanceStatusIncompatibleRestore                          = "incompatible-restore"
	InstanceStatusInsufficentCapacity                          = "insufficient-capacity"
	InstanceStatusMaintenance                                  = "maintenance"
	InstanceStatusModifying                                    = "modifying"
	InstanceStatusMovingToVPC                                  = "moving-to-vpc"
	InstanceStatusRebooting                                    = "rebooting"
	InstanceStatusResettingMasterCredentials                   = "resetting-master-credentials"
	InstanceStatusRenaming                                     = "renaming"
	InstanceStatusRestoreError                                 = "restore-error"
	InstanceStatusStarting                                     = "starting"
	InstanceStatusStopped                                      = "stopped"
	InstanceStatusStopping                                     = "stopping"
	InstanceStatusStorageFull                                  = "storage-full"
	InstanceStatusStorageOptimization                          = "storage-optimization"
	InstanceStatusUpgrading                                    = "upgrading"
)

const (
	GlobalClusterStatusAvailable = "available"
	GlobalClusterStatusCreating  = "creating"
	GlobalClusterStatusDeleting  = "deleting"
	GlobalClusterStatusModifying = "modifying"
	GlobalClusterStatusUpgrading = "upgrading"
)

const (
	EventSubscriptionStatusActive    = "active"
	EventSubscriptionStatusCreating  = "creating"
	EventSubscriptionStatusDeleting  = "deleting"
	EventSubscriptionStatusModifying = "modifying"
)

const (
	DBSnapshotAvailable = "available"
	DBSnapshotCreating  = "creating"
)

const (
	ClusterEngineAuroraMySQL      = "aurora-mysql"
	ClusterEngineAuroraPostgreSQL = "aurora-postgresql"
	ClusterEngineMySQL            = "mysql"
	ClusterEnginePostgres         = "postgres"
	ClusterEngineCustomPrefix     = "custom-"
)

func ClusterEngine_Values() []string {
	return []string{
		ClusterEngineAuroraMySQL,
		ClusterEngineAuroraPostgreSQL,
		ClusterEngineMySQL,
		ClusterEnginePostgres,
	}
}

func ClusterInstanceEngine_Values() []string {
	return []string{
		ClusterEngineAuroraMySQL,
		ClusterEngineAuroraPostgreSQL,
	}
}

const (
	GlobalClusterEngineAurora           = "aurora"
	GlobalClusterEngineAuroraMySQL      = "aurora-mysql"
	GlobalClusterEngineAuroraPostgreSQL = "aurora-postgresql"
)

func GlobalClusterEngine_Values() []string {
	return []string{
		GlobalClusterEngineAurora,
		GlobalClusterEngineAuroraMySQL,
		GlobalClusterEngineAuroraPostgreSQL,
	}
}

const (
	EngineModeGlobal        = "global"
	EngineModeMultiMaster   = "multimaster"
	EngineModeParallelQuery = "parallelquery"
	EngineModeProvisioned   = "provisioned"
	EngineModeServerless    = "serverless"
)

func EngineMode_Values() []string {
	return []string{
		EngineModeGlobal,
		EngineModeMultiMaster,
		EngineModeParallelQuery,
		EngineModeProvisioned,
		EngineModeServerless,
	}
}

const (
	ExportableLogTypeAgent      = "agent"
	ExportableLogTypeAlert      = "alert"
	ExportableLogTypeAudit      = "audit"
	ExportableLogTypeDiagLog    = "diag.log"
	ExportableLogTypeError      = "error"
	ExportableLogTypeGeneral    = "general"
	ExportableLogTypeListener   = "listener"
	ExportableLogTypeNotifyLog  = "notify.log"
	ExportableLogTypeOEMAgent   = "oemagent"
	ExportableLogTypePostgreSQL = "postgresql"
	ExportableLogTypeSlowQuery  = "slowquery"
	ExportableLogTypeTrace      = "trace"
	ExportableLogTypeUpgrade    = "upgrade"
)

func ClusterExportableLogType_Values() []string {
	return []string{
		ExportableLogTypeAudit,
		ExportableLogTypeError,
		ExportableLogTypeGeneral,
		ExportableLogTypePostgreSQL,
		ExportableLogTypeSlowQuery,
		ExportableLogTypeUpgrade,
	}
}

func InstanceExportableLogType_Values() []string {
	return []string{
		ExportableLogTypeAgent,
		ExportableLogTypeAlert,
		ExportableLogTypeAudit,
		ExportableLogTypeDiagLog,
		ExportableLogTypeError,
		ExportableLogTypeGeneral,
		ExportableLogTypeListener,
		ExportableLogTypeNotifyLog,
		ExportableLogTypeOEMAgent,
		ExportableLogTypePostgreSQL,
		ExportableLogTypeSlowQuery,
		ExportableLogTypeTrace,
		ExportableLogTypeUpgrade,
	}
}

const (
	NetworkTypeDual = "DUAL"
	NetworkTypeIPv4 = "IPV4"
)

func NetworkType_Values() []string {
	return []string{
		NetworkTypeDual,
		NetworkTypeIPv4,
	}
}

const (
	RestoreTypeCopyOnWrite = "copy-on-write"
	RestoreTypeFullCopy    = "full-copy"
)

func RestoreType_Values() []string {
	return []string{
		RestoreTypeCopyOnWrite,
		RestoreTypeFullCopy,
	}
}

const (
	TimeoutActionForceApplyCapacityChange = "ForceApplyCapacityChange"
	TimeoutActionRollbackCapacityChange   = "RollbackCapacityChange"
)

func TimeoutAction_Values() []string {
	return []string{
		TimeoutActionForceApplyCapacityChange,
		TimeoutActionRollbackCapacityChange,
	}
}

const (
	backupTargetOutposts = "outposts"
	backupTargetRegion   = "region"
)

func backupTarget_Values() []string {
	return []string{
		backupTargetOutposts,
		backupTargetRegion,
	}
}

const (
	propagationTimeout = 2 * time.Minute
)

const (
	ResNameTags = "Tags"
)

const (
	ReservedInstanceStateActive         = "active"
	ReservedInstanceStateRetired        = "retired"
	ReservedInstanceStatePaymentPending = "payment-pending"
)
