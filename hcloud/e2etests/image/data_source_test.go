package image_test

import (
	"testing"

	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/e2etests"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/image"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/loadbalancer"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testsupport"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testtemplate"
)

const TestImageName = e2etests.TestImage
const TestImageID = "15512617"

func TestAccHcloudDataSourceImageTest(t *testing.T) {
	tmplMan := testtemplate.Manager{}

	imageByName := &image.DData{
		ImageName: TestImageName,
	}
	imageByName.SetRName("image_by_name")
	imageByID := &image.DData{
		ImageID: TestImageID,
	}
	imageByID.SetRName("image_by_id")

	resource.Test(t, resource.TestCase{
		PreCheck:     e2etests.PreCheck(t),
		Providers:    e2etests.Providers(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(loadbalancer.ResourceType, loadbalancer.ByID(t, nil)),
		Steps: []resource.TestStep{
			{
				Config: tmplMan.Render(t,
					"testdata/d/hcloud_image", imageByName,
					"testdata/d/hcloud_image", imageByID,
				),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(imageByName.TFID(),
						"name", TestImageName),
					resource.TestCheckResourceAttr(imageByName.TFID(), "id", TestImageID),

					resource.TestCheckResourceAttr(imageByID.TFID(),
						"name", TestImageName),
					resource.TestCheckResourceAttr(imageByID.TFID(), "id", TestImageID),
				),
			},
		},
	})
}
