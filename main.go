package main

import (
	_ "github.com/aws/aws-sdk-go-v2/aws/ec2metadata"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	VERSION string

	showAll                = kingpin.Flag("all", "Show all metadata information for this host (also default).").Bool()
	showAmiId              = kingpin.Flag("ami-id", "The AMI ID used to launch this instance").Bool()
	showAmiLaunchIndex     = kingpin.Flag("ami-launch-index", "The index of this instance in the reservation (per AMI).").Bool()
	showAmiManifestPath    = kingpin.Flag("ami-manifest-path", "The manifest path of the AMI with which the instance was launched.").Bool()
	showAncestorAmiIds     = kingpin.Flag("ancestor-ami-ids", "The AMI IDs of any instances that were rebundled to create this AMI.").Bool()
	showBlockDeviceMapping = kingpin.Flag("block-device-mapping", "Defines native device names to use when exposing virtual devices.").Bool()
	showInstanceId         = kingpin.Flag("instance-id", "The ID of this instance").Bool()
	showInstanceType       = kingpin.Flag("instance-type", "The type of instance to launch. For more information, see Instance Types.").Bool()
	showLocalHostname      = kingpin.Flag("local-hostname", "The local hostname of the instance.").Bool()
	showLocalIpv4          = kingpin.Flag("local-ipv4", "Public IP address if launched with direct addressing; private IP address if launched with public addressing.").Bool()
	showKernelId           = kingpin.Flag("kernel-id", "The ID of the kernel launched with this instance, if applicable.").Bool()
	showAvailabilityZone   = kingpin.Flag("availability-zone", "The availability zone in which the instance launched. Same as placement").Bool()
	showProductCodes       = kingpin.Flag("product-codes", "Product codes associated with this instance.").Bool()
	showPublicHostname     = kingpin.Flag("public-hostname", "The public hostname of the instance.").Bool()
	showPublicIpv4         = kingpin.Flag("public-ipv4", "NATted public IP Address").Bool()
	showPublicKeys         = kingpin.Flag("public-keys", "Public keys. Only available if supplied at instance launch time").Bool()
	showRamdiskId          = kingpin.Flag("ramdisk-id", "The ID of the RAM disk launched with this instance, if applicable.").Bool()
	showReservationId      = kingpin.Flag("reservation-id", "ID of the reservation.").Bool()
	showSecurityGroups     = kingpin.Flag("security-groups", "Names of the security groups the instance is launched in. Only available if supplied at instance launch time").Bool()
	showUserData           = kingpin.Flag("user-data", "User-supplied data.Only available if supplied at instance launch time.").Bool()
)

func main() {
	kingpin.Version(VERSION)
	kingpin.Parse()
}
