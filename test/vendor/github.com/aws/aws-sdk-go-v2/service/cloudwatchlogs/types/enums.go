// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AnomalyDetectorStatus string

// Enum values for AnomalyDetectorStatus
const (
	AnomalyDetectorStatusInitializing AnomalyDetectorStatus = "INITIALIZING"
	AnomalyDetectorStatusTraining     AnomalyDetectorStatus = "TRAINING"
	AnomalyDetectorStatusAnalyzing    AnomalyDetectorStatus = "ANALYZING"
	AnomalyDetectorStatusFailed       AnomalyDetectorStatus = "FAILED"
	AnomalyDetectorStatusDeleted      AnomalyDetectorStatus = "DELETED"
	AnomalyDetectorStatusPaused       AnomalyDetectorStatus = "PAUSED"
)

// Values returns all known values for AnomalyDetectorStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyDetectorStatus) Values() []AnomalyDetectorStatus {
	return []AnomalyDetectorStatus{
		"INITIALIZING",
		"TRAINING",
		"ANALYZING",
		"FAILED",
		"DELETED",
		"PAUSED",
	}
}

type DataProtectionStatus string

// Enum values for DataProtectionStatus
const (
	DataProtectionStatusActivated DataProtectionStatus = "ACTIVATED"
	DataProtectionStatusDeleted   DataProtectionStatus = "DELETED"
	DataProtectionStatusArchived  DataProtectionStatus = "ARCHIVED"
	DataProtectionStatusDisabled  DataProtectionStatus = "DISABLED"
)

// Values returns all known values for DataProtectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataProtectionStatus) Values() []DataProtectionStatus {
	return []DataProtectionStatus{
		"ACTIVATED",
		"DELETED",
		"ARCHIVED",
		"DISABLED",
	}
}

type DeliveryDestinationType string

// Enum values for DeliveryDestinationType
const (
	DeliveryDestinationTypeS3  DeliveryDestinationType = "S3"
	DeliveryDestinationTypeCwl DeliveryDestinationType = "CWL"
	DeliveryDestinationTypeFh  DeliveryDestinationType = "FH"
)

// Values returns all known values for DeliveryDestinationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryDestinationType) Values() []DeliveryDestinationType {
	return []DeliveryDestinationType{
		"S3",
		"CWL",
		"FH",
	}
}

type Distribution string

// Enum values for Distribution
const (
	DistributionRandom      Distribution = "Random"
	DistributionByLogStream Distribution = "ByLogStream"
)

// Values returns all known values for Distribution. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Distribution) Values() []Distribution {
	return []Distribution{
		"Random",
		"ByLogStream",
	}
}

type EntityRejectionErrorType string

// Enum values for EntityRejectionErrorType
const (
	EntityRejectionErrorTypeInvalidEntity           EntityRejectionErrorType = "InvalidEntity"
	EntityRejectionErrorTypeInvalidTypeValue        EntityRejectionErrorType = "InvalidTypeValue"
	EntityRejectionErrorTypeInvalidKeyAttribute     EntityRejectionErrorType = "InvalidKeyAttributes"
	EntityRejectionErrorTypeInvalidAttributes       EntityRejectionErrorType = "InvalidAttributes"
	EntityRejectionErrorTypeEntitySizeTooLarge      EntityRejectionErrorType = "EntitySizeTooLarge"
	EntityRejectionErrorTypeUnsupportedLogGroupType EntityRejectionErrorType = "UnsupportedLogGroupType"
	EntityRejectionErrorTypeMissingRequiredFields   EntityRejectionErrorType = "MissingRequiredFields"
)

// Values returns all known values for EntityRejectionErrorType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EntityRejectionErrorType) Values() []EntityRejectionErrorType {
	return []EntityRejectionErrorType{
		"InvalidEntity",
		"InvalidTypeValue",
		"InvalidKeyAttributes",
		"InvalidAttributes",
		"EntitySizeTooLarge",
		"UnsupportedLogGroupType",
		"MissingRequiredFields",
	}
}

type EvaluationFrequency string

