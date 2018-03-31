// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pinpointiface provides an interface to enable mocking the Amazon Pinpoint service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pinpointiface

import (
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

// PinpointAPI provides an interface to enable mocking the
// pinpoint.Pinpoint service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Pinpoint.
//    func myFunc(svc pinpointiface.PinpointAPI) bool {
//        // Make svc.CreateApp request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := pinpoint.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPinpointClient struct {
//        pinpointiface.PinpointAPI
//    }
//    func (m *mockPinpointClient) CreateApp(input *pinpoint.CreateAppInput) (*pinpoint.CreateAppOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPinpointClient{}
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
type PinpointAPI interface {
	CreateAppRequest(*pinpoint.CreateAppInput) pinpoint.CreateAppRequest

	CreateCampaignRequest(*pinpoint.CreateCampaignInput) pinpoint.CreateCampaignRequest

	CreateImportJobRequest(*pinpoint.CreateImportJobInput) pinpoint.CreateImportJobRequest

	CreateSegmentRequest(*pinpoint.CreateSegmentInput) pinpoint.CreateSegmentRequest

	DeleteAdmChannelRequest(*pinpoint.DeleteAdmChannelInput) pinpoint.DeleteAdmChannelRequest

	DeleteApnsChannelRequest(*pinpoint.DeleteApnsChannelInput) pinpoint.DeleteApnsChannelRequest

	DeleteApnsSandboxChannelRequest(*pinpoint.DeleteApnsSandboxChannelInput) pinpoint.DeleteApnsSandboxChannelRequest

	DeleteApnsVoipChannelRequest(*pinpoint.DeleteApnsVoipChannelInput) pinpoint.DeleteApnsVoipChannelRequest

	DeleteApnsVoipSandboxChannelRequest(*pinpoint.DeleteApnsVoipSandboxChannelInput) pinpoint.DeleteApnsVoipSandboxChannelRequest

	DeleteAppRequest(*pinpoint.DeleteAppInput) pinpoint.DeleteAppRequest

	DeleteBaiduChannelRequest(*pinpoint.DeleteBaiduChannelInput) pinpoint.DeleteBaiduChannelRequest

	DeleteCampaignRequest(*pinpoint.DeleteCampaignInput) pinpoint.DeleteCampaignRequest

	DeleteEmailChannelRequest(*pinpoint.DeleteEmailChannelInput) pinpoint.DeleteEmailChannelRequest

	DeleteEventStreamRequest(*pinpoint.DeleteEventStreamInput) pinpoint.DeleteEventStreamRequest

	DeleteGcmChannelRequest(*pinpoint.DeleteGcmChannelInput) pinpoint.DeleteGcmChannelRequest

	DeleteSegmentRequest(*pinpoint.DeleteSegmentInput) pinpoint.DeleteSegmentRequest

	DeleteSmsChannelRequest(*pinpoint.DeleteSmsChannelInput) pinpoint.DeleteSmsChannelRequest

	GetAdmChannelRequest(*pinpoint.GetAdmChannelInput) pinpoint.GetAdmChannelRequest

	GetApnsChannelRequest(*pinpoint.GetApnsChannelInput) pinpoint.GetApnsChannelRequest

	GetApnsSandboxChannelRequest(*pinpoint.GetApnsSandboxChannelInput) pinpoint.GetApnsSandboxChannelRequest

	GetApnsVoipChannelRequest(*pinpoint.GetApnsVoipChannelInput) pinpoint.GetApnsVoipChannelRequest

	GetApnsVoipSandboxChannelRequest(*pinpoint.GetApnsVoipSandboxChannelInput) pinpoint.GetApnsVoipSandboxChannelRequest

	GetAppRequest(*pinpoint.GetAppInput) pinpoint.GetAppRequest

	GetApplicationSettingsRequest(*pinpoint.GetApplicationSettingsInput) pinpoint.GetApplicationSettingsRequest

	GetAppsRequest(*pinpoint.GetAppsInput) pinpoint.GetAppsRequest

	GetBaiduChannelRequest(*pinpoint.GetBaiduChannelInput) pinpoint.GetBaiduChannelRequest

	GetCampaignRequest(*pinpoint.GetCampaignInput) pinpoint.GetCampaignRequest

	GetCampaignActivitiesRequest(*pinpoint.GetCampaignActivitiesInput) pinpoint.GetCampaignActivitiesRequest

	GetCampaignVersionRequest(*pinpoint.GetCampaignVersionInput) pinpoint.GetCampaignVersionRequest

	GetCampaignVersionsRequest(*pinpoint.GetCampaignVersionsInput) pinpoint.GetCampaignVersionsRequest

	GetCampaignsRequest(*pinpoint.GetCampaignsInput) pinpoint.GetCampaignsRequest

	GetEmailChannelRequest(*pinpoint.GetEmailChannelInput) pinpoint.GetEmailChannelRequest

	GetEndpointRequest(*pinpoint.GetEndpointInput) pinpoint.GetEndpointRequest

	GetEventStreamRequest(*pinpoint.GetEventStreamInput) pinpoint.GetEventStreamRequest

	GetGcmChannelRequest(*pinpoint.GetGcmChannelInput) pinpoint.GetGcmChannelRequest

	GetImportJobRequest(*pinpoint.GetImportJobInput) pinpoint.GetImportJobRequest

	GetImportJobsRequest(*pinpoint.GetImportJobsInput) pinpoint.GetImportJobsRequest

	GetSegmentRequest(*pinpoint.GetSegmentInput) pinpoint.GetSegmentRequest

	GetSegmentImportJobsRequest(*pinpoint.GetSegmentImportJobsInput) pinpoint.GetSegmentImportJobsRequest

	GetSegmentVersionRequest(*pinpoint.GetSegmentVersionInput) pinpoint.GetSegmentVersionRequest

	GetSegmentVersionsRequest(*pinpoint.GetSegmentVersionsInput) pinpoint.GetSegmentVersionsRequest

	GetSegmentsRequest(*pinpoint.GetSegmentsInput) pinpoint.GetSegmentsRequest

	GetSmsChannelRequest(*pinpoint.GetSmsChannelInput) pinpoint.GetSmsChannelRequest

	PutEventStreamRequest(*pinpoint.PutEventStreamInput) pinpoint.PutEventStreamRequest

	SendMessagesRequest(*pinpoint.SendMessagesInput) pinpoint.SendMessagesRequest

	SendUsersMessagesRequest(*pinpoint.SendUsersMessagesInput) pinpoint.SendUsersMessagesRequest

	UpdateAdmChannelRequest(*pinpoint.UpdateAdmChannelInput) pinpoint.UpdateAdmChannelRequest

	UpdateApnsChannelRequest(*pinpoint.UpdateApnsChannelInput) pinpoint.UpdateApnsChannelRequest

	UpdateApnsSandboxChannelRequest(*pinpoint.UpdateApnsSandboxChannelInput) pinpoint.UpdateApnsSandboxChannelRequest

	UpdateApnsVoipChannelRequest(*pinpoint.UpdateApnsVoipChannelInput) pinpoint.UpdateApnsVoipChannelRequest

	UpdateApnsVoipSandboxChannelRequest(*pinpoint.UpdateApnsVoipSandboxChannelInput) pinpoint.UpdateApnsVoipSandboxChannelRequest

	UpdateApplicationSettingsRequest(*pinpoint.UpdateApplicationSettingsInput) pinpoint.UpdateApplicationSettingsRequest

	UpdateBaiduChannelRequest(*pinpoint.UpdateBaiduChannelInput) pinpoint.UpdateBaiduChannelRequest

	UpdateCampaignRequest(*pinpoint.UpdateCampaignInput) pinpoint.UpdateCampaignRequest

	UpdateEmailChannelRequest(*pinpoint.UpdateEmailChannelInput) pinpoint.UpdateEmailChannelRequest

	UpdateEndpointRequest(*pinpoint.UpdateEndpointInput) pinpoint.UpdateEndpointRequest

	UpdateEndpointsBatchRequest(*pinpoint.UpdateEndpointsBatchInput) pinpoint.UpdateEndpointsBatchRequest

	UpdateGcmChannelRequest(*pinpoint.UpdateGcmChannelInput) pinpoint.UpdateGcmChannelRequest

	UpdateSegmentRequest(*pinpoint.UpdateSegmentInput) pinpoint.UpdateSegmentRequest

	UpdateSmsChannelRequest(*pinpoint.UpdateSmsChannelInput) pinpoint.UpdateSmsChannelRequest
}

var _ PinpointAPI = (*pinpoint.Pinpoint)(nil)
