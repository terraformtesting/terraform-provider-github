package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryDataSource(t *testing.T) {

	t.Run("queries a repository without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_repositories" "test" {
				query = "org:%s"
			}
			data "github_repository" "test" {
				full_name = data.github_repositories.test.full_names.0
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^`+testOrganization)),
			resource.TestMatchResourceAttr("data.github_repository.test", "full_name", regexp.MustCompile(`^`+testOrganization)),
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

	t.Run("raises expected errors when querying for a repository", func(t *testing.T) {

		// config := fmt.Sprintf(`
		// 	data "github_repositories" "test" {
		// 		query = "org:%s"
		// 	}
		// `, testOrganization)
		//
		// check := resource.ComposeTestCheckFunc(
		// 	resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^`+testOrganization)),
		// 	resource.TestCheckResourceAttrSet("data.github_repositories.test", "names.0"),
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

// func TestAccGithubRepositoryDataSource_fullName_noMatchReturnsError(t *testing.T) {
// 	fullName := "klsafj_23434_doesnt_exist/not-exists"
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccCheckGithubRepositoryDataSourceConfig_fullName(fullName),
// 				ExpectError: regexp.MustCompile(`Not Found`),
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubRepositoryDataSource_name_noMatchReturnsError(t *testing.T) {
// 	name := "not-exists"
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccCheckGithubRepositoryDataSourceConfig_name(name),
// 				ExpectError: regexp.MustCompile(`Not Found`),
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubRepositoryDataSource_fullName_existing(t *testing.T) {
// 	fullName := testOwner + "/test-repo"
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckGithubRepositoryDataSourceConfig_fullName(fullName),
// 				Check:  testRepoCheck(),
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubRepositoryDataSource_name_existing(t *testing.T) {
// 	name := "test-repo"
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			testAccPreCheck(t)
// 		},
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckGithubRepositoryDataSourceConfig_name(name),
// 				Check:  testRepoCheck(),
// 			},
// 		},
// 	})
// }
//
// func testRepoCheck() resource.TestCheckFunc {
// 	return resource.ComposeAggregateTestCheckFunc(
// 		resource.TestCheckResourceAttr("data.github_repository.test", "id", "test-repo"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "name", "test-repo"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "private", "false"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "visibility", "public"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "description", "Test description, used in GitHub Terraform provider acceptance test."),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "homepage_url", "http://www.example.com"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "has_issues", "true"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "has_wiki", "true"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "allow_merge_commit", "true"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "allow_squash_merge", "true"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "allow_rebase_merge", "true"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "has_downloads", "true"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "full_name", testOwner+"/test-repo"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "default_branch", "master"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "html_url", "https://github.com/"+testOwner+"/test-repo"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "ssh_clone_url", "git@github.com:"+testOwner+"/test-repo.git"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "svn_url", "https://github.com/"+testOwner+"/test-repo"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "git_clone_url", "git://github.com/"+testOwner+"/test-repo.git"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "http_clone_url", "https://github.com/"+testOwner+"/test-repo.git"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "archived", "false"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "topics.#", "2"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "topics.0", "second-test-topic"),
// 		resource.TestCheckResourceAttr("data.github_repository.test", "topics.1", "test-topic"),
// 	)
// }
//
// func testAccCheckGithubRepositoryDataSourceConfig_fullName(fullName string) string {
// 	return fmt.Sprintf(`
// data "github_repository" "test" {
//   full_name = "%s"
// }
// `, fullName)
// }
//
// func testAccCheckGithubRepositoryDataSourceConfig_name(name string) string {
// 	return fmt.Sprintf(`
// data "github_repository" "test" {
//   name = "%s"
// }
// `, name)
// }
