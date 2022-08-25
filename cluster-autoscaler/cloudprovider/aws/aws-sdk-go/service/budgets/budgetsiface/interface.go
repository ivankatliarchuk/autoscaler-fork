// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package budgetsiface provides an interface to enable mocking the AWS Budgets service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package budgetsiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/budgets"
)

// BudgetsAPI provides an interface to enable mocking the
// budgets.Budgets service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Budgets.
//    func myFunc(svc budgetsiface.BudgetsAPI) bool {
//        // Make svc.CreateBudget request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := budgets.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockBudgetsClient struct {
//        budgetsiface.BudgetsAPI
//    }
//    func (m *mockBudgetsClient) CreateBudget(input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockBudgetsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type BudgetsAPI interface {
	CreateBudget(*budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error)
	CreateBudgetWithContext(aws.Context, *budgets.CreateBudgetInput, ...request.Option) (*budgets.CreateBudgetOutput, error)
	CreateBudgetRequest(*budgets.CreateBudgetInput) (*request.Request, *budgets.CreateBudgetOutput)

	CreateBudgetAction(*budgets.CreateBudgetActionInput) (*budgets.CreateBudgetActionOutput, error)
	CreateBudgetActionWithContext(aws.Context, *budgets.CreateBudgetActionInput, ...request.Option) (*budgets.CreateBudgetActionOutput, error)
	CreateBudgetActionRequest(*budgets.CreateBudgetActionInput) (*request.Request, *budgets.CreateBudgetActionOutput)

	CreateNotification(*budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error)
	CreateNotificationWithContext(aws.Context, *budgets.CreateNotificationInput, ...request.Option) (*budgets.CreateNotificationOutput, error)
	CreateNotificationRequest(*budgets.CreateNotificationInput) (*request.Request, *budgets.CreateNotificationOutput)

	CreateSubscriber(*budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error)
	CreateSubscriberWithContext(aws.Context, *budgets.CreateSubscriberInput, ...request.Option) (*budgets.CreateSubscriberOutput, error)
	CreateSubscriberRequest(*budgets.CreateSubscriberInput) (*request.Request, *budgets.CreateSubscriberOutput)

	DeleteBudget(*budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error)
	DeleteBudgetWithContext(aws.Context, *budgets.DeleteBudgetInput, ...request.Option) (*budgets.DeleteBudgetOutput, error)
	DeleteBudgetRequest(*budgets.DeleteBudgetInput) (*request.Request, *budgets.DeleteBudgetOutput)

	DeleteBudgetAction(*budgets.DeleteBudgetActionInput) (*budgets.DeleteBudgetActionOutput, error)
	DeleteBudgetActionWithContext(aws.Context, *budgets.DeleteBudgetActionInput, ...request.Option) (*budgets.DeleteBudgetActionOutput, error)
	DeleteBudgetActionRequest(*budgets.DeleteBudgetActionInput) (*request.Request, *budgets.DeleteBudgetActionOutput)

	DeleteNotification(*budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error)
	DeleteNotificationWithContext(aws.Context, *budgets.DeleteNotificationInput, ...request.Option) (*budgets.DeleteNotificationOutput, error)
	DeleteNotificationRequest(*budgets.DeleteNotificationInput) (*request.Request, *budgets.DeleteNotificationOutput)

	DeleteSubscriber(*budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error)
	DeleteSubscriberWithContext(aws.Context, *budgets.DeleteSubscriberInput, ...request.Option) (*budgets.DeleteSubscriberOutput, error)
	DeleteSubscriberRequest(*budgets.DeleteSubscriberInput) (*request.Request, *budgets.DeleteSubscriberOutput)

	DescribeBudget(*budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error)
	DescribeBudgetWithContext(aws.Context, *budgets.DescribeBudgetInput, ...request.Option) (*budgets.DescribeBudgetOutput, error)
	DescribeBudgetRequest(*budgets.DescribeBudgetInput) (*request.Request, *budgets.DescribeBudgetOutput)

	DescribeBudgetAction(*budgets.DescribeBudgetActionInput) (*budgets.DescribeBudgetActionOutput, error)
	DescribeBudgetActionWithContext(aws.Context, *budgets.DescribeBudgetActionInput, ...request.Option) (*budgets.DescribeBudgetActionOutput, error)
	DescribeBudgetActionRequest(*budgets.DescribeBudgetActionInput) (*request.Request, *budgets.DescribeBudgetActionOutput)

	DescribeBudgetActionHistories(*budgets.DescribeBudgetActionHistoriesInput) (*budgets.DescribeBudgetActionHistoriesOutput, error)
	DescribeBudgetActionHistoriesWithContext(aws.Context, *budgets.DescribeBudgetActionHistoriesInput, ...request.Option) (*budgets.DescribeBudgetActionHistoriesOutput, error)
	DescribeBudgetActionHistoriesRequest(*budgets.DescribeBudgetActionHistoriesInput) (*request.Request, *budgets.DescribeBudgetActionHistoriesOutput)

	DescribeBudgetActionHistoriesPages(*budgets.DescribeBudgetActionHistoriesInput, func(*budgets.DescribeBudgetActionHistoriesOutput, bool) bool) error
	DescribeBudgetActionHistoriesPagesWithContext(aws.Context, *budgets.DescribeBudgetActionHistoriesInput, func(*budgets.DescribeBudgetActionHistoriesOutput, bool) bool, ...request.Option) error

	DescribeBudgetActionsForAccount(*budgets.DescribeBudgetActionsForAccountInput) (*budgets.DescribeBudgetActionsForAccountOutput, error)
	DescribeBudgetActionsForAccountWithContext(aws.Context, *budgets.DescribeBudgetActionsForAccountInput, ...request.Option) (*budgets.DescribeBudgetActionsForAccountOutput, error)
	DescribeBudgetActionsForAccountRequest(*budgets.DescribeBudgetActionsForAccountInput) (*request.Request, *budgets.DescribeBudgetActionsForAccountOutput)

	DescribeBudgetActionsForAccountPages(*budgets.DescribeBudgetActionsForAccountInput, func(*budgets.DescribeBudgetActionsForAccountOutput, bool) bool) error
	DescribeBudgetActionsForAccountPagesWithContext(aws.Context, *budgets.DescribeBudgetActionsForAccountInput, func(*budgets.DescribeBudgetActionsForAccountOutput, bool) bool, ...request.Option) error

	DescribeBudgetActionsForBudget(*budgets.DescribeBudgetActionsForBudgetInput) (*budgets.DescribeBudgetActionsForBudgetOutput, error)
	DescribeBudgetActionsForBudgetWithContext(aws.Context, *budgets.DescribeBudgetActionsForBudgetInput, ...request.Option) (*budgets.DescribeBudgetActionsForBudgetOutput, error)
	DescribeBudgetActionsForBudgetRequest(*budgets.DescribeBudgetActionsForBudgetInput) (*request.Request, *budgets.DescribeBudgetActionsForBudgetOutput)

	DescribeBudgetActionsForBudgetPages(*budgets.DescribeBudgetActionsForBudgetInput, func(*budgets.DescribeBudgetActionsForBudgetOutput, bool) bool) error
	DescribeBudgetActionsForBudgetPagesWithContext(aws.Context, *budgets.DescribeBudgetActionsForBudgetInput, func(*budgets.DescribeBudgetActionsForBudgetOutput, bool) bool, ...request.Option) error

	DescribeBudgetNotificationsForAccount(*budgets.DescribeBudgetNotificationsForAccountInput) (*budgets.DescribeBudgetNotificationsForAccountOutput, error)
	DescribeBudgetNotificationsForAccountWithContext(aws.Context, *budgets.DescribeBudgetNotificationsForAccountInput, ...request.Option) (*budgets.DescribeBudgetNotificationsForAccountOutput, error)
	DescribeBudgetNotificationsForAccountRequest(*budgets.DescribeBudgetNotificationsForAccountInput) (*request.Request, *budgets.DescribeBudgetNotificationsForAccountOutput)

	DescribeBudgetNotificationsForAccountPages(*budgets.DescribeBudgetNotificationsForAccountInput, func(*budgets.DescribeBudgetNotificationsForAccountOutput, bool) bool) error
	DescribeBudgetNotificationsForAccountPagesWithContext(aws.Context, *budgets.DescribeBudgetNotificationsForAccountInput, func(*budgets.DescribeBudgetNotificationsForAccountOutput, bool) bool, ...request.Option) error

	DescribeBudgetPerformanceHistory(*budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error)
	DescribeBudgetPerformanceHistoryWithContext(aws.Context, *budgets.DescribeBudgetPerformanceHistoryInput, ...request.Option) (*budgets.DescribeBudgetPerformanceHistoryOutput, error)
	DescribeBudgetPerformanceHistoryRequest(*budgets.DescribeBudgetPerformanceHistoryInput) (*request.Request, *budgets.DescribeBudgetPerformanceHistoryOutput)

	DescribeBudgetPerformanceHistoryPages(*budgets.DescribeBudgetPerformanceHistoryInput, func(*budgets.DescribeBudgetPerformanceHistoryOutput, bool) bool) error
	DescribeBudgetPerformanceHistoryPagesWithContext(aws.Context, *budgets.DescribeBudgetPerformanceHistoryInput, func(*budgets.DescribeBudgetPerformanceHistoryOutput, bool) bool, ...request.Option) error

	DescribeBudgets(*budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error)
	DescribeBudgetsWithContext(aws.Context, *budgets.DescribeBudgetsInput, ...request.Option) (*budgets.DescribeBudgetsOutput, error)
	DescribeBudgetsRequest(*budgets.DescribeBudgetsInput) (*request.Request, *budgets.DescribeBudgetsOutput)

	DescribeBudgetsPages(*budgets.DescribeBudgetsInput, func(*budgets.DescribeBudgetsOutput, bool) bool) error
	DescribeBudgetsPagesWithContext(aws.Context, *budgets.DescribeBudgetsInput, func(*budgets.DescribeBudgetsOutput, bool) bool, ...request.Option) error

	DescribeNotificationsForBudget(*budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error)
	DescribeNotificationsForBudgetWithContext(aws.Context, *budgets.DescribeNotificationsForBudgetInput, ...request.Option) (*budgets.DescribeNotificationsForBudgetOutput, error)
	DescribeNotificationsForBudgetRequest(*budgets.DescribeNotificationsForBudgetInput) (*request.Request, *budgets.DescribeNotificationsForBudgetOutput)

	DescribeNotificationsForBudgetPages(*budgets.DescribeNotificationsForBudgetInput, func(*budgets.DescribeNotificationsForBudgetOutput, bool) bool) error
	DescribeNotificationsForBudgetPagesWithContext(aws.Context, *budgets.DescribeNotificationsForBudgetInput, func(*budgets.DescribeNotificationsForBudgetOutput, bool) bool, ...request.Option) error

	DescribeSubscribersForNotification(*budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error)
	DescribeSubscribersForNotificationWithContext(aws.Context, *budgets.DescribeSubscribersForNotificationInput, ...request.Option) (*budgets.DescribeSubscribersForNotificationOutput, error)
	DescribeSubscribersForNotificationRequest(*budgets.DescribeSubscribersForNotificationInput) (*request.Request, *budgets.DescribeSubscribersForNotificationOutput)

	DescribeSubscribersForNotificationPages(*budgets.DescribeSubscribersForNotificationInput, func(*budgets.DescribeSubscribersForNotificationOutput, bool) bool) error
	DescribeSubscribersForNotificationPagesWithContext(aws.Context, *budgets.DescribeSubscribersForNotificationInput, func(*budgets.DescribeSubscribersForNotificationOutput, bool) bool, ...request.Option) error

	ExecuteBudgetAction(*budgets.ExecuteBudgetActionInput) (*budgets.ExecuteBudgetActionOutput, error)
	ExecuteBudgetActionWithContext(aws.Context, *budgets.ExecuteBudgetActionInput, ...request.Option) (*budgets.ExecuteBudgetActionOutput, error)
	ExecuteBudgetActionRequest(*budgets.ExecuteBudgetActionInput) (*request.Request, *budgets.ExecuteBudgetActionOutput)

	UpdateBudget(*budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error)
	UpdateBudgetWithContext(aws.Context, *budgets.UpdateBudgetInput, ...request.Option) (*budgets.UpdateBudgetOutput, error)
	UpdateBudgetRequest(*budgets.UpdateBudgetInput) (*request.Request, *budgets.UpdateBudgetOutput)

	UpdateBudgetAction(*budgets.UpdateBudgetActionInput) (*budgets.UpdateBudgetActionOutput, error)
	UpdateBudgetActionWithContext(aws.Context, *budgets.UpdateBudgetActionInput, ...request.Option) (*budgets.UpdateBudgetActionOutput, error)
	UpdateBudgetActionRequest(*budgets.UpdateBudgetActionInput) (*request.Request, *budgets.UpdateBudgetActionOutput)

	UpdateNotification(*budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error)
	UpdateNotificationWithContext(aws.Context, *budgets.UpdateNotificationInput, ...request.Option) (*budgets.UpdateNotificationOutput, error)
	UpdateNotificationRequest(*budgets.UpdateNotificationInput) (*request.Request, *budgets.UpdateNotificationOutput)

	UpdateSubscriber(*budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error)
	UpdateSubscriberWithContext(aws.Context, *budgets.UpdateSubscriberInput, ...request.Option) (*budgets.UpdateSubscriberOutput, error)
	UpdateSubscriberRequest(*budgets.UpdateSubscriberInput) (*request.Request, *budgets.UpdateSubscriberOutput)
}

var _ BudgetsAPI = (*budgets.Budgets)(nil)