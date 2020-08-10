package github

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubReleaseDataSource(t *testing.T) {

	t.Run("errors when querying with non-existent ID", func(t *testing.T) {

		config := `
			resource "github_repository" "test" {
			  name         = "tf-acc-test-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				auto_init 	 = true
			}

			data "github_release" "test" {
				repository = github_repository.test.id
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
}

// func TestAccGithubReleaseDataSource_latestExisting(t *testing.T) {
// 	if err := testAccCheckOrganization(); err != nil {
// 		t.Skipf("Skipping because %s.", err.Error())
// 	}
//
// 	repo := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")
// 	owner := os.Getenv("GITHUB_OWNER")
// 	retrieveBy := "latest"
// 	expectedUrl := regexp.MustCompile(fmt.Sprintf("%s/%s", owner, repo))
// 	expectedTarball := regexp.MustCompile(fmt.Sprintf("%s/%s/tarball", owner, repo))
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", 0),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr("data.github_release.test", "url", expectedUrl),
// 					resource.TestMatchResourceAttr("data.github_release.test", "tarball_url", expectedTarball),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccGithubReleaseDataSource_fetchByIdExisting(t *testing.T) {
// 	if err := testAccCheckOrganization(); err != nil {
// 		t.Skipf("Skipping because %s.", err.Error())
// 	}
//
// 	repo := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")
// 	owner := os.Getenv("GITHUB_OWNER")
// 	retrieveBy := "id"
// 	expectedUrl := regexp.MustCompile(fmt.Sprintf("%s/%s", owner, repo))
// 	expectedTarball := regexp.MustCompile(fmt.Sprintf("%s/%s/tarball", owner, repo))
// 	id, _ := strconv.ParseInt(os.Getenv("GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID"), 10, 64)
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", id),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.github_release.test", "release_id", strconv.FormatInt(id, 10)),
// 					resource.TestMatchResourceAttr("data.github_release.test", "url", expectedUrl),
// 					resource.TestMatchResourceAttr("data.github_release.test", "tarball_url", expectedTarball),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccGithubReleaseDataSource_fetchByTagNoTagReturnsError(t *testing.T) {
// 	repo := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")
// 	owner := os.Getenv("GITHUB_OWNER")
// 	retrieveBy := "tag"
// 	id := int64(0)
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", id),
// 				ExpectError: regexp.MustCompile("`release_tag` must be set when `retrieve_by` = `tag`"),
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubReleaseDataSource_fetchByTagExisting(t *testing.T) {
// 	if err := testAccCheckOrganization(); err != nil {
// 		t.Skipf("Skipping because %s.", err.Error())
// 	}
//
// 	repo := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")
// 	owner := os.Getenv("GITHUB_OWNER")
// 	retrieveBy := "tag"
// 	tag := "v1.0"
// 	expectedUrl := regexp.MustCompile(fmt.Sprintf("%s/%s", owner, repo))
// 	expectedTarball := regexp.MustCompile(fmt.Sprintf("%s/%s/tarball", owner, repo))
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, tag, 0),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.github_release.test", "release_tag", tag),
// 					resource.TestMatchResourceAttr("data.github_release.test", "url", expectedUrl),
// 					resource.TestMatchResourceAttr("data.github_release.test", "tarball_url", expectedTarball),
// 				),
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubReleaseDataSource_invalidRetrieveMethodReturnsError(t *testing.T) {
// 	retrieveBy := "not valid"
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccCheckGithubReleaseDataSourceConfig("", "", retrieveBy, "", 0),
// 				ExpectError: regexp.MustCompile(`expected retrieve_by to be one of \[latest id tag]`),
// 			},
// 		},
// 	})
//
// }
//
// func testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, tag string, id int64) string {
// 	return fmt.Sprintf(`
// data "github_release" "test" {
// 	repository = "%s"
// 	owner = "%s"
// 	retrieve_by = "%s"
// 	release_tag = "%s"
// 	release_id = %d
// }
// `, repo, owner, retrieveBy, tag, id)
// }
