package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	flagAll                = kingpin.Flag("all", "Show all metadata information for this host (also default).").Bool()
	flagAmiId              = kingpin.Flag("ami-id", "The AMI ID used to launch this instance").Bool()
	flagAmiLaunchIndex     = kingpin.Flag("ami-launch-index", "The index of this instance in the reservation (per AMI).").Bool()
	flagAmiManifestPath    = kingpin.Flag("ami-manifest-path", "The manifest path of the AMI with which the instance was launched.").Bool()
	flagAncestorAmiIds     = kingpin.Flag("ancestor-ami-ids", "The AMI IDs of any instances that were rebundled to create this AMI.").Bool()
	flagBlockDeviceMapping = kingpin.Flag("block-device-mapping", "Defines native device names to use when exposing virtual devices.").Bool()
	flagInstanceId         = kingpin.Flag("instance-id", "The ID of this instance").Bool()
	flagInstanceType       = kingpin.Flag("instance-type", "The type of instance to launch. For more information, see Instance Types.").Bool()
	flagLocalHostname      = kingpin.Flag("local-hostname", "The local hostname of the instance.").Bool()
	flagLocalIpv4          = kingpin.Flag("local-ipv4", "Public IP address if launched with direct addressing; private IP address if launched with public addressing.").Bool()
	flagKernelId           = kingpin.Flag("kernel-id", "The ID of the kernel launched with this instance, if applicable.").Bool()
	flagAvailabilityZone   = kingpin.Flag("availability-zone", "The availability zone in which the instance launched. Same as placement").Bool()
	flagProductCodes       = kingpin.Flag("product-codes", "Product codes associated with this instance.").Bool()
	flagPublicHostname     = kingpin.Flag("public-hostname", "The public hostname of the instance.").Bool()
	flagPublicIpv4         = kingpin.Flag("public-ipv4", "NATted public IP Address").Bool()
	flagPublicKeys         = kingpin.Flag("public-keys", "Public keys. Only available if supplied at instance launch time").Bool()
	flagRamdiskId          = kingpin.Flag("ramdisk-id", "The ID of the RAM disk launched with this instance, if applicable.").Bool()
	flagReservationId      = kingpin.Flag("reservation-id", "ID of the reservation.").Bool()
	flagSecurityGroups     = kingpin.Flag("security-groups", "Names of the security groups the instance is launched in. Only available if supplied at instance launch time").Bool()
	flagUserData           = kingpin.Flag("user-data", "User-supplied data.Only available if supplied at instance launch time.").Bool()
	flagInstanceAction           = kingpin.Flag("instance-action", "display the instance-action").Bool()
	flagMac           = kingpin.Flag("mac", "display the instance mac address").Bool()
	flagProfile           = kingpin.Flag("profile", "display the instance profile").Bool()
)
