package loadbalancer_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/e2etests"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/loadbalancer"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testsupport"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testtemplate"
)

func TestAccHcloudDataSourceLoadBalancerTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	res := &loadbalancer.RData{
		Name:         "some-load-balancer",
		LocationName: e2etests.TestLocationName,
		Labels: map[string]string{
			"key": strconv.Itoa(acctest.RandInt()),
		},
	}
	lbByName := &loadbalancer.DData{
		Name:             "lb_by_name",
		LoadBalancerName: res.TFID() + ".name",
	}
	lbByID := &loadbalancer.DData{
		Name:           "lb_by_id",
		LoadBalancerID: res.TFID() + ".id",
	}
	lbBySel := &loadbalancer.DData{
		Name:          "lb_by_sel",
		LabelSelector: fmt.Sprintf("key=${%s.labels[\"key\"]}", res.TFID()),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     e2etests.PreCheck(t),
		Providers:    e2etests.Providers(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(loadbalancer.ResourceType, loadbalancer.ByID(t, nil)),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_load_balancer", res,
				),
			},
			{
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_load_balancer", res,
					"testdata/d/hcloud_load_balancer", lbByName,
					"testdata/d/hcloud_load_balancer", lbByID,
					"testdata/d/hcloud_load_balancer", lbBySel,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(lbByName.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(lbByName.TFID(), "location", res.LocationName),
					resource.TestCheckResourceAttr(lbByName.TFID(), "target.#", "0"),
					resource.TestCheckResourceAttr(lbByName.TFID(), "service.#", "0"),

					resource.TestCheckResourceAttr(lbByID.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(lbByID.TFID(), "location", res.LocationName),
					resource.TestCheckResourceAttr(lbByID.TFID(), "targets.#", "0"),
					resource.TestCheckResourceAttr(lbByID.TFID(), "service.#", "0"),

					resource.TestCheckResourceAttr(lbBySel.TFID(),
						"name", fmt.Sprintf("%s--%d", res.Name, tmplMan.RandInt)),
					resource.TestCheckResourceAttr(lbBySel.TFID(), "location", res.LocationName),
					resource.TestCheckResourceAttr(lbBySel.TFID(), "targets.#", "0"),
					resource.TestCheckResourceAttr(lbBySel.TFID(), "service.#", "0"),
				),
			},
		},
	})
}
