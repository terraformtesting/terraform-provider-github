package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryWebhook(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates repository webhooks without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "test-%[1]s"
			  description  = "Terraform acceptance tests"
			}

			resource "github_repository_webhook" "test" {
			  depends_on = ["github_repository.test"]
			  repository = "test-%[1]s"

			  configuration {
			    url          = "https://google.de/webhook"
			    content_type = "json"
			    insecure_ssl = true
			  }

			  events = ["pull_request"]
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_webhook.test", "active", "true",
			),
			resource.TestCheckResourceAttr(
				"github_repository_webhook.test", "events.#", "1",
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
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("imports repository webhooks without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "test-%[1]s"
				description  = "Terraform acceptance tests"
			}

			resource "github_repository_webhook" "test" {
				depends_on = ["github_repository.test"]
				repository = "test-%[1]s"
				configuration {
					url          = "https://google.de/webhook"
					content_type = "json"
					insecure_ssl = true
				}
				events = ["pull_request"]
			}
			`, randomID)

		check := resource.ComposeTestCheckFunc()

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
						ResourceName:        "github_repository_webhook.test",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("test-%s/", randomID),
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

	t.Run("updates repository webhooks without error", func(t *testing.T) {

		configs := map[string]string{
			"before": fmt.Sprintf(`
				resource "github_repository" "test" {
				  name         = "test-%[1]s"
				  description  = "Terraform acceptance tests"
				}

				resource "github_repository_webhook" "test" {
				  depends_on = ["github_repository.test"]
				  repository = "test-%[1]s"

				  configuration {
				    url          = "https://google.de/webhook"
				    content_type = "json"
				    insecure_ssl = true
				  }

				  events = ["pull_request"]
				}
			`, randomID),
			"after": fmt.Sprintf(`
				resource "github_repository" "test" {
				  name         = "test-%[1]s"
				  description  = "Terraform acceptance tests"
				}

				resource "github_repository_webhook" "test" {
				  depends_on = ["github_repository.test"]
				  repository = "test-%[1]s"

				  configuration {
				    secret       = "%[1]s"
				    url          = "https://google.de/webhook"
				    content_type = "json"
				    insecure_ssl = true
				  }

				  events = ["pull_request"]
				}
			`, randomID),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.TestCheckResourceAttr(
				"github_repository_webhook.test", "events.#", "1",
			),
			"after": resource.TestCheckResourceAttr(
				"github_repository_webhook.test", "events.#", "2",
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configs["before"],
						Check:  checks["before"],
					},
					{
						Config: configs["after"],
						Check:  checks["after"],
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

// func TestAccGithubRepositoryWebhook_basic(t *testing.T) {
// 	rn := "github_repository_webhook.foo"
// 	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	var hook github.Hook
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckGithubRepositoryWebhookDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccGithubRepositoryWebhookConfig(randString),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryWebhookExists(rn, fmt.Sprintf("foo-%s", randString), &hook),
// 					testAccCheckGithubRepositoryWebhookAttributes(&hook, &testAccGithubRepositoryWebhookExpectedAttributes{
// 						Events: []string{"pull_request"},
// 						Configuration: map[string]interface{}{
// 							"url":          "https://google.de/webhook",
// 							"content_type": "json",
// 							"insecure_ssl": "1",
// 						},
// 						Active: true,
// 					}),
// 				),
// 			},
// 			{
// 				Config: testAccGithubRepositoryWebhookUpdateConfig(randString),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryWebhookExists(rn, fmt.Sprintf("foo-%s", randString), &hook),
// 					testAccCheckGithubRepositoryWebhookAttributes(&hook, &testAccGithubRepositoryWebhookExpectedAttributes{
// 						Events: []string{"issues"},
// 						Configuration: map[string]interface{}{
// 							"url":          "https://google.de/webhooks",
// 							"content_type": "form",
// 							"insecure_ssl": "0",
// 						},
// 						Active: false,
// 					}),
// 				),
// 			},
// 			{
// 				ResourceName:        rn,
// 				ImportState:         true,
// 				ImportStateVerify:   true,
// 				ImportStateIdPrefix: fmt.Sprintf("foo-%s/", randString),
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubRepositoryWebhook_secret(t *testing.T) {
// 	rn := "github_repository_webhook.foo"
// 	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	var hook github.Hook
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckGithubRepositoryWebhookDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccGithubRepositoryWebhookConfig_secret(randString),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryWebhookExists(rn, fmt.Sprintf("foo-%s", randString), &hook),
// 					testAccCheckGithubRepositoryWebhookAttributes(&hook, &testAccGithubRepositoryWebhookExpectedAttributes{
// 						Events: []string{"pull_request"},
// 						Configuration: map[string]interface{}{
// 							"url":          "https://www.terraform.io/webhook",
// 							"content_type": "json",
// 							"secret":       "********",
// 							"insecure_ssl": "0",
// 						},
// 						Active: true,
// 					}),
// 				),
// 			},
// 			{
// 				ResourceName:            rn,
// 				ImportState:             true,
// 				ImportStateVerify:       true,
// 				ImportStateIdPrefix:     fmt.Sprintf("foo-%s/", randString),
// 				ImportStateVerifyIgnore: []string{"configuration.0.secret"},
// 			},
// 		},
// 	})
// }
//
// func testAccCheckGithubRepositoryWebhookExists(n string, repoName string, hook *github.Hook) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not Found: %s", n)
// 		}
//
// 		hookID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
// 		if err != nil {
// 			return unconvertibleIdErr(rs.Primary.ID, err)
// 		}
// 		if hookID == 0 {
// 			return fmt.Errorf("No repository name is set")
// 		}
//
// 		conn := testAccProvider.Meta().(*Owner).v3client
// 		owner := testAccProvider.Meta().(*Owner).name
// 		getHook, _, err := conn.Repositories.GetHook(context.TODO(), owner, repoName, hookID)
// 		if err != nil {
// 			return err
// 		}
// 		*hook = *getHook
// 		return nil
// 	}
// }
//
// type testAccGithubRepositoryWebhookExpectedAttributes struct {
// 	Events        []string
// 	Configuration map[string]interface{}
// 	Active        bool
// }
//
// func testAccCheckGithubRepositoryWebhookAttributes(hook *github.Hook, want *testAccGithubRepositoryWebhookExpectedAttributes) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
//
// 		if *hook.Active != want.Active {
// 			return fmt.Errorf("got hook %t; want %t", *hook.Active, want.Active)
// 		}
// 		if !strings.HasPrefix(*hook.URL, "https://") {
// 			return fmt.Errorf("got http URL %q; want to start with 'https://'", *hook.URL)
// 		}
// 		if !reflect.DeepEqual(hook.Events, want.Events) {
// 			return fmt.Errorf("got hook events %q; want %q", hook.Events, want.Events)
// 		}
// 		if !reflect.DeepEqual(hook.Config, want.Configuration) {
// 			return fmt.Errorf("got hook configuration %q; want %q", hook.Config, want.Configuration)
// 		}
//
// 		return nil
// 	}
// }
//
// func testAccCheckGithubRepositoryWebhookDestroy(s *terraform.State) error {
// 	conn := testAccProvider.Meta().(*Owner).v3client
// 	owner := testAccProvider.Meta().(*Owner).name
//
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "github_repository_webhook" {
// 			continue
// 		}
//
// 		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
// 		if err != nil {
// 			return unconvertibleIdErr(rs.Primary.ID, err)
// 		}
//
// 		gotHook, resp, err := conn.Repositories.GetHook(context.TODO(), owner, rs.Primary.Attributes["repository"], id)
// 		if err == nil {
// 			if gotHook != nil && gotHook.GetID() == id {
// 				return fmt.Errorf("Webhook still exists")
// 			}
// 		}
// 		if resp.StatusCode != 404 {
// 			return err
// 		}
// 		return nil
// 	}
// 	return nil
// }
//
// func testAccGithubRepositoryWebhookConfig(randString string) string {
// 	return fmt.Sprintf(`
// resource "github_repository" "foo" {
//   name         = "foo-%s"
//   description  = "Terraform acceptance tests"
//   homepage_url = "http://example.com/"
//
//   # So that acceptance tests can be run in a github organization
//   # with no billing
//   private = false
//
//   has_issues    = true
//   has_wiki      = true
//   has_downloads = true
// }
//
// resource "github_repository_webhook" "foo" {
//   depends_on = ["github_repository.foo"]
//   repository = "foo-%s"
//
//   configuration {
//     url          = "https://google.de/webhook"
//     content_type = "json"
//     insecure_ssl = true
//   }
//
//   events = ["pull_request"]
// }
// `, randString, randString)
// }
//
// func testAccGithubRepositoryWebhookConfig_secret(randString string) string {
// 	return fmt.Sprintf(`
// resource "github_repository" "foo" {
//   name         = "foo-%s"
//   description  = "Terraform acceptance tests"
//   homepage_url = "http://example.com/"
//
//   # So that acceptance tests can be run in a github organization
//   # with no billing
//   private = false
//
//   has_issues    = true
//   has_wiki      = true
//   has_downloads = true
// }
//
// resource "github_repository_webhook" "foo" {
//   repository = "${github_repository.foo.name}"
//
//   configuration {
//     url          = "https://www.terraform.io/webhook"
//     content_type = "json"
//     secret       = "RandomSecretString"
//     insecure_ssl = false
//   }
//
//   events = ["pull_request"]
// }
// `, randString)
// }
//
// func testAccGithubRepositoryWebhookUpdateConfig(randString string) string {
// 	return fmt.Sprintf(`
// resource "github_repository" "foo" {
//   name         = "foo-%s"
//   description  = "Terraform acceptance tests"
//   homepage_url = "http://example.com/"
//
//   # So that acceptance tests can be run in a github organization
//   # with no billing
//   private = false
//
//   has_issues    = true
//   has_wiki      = true
//   has_downloads = true
// }
//
// resource "github_repository_webhook" "foo" {
//   depends_on = ["github_repository.foo"]
//   repository = "foo-%s"
//
//   configuration {
//     url          = "https://google.de/webhooks"
//     content_type = "form"
//     insecure_ssl = false
//   }
//   active = false
//
//   events = ["issues"]
// }
// `, randString, randString)
// }
