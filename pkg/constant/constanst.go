package constant

import "errors"

const (
	TaxjarTransactionsTopicName      = "taxjar-sync-transactions"
	TaxjarTransactionsRetryTopicName = "taxjar-sync-transactions-retry"
	TaxjarRefundsTopicName           = "taxjar-sync-refunds"
	TaxjarRefundsRetryTopicName      = "taxjar-sync-refunds-retry"
	TaxjarNotificationStatusPayment  = "taxjar-payment"
	TaxjarNotificationStatusRefund   = "taxjar-refund"

	RegistryKubernetes = "kubernetes"

	CollectionSavedCard = "saved_card"

	SavedCardErrorNotFound = "saved card with specified identifier not found"
	SavedCardErrorDatabase = "saved cards database query failed"

	ErrorDatabaseQueryFailed          = "Query to database collection failed"
	ErrorQueryCursorExecutionFailed   = "Execute result from query cursor failed"
	ErrorDatabaseFieldCollection      = "collection"
	ErrorDatabaseFieldDocumentId      = "document_id"
	ErrorDatabaseFieldQuery           = "query"
	ErrorDatabaseFieldSet             = "set"
	ErrorDatabaseFieldSorts           = "sorts"
	ErrorDatabaseFieldLimit           = "limit"
	ErrorDatabaseFieldOffset          = "offset"
	ErrorDatabaseFieldOperation       = "operation"
	ErrorDatabaseFieldOperationCount  = "count"
	ErrorDatabaseFieldOperationInsert = "insert"
	ErrorDatabaseFieldOperationUpdate = "update"
	ErrorDatabaseFieldOperationUpsert = "upsert"
	ErrorDatabaseFieldDocument        = "document"
)

var (
	ErrNotFound = errors.New(SavedCardErrorNotFound)
	ErrDatabase = errors.New(SavedCardErrorDatabase)
)