// Enum values for EvaluationFrequency
const (
	EvaluationFrequencyOneMin     EvaluationFrequency = "ONE_MIN"
	EvaluationFrequencyFiveMin    EvaluationFrequency = "FIVE_MIN"
	EvaluationFrequencyTenMin     EvaluationFrequency = "TEN_MIN"
	EvaluationFrequencyFifteenMin EvaluationFrequency = "FIFTEEN_MIN"
	EvaluationFrequencyThirtyMin  EvaluationFrequency = "THIRTY_MIN"
	EvaluationFrequencyOneHour    EvaluationFrequency = "ONE_HOUR"
)

// Values returns all known values for EvaluationFrequency. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EvaluationFrequency) Values() []EvaluationFrequency {
	return []EvaluationFrequency{
		"ONE_MIN",
		"FIVE_MIN",
		"TEN_MIN",
		"FIFTEEN_MIN",
		"THIRTY_MIN",
		"ONE_HOUR",
	}
}

type ExportTaskStatusCode string

// Enum values for ExportTaskStatusCode
const (
	ExportTaskStatusCodeCancelled     ExportTaskStatusCode = "CANCELLED"
	ExportTaskStatusCodeCompleted     ExportTaskStatusCode = "COMPLETED"
	ExportTaskStatusCodeFailed        ExportTaskStatusCode = "FAILED"
	ExportTaskStatusCodePending       ExportTaskStatusCode = "PENDING"
	ExportTaskStatusCodePendingCancel ExportTaskStatusCode = "PENDING_CANCEL"
	ExportTaskStatusCodeRunning       ExportTaskStatusCode = "RUNNING"
)

// Values returns all known values for ExportTaskStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExportTaskStatusCode) Values() []ExportTaskStatusCode {
	return []ExportTaskStatusCode{
		"CANCELLED",
		"COMPLETED",
		"FAILED",
		"PENDING",
		"PENDING_CANCEL",
		"RUNNING",
	}
}

type FlattenedElement string

// Enum values for FlattenedElement
const (
	FlattenedElementFirst FlattenedElement = "first"
	FlattenedElementLast  FlattenedElement = "last"
)

// Values returns all known values for FlattenedElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlattenedElement) Values() []FlattenedElement {
	return []FlattenedElement{
		"first",
		"last",
	}
}

type IndexSource string

// Enum values for IndexSource
const (
	IndexSourceAccount  IndexSource = "ACCOUNT"
	IndexSourceLogGroup IndexSource = "LOG_GROUP"
)

// Values returns all known values for IndexSource. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IndexSource) Values() []IndexSource {
	return []IndexSource{
		"ACCOUNT",
		"LOG_GROUP",
	}
}

type InheritedProperty string

// Enum values for InheritedProperty
const (
	InheritedPropertyAccountDataProtection InheritedProperty = "ACCOUNT_DATA_PROTECTION"
)

// Values returns all known values for InheritedProperty. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InheritedProperty) Values() []InheritedProperty {
	return []InheritedProperty{
		"ACCOUNT_DATA_PROTECTION",
	}
}

type IntegrationStatus string

// Enum values for IntegrationStatus
const (
	IntegrationStatusProvisioning IntegrationStatus = "PROVISIONING"
	IntegrationStatusActive       IntegrationStatus = "ACTIVE"
	IntegrationStatusFailed       IntegrationStatus = "FAILED"
)

// Values returns all known values for IntegrationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IntegrationStatus) Values() []IntegrationStatus {
	return []IntegrationStatus{
		"PROVISIONING",
		"ACTIVE",
		"FAILED",
	}
}

type IntegrationType string

// Enum values for IntegrationType
const (
	IntegrationTypeOpensearch IntegrationType = "OPENSEARCH"
)

// Values returns all known values for IntegrationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IntegrationType) Values() []IntegrationType {
	return []IntegrationType{
		"OPENSEARCH",
	}
}

type LogGroupClass string

// Enum values for LogGroupClass
const (
	LogGroupClassStandard         LogGroupClass = "STANDARD"
	LogGroupClassInfrequentAccess LogGroupClass = "INFREQUENT_ACCESS"
)

