//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

const (
	moduleName    = "armkeyvault"
	moduleVersion = "v1.3.0"
)

type AccessPolicyUpdateKind string

const (
	AccessPolicyUpdateKindAdd     AccessPolicyUpdateKind = "add"
	AccessPolicyUpdateKindRemove  AccessPolicyUpdateKind = "remove"
	AccessPolicyUpdateKindReplace AccessPolicyUpdateKind = "replace"
)

// PossibleAccessPolicyUpdateKindValues returns the possible values for the AccessPolicyUpdateKind const type.
func PossibleAccessPolicyUpdateKindValues() []AccessPolicyUpdateKind {
	return []AccessPolicyUpdateKind{
		AccessPolicyUpdateKindAdd,
		AccessPolicyUpdateKindRemove,
		AccessPolicyUpdateKindReplace,
	}
}

// ActionsRequired - A message indicating if changes on the service provider require any updates on the consumer.
type ActionsRequired string

const (
	ActionsRequiredNone ActionsRequired = "None"
)

// PossibleActionsRequiredValues returns the possible values for the ActionsRequired const type.
func PossibleActionsRequiredValues() []ActionsRequired {
	return []ActionsRequired{
		ActionsRequiredNone,
	}
}

// ActivationStatus - Activation Status
type ActivationStatus string

const (
	// ActivationStatusActive - The managed HSM Pool is active.
	ActivationStatusActive ActivationStatus = "Active"
	// ActivationStatusFailed - Failed to activate managed hsm.
	ActivationStatusFailed ActivationStatus = "Failed"
	// ActivationStatusNotActivated - The managed HSM Pool is not yet activated.
	ActivationStatusNotActivated ActivationStatus = "NotActivated"
	// ActivationStatusUnknown - An unknown error occurred while activating managed hsm.
	ActivationStatusUnknown ActivationStatus = "Unknown"
)

// PossibleActivationStatusValues returns the possible values for the ActivationStatus const type.
func PossibleActivationStatusValues() []ActivationStatus {
	return []ActivationStatus{
		ActivationStatusActive,
		ActivationStatusFailed,
		ActivationStatusNotActivated,
		ActivationStatusUnknown,
	}
}

type CertificatePermissions string

const (
	CertificatePermissionsAll            CertificatePermissions = "all"
	CertificatePermissionsBackup         CertificatePermissions = "backup"
	CertificatePermissionsCreate         CertificatePermissions = "create"
	CertificatePermissionsDelete         CertificatePermissions = "delete"
	CertificatePermissionsDeleteissuers  CertificatePermissions = "deleteissuers"
	CertificatePermissionsGet            CertificatePermissions = "get"
	CertificatePermissionsGetissuers     CertificatePermissions = "getissuers"
	CertificatePermissionsImport         CertificatePermissions = "import"
	CertificatePermissionsList           CertificatePermissions = "list"
	CertificatePermissionsListissuers    CertificatePermissions = "listissuers"
	CertificatePermissionsManagecontacts CertificatePermissions = "managecontacts"
	CertificatePermissionsManageissuers  CertificatePermissions = "manageissuers"
	CertificatePermissionsPurge          CertificatePermissions = "purge"
	CertificatePermissionsRecover        CertificatePermissions = "recover"
	CertificatePermissionsRestore        CertificatePermissions = "restore"
	CertificatePermissionsSetissuers     CertificatePermissions = "setissuers"
	CertificatePermissionsUpdate         CertificatePermissions = "update"
)

// PossibleCertificatePermissionsValues returns the possible values for the CertificatePermissions const type.
func PossibleCertificatePermissionsValues() []CertificatePermissions {
	return []CertificatePermissions{
		CertificatePermissionsAll,
		CertificatePermissionsBackup,
		CertificatePermissionsCreate,
		CertificatePermissionsDelete,
		CertificatePermissionsDeleteissuers,
		CertificatePermissionsGet,
		CertificatePermissionsGetissuers,
		CertificatePermissionsImport,
		CertificatePermissionsList,
		CertificatePermissionsListissuers,
		CertificatePermissionsManagecontacts,
		CertificatePermissionsManageissuers,
		CertificatePermissionsPurge,
		CertificatePermissionsRecover,
		CertificatePermissionsRestore,
		CertificatePermissionsSetissuers,
		CertificatePermissionsUpdate,
	}
}

