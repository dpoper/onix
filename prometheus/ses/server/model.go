/*
   Onix Config Manager - SeS - Onix Webhook Receiver for AlertManager
   Copyright (c) 2020 by www.gatblau.org

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
package server

import (
	"errors"
	. "github.com/gatblau/oxc"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

type SeSModel struct {
	client *Client
}

const (
	SeSModelType                      = "SES"
	SeSServiceItemType                = "SES_SERVICE"
	SeSServiceItemTypeAttrName        = "SES_SERVICE_ATTR_NAME"
	SeSServiceItemTypeAttrStatus      = "SES_SERVICE_ATTR_STATUS"
	SeSServiceItemTypeAttrURI         = "SES_SERVICE_ATTR_URI"
	SeSServiceItemTypeAttrDescription = "SES_SERVICE_ATTR_DESCRIPTION"
)

// creates a new instance of the SeS model
func NewSeSModel(client *Client) *SeSModel {
	model := new(SeSModel)
	model.client = client
	return model
}

// checks the SeS model is defined in Onix
func (m *SeSModel) exists() (bool, error) {
	model, err := m.client.GetModel(&Model{Key: SeSModelType})
	// if we have an error
	if err != nil {
		// if the error contains 404 not found
		if strings.Contains(err.Error(), "404") {
			// return false and no error
			return false, nil
		} else {
			// there was a problem, the service might not be there
			return false, err
		}
	}
	return model != nil, err
}

// create the SeS model in Onix
func (m *SeSModel) create() error {
	var (
		exist    bool
		attempts int
		interval time.Duration = 30 // the interval to wait for reconnection
		err      error
	)
	// tries and connects to Onix using retry if the service is not there
	for {
		// check if the model exists
		exist, err = m.exists()
		if err == nil {
			// could connect to the Web API therefore breaks the retry loop
			break
		}
		// could not connect so increment the retry count
		attempts = attempts + 1
		// issue a warning to the console output
		log.Warn().Msgf("Can't connect to Onix: %s. Attempt %s, waiting before attempting to connect again.", err, strconv.Itoa(attempts))
		// wait of a second before retrying
		time.Sleep(interval * time.Second)
	}
	// if the model is not defined in Onix
	if !exist {
		// create the model
		log.Trace().Msg("The SeS model is not yet defined in Onix, proceeding to create it.")
		result, err := m.client.PutData(m.getModelData())
		if err != nil {
			log.Error().Msgf("Can't create SeS model: %s", err)
			return err
		}
		if result.Error {
			log.Error().Msgf("Can't create SeS meta-model: %s", result.Message)
			return errors.New(result.Message)
		}
	} else {
		log.Trace().Msg("SeS model found in Onix.")
	}
	return nil
}

// gets the Terra's meta model data
func (m *SeSModel) getModelData() *GraphData {
	return &GraphData{
		Models: []Model{
			Model{
				Key:         SeSModelType,
				Name:        "Service Status Model",
				Description: "Defines the Service item that capture status change information.",
				Managed:     true,
			},
		},
		ItemTypes: []ItemType{
			ItemType{
				Key:         SeSServiceItemType,
				Name:        "SeS Service",
				Description: "Defines a service that is monitored for status changes.",
				Model:       SeSModelType,
			},
		},
		ItemTypeAttributes: []ItemTypeAttribute{
			ItemTypeAttribute{
				Key:         SeSServiceItemTypeAttrName,
				Name:        "name",
				Description: "The name of the service.",
				Type:        "string",
				ItemTypeKey: SeSServiceItemType,
				Required:    true,
			},
			ItemTypeAttribute{
				Key:         SeSServiceItemTypeAttrStatus,
				Name:        "status",
				Description: "The status of the service.",
				Type:        "string",
				ItemTypeKey: SeSServiceItemType,
				Required:    true,
			},
			ItemTypeAttribute{
				Key:         SeSServiceItemTypeAttrURI,
				Name:        "uri",
				Description: "The URI of the service.",
				Type:        "string",
				ItemTypeKey: SeSServiceItemType,
				Required:    true,
			},
			ItemTypeAttribute{
				Key:         SeSServiceItemTypeAttrDescription,
				Name:        "description",
				Description: "The description of the service status.",
				Type:        "string",
				ItemTypeKey: SeSServiceItemType,
				Required:    true,
			},
		},
	}
}
