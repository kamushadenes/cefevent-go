package cefevent

var extensions []*CEFExtension

var extensionShortMap map[string]*CEFExtension
var extensionLongMap map[string]*CEFExtension

func buildMaps() {
	extensionShortMap = make(map[string]*CEFExtension)
	extensionLongMap = make(map[string]*CEFExtension)

	for k := range extensions {
		extensionShortMap[extensions[k].ShortName] = extensions[k]
		extensionLongMap[extensions[k].FullName] = extensions[k]
	}
}

func init() {
	extensions = append(extensions, &CEFExtension{
		ShortName:   "act",
		FullName:    "deviceAction",
		DataType:    "String",
		Length:      63,
		Description: "Action mentioned in the event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "app",
		FullName:    "applicationProtocol",
		DataType:    "String",
		Length:      31,
		Description: "Application level protocol, example values are: HTTP, HTTPS, SSHv2, Telnet, POP, IMAP, IMAPS, etc.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a1",
		FullName:    "deviceCustomIPv6Address1",
		DataType:    "IPv6 Address",
		Length:      0,
		Description: "One of four IPV6 address fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a1Label",
		FullName:    "deviceCustomIPv6Address1Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a2",
		FullName:    "deviceCustomIPv6Address2",
		DataType:    "IPv6 Address",
		Length:      0,
		Description: "One of four IPV6 address fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a2Label",
		FullName:    "deviceCustomIPv6Address2Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a3",
		FullName:    "deviceCustomIPv6Address3",
		DataType:    "IPv6 Address",
		Length:      0,
		Description: "One of four IPV6 address fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a3Label",
		FullName:    "deviceCustomIPv6Address3Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a4",
		FullName:    "deviceCustomIPv6Address4",
		DataType:    "IPv6 Address",
		Length:      0,
		Description: "One of four IPV6 address fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "c6a4Label",
		FullName:    "deviceCustomIPv6Address4Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp1",
		FullName:    "deviceCustomFloatingPoint1",
		DataType:    "Floating Point",
		Length:      0,
		Description: "One of four floating point fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp1Label",
		FullName:    "deviceCustomFloatingPoint1Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp2",
		FullName:    "deviceCustomFloatingPoint2",
		DataType:    "Floating Point",
		Length:      0,
		Description: "One of four floating point fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp2Label",
		FullName:    "deviceCustomFloatingPoint2Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp3",
		FullName:    "deviceCustomFloatingPoint3",
		DataType:    "Floating Point",
		Length:      0,
		Description: "One of four floating point fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp3Label",
		FullName:    "deviceCustomFloatingPoint3Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp4",
		FullName:    "deviceCustomFloatingPoint4",
		DataType:    "Floating Point",
		Length:      0,
		Description: "One of four floating point fields available to map fields that do not apply to any other in this dictionary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cfp4Label",
		FullName:    "deviceCustomFloatingPoint4Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cat",
		FullName:    "deviceEventCategory",
		DataType:    "String",
		Length:      1023,
		Description: "Represents the category assigned by the originating device. Devices oftentimes use their own categorization schema to classify events.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cn1",
		FullName:    "deviceCustomNumber1",
		DataType:    "Long",
		Length:      0,
		Description: "There are three number fields available which can be used to map fields which do not fit into any other field of this dictionary. If possible, \"these fields should not be used, but a more specific field from the dictionary. Also check the guidelines hereafter for hints on how to utilize these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cn1Label",
		FullName:    "deviceCustomNumber1Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cn2",
		FullName:    "deviceCustomNumber2",
		DataType:    "Long",
		Length:      0,
		Description: "There are three number fields available which can be used to map fields which do not fit into any other field of this dictionary. If possible, \"these fields should not be used, but a more specific field from the dictionary. Also check the guidelines hereafter for hints on how to utilize these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cn2Label",
		FullName:    "deviceCustomNumber2Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cn3",
		FullName:    "deviceCustomNumber3",
		DataType:    "Long",
		Length:      0,
		Description: "There are three number fields available which can be used to map fields which do not fit into any other field of this dictionary. If possible, \"these fields should not be used, but a more specific field from the dictionary. Also check the guidelines hereafter for hints on how to utilize these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cn3Label",
		FullName:    "deviceCustomNumber3Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cnt",
		FullName:    "baseEventCount",
		DataType:    "Integer",
		Length:      0,
		Description: "A count associated with this event. How many times was this same event observed?",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs1",
		FullName:    "deviceCustomString1",
		DataType:    "String",
		Length:      1023,
		Description: "There are six strings available which can be used to map fields which do not fit into any other field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs1Label",
		FullName:    "deviceCustomString1Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs2",
		FullName:    "deviceCustomString2",
		DataType:    "String",
		Length:      1023,
		Description: "There are six strings available which can be used to map fields which do not fit into any other field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs2Label",
		FullName:    "deviceCustomString2Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs3",
		FullName:    "deviceCustomString3",
		DataType:    "String",
		Length:      1023,
		Description: "There are six strings available which can be used to map fields which do not fit into any other field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs3Label",
		FullName:    "deviceCustomString3Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs4",
		FullName:    "deviceCustomString4",
		DataType:    "String",
		Length:      1023,
		Description: "There are six strings available which can be used to map fields which do not fit into any other field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs4Label",
		FullName:    "deviceCustomString4Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs5",
		FullName:    "deviceCustomString5",
		DataType:    "String",
		Length:      1023,
		Description: "There are six strings available which can be used to map fields which do not fit into any other field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs5Label",
		FullName:    "deviceCustomString5Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs6",
		FullName:    "deviceCustomString6",
		DataType:    "String",
		Length:      1023,
		Description: "There are six strings available which can be used to map fields which do not fit into any other field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "cs6Label",
		FullName:    "deviceCustomString6Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "destinationDnsDomain",
		FullName:    "destinationDnsDomain",
		DataType:    "String",
		Length:      255,
		Description: "The DNS domain part of the complete fully qualified domain name (FQDN).",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "destinationServiceName",
		FullName:    "destinationServiceName",
		DataType:    "String",
		Length:      1023,
		Description: "The service which is targeted by this event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "destinationTranslatedAddress",
		FullName:    "destinationTranslatedAddress",
		DataType:    "IPv4 Address",
		Length:      0,
		Description: "Identifies the translated destination that the event refers to in an IP network. The format is an IPv4 address. Example: \"192.168.10.1\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "destinationTranslatedPort",
		FullName:    "destinationTranslatedPort",
		DataType:    "Integer",
		Length:      0,
		Description: "Port after it was translated",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceCustomDate1",
		FullName:    "deviceCustomDate1",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "There are two timestamp fields",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceCustomDate1Label",
		FullName:    "deviceCustomDate1Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceCustomDate2",
		FullName:    "deviceCustomDate2",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "There are two timestamp fields available which can be used to map fields which do not fit into any other \"field of this dictionary. If possible, these fields should not be used, but a more specific field from the dictionary. Also check the guidelines later in this document for hints about utilizing these fields.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceCustomDate2Label",
		FullName:    "deviceCustomDate2Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceCustomDate3Label",
		FullName:    "deviceCustomDate3Label",
		DataType:    "String",
		Length:      1023,
		Description: "All custom fields have a corresponding label field where the field itself can be described. Each of the fields is a string describing the purpose of this field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceDirection",
		FullName:    "deviceDirection",
		DataType:    "String",
		Length:      0,
		Description: "Any information about what direction the communication that was observed has taken.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceDnsDomain",
		FullName:    "deviceDnsDomain",
		DataType:    "String",
		Length:      255,
		Description: "The DNS domain part of the complete fully qualified domain name (FQDN).",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceExternalId",
		FullName:    "deviceExternalId",
		DataType:    "String",
		Length:      255,
		Description: "A name that uniquely identifies the device generating this event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceFacility",
		FullName:    "deviceFacility",
		DataType:    "String",
		Length:      1023,
		Description: "The facility generating this event. Syslog for example has an explicit facility associated with every event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceInboundInterface",
		FullName:    "deviceInboundInterface",
		DataType:    "String",
		Length:      15,
		Description: "Interface on which the packet or data entered the device.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceMacAddress",
		FullName:    "deviceMacAddress",
		DataType:    "MAC Address",
		Length:      0,
		Description: "Six colon-separated hexadecimal numbers. Example: 00:0D:60:AF:1B:61",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceNtDomain",
		FullName:    "deviceNtDomain",
		DataType:    "String",
		Length:      255,
		Description: "The Windows domain name of the device address.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceOutboundInterface",
		FullName:    "deviceOutboundInterface",
		DataType:    "String",
		Length:      15,
		Description: "Interface on which the packet or data left the device.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceProcessName",
		FullName:    "deviceProcessName",
		DataType:    "String",
		Length:      1023,
		Description: "Process name associated to the event. In UNIX, the process generating the syslog entry for example.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "deviceTranslatedAddress",
		FullName:    "deviceTranslatedAddress",
		DataType:    "IPv4 Address",
		Length:      0,
		Description: "Identifies the translated device address that the event refers to in an IP network. The format is an IPv4 address. Example: \"192.168.10.1\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dhost",
		FullName:    "DestinationHostName",
		DataType:    "String",
		Length:      1023,
		Description: "Identifies the source that an event refers to in an IP network. The format should be a fully qualified domain name associated with the \"destination node, when a node is available. Examples: \"host.domain.com\" or \"host\".",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dmac",
		FullName:    "destinationMac",
		DataType:    "MAC Address",
		Length:      0,
		Description: "Six colon-separated hexadecimal numbers. Example: \"00:0D:60:AF:1B:61\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dntdom",
		FullName:    "destinationNtDomain",
		DataType:    "String",
		Length:      255,
		Description: "The Windows domain name of the destination address.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dpid",
		FullName:    "destinationProcessId",
		DataType:    "Integer",
		Length:      0,
		Description: "Provides the ID of the destination process associated with the event. For example, if an event contains process ID 105, \"105\" is the process ID.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dpriv",
		FullName:    "destinationUserPrivileges",
		DataType:    "String",
		Length:      1023,
		Description: "The allowed values are: \"Administrator\", \"User\", and \"Guest\". This identifies the destination user's privileges. In UNIX, for example, activity executed on the root user would be identified with destinationUserPrivileges of \"Administrator\". This is an idealized and simplified view on privileges and can be extended in the future.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dproc",
		FullName:    "destinationProcessName",
		DataType:    "String",
		Length:      1023,
		Description: "The name of the process which is the event's destination. For example: \"telnetd\", or \"sshd\".",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dpt",
		FullName:    "destinationPort",
		DataType:    "Integer",
		Length:      0,
		Description: "The valid port numbers are between 0 and 65535.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dst",
		FullName:    "destinationAddress",
		DataType:    "IPv4 Address",
		Length:      0,
		Description: "Identifies destination that the event refers to in an IP network. The format is an IPv4 address. Example: \"192.168.10.1\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "duid",
		FullName:    "destinationUserId",
		DataType:    "String",
		Length:      1023,
		Description: "Identifies the destination user by ID. For example, in UNIX, the root user is generally associated with user ID 0.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "duser",
		FullName:    "destinationUserName",
		DataType:    "String",
		Length:      1023,
		Description: "Identifies the destination user by name. This is the user associated with the event's destination. E-mail addresses are also mapped into the UserName fields. The recipient is a candidate to put into destinationUserName.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dvc",
		FullName:    "deviceAddress",
		DataType:    "IPv4 Address",
		Length:      16,
		Description: "Identifies the device that an event refers to in an IP network. The format is an IPv4 address. Example: \"192.168.10.1\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dvchost",
		FullName:    "deviceHostName",
		DataType:    "String",
		Length:      100,
		Description: "The format should be a fully qualified domain name associated with the device node, when a node is available. Examples: \"host.domain.com\" or \"host\".",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "dvcpid",
		FullName:    "deviceProcessId",
		DataType:    "Integer",
		Length:      0,
		Description: "Provides the ID of the process on the device generating the event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "end",
		FullName:    "endTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "The time at which the activity related to the event ended. The format is MMM dd yyyy HH:mm:ss or milliseconds since epoch (Jan 1st 1970). An example would be reporting the end of a session.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "externalId",
		FullName:    "externalId",
		DataType:    "String",
		Length:      40,
		Description: "An ID used by the originating device.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fileCreateTime",
		FullName:    "fileCreateTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "Time when file was created.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fileHash",
		FullName:    "fileHash",
		DataType:    "String",
		Length:      255,
		Description: "Hash of a file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fileId",
		FullName:    "fileId",
		DataType:    "String",
		Length:      1023,
		Description: "An ID associated with a file could be the inode.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fileModificationTime",
		FullName:    "fileModificationTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "Time when file was last modified.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "filePath",
		FullName:    "filePath",
		DataType:    "String",
		Length:      1023,
		Description: "Full path to the file, including file name itself.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "filePermission",
		FullName:    "filePermission",
		DataType:    "String",
		Length:      1023,
		Description: "Permissions of the file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fileType",
		FullName:    "fileType",
		DataType:    "String",
		Length:      1023,
		Description: "Type of file (pipe, socket, etc.)",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fd1",
		FullName:    "flexDate1",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "A timestamp field available to map a timestamp that does not apply to any other defined timestamp field in this dictionary. Use all flex fields sparingly and seek a more specific, dictionary supplied field when possible. These fields are typically reserved for customer use and should not be set by vendors unless necessary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fd1Label",
		FullName:    "flexDate1Label",
		DataType:    "String",
		Length:      128,
		Description: "The label field is a string and describes the purpose of the flex field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fs1",
		FullName:    "flexString1",
		DataType:    "String",
		Length:      1023,
		Description: "One of four floating point fields available to map fields that do not apply to any other in this dictionary. Use sparingly and seek a more specific, dictionary supplied field when possible. These fields are typically reserved for customer use and should not be set by vendors unless necessary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fs1Label",
		FullName:    "flexString1Label",
		DataType:    "String",
		Length:      128,
		Description: "The label field is a string and describes the purpose of the flex field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fs2",
		FullName:    "flexString2",
		DataType:    "String",
		Length:      1023,
		Description: "One of four floating point fields available to map fields that do not apply to any other in this dictionary. Use sparingly and seek a more specific, dictionary supplied field when possible. These fields are typically reserved for customer use and should not be set by vendors unless necessary.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fs2Label",
		FullName:    "flexString2Label",
		DataType:    "String",
		Length:      128,
		Description: "The label field is a string and describes the purpose of the flex field.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fname",
		FullName:    "fileName",
		DataType:    "String",
		Length:      1023,
		Description: "Name of the file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "fsize",
		FullName:    "fileSize",
		DataType:    "Integer",
		Length:      0,
		Description: "Size of the file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "in",
		FullName:    "bytesIn",
		DataType:    "Integer",
		Length:      0,
		Description: "Number of bytes transferred inbound. Inbound relative to the source to destination relationship, meaning that data was flowing from source to destination.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "msg",
		FullName:    "message",
		DataType:    "String",
		Length:      1023,
		Description: "An arbitrary message giving more details about the event. Multi-line entries can be produced by using \n as the new-line separator.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileCreateTime",
		FullName:    "oldFileCreateTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "Time when old file was created.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileHash",
		FullName:    "oldFileHash",
		DataType:    "String",
		Length:      255,
		Description: "Hash of the old file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileId",
		FullName:    "oldFileId",
		DataType:    "String",
		Length:      1023,
		Description: "An ID associated with the old file could be the inode.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileModificationTime",
		FullName:    "oldFileModificationTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "Time when old file was last modified.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileName",
		FullName:    "oldFileName",
		DataType:    "String",
		Length:      1023,
		Description: "Name of the old file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFilePath",
		FullName:    "oldFilePath",
		DataType:    "String",
		Length:      1023,
		Description: "Full path to the old file, including file name itself.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFilePermission",
		FullName:    "oldFilePermission",
		DataType:    "String",
		Length:      1023,
		Description: "Permissions of the old file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileSize",
		FullName:    "oldFileSize",
		DataType:    "Integer",
		Length:      0,
		Description: "Size of the old file.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "oldFileType",
		FullName:    "oldFileType",
		DataType:    "String",
		Length:      1023,
		Description: "Type of the old file (pipe, socket, etc.)",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "out",
		FullName:    "bytesOut",
		DataType:    "Integer",
		Length:      0,
		Description: "Number of bytes transferred outbound. Outbound relative to the source to destination relationship, meaning that data was flowing from destination to source.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "proto",
		FullName:    "transportProtocol",
		DataType:    "String",
		Length:      31,
		Description: "Identifies the Layer-4 protocol used. The possible values are protocol names such as TCP or UDP.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "reason",
		FullName:    "reason",
		DataType:    "String",
		Length:      1023,
		Description: "The reason an audit event was generated. For example \"Bad password\" or \"Unknown User\". This could also be an error or return code. Example: \"0x1234\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "requestClientApplication",
		FullName:    "requestClientApplication",
		DataType:    "String",
		Length:      1023,
		Description: "The User-Agent associated with the request.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "requestCookies",
		FullName:    "requestCookies",
		DataType:    "String",
		Length:      1023,
		Description: "Cookies associated with the request.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "requestMethod",
		FullName:    "requestMethod",
		DataType:    "String",
		Length:      1023,
		Description: "The method used to access a URL. Possible values: \"POST\", \"GET\", ...",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "request",
		FullName:    "requestURL",
		DataType:    "String",
		Length:      1023,
		Description: "In the case of an HTTP request, this field contains the URL accessed. The URL should contain the protocol as well, e.g., \"http://www.security.com\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "rt",
		FullName:    "receiptTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "The time at which the event related to the activity was received. The format is MMM dd yyyy HH:mm:ss or milliseconds since epoch (Jan 1st 1970).",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "shost",
		FullName:    "sourceHostName",
		DataType:    "String",
		Length:      1023,
		Description: "Identifies the source that an event refers to in an IP network. The format should be a fully qualified domain name associated with the source node, when a node is available. Examples: \"host.domain.com\" or \"host\".",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "smac",
		FullName:    "sourceMacAddress",
		DataType:    "MAC Address",
		Length:      0,
		Description: "Six colon-separated hexadecimal numbers. Example: \"00:0D:60:AF:1B:61\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "sntdom",
		FullName:    "sourceNtDomain",
		DataType:    "String",
		Length:      255,
		Description: "The Windows domain name for the source address.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "sourceDnsDomain",
		FullName:    "sourceDnsDomain",
		DataType:    "String",
		Length:      255,
		Description: "The DNS domain part of the complete fully qualified domain name (FQDN).",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "sourceServiceName",
		FullName:    "sourceServiceName",
		DataType:    "String",
		Length:      1023,
		Description: "The service which is responsible for generating this event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "sourceTranslatedAddress",
		FullName:    "sourceTranslatedAddress",
		DataType:    "IPv4 Address",
		Length:      0,
		Description: "Identifies the translated source that the event refers to in an IP network. The format is an IPv4 address. Example: \"192.168.10.1\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "sourceTranslatedPort",
		FullName:    "sourceTranslatedPort",
		DataType:    "Integer",
		Length:      0,
		Description: "Port after it was translated by for example a firewall. Valid port numbers are 0 to 65535.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "spid",
		FullName:    "sourceProcessId",
		DataType:    "Integer",
		Length:      0,
		Description: "The ID of the source process associated with the event.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "spriv",
		FullName:    "sourceUserPrivileges",
		DataType:    "String",
		Length:      1023,
		Description: "The allowed values are: \"Administrator\", \"User\", and \"Guest\". It identifies the source user's privileges. In UNIX, for example, activity executed by the root user would be identified with sourceUserPrivileges of \"Administrator\". This is an idealized and simplified view on privileges and can be extended in the future.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "spt",
		FullName:    "sourcePort",
		DataType:    "Integer",
		Length:      0,
		Description: "The valid port numbers are 0 to 65535.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "src",
		FullName:    "sourceAddress",
		DataType:    "IPv4 Address",
		Length:      0,
		Description: "Identifies the source that an event refers to in an IP network. The format is an IPv4 address. Example: \"192.168.10.1\"",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "start",
		FullName:    "startTime",
		DataType:    "TimeStamp",
		Length:      0,
		Description: "The time when the activity the event referred to started. The format is MMM dd yyyy HH:mm:ss or milliseconds since epoch (Jan 1st 1970).",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "suid",
		FullName:    "sourceUserId",
		DataType:    "String",
		Length:      1023,
		Description: "Identifies the source user by ID. This is the user associated with the source of the event. For example, in \"UNIX, the root user is generally associated with user ID 0.",
	})
	extensions = append(extensions, &CEFExtension{
		ShortName:   "suser",
		FullName:    "sourceUserName",
		DataType:    "String",
		Length:      1023,
		Description: "Identifies the source user by name. E-mail addresses are also mapped into the UserName fields. The sender is a candidate to put into sourceUserName.",
	})

	buildMaps()
}