// CreateMode - The vault's create mode to indicate whether the vault need to be recovered or not.
type CreateMode string

const (
	CreateModeDefault CreateMode = "default"
	CreateModeRecover CreateMode = "recover"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeRecover,
	}
}

// DeletionRecoveryLevel - The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then
// the object can be permanently deleted by a privileged user; otherwise, only the system can purge the
// object at the end of the retention interval.
type DeletionRecoveryLevel string

const (
	DeletionRecoveryLevelPurgeable                        DeletionRecoveryLevel = "Purgeable"
	DeletionRecoveryLevelRecoverable                      DeletionRecoveryLevel = "Recoverable"
	DeletionRecoveryLevelRecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
	DeletionRecoveryLevelRecoverablePurgeable             DeletionRecoveryLevel = "Recoverable+Purgeable"
)

// PossibleDeletionRecoveryLevelValues returns the possible values for the DeletionRecoveryLevel const type.
func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
	return []DeletionRecoveryLevel{
		DeletionRecoveryLevelPurgeable,
		DeletionRecoveryLevelRecoverable,
		DeletionRecoveryLevelRecoverableProtectedSubscription,
		DeletionRecoveryLevelRecoverablePurgeable,
	}
}

// GeoReplicationRegionProvisioningState - The current provisioning state.
type GeoReplicationRegionProvisioningState string

const (
	GeoReplicationRegionProvisioningStateCleanup         GeoReplicationRegionProvisioningState = "Cleanup"
	GeoReplicationRegionProvisioningStateDeleting        GeoReplicationRegionProvisioningState = "Deleting"
	GeoReplicationRegionProvisioningStateFailed          GeoReplicationRegionProvisioningState = "Failed"
	GeoReplicationRegionProvisioningStatePreprovisioning GeoReplicationRegionProvisioningState = "Preprovisioning"
	GeoReplicationRegionProvisioningStateProvisioning    GeoReplicationRegionProvisioningState = "Provisioning"
	GeoReplicationRegionProvisioningStateSucceeded       GeoReplicationRegionProvisioningState = "Succeeded"
)

// PossibleGeoReplicationRegionProvisioningStateValues returns the possible values for the GeoReplicationRegionProvisioningState const type.
func PossibleGeoReplicationRegionProvisioningStateValues() []GeoReplicationRegionProvisioningState {
	return []GeoReplicationRegionProvisioningState{
		GeoReplicationRegionProvisioningStateCleanup,
		GeoReplicationRegionProvisioningStateDeleting,
		GeoReplicationRegionProvisioningStateFailed,
		GeoReplicationRegionProvisioningStatePreprovisioning,
		GeoReplicationRegionProvisioningStateProvisioning,
		GeoReplicationRegionProvisioningStateSucceeded,
	}
}

// IdentityType - The type of identity.
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "Application"
	IdentityTypeKey             IdentityType = "Key"
	IdentityTypeManagedIdentity IdentityType = "ManagedIdentity"
	IdentityTypeUser            IdentityType = "User"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// JSONWebKeyCurveName - The elliptic curve name. For valid values, see JsonWebKeyCurveName.
type JSONWebKeyCurveName string

const (
	JSONWebKeyCurveNameP256  JSONWebKeyCurveName = "P-256"
	JSONWebKeyCurveNameP256K JSONWebKeyCurveName = "P-256K"
	JSONWebKeyCurveNameP384  JSONWebKeyCurveName = "P-384"
	JSONWebKeyCurveNameP521  JSONWebKeyCurveName = "P-521"
)

// PossibleJSONWebKeyCurveNameValues returns the possible values for the JSONWebKeyCurveName const type.
func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
	return []JSONWebKeyCurveName{
		JSONWebKeyCurveNameP256,
		JSONWebKeyCurveNameP256K,
		JSONWebKeyCurveNameP384,
		JSONWebKeyCurveNameP521,
	}
}

// JSONWebKeyOperation - The permitted JSON web key operations of the key. For more information, see JsonWebKeyOperation.
type JSONWebKeyOperation string

