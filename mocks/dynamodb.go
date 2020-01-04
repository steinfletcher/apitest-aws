package mocks

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var _ dynamodbiface.DynamoDBAPI = &DynamoDB{}

type DynamoDB struct {
	DescribeTableReplicaAutoScalingFunc               func(*dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeTableReplicaAutoScalingInvoked            bool
	DescribeTableReplicaAutoScalingWithContextFunc    func(aws.Context, *dynamodb.DescribeTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeTableReplicaAutoScalingWithContextInvoked bool
	DescribeTableReplicaAutoScalingRequestFunc        func(*dynamodb.DescribeTableReplicaAutoScalingInput) *dynamodb.DescribeTableReplicaAutoScalingOutput
	DescribeTableReplicaAutoScalingRequestInvoked     bool
	UpdateTableReplicaAutoScalingFunc                 func(*dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	UpdateTableReplicaAutoScalingInvoked              bool
	UpdateTableReplicaAutoScalingWithContextFunc      func(aws.Context, *dynamodb.UpdateTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	UpdateTableReplicaAutoScalingWithContextInvoked   bool
	UpdateTableReplicaAutoScalingRequestFunc          func(*dynamodb.UpdateTableReplicaAutoScalingInput) *dynamodb.UpdateTableReplicaAutoScalingOutput
	UpdateTableReplicaAutoScalingRequestInvoked       bool
	BatchGetItemFunc                                  func(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
	BatchGetItemInvoked                               bool
	BatchGetItemWithContextFunc                       func(aws.Context, *dynamodb.BatchGetItemInput, ...request.Option) (*dynamodb.BatchGetItemOutput, error)
	BatchGetItemWithContextInvoked                    bool
	BatchGetItemRequestFunc                           func(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput)
	BatchGetItemRequestInvoked                        bool
	BatchGetItemPagesFunc                             func(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error
	BatchGetItemPagesInvoked                          bool
	BatchGetItemPagesWithContextFunc                  func(aws.Context, *dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool, ...request.Option) error
	BatchGetItemPagesWithContextInvoked               bool
	BatchWriteItemFunc                                func(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
	BatchWriteItemInvoked                             bool
	BatchWriteItemWithContextFunc                     func(aws.Context, *dynamodb.BatchWriteItemInput, ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
	BatchWriteItemWithContextInvoked                  bool
	BatchWriteItemRequestFunc                         func(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput)
	BatchWriteItemRequestInvoked                      bool
	CreateBackupFunc                                  func(*dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error)
	CreateBackupInvoked                               bool
	CreateBackupWithContextFunc                       func(aws.Context, *dynamodb.CreateBackupInput, ...request.Option) (*dynamodb.CreateBackupOutput, error)
	CreateBackupWithContextInvoked                    bool
	CreateBackupRequestFunc                           func(*dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput)
	CreateBackupRequestInvoked                        bool
	CreateGlobalTableFunc                             func(*dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error)
	CreateGlobalTableInvoked                          bool
	CreateGlobalTableWithContextFunc                  func(aws.Context, *dynamodb.CreateGlobalTableInput, ...request.Option) (*dynamodb.CreateGlobalTableOutput, error)
	CreateGlobalTableWithContextInvoked               bool
	CreateGlobalTableRequestFunc                      func(*dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput)
	CreateGlobalTableRequestInvoked                   bool
	CreateTableFunc                                   func(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
	CreateTableInvoked                                bool
	CreateTableWithContextFunc                        func(aws.Context, *dynamodb.CreateTableInput, ...request.Option) (*dynamodb.CreateTableOutput, error)
	CreateTableWithContextInvoked                     bool
	CreateTableRequestFunc                            func(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput)
	CreateTableRequestInvoked                         bool
	DeleteBackupFunc                                  func(*dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error)
	DeleteBackupInvoked                               bool
	DeleteBackupWithContextFunc                       func(aws.Context, *dynamodb.DeleteBackupInput, ...request.Option) (*dynamodb.DeleteBackupOutput, error)
	DeleteBackupWithContextInvoked                    bool
	DeleteBackupRequestFunc                           func(*dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput)
	DeleteBackupRequestInvoked                        bool
	DeleteItemFunc                                    func(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
	DeleteItemInvoked                                 bool
	DeleteItemWithContextFunc                         func(aws.Context, *dynamodb.DeleteItemInput, ...request.Option) (*dynamodb.DeleteItemOutput, error)
	DeleteItemWithContextInvoked                      bool
	DeleteItemRequestFunc                             func(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput)
	DeleteItemRequestInvoked                          bool
	DeleteTableFunc                                   func(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)
	DeleteTableInvoked                                bool
	DeleteTableWithContextFunc                        func(aws.Context, *dynamodb.DeleteTableInput, ...request.Option) (*dynamodb.DeleteTableOutput, error)
	DeleteTableWithContextInvoked                     bool
	DeleteTableRequestFunc                            func(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput)
	DeleteTableRequestInvoked                         bool
	DescribeBackupFunc                                func(*dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error)
	DescribeBackupInvoked                             bool
	DescribeBackupWithContextFunc                     func(aws.Context, *dynamodb.DescribeBackupInput, ...request.Option) (*dynamodb.DescribeBackupOutput, error)
	DescribeBackupWithContextInvoked                  bool
	DescribeBackupRequestFunc                         func(*dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput)
	DescribeBackupRequestInvoked                      bool
	DescribeContinuousBackupsFunc                     func(*dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeContinuousBackupsInvoked                  bool
	DescribeContinuousBackupsWithContextFunc          func(aws.Context, *dynamodb.DescribeContinuousBackupsInput, ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeContinuousBackupsWithContextInvoked       bool
	DescribeContinuousBackupsRequestFunc              func(*dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput)
	DescribeContinuousBackupsRequestInvoked           bool
	DescribeEndpointsFunc                             func(*dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeEndpointsInvoked                          bool
	DescribeEndpointsWithContextFunc                  func(aws.Context, *dynamodb.DescribeEndpointsInput, ...request.Option) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeEndpointsWithContextInvoked               bool
	DescribeEndpointsRequestFunc                      func(*dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput)
	DescribeEndpointsRequestInvoked                   bool
	DescribeGlobalTableFunc                           func(*dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableInvoked                        bool
	DescribeGlobalTableWithContextFunc                func(aws.Context, *dynamodb.DescribeGlobalTableInput, ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableWithContextInvoked             bool
	DescribeGlobalTableRequestFunc                    func(*dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput)
	DescribeGlobalTableRequestInvoked                 bool
	DescribeGlobalTableSettingsFunc                   func(*dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeGlobalTableSettingsInvoked                bool
	DescribeGlobalTableSettingsWithContextFunc        func(aws.Context, *dynamodb.DescribeGlobalTableSettingsInput, ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeGlobalTableSettingsWithContextInvoked     bool
	DescribeGlobalTableSettingsRequestFunc            func(*dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput)
	DescribeGlobalTableSettingsRequestInvoked         bool
	DescribeLimitsFunc                                func(*dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error)
	DescribeLimitsInvoked                             bool
	DescribeLimitsWithContextFunc                     func(aws.Context, *dynamodb.DescribeLimitsInput, ...request.Option) (*dynamodb.DescribeLimitsOutput, error)
	DescribeLimitsWithContextInvoked                  bool
	DescribeLimitsRequestFunc                         func(*dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput)
	DescribeLimitsRequestInvoked                      bool
	DescribeTableFunc                                 func(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
	DescribeTableInvoked                              bool
	DescribeTableWithContextFunc                      func(aws.Context, *dynamodb.DescribeTableInput, ...request.Option) (*dynamodb.DescribeTableOutput, error)
	DescribeTableWithContextInvoked                   bool
	DescribeTableRequestFunc                          func(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput)
	DescribeTableRequestInvoked                       bool
	DescribeTimeToLiveFunc                            func(*dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error)
	DescribeTimeToLiveInvoked                         bool
	DescribeTimeToLiveWithContextFunc                 func(aws.Context, *dynamodb.DescribeTimeToLiveInput, ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error)
	DescribeTimeToLiveWithContextInvoked              bool
	DescribeTimeToLiveRequestFunc                     func(*dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput)
	DescribeTimeToLiveRequestInvoked                  bool
	GetItemFunc                                       func(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	GetItemInvoked                                    bool
	GetItemWithContextFunc                            func(aws.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error)
	GetItemWithContextInvoked                         bool
	GetItemRequestFunc                                func(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput)
	GetItemRequestInvoked                             bool
	ListBackupsFunc                                   func(*dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error)
	ListBackupsInvoked                                bool
	ListBackupsWithContextFunc                        func(aws.Context, *dynamodb.ListBackupsInput, ...request.Option) (*dynamodb.ListBackupsOutput, error)
	ListBackupsWithContextInvoked                     bool
	ListBackupsRequestFunc                            func(*dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput)
	ListBackupsRequestInvoked                         bool
	ListGlobalTablesFunc                              func(*dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error)
	ListGlobalTablesInvoked                           bool
	ListGlobalTablesWithContextFunc                   func(aws.Context, *dynamodb.ListGlobalTablesInput, ...request.Option) (*dynamodb.ListGlobalTablesOutput, error)
	ListGlobalTablesWithContextInvoked                bool
	ListGlobalTablesRequestFunc                       func(*dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput)
	ListGlobalTablesRequestInvoked                    bool
	ListTablesFunc                                    func(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)
	ListTablesInvoked                                 bool
	ListTablesWithContextFunc                         func(aws.Context, *dynamodb.ListTablesInput, ...request.Option) (*dynamodb.ListTablesOutput, error)
	ListTablesWithContextInvoked                      bool
	ListTablesRequestFunc                             func(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput)
	ListTablesRequestInvoked                          bool
	ListTablesPagesFunc                               func(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error
	ListTablesPagesInvoked                            bool
	ListTablesPagesWithContextFunc                    func(aws.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...request.Option) error
	ListTablesPagesWithContextInvoked                 bool
	ListTagsOfResourceFunc                            func(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error)
	ListTagsOfResourceInvoked                         bool
	ListTagsOfResourceWithContextFunc                 func(aws.Context, *dynamodb.ListTagsOfResourceInput, ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error)
	ListTagsOfResourceWithContextInvoked              bool
	ListTagsOfResourceRequestFunc                     func(*dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput)
	ListTagsOfResourceRequestInvoked                  bool
	PutItemFunc                                       func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	PutItemInvoked                                    bool
	PutItemWithContextFunc                            func(aws.Context, *dynamodb.PutItemInput, ...request.Option) (*dynamodb.PutItemOutput, error)
	PutItemWithContextInvoked                         bool
	PutItemRequestFunc                                func(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput)
	PutItemRequestInvoked                             bool
	QueryFunc                                         func(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	QueryInvoked                                      bool
	QueryWithContextFunc                              func(aws.Context, *dynamodb.QueryInput, ...request.Option) (*dynamodb.QueryOutput, error)
	QueryWithContextInvoked                           bool
	QueryRequestFunc                                  func(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput)
	QueryRequestInvoked                               bool
	QueryPagesFunc                                    func(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error
	QueryPagesInvoked                                 bool
	QueryPagesWithContextFunc                         func(aws.Context, *dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool, ...request.Option) error
	QueryPagesWithContextInvoked                      bool
	RestoreTableFromBackupFunc                        func(*dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error)
	RestoreTableFromBackupInvoked                     bool
	RestoreTableFromBackupWithContextFunc             func(aws.Context, *dynamodb.RestoreTableFromBackupInput, ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error)
	RestoreTableFromBackupWithContextInvoked          bool
	RestoreTableFromBackupRequestFunc                 func(*dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput)
	RestoreTableFromBackupRequestInvoked              bool
	RestoreTableToPointInTimeFunc                     func(*dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	RestoreTableToPointInTimeInvoked                  bool
	RestoreTableToPointInTimeWithContextFunc          func(aws.Context, *dynamodb.RestoreTableToPointInTimeInput, ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	RestoreTableToPointInTimeWithContextInvoked       bool
	RestoreTableToPointInTimeRequestFunc              func(*dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput)
	RestoreTableToPointInTimeRequestInvoked           bool
	ScanFunc                                          func(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	ScanInvoked                                       bool
	ScanWithContextFunc                               func(aws.Context, *dynamodb.ScanInput, ...request.Option) (*dynamodb.ScanOutput, error)
	ScanWithContextInvoked                            bool
	ScanRequestFunc                                   func(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput)
	ScanRequestInvoked                                bool
	ScanPagesFunc                                     func(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error
	ScanPagesInvoked                                  bool
	ScanPagesWithContextFunc                          func(aws.Context, *dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool, ...request.Option) error
	ScanPagesWithContextInvoked                       bool
	TagResourceFunc                                   func(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error)
	TagResourceInvoked                                bool
	TagResourceWithContextFunc                        func(aws.Context, *dynamodb.TagResourceInput, ...request.Option) (*dynamodb.TagResourceOutput, error)
	TagResourceWithContextInvoked                     bool
	TagResourceRequestFunc                            func(*dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput)
	TagResourceRequestInvoked                         bool
	TransactGetItemsFunc                              func(*dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error)
	TransactGetItemsInvoked                           bool
	TransactGetItemsWithContextFunc                   func(aws.Context, *dynamodb.TransactGetItemsInput, ...request.Option) (*dynamodb.TransactGetItemsOutput, error)
	TransactGetItemsWithContextInvoked                bool
	TransactGetItemsRequestFunc                       func(*dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput)
	TransactGetItemsRequestInvoked                    bool
	TransactWriteItemsFunc                            func(*dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
	TransactWriteItemsInvoked                         bool
	TransactWriteItemsWithContextFunc                 func(aws.Context, *dynamodb.TransactWriteItemsInput, ...request.Option) (*dynamodb.TransactWriteItemsOutput, error)
	TransactWriteItemsWithContextInvoked              bool
	TransactWriteItemsRequestFunc                     func(*dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput)
	TransactWriteItemsRequestInvoked                  bool
	UntagResourceFunc                                 func(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error)
	UntagResourceInvoked                              bool
	UntagResourceWithContextFunc                      func(aws.Context, *dynamodb.UntagResourceInput, ...request.Option) (*dynamodb.UntagResourceOutput, error)
	UntagResourceWithContextInvoked                   bool
	UntagResourceRequestFunc                          func(*dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput)
	UntagResourceRequestInvoked                       bool
	UpdateContinuousBackupsFunc                       func(*dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error)
	UpdateContinuousBackupsInvoked                    bool
	UpdateContinuousBackupsWithContextFunc            func(aws.Context, *dynamodb.UpdateContinuousBackupsInput, ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error)
	UpdateContinuousBackupsWithContextInvoked         bool
	UpdateContinuousBackupsRequestFunc                func(*dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput)
	UpdateContinuousBackupsRequestInvoked             bool
	UpdateGlobalTableFunc                             func(*dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error)
	UpdateGlobalTableInvoked                          bool
	UpdateGlobalTableWithContextFunc                  func(aws.Context, *dynamodb.UpdateGlobalTableInput, ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error)
	UpdateGlobalTableWithContextInvoked               bool
	UpdateGlobalTableRequestFunc                      func(*dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput)
	UpdateGlobalTableRequestInvoked                   bool
	UpdateGlobalTableSettingsFunc                     func(*dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	UpdateGlobalTableSettingsInvoked                  bool
	UpdateGlobalTableSettingsWithContextFunc          func(aws.Context, *dynamodb.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	UpdateGlobalTableSettingsWithContextInvoked       bool
	UpdateGlobalTableSettingsRequestFunc              func(*dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput)
	UpdateGlobalTableSettingsRequestInvoked           bool
	UpdateItemFunc                                    func(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	UpdateItemInvoked                                 bool
	UpdateItemWithContextFunc                         func(aws.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error)
	UpdateItemWithContextInvoked                      bool
	UpdateItemRequestFunc                             func(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput)
	UpdateItemRequestInvoked                          bool
	UpdateTableFunc                                   func(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
	UpdateTableInvoked                                bool
	UpdateTableWithContextFunc                        func(aws.Context, *dynamodb.UpdateTableInput, ...request.Option) (*dynamodb.UpdateTableOutput, error)
	UpdateTableWithContextInvoked                     bool
	UpdateTableRequestFunc                            func(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput)
	UpdateTableRequestInvoked                         bool
	UpdateTimeToLiveFunc                              func(*dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error)
	UpdateTimeToLiveInvoked                           bool
	UpdateTimeToLiveWithContextFunc                   func(aws.Context, *dynamodb.UpdateTimeToLiveInput, ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error)
	UpdateTimeToLiveWithContextInvoked                bool
	UpdateTimeToLiveRequestFunc                       func(*dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput)
	UpdateTimeToLiveRequestInvoked                    bool
	WaitUntilTableExistsFunc                          func(*dynamodb.DescribeTableInput) error
	WaitUntilTableExistsInvoked                       bool
	WaitUntilTableExistsWithContextFunc               func(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error
	WaitUntilTableExistsWithContextInvoked            bool
	WaitUntilTableNotExistsFunc                       func(*dynamodb.DescribeTableInput) error
	WaitUntilTableNotExistsInvoked                    bool
	WaitUntilTableNotExistsWithContextFunc            func(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error
	WaitUntilTableNotExistsWithContextInvoked         bool
}

func (d *DynamoDB) DescribeTableReplicaAutoScaling(input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	d.DescribeTableReplicaAutoScalingInvoked = true
	return d.DescribeTableReplicaAutoScaling(input)
}

func (d *DynamoDB) DescribeTableReplicaAutoScalingWithContext(ctx aws.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput, opts ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	d.DescribeTableReplicaAutoScalingWithContextInvoked = true
	return d.DescribeTableReplicaAutoScalingWithContext(ctx, input, opts...)
}

func (d *DynamoDB) DescribeTableReplicaAutoScalingRequest(input *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	d.DescribeTableReplicaAutoScalingRequestInvoked = true
	return d.DescribeTableReplicaAutoScalingRequest(input)
}

func (d *DynamoDB) UpdateTableReplicaAutoScaling(input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	d.UpdateTableReplicaAutoScalingInvoked = true
	return d.UpdateTableReplicaAutoScaling(input)
}

func (d *DynamoDB) UpdateTableReplicaAutoScalingWithContext(ctx aws.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput, opts ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	d.UpdateTableReplicaAutoScalingWithContextInvoked = true
	return d.UpdateTableReplicaAutoScalingWithContext(ctx, input, opts...)
}

func (d *DynamoDB) UpdateTableReplicaAutoScalingRequest(input *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	d.UpdateTableReplicaAutoScalingRequestInvoked = true
	return d.UpdateTableReplicaAutoScalingRequest(input)
}

func (d *DynamoDB) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	d.BatchGetItemInvoked = true
	return d.BatchGetItemFunc(input)
}

func (d *DynamoDB) BatchGetItemWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	d.BatchGetItemWithContextInvoked = true
	return d.BatchGetItemWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	d.BatchGetItemRequestInvoked = true
	return d.BatchGetItemRequestFunc(input)
}

func (d *DynamoDB) BatchGetItemPages(input *dynamodb.BatchGetItemInput, f func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	d.BatchGetItemPagesInvoked = true
	return d.BatchGetItemPagesFunc(input, f)
}

func (d *DynamoDB) BatchGetItemPagesWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, f func(*dynamodb.BatchGetItemOutput, bool) bool, opts ...request.Option) error {
	d.BatchGetItemPagesWithContextInvoked = true
	return d.BatchGetItemPagesWithContextFunc(ctx, input, f, opts...)
}

func (d *DynamoDB) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	d.BatchWriteItemInvoked = true
	return d.BatchWriteItemFunc(input)
}

func (d *DynamoDB) BatchWriteItemWithContext(ctx aws.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	d.BatchWriteItemWithContextInvoked = true
	return d.BatchWriteItemWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	d.BatchWriteItemRequestInvoked = true
	return d.BatchWriteItemRequestFunc(input)
}

func (d *DynamoDB) CreateBackup(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	d.CreateBackupInvoked = true
	return d.CreateBackupFunc(input)
}

func (d *DynamoDB) CreateBackupWithContext(ctx aws.Context, input *dynamodb.CreateBackupInput, opts ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	d.CreateBackupWithContextInvoked = true
	return d.CreateBackupWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) CreateBackupRequest(input *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	d.CreateBackupRequestInvoked = true
	return d.CreateBackupRequestFunc(input)
}

func (d *DynamoDB) CreateGlobalTable(input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	d.CreateGlobalTableInvoked = true
	return d.CreateGlobalTableFunc(input)
}

func (d *DynamoDB) CreateGlobalTableWithContext(ctx aws.Context, input *dynamodb.CreateGlobalTableInput, opts ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	d.CreateGlobalTableWithContextInvoked = true
	return d.CreateGlobalTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) CreateGlobalTableRequest(input *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	d.CreateGlobalTableRequestInvoked = true
	return d.CreateGlobalTableRequestFunc(input)
}

func (d *DynamoDB) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	d.CreateTableInvoked = true
	return d.CreateTableFunc(input)
}

func (d *DynamoDB) CreateTableWithContext(ctx aws.Context, input *dynamodb.CreateTableInput, opts ...request.Option) (*dynamodb.CreateTableOutput, error) {
	d.CreateTableWithContextInvoked = true
	return d.CreateTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) CreateTableRequest(input *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	d.CreateTableRequestInvoked = true
	return d.CreateTableRequestFunc(input)
}

func (d *DynamoDB) DeleteBackup(input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	d.DeleteBackupInvoked = true
	return d.DeleteBackupFunc(input)
}

func (d *DynamoDB) DeleteBackupWithContext(ctx aws.Context, input *dynamodb.DeleteBackupInput, opts ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	d.DeleteBackupWithContextInvoked = true
	return d.DeleteBackupWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DeleteBackupRequest(input *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	d.DeleteBackupRequestInvoked = true
	return d.DeleteBackupRequestFunc(input)
}

func (d *DynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	d.DeleteItemInvoked = true
	return d.DeleteItemFunc(input)
}

func (d *DynamoDB) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	d.DeleteItemWithContextInvoked = true
	return d.DeleteItemWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	d.DeleteItemRequestInvoked = true
	return d.DeleteItemRequestFunc(input)
}

func (d *DynamoDB) DeleteTable(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	d.DeleteTableInvoked = true
	return d.DeleteTableFunc(input)
}

func (d *DynamoDB) DeleteTableWithContext(ctx aws.Context, input *dynamodb.DeleteTableInput, opts ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	d.DeleteTableWithContextInvoked = true
	return d.DeleteTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DeleteTableRequest(input *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	d.DeleteTableRequestInvoked = true
	return d.DeleteTableRequestFunc(input)
}

func (d *DynamoDB) DescribeBackup(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	d.DescribeBackupInvoked = true
	return d.DescribeBackupFunc(input)
}

func (d *DynamoDB) DescribeBackupWithContext(ctx aws.Context, input *dynamodb.DescribeBackupInput, opts ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	d.DescribeBackupWithContextInvoked = true
	return d.DescribeBackupWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeBackupRequest(input *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	d.DescribeBackupRequestInvoked = true
	return d.DescribeBackupRequestFunc(input)
}

func (d *DynamoDB) DescribeContinuousBackups(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	d.DescribeContinuousBackupsInvoked = true
	return d.DescribeContinuousBackupsFunc(input)
}

func (d *DynamoDB) DescribeContinuousBackupsWithContext(ctx aws.Context, input *dynamodb.DescribeContinuousBackupsInput, opts ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	d.DescribeContinuousBackupsWithContextInvoked = true
	return d.DescribeContinuousBackupsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeContinuousBackupsRequest(input *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	d.DescribeContinuousBackupsRequestInvoked = true
	return d.DescribeContinuousBackupsRequestFunc(input)
}

func (d *DynamoDB) DescribeEndpoints(input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	d.DescribeEndpointsInvoked = true
	return d.DescribeEndpointsFunc(input)
}

func (d *DynamoDB) DescribeEndpointsWithContext(ctx aws.Context, input *dynamodb.DescribeEndpointsInput, opts ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	d.DescribeEndpointsWithContextInvoked = true
	return d.DescribeEndpointsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeEndpointsRequest(input *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	d.DescribeEndpointsRequestInvoked = true
	return d.DescribeEndpointsRequestFunc(input)
}

func (d *DynamoDB) DescribeGlobalTable(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	d.DescribeGlobalTableInvoked = true
	return d.DescribeGlobalTableFunc(input)
}

func (d *DynamoDB) DescribeGlobalTableWithContext(ctx aws.Context, input *dynamodb.DescribeGlobalTableInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	d.DescribeGlobalTableWithContextInvoked = true
	return d.DescribeGlobalTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeGlobalTableRequest(input *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	d.DescribeGlobalTableRequestInvoked = true
	return d.DescribeGlobalTableRequestFunc(input)
}

func (d *DynamoDB) DescribeGlobalTableSettings(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	d.DescribeGlobalTableSettingsInvoked = true
	return d.DescribeGlobalTableSettingsFunc(input)
}

func (d *DynamoDB) DescribeGlobalTableSettingsWithContext(ctx aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	d.DescribeGlobalTableSettingsWithContextInvoked = true
	return d.DescribeGlobalTableSettingsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeGlobalTableSettingsRequest(input *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	d.DescribeGlobalTableSettingsRequestInvoked = true
	return d.DescribeGlobalTableSettingsRequestFunc(input)
}

func (d *DynamoDB) DescribeLimits(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	d.DescribeLimitsInvoked = true
	return d.DescribeLimitsFunc(input)
}

func (d *DynamoDB) DescribeLimitsWithContext(ctx aws.Context, input *dynamodb.DescribeLimitsInput, opts ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	d.DescribeLimitsWithContextInvoked = true
	return d.DescribeLimitsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeLimitsRequest(input *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	d.DescribeLimitsRequestInvoked = true
	return d.DescribeLimitsRequestFunc(input)
}

func (d *DynamoDB) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	d.DescribeTableInvoked = true
	return d.DescribeTableFunc(input)
}

func (d *DynamoDB) DescribeTableWithContext(ctx aws.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	d.DescribeTableWithContextInvoked = true
	return d.DescribeTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeTableRequest(input *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	d.DescribeTableRequestInvoked = true
	return d.DescribeTableRequestFunc(input)
}

func (d *DynamoDB) DescribeTimeToLive(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	d.DescribeTimeToLiveInvoked = true
	return d.DescribeTimeToLiveFunc(input)
}

func (d *DynamoDB) DescribeTimeToLiveWithContext(ctx aws.Context, input *dynamodb.DescribeTimeToLiveInput, opts ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	d.DescribeTimeToLiveWithContextInvoked = true
	return d.DescribeTimeToLiveWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) DescribeTimeToLiveRequest(input *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	d.DescribeTimeToLiveRequestInvoked = true
	return d.DescribeTimeToLiveRequestFunc(input)
}

func (d *DynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	d.GetItemInvoked = true
	return d.GetItemFunc(input)
}

func (d *DynamoDB) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	d.GetItemWithContextInvoked = true
	return d.GetItemWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	d.GetItemRequestInvoked = true
	return d.GetItemRequestFunc(input)
}

func (d *DynamoDB) ListBackups(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	d.ListBackupsInvoked = true
	return d.ListBackupsFunc(input)
}

func (d *DynamoDB) ListBackupsWithContext(ctx aws.Context, input *dynamodb.ListBackupsInput, opts ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	d.ListBackupsWithContextInvoked = true
	return d.ListBackupsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) ListBackupsRequest(input *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	d.ListBackupsRequestInvoked = true
	return d.ListBackupsRequestFunc(input)
}

func (d *DynamoDB) ListGlobalTables(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	d.ListGlobalTablesInvoked = true
	return d.ListGlobalTablesFunc(input)
}

func (d *DynamoDB) ListGlobalTablesWithContext(ctx aws.Context, input *dynamodb.ListGlobalTablesInput, opts ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	d.ListGlobalTablesWithContextInvoked = true
	return d.ListGlobalTablesWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) ListGlobalTablesRequest(input *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	d.ListGlobalTablesRequestInvoked = true
	return d.ListGlobalTablesRequestFunc(input)
}

func (d *DynamoDB) ListTables(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	d.ListTablesInvoked = true
	return d.ListTablesFunc(input)
}

func (d *DynamoDB) ListTablesWithContext(ctx aws.Context, input *dynamodb.ListTablesInput, opts ...request.Option) (*dynamodb.ListTablesOutput, error) {
	d.ListTablesWithContextInvoked = true
	return d.ListTablesWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) ListTablesRequest(input *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	d.ListTablesRequestInvoked = true
	return d.ListTablesRequestFunc(input)
}

func (d *DynamoDB) ListTablesPages(input *dynamodb.ListTablesInput, f func(*dynamodb.ListTablesOutput, bool) bool) error {
	d.ListTablesPagesInvoked = true
	return d.ListTablesPagesFunc(input, f)
}

func (d *DynamoDB) ListTablesPagesWithContext(ctx aws.Context, input *dynamodb.ListTablesInput, f func(*dynamodb.ListTablesOutput, bool) bool, opts ...request.Option) error {
	d.ListTablesPagesWithContextInvoked = true
	return d.ListTablesPagesWithContextFunc(ctx, input, f, opts...)
}

func (d *DynamoDB) ListTagsOfResource(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	d.ListTagsOfResourceInvoked = true
	return d.ListTagsOfResourceFunc(input)
}

func (d *DynamoDB) ListTagsOfResourceWithContext(ctx aws.Context, input *dynamodb.ListTagsOfResourceInput, opts ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	d.ListTagsOfResourceWithContextInvoked = true
	return d.ListTagsOfResourceWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) ListTagsOfResourceRequest(input *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	d.ListTagsOfResourceRequestInvoked = true
	return d.ListTagsOfResourceRequestFunc(input)
}

func (d *DynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	d.PutItemInvoked = true
	return d.PutItemFunc(input)
}

func (d *DynamoDB) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	d.PutItemWithContextInvoked = true
	return d.PutItemWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	d.PutItemRequestInvoked = true
	return d.PutItemRequestFunc(input)
}

func (d *DynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	d.QueryInvoked = true
	return d.QueryFunc(input)
}

func (d *DynamoDB) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	d.QueryWithContextInvoked = true
	return d.QueryWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	d.QueryRequestInvoked = true
	return d.QueryRequestFunc(input)
}

func (d *DynamoDB) QueryPages(input *dynamodb.QueryInput, f func(*dynamodb.QueryOutput, bool) bool) error {
	d.QueryPagesInvoked = true
	return d.QueryPagesFunc(input, f)
}

func (d *DynamoDB) QueryPagesWithContext(ctx aws.Context, input *dynamodb.QueryInput, f func(*dynamodb.QueryOutput, bool) bool, opts ...request.Option) error {
	d.QueryPagesWithContextInvoked = true
	return d.QueryPagesWithContextFunc(ctx, input, f, opts...)
}

func (d *DynamoDB) RestoreTableFromBackup(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	d.RestoreTableFromBackupInvoked = true
	return d.RestoreTableFromBackupFunc(input)
}

func (d *DynamoDB) RestoreTableFromBackupWithContext(ctx aws.Context, input *dynamodb.RestoreTableFromBackupInput, opts ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	d.RestoreTableFromBackupWithContextInvoked = true
	return d.RestoreTableFromBackupWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) RestoreTableFromBackupRequest(input *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	d.RestoreTableFromBackupRequestInvoked = true
	return d.RestoreTableFromBackupRequestFunc(input)
}

func (d *DynamoDB) RestoreTableToPointInTime(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	d.RestoreTableToPointInTimeInvoked = true
	return d.RestoreTableToPointInTimeFunc(input)
}

func (d *DynamoDB) RestoreTableToPointInTimeWithContext(ctx aws.Context, input *dynamodb.RestoreTableToPointInTimeInput, opts ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	d.RestoreTableToPointInTimeWithContextInvoked = true
	return d.RestoreTableToPointInTimeWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) RestoreTableToPointInTimeRequest(input *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	d.RestoreTableToPointInTimeRequestInvoked = true
	return d.RestoreTableToPointInTimeRequestFunc(input)
}

func (d *DynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	d.ScanInvoked = true
	return d.ScanFunc(input)
}

func (d *DynamoDB) ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	d.ScanWithContextInvoked = true
	return d.ScanWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	d.ScanRequestInvoked = true
	return d.ScanRequestFunc(input)
}

func (d *DynamoDB) ScanPages(input *dynamodb.ScanInput, f func(*dynamodb.ScanOutput, bool) bool) error {
	d.ScanPagesInvoked = true
	return d.ScanPagesFunc(input, f)
}

func (d *DynamoDB) ScanPagesWithContext(ctx aws.Context, input *dynamodb.ScanInput, f func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error {
	d.ScanPagesWithContextInvoked = true
	return d.ScanPagesWithContextFunc(ctx, input, f, opts...)
}

func (d *DynamoDB) TagResource(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	d.TagResourceInvoked = true
	return d.TagResourceFunc(input)
}

func (d *DynamoDB) TagResourceWithContext(ctx aws.Context, input *dynamodb.TagResourceInput, opts ...request.Option) (*dynamodb.TagResourceOutput, error) {
	d.TagResourceWithContextInvoked = true
	return d.TagResourceWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) TagResourceRequest(input *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	d.TagResourceRequestInvoked = true
	return d.TagResourceRequestFunc(input)
}

func (d *DynamoDB) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	d.TransactGetItemsInvoked = true
	return d.TransactGetItemsFunc(input)
}

func (d *DynamoDB) TransactGetItemsWithContext(ctx aws.Context, input *dynamodb.TransactGetItemsInput, opts ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	d.TransactGetItemsWithContextInvoked = true
	return d.TransactGetItemsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) TransactGetItemsRequest(input *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	d.TransactGetItemsRequestInvoked = true
	return d.TransactGetItemsRequestFunc(input)
}

func (d *DynamoDB) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	d.TransactWriteItemsInvoked = true
	return d.TransactWriteItemsFunc(input)
}

func (d *DynamoDB) TransactWriteItemsWithContext(ctx aws.Context, input *dynamodb.TransactWriteItemsInput, opts ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	d.TransactWriteItemsWithContextInvoked = true
	return d.TransactWriteItemsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) TransactWriteItemsRequest(input *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	d.TransactWriteItemsRequestInvoked = true
	return d.TransactWriteItemsRequestFunc(input)
}

func (d *DynamoDB) UntagResource(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	d.UntagResourceInvoked = true
	return d.UntagResourceFunc(input)
}

func (d *DynamoDB) UntagResourceWithContext(ctx aws.Context, input *dynamodb.UntagResourceInput, opts ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	d.UntagResourceWithContextInvoked = true
	return d.UntagResourceWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UntagResourceRequest(input *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	d.UntagResourceRequestInvoked = true
	return d.UntagResourceRequestFunc(input)
}

func (d *DynamoDB) UpdateContinuousBackups(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	d.UpdateContinuousBackupsInvoked = true
	return d.UpdateContinuousBackupsFunc(input)
}

func (d *DynamoDB) UpdateContinuousBackupsWithContext(ctx aws.Context, input *dynamodb.UpdateContinuousBackupsInput, opts ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	d.UpdateContinuousBackupsWithContextInvoked = true
	return d.UpdateContinuousBackupsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UpdateContinuousBackupsRequest(input *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	d.UpdateContinuousBackupsRequestInvoked = true
	return d.UpdateContinuousBackupsRequestFunc(input)
}

func (d *DynamoDB) UpdateGlobalTable(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	d.UpdateGlobalTableInvoked = true
	return d.UpdateGlobalTableFunc(input)
}

func (d *DynamoDB) UpdateGlobalTableWithContext(ctx aws.Context, input *dynamodb.UpdateGlobalTableInput, opts ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	d.UpdateGlobalTableWithContextInvoked = true
	return d.UpdateGlobalTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UpdateGlobalTableRequest(input *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	d.UpdateGlobalTableRequestInvoked = true
	return d.UpdateGlobalTableRequestFunc(input)
}

func (d *DynamoDB) UpdateGlobalTableSettings(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	d.UpdateGlobalTableSettingsInvoked = true
	return d.UpdateGlobalTableSettingsFunc(input)
}

func (d *DynamoDB) UpdateGlobalTableSettingsWithContext(ctx aws.Context, input *dynamodb.UpdateGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	d.UpdateGlobalTableSettingsWithContextInvoked = true
	return d.UpdateGlobalTableSettingsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UpdateGlobalTableSettingsRequest(input *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	d.UpdateGlobalTableSettingsRequestInvoked = true
	return d.UpdateGlobalTableSettingsRequestFunc(input)
}

func (d *DynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	d.UpdateItemInvoked = true
	return d.UpdateItemFunc(input)
}

func (d *DynamoDB) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	d.UpdateItemWithContextInvoked = true
	return d.UpdateItemWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	d.UpdateItemRequestInvoked = true
	return d.UpdateItemRequestFunc(input)
}

func (d *DynamoDB) UpdateTable(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	d.UpdateTableInvoked = true
	return d.UpdateTableFunc(input)
}

func (d *DynamoDB) UpdateTableWithContext(ctx aws.Context, input *dynamodb.UpdateTableInput, opts ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	d.UpdateTableWithContextInvoked = true
	return d.UpdateTableWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UpdateTableRequest(input *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	d.UpdateTableRequestInvoked = true
	return d.UpdateTableRequestFunc(input)
}

func (d *DynamoDB) UpdateTimeToLive(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	d.UpdateTimeToLiveInvoked = true
	return d.UpdateTimeToLiveFunc(input)
}

func (d *DynamoDB) UpdateTimeToLiveWithContext(ctx aws.Context, input *dynamodb.UpdateTimeToLiveInput, opts ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	d.UpdateTimeToLiveWithContextInvoked = true
	return d.UpdateTimeToLiveWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) UpdateTimeToLiveRequest(input *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	d.UpdateTimeToLiveRequestInvoked = true
	return d.UpdateTimeToLiveRequestFunc(input)
}

func (d *DynamoDB) WaitUntilTableExists(input *dynamodb.DescribeTableInput) error {
	d.WaitUntilTableExistsInvoked = true
	return d.WaitUntilTableExistsFunc(input)
}

func (d *DynamoDB) WaitUntilTableExistsWithContext(ctx aws.Context, input *dynamodb.DescribeTableInput, opts ...request.WaiterOption) error {
	d.WaitUntilTableExistsWithContextInvoked = true
	return d.WaitUntilTableExistsWithContextFunc(ctx, input, opts...)
}

func (d *DynamoDB) WaitUntilTableNotExists(input *dynamodb.DescribeTableInput) error {
	d.WaitUntilTableNotExistsInvoked = true
	return d.WaitUntilTableNotExistsFunc(input)
}

func (d *DynamoDB) WaitUntilTableNotExistsWithContext(ctx aws.Context, input *dynamodb.DescribeTableInput, opts ...request.WaiterOption) error {
	d.WaitUntilTableNotExistsWithContextInvoked = true
	return d.WaitUntilTableNotExistsWithContextFunc(ctx, input, opts...)
}
