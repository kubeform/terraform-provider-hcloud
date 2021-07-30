package hcloud

import (
	"testing"

	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/certificate"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/datacenter"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/firewall"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/floatingip"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/image"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/loadbalancer"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/location"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/network"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/rdns"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/server"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/servertype"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/snapshot"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/sshkey"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/volume"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_Resources(t *testing.T) {
	var provider = Provider()
	expectedResources := []string{
		certificate.ResourceType,
		firewall.ResourceType,
		certificate.UploadedResourceType,
		certificate.ManagedResourceType,
		floatingip.AssignmentResourceType,
		floatingip.ResourceType,
		loadbalancer.NetworkResourceType,
		loadbalancer.ResourceType,
		loadbalancer.ServiceResourceType,
		loadbalancer.TargetResourceType,
		network.ResourceType,
		network.RouteResourceType,
		network.SubnetResourceType,
		rdns.ResourceType,
		server.NetworkResourceType,
		server.ResourceType,
		snapshot.ResourceType,
		sshkey.ResourceType,
		volume.AttachmentResourceType,
		volume.ResourceType,
	}

	resources := provider.Resources()
	assert.Len(t, resources, len(expectedResources))

	for _, datasource := range resources {
		assert.Contains(t, expectedResources, datasource.Name)
	}
}

func TestProvider_DataSources(t *testing.T) {
	var provider = Provider()
	expectedDataSources := []string{
		certificate.DataSourceType,
		datacenter.DatacentersDataSourceType,
		datacenter.DataSourceType,
		firewall.DataSourceType,
		floatingip.DataSourceType,
		image.DataSourceType,
		loadbalancer.DataSourceType,
		location.DataSourceType,
		location.LocationsDataSourceType,
		network.DataSourceType,
		server.DataSourceType,
		servertype.DataSourceType,
		servertype.ServerTypesDataSourceType,
		sshkey.DataSourceType,
		sshkey.SSHKeysDataSourceType,
		volume.DataSourceType,
	}

	datasources := provider.DataSources()
	assert.Len(t, datasources, len(expectedDataSources))

	for _, datasource := range datasources {
		assert.Contains(t, expectedDataSources, datasource.Name)
	}
}