const (
	JSONWebKeyOperationDecrypt   JSONWebKeyOperation = "decrypt"
	JSONWebKeyOperationEncrypt   JSONWebKeyOperation = "encrypt"
	JSONWebKeyOperationImport    JSONWebKeyOperation = "import"
	JSONWebKeyOperationRelease   JSONWebKeyOperation = "release"
	JSONWebKeyOperationSign      JSONWebKeyOperation = "sign"
	JSONWebKeyOperationUnwrapKey JSONWebKeyOperation = "unwrapKey"
	JSONWebKeyOperationVerify    JSONWebKeyOperation = "verify"
	JSONWebKeyOperationWrapKey   JSONWebKeyOperation = "wrapKey"
)

// PossibleJSONWebKeyOperationValues returns the possible values for the JSONWebKeyOperation const type.
func PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation {
	return []JSONWebKeyOperation{
		JSONWebKeyOperationDecrypt,
		JSONWebKeyOperationEncrypt,
		JSONWebKeyOperationImport,
		JSONWebKeyOperationRelease,
		JSONWebKeyOperationSign,
		JSONWebKeyOperationUnwrapKey,
		JSONWebKeyOperationVerify,
		JSONWebKeyOperationWrapKey,
	}
}

// JSONWebKeyType - The type of the key. For valid values, see JsonWebKeyType.
type JSONWebKeyType string

const (
	JSONWebKeyTypeEC     JSONWebKeyType = "EC"
	JSONWebKeyTypeECHSM  JSONWebKeyType = "EC-HSM"
	JSONWebKeyTypeRSA    JSONWebKeyType = "RSA"
	JSONWebKeyTypeRSAHSM JSONWebKeyType = "RSA-HSM"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeEC,
		JSONWebKeyTypeECHSM,
		JSONWebKeyTypeRSA,
		JSONWebKeyTypeRSAHSM,
	}
}

type KeyPermissions string

const (
	KeyPermissionsAll               KeyPermissions = "all"
	KeyPermissionsBackup            KeyPermissions = "backup"
	KeyPermissionsCreate            KeyPermissions = "create"
	KeyPermissionsDecrypt           KeyPermissions = "decrypt"
	KeyPermissionsDelete            KeyPermissions = "delete"
	KeyPermissionsEncrypt           KeyPermissions = "encrypt"
	KeyPermissionsGet               KeyPermissions = "get"
	KeyPermissionsGetrotationpolicy KeyPermissions = "getrotationpolicy"
	KeyPermissionsImport            KeyPermissions = "import"
	KeyPermissionsList              KeyPermissions = "list"
	KeyPermissionsPurge             KeyPermissions = "purge"
	KeyPermissionsRecover           KeyPermissions = "recover"
	KeyPermissionsRelease           KeyPermissions = "release"
	KeyPermissionsRestore           KeyPermissions = "restore"
	KeyPermissionsRotate            KeyPermissions = "rotate"
	KeyPermissionsSetrotationpolicy KeyPermissions = "setrotationpolicy"
	KeyPermissionsSign              KeyPermissions = "sign"
	KeyPermissionsUnwrapKey         KeyPermissions = "unwrapKey"
	KeyPermissionsUpdate            KeyPermissions = "update"
	KeyPermissionsVerify            KeyPermissions = "verify"
	KeyPermissionsWrapKey           KeyPermissions = "wrapKey"
)

// PossibleKeyPermissionsValues returns the possible values for the KeyPermissions const type.
func PossibleKeyPermissionsValues() []KeyPermissions {
	return []KeyPermissions{
		KeyPermissionsAll,
		KeyPermissionsBackup,
		KeyPermissionsCreate,
		KeyPermissionsDecrypt,
		KeyPermissionsDelete,
		KeyPermissionsEncrypt,
		KeyPermissionsGet,
		KeyPermissionsGetrotationpolicy,
		KeyPermissionsImport,
		KeyPermissionsList,
		KeyPermissionsPurge,
		KeyPermissionsRecover,
		KeyPermissionsRelease,
		KeyPermissionsRestore,
		KeyPermissionsRotate,
		KeyPermissionsSetrotationpolicy,
		KeyPermissionsSign,
		KeyPermissionsUnwrapKey,
		KeyPermissionsUpdate,
		KeyPermissionsVerify,
		KeyPermissionsWrapKey,
	}
}

