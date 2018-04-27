package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/ec2metadata"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
  writeHeader bool = false
  svc *ec2metadata.EC2Metadata
  doc ec2metadata.EC2InstanceIdentityDocument
)


func main() {
	kingpin.Version(VERSION)
	kingpin.Parse()

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		kingpin.Errorf("Could not retrieve default AWS config.")
	}

	svc = ec2metadata.New(cfg)

	// TODO add a timeout
	if !svc.Available() {
		kingpin.Errorf("[ERROR] Command not valid outside EC2 instance. Please run this command within a running EC2 instance.")
	}

	doc, err = svc.GetInstanceIdentityDocument()
	if err != nil {
		kingpin.Errorf("not available")
	}

	if *flagAmiId {
		showAmiId()
	}

	if *flagAmiLaunchIndex {
		showAmiLaunchIndex()
	}

	if *flagAmiManifestPath {
		showAmiManifestPath()
	}

	if *flagAncestorAmiIds {
		showAncestorAmiIds()
	}

	if *flagBlockDeviceMapping {
		showBlockDeviceMapping()
	}

	if *flagInstanceId {
		showInstanceId()
	}

	if *flagInstanceType {
		showInstanceType()
	}

	if *flagLocalHostname {
		showLocalHostname()
	}

	if *flagLocalIpv4 {
		showLocalIpv4()
	}

	if *flagKernelId {
		showKernelId()
	}

	if *flagAvailabilityZone {
		showAvailabilityZone()
	}

	if *flagProductCodes {
		showProductCodes()
	}

	if *flagPublicHostname {
		showPublicHostname()
	}

	if *flagPublicIpv4 {
		showPublicIpv4()
	}

	if *flagPublicKeys {
		showPublicKeys()
	}

	if *flagRamdiskId {
		showRamdiskId()
	}

	if *flagReservationId {
		showReservationId()
	}

	if *flagSecurityGroups {
		showSecurityGroups()
	}

	if *flagUserData {
		showUserData()
	}

  if *flagProfile {
    showProfile()
  }

  if *flagMac {
    showMac()
  }

  if *flagInstanceAction {
    showInstanceAction()
  }

	if *flagAll {
    writeHeader = true
		showAmiId()
		showAmiLaunchIndex()
		showAmiManifestPath()
		showAncestorAmiIds()
		showAvailabilityZone()
		showBlockDeviceMapping()
    showInstanceAction()
		showInstanceId()
		showInstanceType()
		showLocalHostname()
		showLocalIpv4()
		showKernelId()
    showMac()
		showProductCodes()
    showProfile()
		showPublicHostname()
		showPublicIpv4()
		showPublicKeys()
		showRamdiskId()
		showReservationId()
		showSecurityGroups()
		showUserData()
	}
}
