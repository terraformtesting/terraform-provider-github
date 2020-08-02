package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationDataSource(t *testing.T) {

	t.Run("queries for an organization without error", func(t *testing.T) {

		organizationConfiguration := fmt.Sprintf(`
			provider "github" {
				organization = "%s"
				token = "%s"
			}
			data "github_organization" "test" { name = "%s" }
		`, testOrganization, testToken, testOrganization)

		organizationCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization.test", "login"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "description"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
		)

		testCase := func(mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: organizationConfiguration,
						Check:  organizationCheck,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase("individual")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase("organization")
		})

	})
}
