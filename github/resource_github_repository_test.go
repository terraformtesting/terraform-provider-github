package github

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositories(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates repositories without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {

			  name         = "tf-acc-test-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"

			  has_issues         = true
			  has_wiki           = true
			  has_downloads      = true
			  allow_merge_commit = true
			  allow_squash_merge = false
			  allow_rebase_merge = false
			  auto_init          = false

			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository.test", "has_issues", "true"),
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

	t.Run("imports repositories without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				auto_init 	 = false
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_repository.test", "name"),
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
					{
						ResourceName:      "github_repository.test",
						ImportState:       true,
						ImportStateVerify: true,
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

	t.Run("archives repositories without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				archived     = false
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_repository.test", "archived", "false"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_repository.test", "archived", "true"),
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
							`archived     = false`,
							`archived     = true`, 1),
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

	t.Run("manages the project feature for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				has_projects = false
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_repository.test", "has_projects", "false"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_repository.test", "has_projects", "true"),
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
							`has_projects = false`,
							`has_projects = true`, 1),
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

	t.Run("manages the default branch feature for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name           = "tf-acc-test-%[1]s"
			  description    = "Terraform acceptance tests %[1]s"
				default_branch = "development"
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_repository.test", "default_branch", "development"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_repository.test", "default_branch", "staging"),
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
							`default_branch = "development"`,
							`default_branch = "staging"`, 1),
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

	t.Run("manages the license and gitignore feature for a repository", func(t *testing.T) {

		// config := fmt.Sprintf(`
		// 	data "github_repositories" "test" {
		// 		query = "org:%s repository:test-repo"
		// 	}
		// `, testOrganization)
		//
		// check := resource.ComposeTestCheckFunc(
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^`+testOrganization)),
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "names.0", regexp.MustCompile(`^test`)),
		// 	resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "updated"),
		// )
		//
		// testCase := func(t *testing.T, mode string) {
		// 	resource.Test(t, resource.TestCase{
		// 		PreCheck:  func() { skipUnlessMode(t, mode) },
		// 		Providers: testAccProviders,
		// 		Steps: []resource.TestStep{
		// 			{
		// 				Config: config,
		// 				Check:  check,
		// 			},
		// 		},
		// 	})
		// }
		//
		// t.Run("with an anonymous account", func(t *testing.T) {
		// 	testCase(t, anonymous)
		// })
		//
		// t.Run("with an individual account", func(t *testing.T) {
		// 	testCase(t, individual)
		// })
		//
		// t.Run("with an organization account", func(t *testing.T) {
		// 	testCase(t, organization)
		// })

	})

	t.Run("configures topics for a repository", func(t *testing.T) {

		// config := fmt.Sprintf(`
		// 	data "github_repositories" "test" {
		// 		query = "org:%s repository:test-repo"
		// 	}
		// `, testOrganization)
		//
		// check := resource.ComposeTestCheckFunc(
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^`+testOrganization)),
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "names.0", regexp.MustCompile(`^test`)),
		// 	resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "updated"),
		// )
		//
		// testCase := func(t *testing.T, mode string) {
		// 	resource.Test(t, resource.TestCase{
		// 		PreCheck:  func() { skipUnlessMode(t, mode) },
		// 		Providers: testAccProviders,
		// 		Steps: []resource.TestStep{
		// 			{
		// 				Config: config,
		// 				Check:  check,
		// 			},
		// 		},
		// 	})
		// }
		//
		// t.Run("with an anonymous account", func(t *testing.T) {
		// 	testCase(t, anonymous)
		// })
		//
		// t.Run("with an individual account", func(t *testing.T) {
		// 	testCase(t, individual)
		// })
		//
		// t.Run("with an organization account", func(t *testing.T) {
		// 	testCase(t, organization)
		// })

	})

	t.Run("creates a repository using a template", func(t *testing.T) {

		// config := fmt.Sprintf(`
		// 	data "github_repositories" "test" {
		// 		query = "org:%s repository:test-repo"
		// 	}
		// `, testOrganization)
		//
		// check := resource.ComposeTestCheckFunc(
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^`+testOrganization)),
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "names.0", regexp.MustCompile(`^test`)),
		// 	resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "updated"),
		// )
		//
		// testCase := func(t *testing.T, mode string) {
		// 	resource.Test(t, resource.TestCase{
		// 		PreCheck:  func() { skipUnlessMode(t, mode) },
		// 		Providers: testAccProviders,
		// 		Steps: []resource.TestStep{
		// 			{
		// 				Config: config,
		// 				Check:  check,
		// 			},
		// 		},
		// 	})
		// }
		//
		// t.Run("with an anonymous account", func(t *testing.T) {
		// 	testCase(t, anonymous)
		// })
		//
		// t.Run("with an individual account", func(t *testing.T) {
		// 	testCase(t, individual)
		// })
		//
		// t.Run("with an organization account", func(t *testing.T) {
		// 	testCase(t, organization)
		// })

	})

}

func testSweepRepositories(region string) error {
	meta, err := sharedConfigForRegion(region)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client

	repos, _, err := client.Repositories.List(context.TODO(), meta.(*Owner).name, nil)
	if err != nil {
		return err
	}

	for _, r := range repos {
		if name := r.GetName(); strings.HasPrefix(name, "tf-acc-") || strings.HasPrefix(name, "foo-") {
			log.Printf("Destroying Repository %s", name)

			if _, err := client.Repositories.Delete(context.TODO(), meta.(*Owner).name, name); err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	resource.AddTestSweepers("github_repository", &resource.Sweeper{
		Name: "github_repository",
		F:    testSweepRepositories,
	})
}
