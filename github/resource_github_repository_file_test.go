package github

import (
	"strings"

	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccGithubRepositoryFile(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates file content without error", func(t *testing.T) {

		fileContent := "file_content_value"
		updatedFileContent := "updated_file_content_value"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_file" "test" {
			  repository = github_repository.test.id
			  file       = "test"
			  content    = "%s"
			}
		`, randomID, fileContent)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_file.test", "sha",
					"deee258b7c807901aad79d01da020d993739160a",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_file.test", "sha",
					"ec9aad0ba478cdd7349faabbeac2a64e5ce72ddb",
				),
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
							fileContent,
							updatedFileContent, 1),
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

	t.Run("manages file content for a specified branch", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_branch" "test" {
			  repository = github_repository.test.id
			  branch     = "tf-acc-test-%[1]s"
			}

			resource "github_repository_file" "test" {
			  repository = github_repository.test.id
			  branch     = github_branch.test.branch
			  file       = "test"
			  content    = "test"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "sha",
				"30d74d258442c7c65512eafab474568dd706c430",
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

	t.Run("commits with custom message, author and e-mail", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_file" "test" {
			  repository = github_repository.test.id
			  file       = "test"
			  content    = "test"
			  commit_message = "Managed by Terraform"
			  commit_author  = "Terraform User"
			  commit_email   = "terraform@example.com"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "commit_message",
				"Managed by Terraform",
			),
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "commit_author",
				"Terraform User",
			),
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "commit_email",
				"terraform@example.com",
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
}

// // The authenticated user's name used for commits should be exported as GITHUB_TEST_USER_NAME
// var userName string = os.Getenv("GITHUB_TEST_USER_NAME")
//
// // The authenticated user's email address used for commits should be exported as GITHUB_TEST_USER_EMAIL
// var userEmail string = os.Getenv("GITHUB_TEST_USER_EMAIL")
//
// func init() {
// 	resource.AddTestSweepers("github_repository_file", &resource.Sweeper{
// 		Name: "github_repository_file",
// 		F:    testSweepRepositoryFiles,
// 	})
//
// }
//
// func testSweepRepositoryFiles(region string) error {
// 	meta, err := sharedConfigForRegion(region)
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := testSweepDeleteRepositoryFiles(meta, "master"); err != nil {
// 		return err
// 	}
//
// 	if err := testSweepDeleteRepositoryFiles(meta, "test-branch"); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func testSweepDeleteRepositoryFiles(meta interface{}, branch string) error {
// 	client := meta.(*Owner).v3client
// 	owner := meta.(*Owner).name
//
// 	_, files, _, err := client.Repositories.GetContents(
// 		context.TODO(), owner, "test-repo", "", &github.RepositoryContentGetOptions{Ref: branch})
// 	if err != nil {
// 		return err
// 	}
//
// 	for _, f := range files {
// 		if name := f.GetName(); strings.HasPrefix(name, "tf-acc-") {
// 			log.Printf("Deleting repository file: %s, repo: %s/test-repo, branch: %s", name, owner, branch)
// 			opts := &github.RepositoryContentFileOptions{Branch: github.String(branch)}
// 			if _, _, err := client.Repositories.DeleteFile(context.TODO(), owner, "test-repo", name, opts); err != nil {
// 				return err
// 			}
// 		}
// 	}
//
// 	return nil
// }
//
// func TestAccGithubRepositoryFile_basic(t *testing.T) {
// 	if userName == "" {
// 		t.Skip("This test requires you to set the test user's name (set it by exporting GITHUB_TEST_USER_NAME)")
// 	}
//
// 	if userEmail == "" {
// 		t.Skip("This test requires you to set the test user's email address (set it by exporting GITHUB_TEST_USER_EMAIL)")
// 	}
//
// 	var content github.RepositoryContent
// 	var commit github.RepositoryCommit
//
// 	rn := "github_repository_file.foo"
// 	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	path := fmt.Sprintf("tf-acc-test-file-%s", randString)
//
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccGithubRepositoryFileConfig(
// 					path, "Terraform acceptance test file"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
// 					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
// 						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file")) + "\n",
// 					}),
// 					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
// 						Branch:        "master",
// 						CommitAuthor:  userName,
// 						CommitEmail:   userEmail,
// 						CommitMessage: fmt.Sprintf("Add %s", path),
// 						Filename:      path,
// 					}),
// 					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
// 					resource.TestCheckResourceAttr(rn, "branch", "master"),
// 					resource.TestCheckResourceAttr(rn, "file", path),
// 					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file"),
// 					resource.TestCheckResourceAttr(rn, "commit_author", userName),
// 					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
// 					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Add %s", path)),
// 				),
// 			},
// 			{
// 				Config: testAccGithubRepositoryFileConfig(
// 					path, "Terraform acceptance test file updated"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
// 					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
// 						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file updated")) + "\n",
// 					}),
// 					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
// 						Branch:        "master",
// 						CommitAuthor:  userName,
// 						CommitEmail:   userEmail,
// 						CommitMessage: fmt.Sprintf("Update %s", path),
// 						Filename:      path,
// 					}),
// 					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
// 					resource.TestCheckResourceAttr(rn, "branch", "master"),
// 					resource.TestCheckResourceAttr(rn, "file", path),
// 					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file updated"),
// 					resource.TestCheckResourceAttr(rn, "commit_author", userName),
// 					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
// 					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Update %s", path)),
// 				),
// 			},
// 			{
// 				ResourceName:      rn,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubRepositoryFile_branch(t *testing.T) {
// 	if userName == "" {
// 		t.Skip("This test requires you to set the test user's name (set it by exporting GITHUB_TEST_USER_NAME)")
// 	}
//
// 	if userEmail == "" {
// 		t.Skip("This test requires you to set the test user's email address (set it by exporting GITHUB_TEST_USER_EMAIL)")
// 	}
//
// 	var content github.RepositoryContent
// 	var commit github.RepositoryCommit
//
// 	rn := "github_repository_file.foo"
// 	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	path := fmt.Sprintf("tf-acc-test-file-%s", randString)
//
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccGithubRepositoryFileBranchConfig(
// 					path, "Terraform acceptance test file"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryFileExists(rn, path, "test-branch", &content, &commit),
// 					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
// 						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file")) + "\n",
// 					}),
// 					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
// 						Branch:        "test-branch",
// 						CommitAuthor:  userName,
// 						CommitEmail:   userEmail,
// 						CommitMessage: fmt.Sprintf("Add %s", path),
// 						Filename:      path,
// 					}),
// 					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
// 					resource.TestCheckResourceAttr(rn, "branch", "test-branch"),
// 					resource.TestCheckResourceAttr(rn, "file", path),
// 					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file"),
// 					resource.TestCheckResourceAttr(rn, "commit_author", userName),
// 					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
// 					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Add %s", path)),
// 				),
// 			},
// 			{
// 				Config: testAccGithubRepositoryFileBranchConfig(
// 					path, "Terraform acceptance test file updated"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryFileExists(rn, path, "test-branch", &content, &commit),
// 					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
// 						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file updated")) + "\n",
// 					}),
// 					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
// 						Branch:        "test-branch",
// 						CommitAuthor:  userName,
// 						CommitEmail:   userEmail,
// 						CommitMessage: fmt.Sprintf("Update %s", path),
// 						Filename:      path,
// 					}),
// 					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
// 					resource.TestCheckResourceAttr(rn, "branch", "test-branch"),
// 					resource.TestCheckResourceAttr(rn, "file", path),
// 					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file updated"),
// 					resource.TestCheckResourceAttr(rn, "commit_author", userName),
// 					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
// 					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Update %s", path)),
// 				),
// 			},
// 			{
// 				ResourceName:      rn,
// 				ImportState:       true,
// 				ImportStateId:     fmt.Sprintf("test-repo/%s:test-branch", path),
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }
//
// func TestAccGithubRepositoryFile_committer(t *testing.T) {
// 	var content github.RepositoryContent
// 	var commit github.RepositoryCommit
//
// 	rn := "github_repository_file.foo"
// 	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	path := fmt.Sprintf("tf-acc-test-file-%s", randString)
//
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccGithubRepositoryFileCommitterConfig(
// 					path, "Terraform acceptance test file"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
// 					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
// 						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file")) + "\n",
// 					}),
// 					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
// 						Branch:        "master",
// 						CommitAuthor:  "Terraform User",
// 						CommitEmail:   "terraform@example.com",
// 						CommitMessage: "Managed by Terraform",
// 						Filename:      path,
// 					}),
// 					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
// 					resource.TestCheckResourceAttr(rn, "branch", "master"),
// 					resource.TestCheckResourceAttr(rn, "file", path),
// 					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file"),
// 					resource.TestCheckResourceAttr(rn, "commit_author", "Terraform User"),
// 					resource.TestCheckResourceAttr(rn, "commit_email", "terraform@example.com"),
// 					resource.TestCheckResourceAttr(rn, "commit_message", "Managed by Terraform"),
// 				),
// 			},
// 			{
// 				Config: testAccGithubRepositoryFileCommitterConfig(
// 					path, "Terraform acceptance test file updated"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
// 					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
// 						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file updated")) + "\n",
// 					}),
// 					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
// 						Branch:        "master",
// 						CommitAuthor:  "Terraform User",
// 						CommitEmail:   "terraform@example.com",
// 						CommitMessage: "Managed by Terraform",
// 						Filename:      path,
// 					}),
// 					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
// 					resource.TestCheckResourceAttr(rn, "branch", "master"),
// 					resource.TestCheckResourceAttr(rn, "file", path),
// 					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file updated"),
// 					resource.TestCheckResourceAttr(rn, "commit_author", "Terraform User"),
// 					resource.TestCheckResourceAttr(rn, "commit_email", "terraform@example.com"),
// 					resource.TestCheckResourceAttr(rn, "commit_message", "Managed by Terraform"),
// 				),
// 			},
// 			{
// 				ResourceName:      rn,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }
//
// func testAccCheckGithubRepositoryFileExists(n, path, branch string, content *github.RepositoryContent, commit *github.RepositoryCommit) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
//
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not Found: %s", n)
// 		}
//
// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("No repository file path set")
// 		}
//
// 		conn := testAccProvider.Meta().(*Owner).v3client
// 		owner := testAccProvider.Meta().(*Owner).name
//
// 		opts := &github.RepositoryContentGetOptions{Ref: branch}
// 		gotContent, _, _, err := conn.Repositories.GetContents(context.TODO(), owner, "test-repo", path, opts)
// 		if err != nil {
// 			return err
// 		}
//
// 		gotCommit, err := getFileCommit(conn, owner, "test-repo", path, branch)
// 		if err != nil {
// 			return err
// 		}
//
// 		*content = *gotContent
// 		*commit = *gotCommit
//
// 		return nil
// 	}
// }
//
// type testAccGithubRepositoryFileExpectedAttributes struct {
// 	Content string
// }
//
// func testAccCheckGithubRepositoryFileAttributes(content *github.RepositoryContent, want *testAccGithubRepositoryFileExpectedAttributes) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
//
// 		if *content.Content != want.Content {
// 			return fmt.Errorf("got content %q; want %q", *content.Content, want.Content)
// 		}
//
// 		return nil
// 	}
// }
//
// type testAccGithubRepositoryFileExpectedCommitAttributes struct {
// 	Branch        string
// 	CommitAuthor  string
// 	CommitEmail   string
// 	CommitMessage string
// 	Filename      string
// }
//
// func testAccCheckGithubRepositoryFileCommitAttributes(commit *github.RepositoryCommit, want *testAccGithubRepositoryFileExpectedCommitAttributes) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
//
// 		if name := commit.GetCommit().GetCommitter().GetName(); name != want.CommitAuthor {
// 			return fmt.Errorf("got committer author name %q; want %q", name, want.CommitAuthor)
// 		}
//
// 		if email := commit.GetCommit().GetCommitter().GetEmail(); email != want.CommitEmail {
// 			return fmt.Errorf("got committer author email %q; want %q", email, want.CommitEmail)
// 		}
//
// 		if message := commit.GetCommit().GetMessage(); message != want.CommitMessage {
// 			return fmt.Errorf("got commit message %q; want %q", message, want.CommitMessage)
// 		}
//
// 		if len(commit.Files) != 1 {
// 			return fmt.Errorf("got multiple files in commit (%q); expected 1", len(commit.Files))
// 		}
//
// 		file := commit.Files[0]
// 		if filename := file.GetFilename(); filename != want.Filename {
// 			return fmt.Errorf("got filename %q; want %q", filename, want.Filename)
// 		}
//
// 		return nil
// 	}
// }
//
// func testAccCheckGithubRepositoryFileDestroy(s *terraform.State) error {
// 	conn := testAccProvider.Meta().(*Owner).v3client
// 	owner := testAccProvider.Meta().(*Owner).name
//
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "github_repository_file" {
// 			continue
// 		}
//
// 		repo, file := splitRepoFilePath(rs.Primary.ID)
// 		opts := &github.RepositoryContentGetOptions{Ref: rs.Primary.Attributes["branch"]}
//
// 		fc, _, resp, err := conn.Repositories.GetContents(context.TODO(), owner, repo, file, opts)
// 		if err == nil {
// 			if fc != nil {
// 				return fmt.Errorf("Repository file %s/%s/%s still exists", owner, repo, file)
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
// func testAccGithubRepositoryFileConfig(file, content string) string {
// 	return fmt.Sprintf(`
// resource "github_repository_file" "foo" {
//   repository = "test-repo"
//   file       = "%s"
//   content    = "%s"
// }
// `, file, content)
// }
//
// func testAccGithubRepositoryFileBranchConfig(file, content string) string {
// 	return fmt.Sprintf(`
// resource "github_repository_file" "foo" {
//   repository = "test-repo"
//   branch     = "test-branch"
//   file       = "%s"
//   content    = "%s"
// }
// `, file, content)
// }
//
// func testAccGithubRepositoryFileCommitterConfig(file, content string) string {
// 	return fmt.Sprintf(`
// resource "github_repository_file" "foo" {
//   repository     = "test-repo"
//   file           = "%s"
//   content        = "%s"
//   commit_message = "Managed by Terraform"
//   commit_author  = "Terraform User"
//   commit_email   = "terraform@example.com"
// }
// `, file, content)
// }
