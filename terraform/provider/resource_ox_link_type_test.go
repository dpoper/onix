/*
   Onix Config Manager - Terraform Provider
   Copyright (c) 2018-2020 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/
package main

import (
	"fmt"
	"github.com/gatblau/oxc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strconv"
	"testing"
)

// all the constants for the item test
const (
	LinkTypeResourceName  = "ox_link_type.ox_link_type_1"
	LinkTypeRsModelKey    = "test_acc_ox_link_type_model_1"
	LinkTypeRsModelName   = "ox_link_type_model_1"
	LinkTypeRsModelDesc   = "ox_link_type_model_1 Description"
	LinkTypeRsKey         = "test_acc_ox_link_type_1"
	LinkTypeRsName        = "ox_link_type_1_name"
	LinkTypeRsDesc        = "ox_link_type_1_description"
	LinkTypeRsEncryptMeta = false
	LinkTypeRsEncryptTxt  = false
	LinkTypeRsManaged     = true
)

func TestLinkTypeResource(t *testing.T) {
	resourceName := LinkTypeResourceName
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { prepareLinkTypeResourceTest(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			//create
			{
				Config: oxLinkTypeResource(),
				Check: resource.ComposeTestCheckFunc(
					// check link type resource attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "key", LinkTypeRsKey),
					resource.TestCheckResourceAttr(resourceName, "name", LinkTypeRsName),
					resource.TestCheckResourceAttr(resourceName, "description", LinkTypeRsDesc),
					resource.TestCheckResourceAttr(resourceName, "model_key", LinkTypeRsModelKey),
					resource.TestCheckResourceAttr(resourceName, "encrypt_meta", strconv.FormatBool(LinkTypeRsEncryptMeta)),
					resource.TestCheckResourceAttr(resourceName, "encrypt_txt", strconv.FormatBool(LinkTypeRsEncryptTxt)),

					// check for side effects in Onix database
					checkLinkTypeResourceCreated(LinkTypeResourceName),
				),
			},
		},
		CheckDestroy: checkLinkTypeResourceDestroyed,
	})
}

// create supporting database entities
func prepareLinkTypeResourceTest(t *testing.T) {
	// need a model first
	result, err := cfg.Client.PutModel(
		&oxc.Model{
			Key:         LinkTypeRsModelKey,
			Name:        LinkTypeRsModelName,
			Description: LinkTypeRsModelDesc,
		})
	if err != nil {
		t.Error(err)
	}
	if result.Error {
		t.Error(result.Message)
	}
}

// check the item has been created in the database
func checkLinkTypeResourceCreated(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		// retrieve the item from the database
		key := rs.Primary.Attributes["key"]
		linkType, err := cfg.Client.GetLinkType(&oxc.LinkType{Key: key})
		if err != nil {
			return fmt.Errorf("can't read link type: %s", key)
		}

		// compares the state with the database values
		if err := checkEntityAttr(rs, "name", linkType.Name); err != nil {
			return err
		}
		if err := checkEntityAttr(rs, "description", linkType.Description); err != nil {
			return err
		}
		if err := checkEntityAttr(rs, "model_key", linkType.Model); err != nil {
			return err
		}
		if err := checkEntityAttr(rs, "encrypt_meta", strconv.FormatBool(linkType.EncryptMeta)); err != nil {
			return err
		}
		if err := checkEntityAttr(rs, "encrypt_txt", strconv.FormatBool(linkType.EncryptTxt)); err != nil {
			return err
		}
		return nil
	}
}

// check the item has been destroyed destroyed in the database
func checkLinkTypeResourceDestroyed(s *terraform.State) error {
	itemType, _ := cfg.Client.GetLinkType(&oxc.LinkType{Key: LinkTypeRsKey})
	// should get an error as item should not exist anymore
	if itemType != nil {
		return fmt.Errorf("item type %s still exists", LinkRsKey)
	}
	// clean up the database after the test
	return cleanUpItemTypeRs()
}

// remove supporting database entities
func cleanUpLinkTypeRs() error {
	_, err := cfg.Client.DeleteLinkType(&oxc.LinkType{Key: LinkTypeRsKey})
	_, err = cfg.Client.DeleteModel(&oxc.Model{Key: LinkTypeRsModelKey})
	return err
}

func oxLinkTypeResource() string {
	return fmt.Sprintf(
		`resource "ox_link_type" "ox_link_type_1" {
  key         	= "%s"
  name        	= "%s"
  description 	= "%s"
  model_key   	= "%s"
  encrypt_txt 	= %s
  encrypt_meta 	= "%s"
}`, LinkTypeRsKey,
		LinkTypeRsName,
		LinkTypeRsDesc,
		LinkTypeRsModelKey,
		strconv.FormatBool(LinkTypeRsEncryptTxt),
		strconv.FormatBool(LinkTypeRsEncryptMeta))
}
