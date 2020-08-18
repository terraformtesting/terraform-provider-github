package github

import (
	"context"
	"crypto/rand"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"golang.org/x/crypto/ed25519"
)

func TestAccGithubUserSshKey(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and destroys a user SSH key without error", func(t *testing.T) {

		publicKey, _, err := ed25519.GenerateKey(rand.Reader)
		testKey := fmt.Sprintf("%v", publicKey)

		config := fmt.Sprintf(`
			resource "github_user_ssh_key" "test" {
				title = "tf-acc-test-%s"
				key   = "%s"
			}
		`, randomID, testKey)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_user_ssh_key.test", "title",
				regexp.MustCompile(randomID),
			),
			resource.TestMatchResourceAttr(
				"github_user_ssh_key.test", "key",
				regexp.MustCompile("^ecdsa-sha2-nistp384 "),
			),
			resource.TestMatchResourceAttr(
				"github_user_ssh_key.test", "url",
				regexp.MustCompile("^https://api.github.com/[a-z0-9]+/keys/"),
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:     func() { skipUnlessMode(t, mode) },
				Providers:    testAccProviders,
				CheckDestroy: testAccCheckGithubUserSshKeyDestroy,
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

	// t.Run("imports an individual account SSH key without error", func(t *testing.T) {
	//
	// 	title := fmt.Sprintf("tf-acc-test-%s",
	// 		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	// 	config := fmt.Sprintf(`
	// 		resource "github_user_ssh_key" "test" {
	// 			title = "%s"
	// 			key   = "%s""
	// 		}
	// 	`, title, testKey)
	//
	// 	check := resource.ComposeTestCheckFunc(
	// 		resource.TestCheckResourceAttrSet("github_user_ssh_key.test", "title"),
	// 		resource.TestCheckResourceAttrSet("github_user_ssh_key.test", "key"),
	// 		resource.TestCheckResourceAttrSet("github_user_ssh_key.test", "url"),
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
	// 				{
	// 					ResourceName:      "github_user_ssh_key.test",
	// 					ImportState:       true,
	// 					ImportStateVerify: true,
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

func testAccCheckGithubUserSshKeyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Owner).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_user_ssh_key" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		_, resp, err := conn.Users.GetKey(context.TODO(), id)
		if err == nil {
			return fmt.Errorf("SSH key %s still exists", rs.Primary.ID)
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

// const testKey = "ecdsa-sha2-nistp384 AAAAE2VjZHNhLXNoYTItbmlzdHAzODQAAAAIbmlzdHAzODQAAABhBM3cbPV+J02cSXUJ5pfUfQ839WfYbhmM44J8xCslmZeyGVvql+wdfVoKCToh4N6zokCVkBDgnPL2oWnuyqYL7W2vOUiZLt5USunQ/Ywg7ZVkT1ULiGslF2P72AZVrkoq9Q=="