// KeyRotationPolicyActionType - The type of action.
type KeyRotationPolicyActionType string

const (
	KeyRotationPolicyActionTypeNotify KeyRotationPolicyActionType = "notify"
	KeyRotationPolicyActionTypeRotate KeyRotationPolicyActionType = "rotate"
)

// PossibleKeyRotationPolicyActionTypeValues returns the possible values for the KeyRotationPolicyActionType const type.
func PossibleKeyRotationPolicyActionTypeValues() []KeyRotationPolicyActionType {
	return []KeyRotationPolicyActionType{
		KeyRotationPolicyActionTypeNotify,
		KeyRotationPolicyActionTypeRotate,
	}
}

// ManagedHsmSKUFamily - SKU Family of the managed HSM Pool
type ManagedHsmSKUFamily string

const (
	ManagedHsmSKUFamilyB ManagedHsmSKUFamily = "B"
)

// PossibleManagedHsmSKUFamilyValues returns the possible values for the ManagedHsmSKUFamily const type.
func PossibleManagedHsmSKUFamilyValues() []ManagedHsmSKUFamily {
	return []ManagedHsmSKUFamily{
		ManagedHsmSKUFamilyB,
	}
}

// ManagedHsmSKUName - SKU of the managed HSM Pool
type ManagedHsmSKUName string

const (
	ManagedHsmSKUNameCustomB32  ManagedHsmSKUName = "Custom_B32"
	ManagedHsmSKUNameCustomB6   ManagedHsmSKUName = "Custom_B6"
	ManagedHsmSKUNameStandardB1 ManagedHsmSKUName = "Standard_B1"
)

// PossibleManagedHsmSKUNameValues returns the possible values for the ManagedHsmSKUName const type.
func PossibleManagedHsmSKUNameValues() []ManagedHsmSKUName {
	return []ManagedHsmSKUName{
		ManagedHsmSKUNameCustomB32,
		ManagedHsmSKUNameCustomB6,
		ManagedHsmSKUNameStandardB1,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// NetworkRuleAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used
// after the bypass property has been evaluated.
type NetworkRuleAction string

const (
	NetworkRuleActionAllow NetworkRuleAction = "Allow"
	NetworkRuleActionDeny  NetworkRuleAction = "Deny"
)

// PossibleNetworkRuleActionValues returns the possible values for the NetworkRuleAction const type.
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return []NetworkRuleAction{
		NetworkRuleActionAllow,
		NetworkRuleActionDeny,
	}
}

// NetworkRuleBypassOptions - Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'. If not specified
// the default is 'AzureServices'.
type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

// PossibleNetworkRuleBypassOptionsValues returns the possible values for the NetworkRuleBypassOptions const type.
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return []NetworkRuleBypassOptions{
		NetworkRuleBypassOptionsAzureServices,
		NetworkRuleBypassOptionsNone,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating     PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting     PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDisconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating     PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateDisconnected,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusDisconnected,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - Provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateActivated - The managed HSM pool is ready for normal use.
	ProvisioningStateActivated ProvisioningState = "Activated"
	// ProvisioningStateDeleting - The managed HSM Pool is currently being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Provisioning of the managed HSM Pool has failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - The managed HSM Pool is currently being provisioned.
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateRestoring - The managed HSM pool is being restored from full HSM backup.
	ProvisioningStateRestoring ProvisioningState = "Restoring"
	// ProvisioningStateSecurityDomainRestore - The managed HSM pool is waiting for a security domain restore action.
	ProvisioningStateSecurityDomainRestore ProvisioningState = "SecurityDomainRestore"
	// ProvisioningStateSucceeded - The managed HSM Pool has been full provisioned.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The managed HSM Pool is currently being updated.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateActivated,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateRestoring,
		ProvisioningStateSecurityDomainRestore,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Control permission to the managed HSM from public networks.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// Reason - The reason that a vault name could not be used. The Reason element is only returned if NameAvailable is false.
type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAccountNameInvalid,
		ReasonAlreadyExists,
	}
}

