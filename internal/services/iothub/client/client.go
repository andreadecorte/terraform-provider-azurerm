package client

import (
	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-03-31/devices"
	"github.com/Azure/azure-sdk-for-go/services/provisioningservices/mgmt/2021-10-15/iothub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	ResourceClient       *devices.IotHubResourceClient
	DPSResourceClient    *iothub.IotDpsResourceClient
	DPSCertificateClient *iothub.DpsCertificateClient
}

func NewClient(o *common.ClientOptions) *Client {
	ResourceClient := devices.NewIotHubResourceClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&ResourceClient.Client, o.ResourceManagerAuthorizer)

	DPSResourceClient := iothub.NewIotDpsResourceClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DPSResourceClient.Client, o.ResourceManagerAuthorizer)

	DPSCertificateClient := iothub.NewDpsCertificateClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&DPSCertificateClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ResourceClient:       &ResourceClient,
		DPSResourceClient:    &DPSResourceClient,
		DPSCertificateClient: &DPSCertificateClient,
	}
}
