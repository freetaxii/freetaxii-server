// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package config

import (
	"encoding/json"
	"os"

	"github.com/freetaxii/libstix2/resources"
	"github.com/gologme/log"
	"github.com/gorilla/mux"
)

/*
ServerConfigType - This type defines the configuration for the entire server.
*/
type ServerConfigType struct {
	Router *mux.Router
	Global struct {
		Prefix             string
		Listen             string
		Protocol           string
		TLSKey             string
		TLSCrt             string
		DbConfig           bool
		DbType             string
		DbFile             string
		HTMLTemplateDir    string
		MaxNumberOfObjects int
	}
	Logging struct {
		Enabled bool
		LogFile string
	}
	DiscoveryServer struct {
		Enabled      bool
		HTMLEnabled  bool
		HTMLBranding HTMLTemplateFilesType
		Services     []DiscoveryServiceType
	}
	APIRootServer struct {
		Enabled      bool
		HTMLEnabled  bool
		HTMLBranding HTMLTemplateFilesType
		Services     []APIRootServiceType
	}
	DiscoveryResources  map[string]resources.DiscoveryType
	APIRootResources    map[string]resources.APIRootType
	CollectionResources map[string]resources.CollectionType
}

/*
DiscoveryServiceType - This struct represents an instance of a Discovery server.
If someone tries to set the 'resourcepath' directive in the configuration file it
will get overwritten in code.
*/
type DiscoveryServiceType struct {
	Enabled          bool                  // User defined in configuration file
	Name             string                // User defined in configuration file
	ResourcePath     string                // Set in verifyDiscoveryConfig()
	HTMLEnabled      bool                  // Set in verifyDiscoveryHTMLConfig()
	HTMLBranding     HTMLTemplateFilesType // User defined in configuration file or set to DiscoveryServer value in verifyDiscoveryHTMLConfig()
	HTMLTemplatePath string                // Set in verifyDiscoveryHTMLConfig() = Prefix + HTMLTemplateDir
	ResourceID       string                // User defined in configuration file
}

/*
APIRootServiceType - This struct represents an instance of an API Root server.
If someone tries to set the 'path' directive in the configuration file it
will just get overwritten in code.
*/
type APIRootServiceType struct {
	Enabled          bool                  // User defined in configuration file
	Name             string                // User defined in configuration file
	ResourcePath     string                // Set in verifyAPIRootConfig()
	HTMLEnabled      bool                  // Set in verifyAPIRootHTMLConfig()
	HTMLBranding     HTMLTemplateFilesType // User defined in configuration file or set to APIRootServer value in verifyAPIRootHTMLConfig()
	HTMLTemplatePath string                // Set in verifyAPIRootHTMLConfig() = Prefix + HTMLTemplateDir
	ResourceID       string                // User defined in configuration file
	Collections      struct {
		Enabled      bool
		ResourcePath string // Set in verifyAPIRootConfig()
		Members      []string
	}
}

/*
HTMLTemplateFilesType - This struct holds the individual template filenames for
each type of resource.
*/
type HTMLTemplateFilesType struct {
	Discovery   string
	APIRoot     string
	Collections string
	Collection  string
	Objects     string
	Manifest    string
}

// --------------------------------------------------
//
// Load Configuration File and Parse JSON
//
// --------------------------------------------------

// LoadServerConfig - This methods takes in a string value representing a
// filename of the configuration file and loads the configuration into memory.
func (config *ServerConfigType) LoadServerConfig(filename string) {
	// TODO - Need to make make a validation check for the configuration file

	// Open and read configuration file
	sysConfigFileData, err := os.Open(filename)
	defer sysConfigFileData.Close()
	if err != nil {
		log.Fatalf("ERROR: A problem occurred opening the configuration file: %v", err)
	}

	// --------------------------------------------------
	// Decode JSON configuration file
	// --------------------------------------------------
	// Use decoder instead of unmarshal so we can handle stream data
	decoder := json.NewDecoder(sysConfigFileData)
	err = decoder.Decode(config)

	if err != nil {
		log.Fatalf("ERROR: A problem occurred parsing the configuration file %v", err)
	}

	log.Debugln("DEBUG LoadServerConfig(): System Configuration Dump")
	log.Debugf("%+v\n", config)
}