// Values returns all known values for LogGroupClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogGroupClass) Values() []LogGroupClass {
	return []LogGroupClass{
		"STANDARD",
		"INFREQUENT_ACCESS",
	}
}

type OpenSearchResourceStatusType string

// Enum values for OpenSearchResourceStatusType
const (
	OpenSearchResourceStatusTypeActive   OpenSearchResourceStatusType = "ACTIVE"
	OpenSearchResourceStatusTypeNotFound OpenSearchResourceStatusType = "NOT_FOUND"
	OpenSearchResourceStatusTypeError    OpenSearchResourceStatusType = "ERROR"
)

// Values returns all known values for OpenSearchResourceStatusType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OpenSearchResourceStatusType) Values() []OpenSearchResourceStatusType {
	return []OpenSearchResourceStatusType{
		"ACTIVE",
		"NOT_FOUND",
		"ERROR",
	}
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByLogStreamName OrderBy = "LogStreamName"
	OrderByLastEventTime OrderBy = "LastEventTime"
)

// Values returns all known values for OrderBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrderBy) Values() []OrderBy {
	return []OrderBy{
		"LogStreamName",
		"LastEventTime",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatJson    OutputFormat = "json"
	OutputFormatPlain   OutputFormat = "plain"
	OutputFormatW3c     OutputFormat = "w3c"
	OutputFormatRaw     OutputFormat = "raw"
	OutputFormatParquet OutputFormat = "parquet"
)

// Values returns all known values for OutputFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"json",
		"plain",
		"w3c",
		"raw",
		"parquet",
	}
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeDataProtectionPolicy     PolicyType = "DATA_PROTECTION_POLICY"
	PolicyTypeSubscriptionFilterPolicy PolicyType = "SUBSCRIPTION_FILTER_POLICY"
	PolicyTypeFieldIndexPolicy         PolicyType = "FIELD_INDEX_POLICY"
	PolicyTypeTransformerPolicy        PolicyType = "TRANSFORMER_POLICY"
)

// Values returns all known values for PolicyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PolicyType) Values() []PolicyType {
	return []PolicyType{
		"DATA_PROTECTION_POLICY",
		"SUBSCRIPTION_FILTER_POLICY",
		"FIELD_INDEX_POLICY",
		"TRANSFORMER_POLICY",
	}
}

type QueryLanguage string

// Enum values for QueryLanguage
const (
	QueryLanguageCwli QueryLanguage = "CWLI"
	QueryLanguageSql  QueryLanguage = "SQL"
	QueryLanguagePpl  QueryLanguage = "PPL"
)

// Values returns all known values for QueryLanguage. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryLanguage) Values() []QueryLanguage {
	return []QueryLanguage{
		"CWLI",
		"SQL",
		"PPL",
	}
}

type QueryStatus string

// Enum values for QueryStatus
const (
	QueryStatusScheduled QueryStatus = "Scheduled"
	QueryStatusRunning   QueryStatus = "Running"
	QueryStatusComplete  QueryStatus = "Complete"
	QueryStatusFailed    QueryStatus = "Failed"
	QueryStatusCancelled QueryStatus = "Cancelled"
	QueryStatusTimeout   QueryStatus = "Timeout"
	QueryStatusUnknown   QueryStatus = "Unknown"
)

// Values returns all known values for QueryStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryStatus) Values() []QueryStatus {
	return []QueryStatus{
		"Scheduled",
		"Running",
		"Complete",
		"Failed",
		"Cancelled",
		"Timeout",
		"Unknown",
	}
}

type Scope string

// Enum values for Scope
const (
	ScopeAll Scope = "ALL"
)

// Values returns all known values for Scope. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Scope) Values() []Scope {
	return []Scope{
		"ALL",
	}
}

type StandardUnit string

