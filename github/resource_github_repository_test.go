package github

import (
	"testing"
)

func TestAccGithubRepositories(t *testing.T) {

	t.Run("creates and updates repositories without error", func(t *testing.T) {

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

	t.Run("imports repositories without error", func(t *testing.T) {

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

	t.Run("archives repositories without error", func(t *testing.T) {

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

	t.Run("manages the project feature for a repository", func(t *testing.T) {

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

	t.Run("manages the default branch feature for a repository", func(t *testing.T) {

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
