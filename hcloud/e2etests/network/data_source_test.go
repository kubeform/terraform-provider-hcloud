package network_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/e2etests"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/network"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/loadbalancer"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testsupport"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testtemplate"
)

func TestAccHcloudDataSourceNetworkTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	res := &network.RData{
		Name:    "network-ds-test",
		IPRange: "10.0.0.0/16",
		Labels: map[string]string{
			"key": strconv.Itoa(acctest.RandInt()),
		},
	}
	res.SetRName("network-ds-test")
	networkByName := &network.DData{
		NetworkName: res.TFID() + ".name",
	}
	networkByName.SetRName("network_by_name")
	networkByID := &network.DData{
		NetworkID: res.TFID() + ".id",
	}
	networkByID.SetRName("network_by_id")
	networkBySel := &network.DData{
		LabelSelector: fmt.Sprintf("key=${%s.labels[\"key\"]}", res.TFID()),
	}
	networkBySel.SetRName("network_by_sel")

	resource.Test(t, resource.TestCase{
		PreCheck:     e2etests.PreCheck(t),
		Providers:    e2etests.Providers(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(loadbalancer.ResourceType, loadbalancer.ByID(t, nil)),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_network", res,
				),
			},
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_network", res,
					"testdata/d/hcloud_network", networkByName,
					"testdata/d/hcloud_network", networkByID,
					"testdata/d/hcloud_network", networkBySel,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(networkByName.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(networkByName.TFID(), "ip_range", res.IPRange),

					resource.TestCheckResourceAttr(networkByID.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(networkByID.TFID(), "ip_range", res.IPRange),

					resource.TestCheckResourceAttr(networkBySel.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(networkBySel.TFID(), "ip_range", res.IPRange),
				),
			},
		},
	})
}
