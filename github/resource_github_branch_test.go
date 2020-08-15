package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubBranch(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a branch without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
		// resource.TestCheckResourceAttr("github_actions_secret.test_secret", "plaintext_value", secretValue),
		// resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "created_at"),
		// resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "updated_at"),
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
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("deletes a branch without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:             config,
						Destroy:            true,
						ExpectNonEmptyPlan: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

}
