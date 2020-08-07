package github

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"github": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"github": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}

func TestProvider(t *testing.T) {

	t.Run("runs internal validation without error", func(t *testing.T) {

		if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
			t.Fatalf("err: %s", err)
		}

	})

	t.Run("has an implementation", func(t *testing.T) {
		// FIXME: unsure if this is useful; refactored from:
		// func TestProvider_impl(t *testing.T) {
		// 	var _ terraform.ResourceProvider = Provider()
		// }

		var _ terraform.ResourceProvider = Provider()
	})

}

func TestAccProviderConfigure(t *testing.T) {

	t.Run("can be configured to run insecurely", func(t *testing.T) {

		// Use ephemeral port range (49152–65535)
		port := fmt.Sprintf("%d", 49152+rand.Intn(16382))

		// Use self-signed certificate
		certFile := filepath.Join("test-fixtures", "cert.pem")
		keyFile := filepath.Join("test-fixtures", "key.pem")

		url, closeFunc := githubTLSApiMock(port, certFile, keyFile, t)
		defer func() {
			err := closeFunc()
			if err != nil {
				t.Fatal(err)
			}
		}()

		oldBaseUrl := os.Getenv("GITHUB_BASE_URL")
		defer os.Setenv("GITHUB_BASE_URL", oldBaseUrl)

		// Point provider to mock API with self-signed cert
		os.Setenv("GITHUB_BASE_URL", url)

		config := fmt.Sprintf(`
			data "github_user" "test" {
				username = "%s"
			}
		`, "hashibot")

		testCase := func(mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("x509: certificate is valid for untrusted, not localhost"),
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

	t.Run("can be configured to run anonymously", func(t *testing.T) {

		anonymousConfiguration := `
			provider "github" {}
			data "github_ip_ranges" "test" {}
		`
		anonymousCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "hooks.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "git.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "pages.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "importer.#"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, "anonymous") },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: anonymousConfiguration,
					Check:  anonymousCheck,
				},
			},
		})
	})

	t.Run("can be configured with an individual account", func(t *testing.T) {

		individualConfiguration := fmt.Sprintf(`
			provider "github" {
				token = "%s"
			}
			data "github_user" "test" { username = "%s" }
		`,
			os.Getenv("GITHUB_TOKEN"),
			os.Getenv("GITHUB_OWNER"),
		)

		individualCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_user.test", "username"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessMode(t, "individual")
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: individualConfiguration,
					Check:  individualCheck,
				},
			},
		})

	})

	t.Run("can be configured with an organization account", func(t *testing.T) {

		organizationConfiguration := fmt.Sprintf(`
			provider "github" {
				organization = "%[1]s"
				token = "%[2]s"
			}
			data "github_organization" "test" { name = "%[1]s" }
		`,
			os.Getenv("GITHUB_ORGANIZATION"),
			os.Getenv("GITHUB_TOKEN"),
		)

		organizationCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessMode(t, "organization")
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: organizationConfiguration,
					Check:  organizationCheck,
				},
			},
		})
	})

	t.Run("can be configured with a GHES deployment", func(t *testing.T) {

		organizationConfiguration := fmt.Sprintf(`
				provider "github" {
					base_url = "%s"
					token = "%s"
				}
				data "github_organization" "test" { name = "%s" }
			`,
			os.Getenv("GHES_BASE_URL"),
			os.Getenv("GHES_TOKEN"),
			os.Getenv("GHES_ORGANIZATION"),
		)

		organizationCheck := resource.ComposeTestCheckFunc(
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
			testCase("anonymous")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase("individual")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase("organization")
		})

	})
}
