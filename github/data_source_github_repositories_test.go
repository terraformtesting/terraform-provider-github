package github

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoriesDataSource(t *testing.T) {

	t.Run("queries a list of repositories without error", func(t *testing.T) {

		config := `
			data "github_repositories" "test" {
				query = "repository:test-repo"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^`+testOrganization)),
			resource.TestCheckResourceAttrSet("data.github_repositories.test", "names.0"),
			resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "updated"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("returns an empty list given an invalid query", func(t *testing.T) {

		config := `
			data "github_repositories" "test" {
				query = "klsafj_23434_doesnt_exist"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_repositories.test", "full_names.#", "0"),
			resource.TestCheckResourceAttr("data.github_repositories.test", "names.#", "0"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}
