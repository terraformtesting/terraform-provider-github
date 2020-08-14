package github

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubReleaseDataSource(t *testing.T) {

	t.Run("queries latest release", func(t *testing.T) {

		config := `
			data "github_release" "test" {
				repository = "terraform-template-module"
				owner = "terraformtesting"
				retrieve_by = "latest"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_release.test", "active", "true",
			),
			resource.TestCheckResourceAttr(
				"github_release.test", "events.#", "1",
			),
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

	// t.Run("queries release by ID", func(t *testing.T) {
	//
	// 	config := `
	// 		data "github_release" "test" {
	// 			repository = "torvalds/linux"
	// 			owner = "owner"
	// 			retrieve_by = "id"
	// 			release_id = "0"
	// 		}
	// 	`
	//
	// 	check := resource.ComposeTestCheckFunc(
	// 	// resource.TestCheckResourceAttr(
	// 	// 	"github_repository_webhook.test", "active", "true",
	// 	// ),
	// 	// resource.TestCheckResourceAttr(
	// 	// 	"github_repository_webhook.test", "events.#", "1",
	// 	// ),
	// 	)
	//
	// 	testCase := func(t *testing.T, mode string) {
	// 		resource.Test(t, resource.TestCase{
	// 			PreCheck:  func() { skipUnlessMode(t, mode) },
	// 			Providers: testAccProviders,
	// 			Steps: []resource.TestStep{
	// 				{
	// 					Config: config,
	// 					Check:  check,
	// 				},
	// 			},
	// 		})
	// 	}
	//
	// 	t.Run("with an anonymous account", func(t *testing.T) {
	// 		testCase(t, anonymous)
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
	//
	// t.Run("queries release by tag", func(t *testing.T) {
	//
	// 	config := `
	// 		data "github_release" "test" {
	// 			repository = "torvalds/linux"
	// 			owner = "owner"
	// 			retrieve_by = "id"
	// 			release_id = "0"
	// 		}
	// 	`
	//
	// 	check := resource.ComposeTestCheckFunc(
	// 	// resource.TestCheckResourceAttr(
	// 	// 	"github_repository_webhook.test", "active", "true",
	// 	// ),
	// 	// resource.TestCheckResourceAttr(
	// 	// 	"github_repository_webhook.test", "events.#", "1",
	// 	// ),
	// 	)
	//
	// 	testCase := func(t *testing.T, mode string) {
	// 		resource.Test(t, resource.TestCase{
	// 			PreCheck:  func() { skipUnlessMode(t, mode) },
	// 			Providers: testAccProviders,
	// 			Steps: []resource.TestStep{
	// 				{
	// 					Config: config,
	// 					Check:  check,
	// 				},
	// 			},
	// 		})
	// 	}
	//
	// 	t.Run("with an anonymous account", func(t *testing.T) {
	// 		testCase(t, anonymous)
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

	t.Run("errors when querying with non-existent ID", func(t *testing.T) {

		config := `
			data "github_release" "test" {
				repository = "torvalds/linux"
				owner = "owner"
				retrieve_by = "id"
			}
		`

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("`release_id` must be set when `retrieve_by` = `id`"),
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

	t.Run("errors when querying with non-existent repository", func(t *testing.T) {

		config := `
			data "github_release" "test" {
				repository = "test"
				owner = "test"
				retrieve_by = "latest"
			}
		`
		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile(`Not Found`),
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

	t.Run("errors when querying with non-existent tag", func(t *testing.T) {

		config := `
			data "github_release" "test" {
				repository = "test"
				owner = "test"
				retrieve_by = "tag"
			}
		`
		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("`release_tag` must be set when `retrieve_by` = `tag`"),
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
