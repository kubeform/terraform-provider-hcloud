package server_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/e2etests"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/server"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/loadbalancer"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testsupport"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testtemplate"
)

func TestAccHcloudDataSourceServerTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	res := &server.RData{
		Name:  "server-ds-test",
		Type:  e2etests.TestServerType,
		Image: e2etests.TestImage,
		Labels: map[string]string{
			"key": strconv.Itoa(acctest.RandInt()),
		},
	}
	res.SetRName("server-ds-test")
	serverByName := &server.DData{
		ServerName: res.TFID() + ".name",
	}
	serverByName.SetRName("server_by_name")
	serverByID := &server.DData{
		ServerID: res.TFID() + ".id",
	}
	serverByID.SetRName("server_by_id")

	serverBySel := &server.DData{
		LabelSelector: fmt.Sprintf("key=${%s.labels[\"key\"]}", res.TFID()),
	}
	serverBySel.SetRName("server_by_sel")
	resource.Test(t, resource.TestCase{
		PreCheck:     e2etests.PreCheck(t),
		Providers:    e2etests.Providers(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(loadbalancer.ResourceType, loadbalancer.ByID(t, nil)),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_server", res,
				),
			},
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_server", res,
					"testdata/d/hcloud_server", serverByName,
					"testdata/d/hcloud_server", serverByID,
					"testdata/d/hcloud_server", serverBySel,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(serverByName.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),

					resource.TestCheckResourceAttr(serverByID.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(serverBySel.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
				),
			},
		},
	})
}
