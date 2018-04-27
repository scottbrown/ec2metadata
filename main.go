package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/ec2metadata"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version(VERSION)
	kingpin.Parse()

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		kingpin.Errorf("Could not retrieve default AWS config.")
	}

	metadata := ec2metadata.New(cfg)

	// TODO add a timeout
	if !metadata.Available() {
		kingpin.Errorf("[ERROR] Command not valid outside EC2 instance. Please run this command within a running EC2 instance.")
	}

	doc, err := metadata.GetInstanceIdentityDocument()
	if err != nil {
		kingpin.Errorf("not available")
	}

	if *flagAmiId {
		showAmiId(doc, false)
	}

	if *flagAmiLaunchIndex {
		showAmiLaunchIndex(metadata, false)
	}

	if *flagAmiManifestPath {
		showAmiManifestPath(metadata, false)
	}

	if *flagAncestorAmiIds {
		showAncestorAmiIds(metadata, false)
	}

	if *flagBlockDeviceMapping {
		showBlockDeviceMapping(metadata, false)
	}

	if *flagInstanceId {
		showInstanceId(doc, false)
	}

	if *flagInstanceType {
		showInstanceType(doc, false)
	}

	if *flagLocalHostname {
		showLocalHostname(metadata, false)
	}

	if *flagLocalIpv4 {
		showLocalIpv4(doc, false)
	}

	if *flagKernelId {
		showKernelId(doc, false)
	}

	if *flagAvailabilityZone {
		showAvailabilityZone(doc, false)
	}

	if *flagProductCodes {
		showProductCodes(metadata, false)
	}

	if *flagPublicHostname {
		showPublicHostname(metadata, false)
	}

	if *flagPublicIpv4 {
		showPublicIpv4(metadata, false)
	}

	if *flagPublicKeys {
		showPublicKeys(metadata, false)
	}

	if *flagRamdiskId {
		showRamdiskId(doc, false)
	}

	if *flagReservationId {
		showReservationId(metadata, false)
	}

	if *flagSecurityGroups {
		showSecurityGroups(metadata, false)
	}

	if *flagUserData {
		showUserData(metadata, false)
	}

  if *flagProfile {
  }

  if *flagMac {
  }

  if *flagInstanceAction {
  }

	if *flagAll {
		showAmiId(doc, true)
		showAmiLaunchIndex(metadata, true)
		showAmiManifestPath(metadata, true)
		showAncestorAmiIds(metadata, true)
		showBlockDeviceMapping(metadata, true)
		showInstanceId(doc, true)
		showInstanceType(doc, true)
		showLocalHostname(metadata, true)
		showLocalIpv4(doc, true)
		showKernelId(doc, true)
		showAvailabilityZone(doc, true)
		showProductCodes(metadata, true)
		showPublicHostname(metadata, true)
		showPublicIpv4(metadata, true)
		showPublicKeys(metadata, true)
		showRamdiskId(doc, true)
		showReservationId(metadata, true)
		showSecurityGroups(metadata, true)
		showUserData(metadata, true)
	}
}