// Enum values for StandardUnit
const (
	StandardUnitSeconds         StandardUnit = "Seconds"
	StandardUnitMicroseconds    StandardUnit = "Microseconds"
	StandardUnitMilliseconds    StandardUnit = "Milliseconds"
	StandardUnitBytes           StandardUnit = "Bytes"
	StandardUnitKilobytes       StandardUnit = "Kilobytes"
	StandardUnitMegabytes       StandardUnit = "Megabytes"
	StandardUnitGigabytes       StandardUnit = "Gigabytes"
	StandardUnitTerabytes       StandardUnit = "Terabytes"
	StandardUnitBits            StandardUnit = "Bits"
	StandardUnitKilobits        StandardUnit = "Kilobits"
	StandardUnitMegabits        StandardUnit = "Megabits"
	StandardUnitGigabits        StandardUnit = "Gigabits"
	StandardUnitTerabits        StandardUnit = "Terabits"
	StandardUnitPercent         StandardUnit = "Percent"
	StandardUnitCount           StandardUnit = "Count"
	StandardUnitBytesSecond     StandardUnit = "Bytes/Second"
	StandardUnitKilobytesSecond StandardUnit = "Kilobytes/Second"
	StandardUnitMegabytesSecond StandardUnit = "Megabytes/Second"
	StandardUnitGigabytesSecond StandardUnit = "Gigabytes/Second"
	StandardUnitTerabytesSecond StandardUnit = "Terabytes/Second"
	StandardUnitBitsSecond      StandardUnit = "Bits/Second"
	StandardUnitKilobitsSecond  StandardUnit = "Kilobits/Second"
	StandardUnitMegabitsSecond  StandardUnit = "Megabits/Second"
	StandardUnitGigabitsSecond  StandardUnit = "Gigabits/Second"
	StandardUnitTerabitsSecond  StandardUnit = "Terabits/Second"
	StandardUnitCountSecond     StandardUnit = "Count/Second"
	StandardUnitNone            StandardUnit = "None"
)

// Values returns all known values for StandardUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StandardUnit) Values() []StandardUnit {
	return []StandardUnit{
		"Seconds",
		"Microseconds",
		"Milliseconds",
		"Bytes",
		"Kilobytes",
		"Megabytes",
		"Gigabytes",
		"Terabytes",
		"Bits",
		"Kilobits",
		"Megabits",
		"Gigabits",
		"Terabits",
		"Percent",
		"Count",
		"Bytes/Second",
		"Kilobytes/Second",
		"Megabytes/Second",
		"Gigabytes/Second",
		"Terabytes/Second",
		"Bits/Second",
		"Kilobits/Second",
		"Megabits/Second",
		"Gigabits/Second",
		"Terabits/Second",
		"Count/Second",
		"None",
	}
}

type State string

// Enum values for State
const (
	StateActive     State = "Active"
	StateSuppressed State = "Suppressed"
	StateBaseline   State = "Baseline"
)

// Values returns all known values for State. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"Active",
		"Suppressed",
		"Baseline",
	}
}

type SuppressionState string

// Enum values for SuppressionState
const (
	SuppressionStateSuppressed   SuppressionState = "SUPPRESSED"
	SuppressionStateUnsuppressed SuppressionState = "UNSUPPRESSED"
)

// Values returns all known values for SuppressionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionState) Values() []SuppressionState {
	return []SuppressionState{
		"SUPPRESSED",
		"UNSUPPRESSED",
	}
}

type SuppressionType string

// Enum values for SuppressionType
const (
	SuppressionTypeLimited  SuppressionType = "LIMITED"
	SuppressionTypeInfinite SuppressionType = "INFINITE"
)

// Values returns all known values for SuppressionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionType) Values() []SuppressionType {
	return []SuppressionType{
		"LIMITED",
		"INFINITE",
	}
}

type SuppressionUnit string

// Enum values for SuppressionUnit
const (
	SuppressionUnitSeconds SuppressionUnit = "SECONDS"
	SuppressionUnitMinutes SuppressionUnit = "MINUTES"
	SuppressionUnitHours   SuppressionUnit = "HOURS"
)

// Values returns all known values for SuppressionUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionUnit) Values() []SuppressionUnit {
	return []SuppressionUnit{
		"SECONDS",
		"MINUTES",
		"HOURS",
	}
}

type Type string

// Enum values for Type
const (
	TypeBoolean Type = "boolean"
	TypeInteger Type = "integer"
	TypeDouble  Type = "double"
	TypeString  Type = "string"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"boolean",
		"integer",
		"double",
		"string",
	}
}