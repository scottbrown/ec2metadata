package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/ec2metadata"
	"strings"
)

const (
	VAL_NOT_AVAILABLE = "unavailable"

  KEY_AMI_ID = "ami-id"
  KEY_AMI_LAUNCH_INDEX = "ami-launch-index"
  KEY_AMI_MANIFEST_PATH = "ami-manifest-path"
  KEY_ANCESTOR_AMI_IDS = "ancestor-ami-ids"
  KEY_BLOCK_DEVICE_MAPPING = "block-device-mapping"
  KEY_INSTANCE_ID = "instance-id"
  KEY_INSTANCE_TYPE = "instance-type"
  KEY_LOCAL_HOSTNAME = "local-hostname"
  KEY_LOCAL_IPV4 = "local-ipv4"
  KEY_KERNEL_ID = "kernel-id"
  KEY_AVAILABILITY_ZONE = "availability-zone"
  KEY_PRODUCT_CODES = "product-codes"
  KEY_PUBLIC_HOSTNAME = "public-hostname"
  KEY_PUBLIC_IPV4 = "public-ipv4"
  KEY_PUBLIC_KEYS = "public-keys"
  KEY_RAMDISK_ID = "ramdisk-id"
  KEY_RESERVATION_ID = "reservation-id"
  KEY_SECURITY_GROUPS = "security-groups"
  KEY_USER_DATA = "user-data"
  KEY_MAC = "mac"
  KEY_INSTANCE_ACTION = "instance-action"
  KEY_PROFILE = "profile"
)

func showAmiId(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_AMI_ID)
	}
	fmt.Println(doc.ImageID)
}

func showAmiLaunchIndex(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_AMI_LAUNCH_INDEX)
	}
	val, err := svc.GetMetadata(KEY_AMI_LAUNCH_INDEX)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showAmiManifestPath(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_AMI_MANIFEST_PATH)
	}
	val, err := svc.GetMetadata(KEY_AMI_MANIFEST_PATH)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showAncestorAmiIds(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_ANCESTOR_AMI_IDS)
	}
	val, err := svc.GetMetadata(KEY_ANCESTOR_AMI_IDS)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showBlockDeviceMapping(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_BLOCK_DEVICE_MAPPING)
	}
	mappings, err := svc.GetMetadata(KEY_BLOCK_DEVICE_MAPPING)
	if err != nil {
		mappings = VAL_NOT_AVAILABLE
		return
	}

	blocks := strings.Split(mappings, "\n")
	for _, b := range blocks {
		device, err := svc.GetMetadata(fmt.Sprintf("%s/%s", KEY_BLOCK_DEVICE_MAPPING, b))
		if err != nil {
			device = VAL_NOT_AVAILABLE
		}
		fmt.Printf("%s: %s\n", b, device)
	}
}

func showInstanceId(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_INSTANCE_ID)
	}
	fmt.Println(doc.InstanceID)
}

func showInstanceType(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_INSTANCE_TYPE)
	}
	fmt.Println(doc.InstanceType)
}

func showLocalHostname(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_LOCAL_HOSTNAME)
	}
	val, err := svc.GetMetadata(KEY_LOCAL_HOSTNAME)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showLocalIpv4(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_LOCAL_IPV4)
	}
	fmt.Println(doc.PrivateIP)
}

func showKernelId(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_KERNEL_ID)
	}
	val := doc.KernelID
	if val == "" {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showAvailabilityZone(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_AVAILABILITY_ZONE)
	}
	fmt.Println(doc.AvailabilityZone)
}

func showProductCodes(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_PRODUCT_CODES)
	}
	val, err := svc.GetMetadata(KEY_PRODUCT_CODES)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showPublicHostname(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_PUBLIC_HOSTNAME)
	}
	val, err := svc.GetMetadata(KEY_PUBLIC_HOSTNAME)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showPublicIpv4(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_PUBLIC_IPV4)
	}
	val, err := svc.GetMetadata(KEY_PUBLIC_IPV4)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showPublicKeys(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_PUBLIC_KEYS)
	}
	mapping, err := svc.GetMetadata(KEY_PUBLIC_KEYS)
	if err != nil {
		fmt.Println(VAL_NOT_AVAILABLE)
    return
	}

	parts := strings.Split(mapping, "=")
	keyType, err := svc.GetMetadata(fmt.Sprintf("%s/%s/", KEY_PUBLIC_KEYS, parts[0]))
	if err != nil {
		fmt.Println(VAL_NOT_AVAILABLE)
    return
	}
  val, err := svc.GetMetadata(fmt.Sprintf("%s/%s/%s/", KEY_PUBLIC_KEYS, parts[0], keyType))
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
  fmt.Println(val)
}

func showRamdiskId(doc ec2metadata.EC2InstanceIdentityDocument, header bool) {
	if header {
		printHeader(KEY_RAMDISK_ID)
	}
	val := doc.RamdiskID
	if val == "" {
		val = VAL_NOT_AVAILABLE
	}

	fmt.Println(val)
}

func showReservationId(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_RESERVATION_ID)
	}
	val, err := svc.GetMetadata(KEY_RESERVATION_ID)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showSecurityGroups(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_SECURITY_GROUPS)
	}
	val, err := svc.GetMetadata(KEY_SECURITY_GROUPS)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showUserData(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_USER_DATA)
	}
	val, err := svc.GetUserData()
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showMac(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_MAC)
	}
  val, err := svc.GetMetadata(KEY_MAC)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showProfile(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_PROFILE)
	}
  val, err := svc.GetMetadata(KEY_PROFILE)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func showInstanceAction(svc *ec2metadata.EC2Metadata, header bool) {
	if header {
		printHeader(KEY_INSTANCE_ACTION)
	}
  val, err := svc.GetMetadata(KEY_INSTANCE_ACTION)
	if err != nil {
		val = VAL_NOT_AVAILABLE
	}
	fmt.Println(val)
}

func printHeader(key string) {
	fmt.Printf("%s: ", key)
}
