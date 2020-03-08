/*
   Onix Config Manager - Terraform Provider
   Copyright (c) 2018-2019 by www.gatblau.org

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
	"github.com/hashicorp/terraform/helper/schema"
)

func LinkResource() *schema.Resource {
	return &schema.Resource{
		Create: createLink,
		Read:   readLink,
		Update: updateLink,
		Delete: deleteLink,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"meta": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"start_item_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end_item_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func LinkDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLink,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"meta": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"start_item_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end_item_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createLink(data *schema.ResourceData, meta interface{}) error {
	// read the resource data into a Link
	link := newLink(data)

	// put the Link to the Web API
	err := link.put(meta)

	// set Link Id key
	data.SetId(link.Key)

	return err
}

func readLink(data *schema.ResourceData, meta interface{}) error {
	// read the resource data into a link
	link := newLink(data)

	// get the resource
	link, err := link.get(meta)

	return err
}

func updateLink(data *schema.ResourceData, meta interface{}) error {
	return createLink(data, meta)
}

func deleteLink(data *schema.ResourceData, meta interface{}) error {
	// read the resource data into an Link
	link := newLink(data)

	// delete the link
	return link.delete(meta)
}