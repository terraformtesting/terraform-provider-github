package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsSecret(t *testing.T) {

	testRepo := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(5,
		acctest.CharSetAlphaNum))

	t.Run("reads a repository public key without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_actions_public_key" "test_pk" {
			  repository = github_repository.test.name
			}

			resource "github_repository" "test" {
			  name = "%s"
			}
		`, testRepo)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_actions_public_key.test_pk", "key_id"),
			resource.TestCheckResourceAttrSet("data.github_actions_public_key.test_pk", "key"),
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

	t.Run("creates and updates secrets without error", func(t *testing.T) {

		secretValue := "super_secret_value"
		updatedSecretValue := "updated_super_secret_value"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_actions_secret" "test_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_secret_name"
			  plaintext_value  = "%s"
			}
		`, testRepo, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_actions_secret.test_secret", "plaintext_value", secretValue),
				resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "created_at"),
				resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "updated_at"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_actions_secret.test_secret", "plaintext_value", updatedSecretValue),
				resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "created_at"),
				resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "updated_at"),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							secretValue,
							updatedSecretValue, 1),
						Check: checks["after"],
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

	// t.Run("deletes secrets without error", func(t *testing.T) {
	//
	// 	secretValue := "super_secret_value"
	//
	// 	config := fmt.Sprintf(`
	// 			resource "github_repository" "test" {
	// 				name = "%s"
	// 			}
	//
	// 			resource "github_actions_secret" "test_secret" {
	// 				repository       = github_repository.test.name
	// 				secret_name      = "test_secret_name"
	// 				plaintext_value  = "%s"
	// 			}
	// 		`, testRepo, secretValue)
	//
	// 	testCase := func(t *testing.T, mode string) {
	// 		resource.Test(t, resource.TestCase{
	// 			PreCheck:  func() { skipUnlessMode(t, mode) },
	// 			Providers: testAccProviders,
	// 			Steps: []resource.TestStep{
	// 				{
	// 					Config:             config,
	// 					Destroy:            true,
	// 					ExpectNonEmptyPlan: true,
	// 				},
	// 			},
	// 		})
	// 	}
	//
	// 	t.Run("with an anonymous account", func(t *testing.T) {
	// 		t.Skip("anonymous account not supported for this operation")
	// 	})
	//
	// 	t.Run("with an individual account", func(t *testing.T) {
	// 		testCase(t, individual)
	// 	})
	//
	// 	t.Run("with an organization account", func(t *testing.T) {
	// 		testCase(t, organization)
	// 	})
	//
	// })

}
