// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/datastore/sqlite3"
	"github.com/gologme/log"
	"github.com/pborman/getopt"
)

// These global variables hold build information. The Build variable will be
// populated by the Makefile and uses the Git Head hash as its identifier.
// These variables are used in the console output for --version and --help.
var (
	Version = "0.0.1"
	Build   string
)

// These global variables are for dealing with command line options
var (
	defaultDatabaseFilename = "freetaxii.db"
	sOptDatabaseFilename    = getopt.StringLong("filename", 'f', defaultDatabaseFilename, "Database Filename", "string")
	sOptSTIXID              = getopt.StringLong("stixid", 's', "", "Object ID", "string")
	sOptVersion             = getopt.StringLong("stixversion", 'v', "", "Version", "string")
	bOptHelp                = getopt.BoolLong("help", 0, "Help")
	bOptVer                 = getopt.BoolLong("version", 0, "Version")
)

func main() {
	processCommandLineFlags()
	var stixid, version string

	stixid = *sOptSTIXID
	version = *sOptVersion

	var ds datastore.Datastorer

	ds = sqlite3.New(*sOptDatabaseFilename)
	o, err := ds.GetSTIXObject(stixid, version)
	if err != nil {
		log.Fatalln(err)
	}

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	fmt.Println(string(data))
}

// --------------------------------------------------
// Private functions
// --------------------------------------------------

// processCommandLineFlags - This function will process the command line flags
// and will print the version or help information as needed.
func processCommandLineFlags() {
	getopt.HelpColumn = 35
	getopt.DisplayWidth = 120
	getopt.SetParameters("")
	getopt.Parse()

	// Lets check to see if the version command line flag was given. If it is
	// lets print out the version infomration and exit.
	if *bOptVer {
		printOutputHeader()
		os.Exit(0)
	}

	// Lets check to see if the help command line flag was given. If it is lets
	// print out the help information and exit.
	if *bOptHelp {
		printOutputHeader()
		getopt.Usage()
		os.Exit(0)
	}

	if *sOptSTIXID == "" {
		log.Fatalln("STIX ID must not be empty")
	}

	if *sOptVersion == "" {
		log.Fatalln("STIX Version must not be empty")
	}

}

// printOutputHeader - This function will print a header for all console output
func printOutputHeader() {
	fmt.Println("")
	fmt.Println("FreeTAXII - TAXII Get Collections")
	fmt.Println("Copyright: Bret Jordan")
	fmt.Println("Version:", Version)
	if Build != "" {
		fmt.Println("Build:", Build)
	}
	fmt.Println("")
}