// SKUFamily - SKU family name
type SKUFamily string

const (
	SKUFamilyA SKUFamily = "A"
)

// PossibleSKUFamilyValues returns the possible values for the SKUFamily const type.
func PossibleSKUFamilyValues() []SKUFamily {
	return []SKUFamily{
		SKUFamilyA,
	}
}

// SKUName - SKU name to specify whether the key vault is a standard vault or a premium vault.
type SKUName string

const (
	SKUNamePremium  SKUName = "premium"
	SKUNameStandard SKUName = "standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePremium,
		SKUNameStandard,
	}
}

type SecretPermissions string

const (
	SecretPermissionsAll     SecretPermissions = "all"
	SecretPermissionsBackup  SecretPermissions = "backup"
	SecretPermissionsDelete  SecretPermissions = "delete"
	SecretPermissionsGet     SecretPermissions = "get"
	SecretPermissionsList    SecretPermissions = "list"
	SecretPermissionsPurge   SecretPermissions = "purge"
	SecretPermissionsRecover SecretPermissions = "recover"
	SecretPermissionsRestore SecretPermissions = "restore"
	SecretPermissionsSet     SecretPermissions = "set"
)

// PossibleSecretPermissionsValues returns the possible values for the SecretPermissions const type.
func PossibleSecretPermissionsValues() []SecretPermissions {
	return []SecretPermissions{
		SecretPermissionsAll,
		SecretPermissionsBackup,
		SecretPermissionsDelete,
		SecretPermissionsGet,
		SecretPermissionsList,
		SecretPermissionsPurge,
		SecretPermissionsRecover,
		SecretPermissionsRestore,
		SecretPermissionsSet,
	}
}

type StoragePermissions string

const (
	StoragePermissionsAll           StoragePermissions = "all"
	StoragePermissionsBackup        StoragePermissions = "backup"
	StoragePermissionsDelete        StoragePermissions = "delete"
	StoragePermissionsDeletesas     StoragePermissions = "deletesas"
	StoragePermissionsGet           StoragePermissions = "get"
	StoragePermissionsGetsas        StoragePermissions = "getsas"
	StoragePermissionsList          StoragePermissions = "list"
	StoragePermissionsListsas       StoragePermissions = "listsas"
	StoragePermissionsPurge         StoragePermissions = "purge"
	StoragePermissionsRecover       StoragePermissions = "recover"
	StoragePermissionsRegeneratekey StoragePermissions = "regeneratekey"
	StoragePermissionsRestore       StoragePermissions = "restore"
	StoragePermissionsSet           StoragePermissions = "set"
	StoragePermissionsSetsas        StoragePermissions = "setsas"
	StoragePermissionsUpdate        StoragePermissions = "update"
)

// PossibleStoragePermissionsValues returns the possible values for the StoragePermissions const type.
func PossibleStoragePermissionsValues() []StoragePermissions {
	return []StoragePermissions{
		StoragePermissionsAll,
		StoragePermissionsBackup,
		StoragePermissionsDelete,
		StoragePermissionsDeletesas,
		StoragePermissionsGet,
		StoragePermissionsGetsas,
		StoragePermissionsList,
		StoragePermissionsListsas,
		StoragePermissionsPurge,
		StoragePermissionsRecover,
		StoragePermissionsRegeneratekey,
		StoragePermissionsRestore,
		StoragePermissionsSet,
		StoragePermissionsSetsas,
		StoragePermissionsUpdate,
	}
}

// VaultProvisioningState - Provisioning state of the vault.
type VaultProvisioningState string

const (
	VaultProvisioningStateRegisteringDNS VaultProvisioningState = "RegisteringDns"
	VaultProvisioningStateSucceeded      VaultProvisioningState = "Succeeded"
)

// PossibleVaultProvisioningStateValues returns the possible values for the VaultProvisioningState const type.
func PossibleVaultProvisioningStateValues() []VaultProvisioningState {
	return []VaultProvisioningState{
		VaultProvisioningStateRegisteringDNS,
		VaultProvisioningStateSucceeded,
	}
}
