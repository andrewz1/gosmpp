package Data

import (
	"sync"
	"time"
)

const (
	SM_CONNID_LEN        uint32 = 16
	SM_MSG_LEN           uint32 = 254
	SM_SYSID_LEN         uint32 = 16
	SM_MSGID_LEN         uint32 = 64
	SM_PASS_LEN          uint32 = 9
	SM_DATE_LEN          uint32 = 17
	SM_SRVTYPE_LEN       uint32 = 6
	SM_SYSTYPE_LEN       uint32 = 13
	SM_ADDR_LEN          uint32 = 21
	SM_DATA_ADDR_LEN     uint32 = 65
	SM_ADDR_RANGE_LEN    uint32 = 41
	SM_TYPE_LEN          uint32 = 13
	SM_DL_NAME_LEN       uint32 = 21
	SM_PARAM_NAME_LEN    uint32 = 10
	SM_PARAM_VALUE_LEN   uint32 = 10
	SM_MAX_CNT_DEST_ADDR uint32 = 254

	CONNECTION_CLOSED uint32 = 0
	CONNECTION_OPENED uint32 = 1

	SM_ACK            uint32 = 1
	SM_NO_ACK         uint32 = 0
	SM_RESPONSE_ACK   uint32 = 0
	SM_RESPONSE_TNACK uint32 = 1
	SM_RESPONSE_PNACK uint32 = 2

	//SMPP Command Set
	GENERIC_NACK          uint32 = 0x80000000
	BIND_RECEIVER         uint32 = 0x00000001
	BIND_RECEIVER_RESP    uint32 = 0x80000001
	BIND_TRANSMITTER      uint32 = 0x00000002
	BIND_TRANSMITTER_RESP uint32 = 0x80000002
	QUERY_SM              uint32 = 0x00000003
	QUERY_SM_RESP         uint32 = 0x80000003
	SUBMIT_SM             uint32 = 0x00000004
	SUBMIT_SM_RESP        uint32 = 0x80000004
	DELIVER_SM            uint32 = 0x00000005
	DELIVER_SM_RESP       uint32 = 0x80000005
	UNBIND                uint32 = 0x00000006
	UNBIND_RESP           uint32 = 0x80000006
	REPLACE_SM            uint32 = 0x00000007
	REPLACE_SM_RESP       uint32 = 0x80000007
	CANCEL_SM             uint32 = 0x00000008
	CANCEL_SM_RESP        uint32 = 0x80000008
	BIND_TRANSCEIVER      uint32 = 0x00000009
	BIND_TRANSCEIVER_RESP uint32 = 0x80000009
	OUTBIND               uint32 = 0x0000000B
	ENQUIRE_LINK          uint32 = 0x00000015
	ENQUIRE_LINK_RESP     uint32 = 0x80000015
	SUBMIT_MULTI          uint32 = 0x00000021
	SUBMIT_MULTI_RESP     uint32 = 0x80000021
	ALERT_NOTIFICATION    uint32 = 0x00000102
	DATA_SM               uint32 = 0x00000103
	DATA_SM_RESP          uint32 = 0x80000103

	//Command_Status Error Codes
	ESME_ROK           uint32 = 0x00000000
	ESME_RINVMSGLEN    uint32 = 0x00000001
	ESME_RINVCMDLEN    uint32 = 0x00000002
	ESME_RINVCMDID     uint32 = 0x00000003
	ESME_RINVBNDSTS    uint32 = 0x00000004
	ESME_RALYBND       uint32 = 0x00000005
	ESME_RINVPRTFLG    uint32 = 0x00000006
	ESME_RINVREGDLVFLG uint32 = 0x00000007
	ESME_RSYSERR       uint32 = 0x00000008
	ESME_RINVSRCADR    uint32 = 0x0000000A
	ESME_RINVDSTADR    uint32 = 0x0000000B
	ESME_RINVMSGID     uint32 = 0x0000000C
	ESME_RBINDFAIL     uint32 = 0x0000000D
	ESME_RINVPASWD     uint32 = 0x0000000E
	ESME_RINVSYSID     uint32 = 0x0000000F
	ESME_RCANCELFAIL   uint32 = 0x00000011
	ESME_RREPLACEFAIL  uint32 = 0x00000013
	ESME_RMSGQFUL      uint32 = 0x00000014
	ESME_RINVSERTYP    uint32 = 0x00000015

	ESME_RADDCUSTFAIL  uint32 = 0x00000019 // Failed to Add Customer
	ESME_RDELCUSTFAIL  uint32 = 0x0000001A // Failed to delete Customer
	ESME_RMODCUSTFAIL  uint32 = 0x0000001B // Failed to modify customer
	ESME_RENQCUSTFAIL  uint32 = 0x0000001C // Failed to Enquire Customer
	ESME_RINVCUSTID    uint32 = 0x0000001D // Invalid Customer ID
	ESME_RINVCUSTNAME  uint32 = 0x0000001F // Invalid Customer Name
	ESME_RINVCUSTADR   uint32 = 0x00000021 // Invalid Customer Address
	ESME_RINVADR       uint32 = 0x00000022 // Invalid Address
	ESME_RCUSTEXIST    uint32 = 0x00000023 // Customer Exists
	ESME_RCUSTNOTEXIST uint32 = 0x00000024 // Customer does not exist
	ESME_RADDDLFAIL    uint32 = 0x00000026 // Failed to Add DL
	ESME_RMODDLFAIL    uint32 = 0x00000027 // Failed to modify DL
	ESME_RDELDLFAIL    uint32 = 0x00000028 // Failed to Delete DL
	ESME_RVIEWDLFAIL   uint32 = 0x00000029 // Failed to View DL
	ESME_RLISTDLSFAIL  uint32 = 0x00000030 // Failed to list DLs
	ESME_RPARAMRETFAIL uint32 = 0x00000031 // Param Retrieve Failed
	ESME_RINVPARAM     uint32 = 0x00000032 // Invalid Param

	ESME_RINVNUMDESTS uint32 = 0x00000033
	ESME_RINVDLNAME   uint32 = 0x00000034

	ESME_RINVDLMEMBDESC uint32 = 0x00000035 // Invalid DL Member Description
	ESME_RINVDLMEMBTYP  uint32 = 0x00000038 // Invalid DL Member Type
	ESME_RINVDLMODOPT   uint32 = 0x00000039 // Invalid DL Modify Option

	ESME_RINVDESTFLAG uint32 = 0x00000040
	ESME_RINVSUBREP   uint32 = 0x00000042
	ESME_RINVESMCLASS uint32 = 0x00000043
	ESME_RCNTSUBDL    uint32 = 0x00000044
	ESME_RSUBMITFAIL  uint32 = 0x00000045
	ESME_RINVSRCTON   uint32 = 0x00000048
	ESME_RINVSRCNPI   uint32 = 0x00000049
	ESME_RINVDSTTON   uint32 = 0x00000050
	ESME_RINVDSTNPI   uint32 = 0x00000051
	ESME_RINVSYSTYP   uint32 = 0x00000053
	ESME_RINVREPFLAG  uint32 = 0x00000054
	ESME_RINVNUMMSGS  uint32 = 0x00000055
	ESME_RTHROTTLED   uint32 = 0x00000058

	ESME_RPROVNOTALLWD uint32 = 0x00000059 // Provisioning Not Allowed

	ESME_RINVSCHED    uint32 = 0x00000061
	ESME_RINVEXPIRY   uint32 = 0x00000062
	ESME_RINVDFTMSGID uint32 = 0x00000063
	ESME_RX_T_APPN    uint32 = 0x00000064
	ESME_RX_P_APPN    uint32 = 0x00000065
	ESME_RX_R_APPN    uint32 = 0x00000066
	ESME_RQUERYFAIL   uint32 = 0x00000067

	ESME_RINVPGCUSTID      uint32 = 0x00000080 // Paging Customer ID Invalid No such subscriber
	ESME_RINVPGCUSTIDLEN   uint32 = 0x00000081 // Paging Customer ID length Invalid
	ESME_RINVCITYLEN       uint32 = 0x00000082 // City Length Invalid
	ESME_RINVSTATELEN      uint32 = 0x00000083 // State Length Invalid
	ESME_RINVZIPPREFIXLEN  uint32 = 0x00000084 // Zip Prefix Length Invalid
	ESME_RINVZIPPOSTFIXLEN uint32 = 0x00000085 // Zip Postfix Length Invalid
	ESME_RINVMINLEN        uint32 = 0x00000086 // MIN Length Invalid
	ESME_RINVMIN           uint32 = 0x00000087 // MIN Invalid (i.e. No such MIN)
	ESME_RINVPINLEN        uint32 = 0x00000088 // PIN Length Invalid
	ESME_RINVTERMCODELEN   uint32 = 0x00000089 // Terminal Code Length Invalid
	ESME_RINVCHANNELLEN    uint32 = 0x0000008A // Channel Length Invalid
	ESME_RINVCOVREGIONLEN  uint32 = 0x0000008B // Coverage Region Length Invalid
	ESME_RINVCAPCODELEN    uint32 = 0x0000008C // Cap Code Length Invalid
	ESME_RINVMDTLEN        uint32 = 0x0000008D // Message delivery time Length Invalid
	ESME_RINVPRIORMSGLEN   uint32 = 0x0000008E // Priority Message Length Invalid
	ESME_RINVPERMSGLEN     uint32 = 0x0000008F // Periodic Messages Length Invalid
	ESME_RINVPGALERTLEN    uint32 = 0x00000090 // Paging Alerts Length Invalid
	ESME_RINVSMUSERLEN     uint32 = 0x00000091 // int16 Message User Group Length Invalid
	ESME_RINVRTDBLEN       uint32 = 0x00000092 // Real Time Data broadcasts Length Invalid
	ESME_RINVREGDELLEN     uint32 = 0x00000093 // Registered Delivery Lenght Invalid
	ESME_RINVMSGDISTLEN    uint32 = 0x00000094 // Message Distribution Lenght Invalid
	ESME_RINVPRIORMSG      uint32 = 0x00000095 // Priority Message Length Invalid
	ESME_RINVMDT           uint32 = 0x00000096 // Message delivery time Invalid
	ESME_RINVPERMSG        uint32 = 0x00000097 // Periodic Messages Invalid
	ESME_RINVMSGDIST       uint32 = 0x00000098 // Message Distribution Invalid
	ESME_RINVPGALERT       uint32 = 0x00000099 // Paging Alerts Invalid
	ESME_RINVSMUSER        uint32 = 0x0000009A // int16 Message User Group Invalid
	ESME_RINVRTDB          uint32 = 0x0000009B // Real Time Data broadcasts Invalid
	ESME_RINVREGDEL        uint32 = 0x0000009C // Registered Delivery Invalid
	//public static final uint32 ESME_RINVOPTPARSTREAM = 0x0000009D  // KIF IW Field out of data
	//public static final uint32 ESME_ROPTPARNOTALLWD = 0x0000009E  // Optional Parameter not allowed
	ESME_RINVOPTPARLEN uint32 = 0x0000009F // Invalid Optional Parameter Length

	ESME_RINVOPTPARSTREAM uint32 = 0x000000C0
	ESME_ROPTPARNOTALLWD  uint32 = 0x000000C1
	ESME_RINVPARLEN       uint32 = 0x000000C2
	ESME_RMISSINGOPTPARAM uint32 = 0x000000C3
	ESME_RINVOPTPARAMVAL  uint32 = 0x000000C4
	ESME_RDELIVERYFAILURE uint32 = 0x000000FE
	ESME_RUNKNOWNERR      uint32 = 0x000000FF

	ESME_LAST_ERROR uint32 = 0x0000012C // the value of the last error code

	//Interface_Version
	SMPP_V33 byte = (0x33 | 0x80)
	SMPP_V34 byte = 0x34

	//Address_TON
	GSM_TON_UNKNOWN       byte = 0x00
	GSM_TON_INTERNATIONAL byte = 0x01
	GSM_TON_NATIONAL      byte = 0x02
	GSM_TON_NETWORK       byte = 0x03
	GSM_TON_SUBSCRIBER    byte = 0x04
	GSM_TON_ALPHANUMERIC  byte = 0x05
	GSM_TON_ABBREVIATED   byte = 0x06
	GSM_TON_RESERVED_EXTN byte = 0x07

	//Address_NPI
	GSM_NPI_UNKNOWN       byte = 0x00
	GSM_NPI_E164          byte = 0x01
	GSM_NPI_ISDN               = GSM_NPI_E164
	GSM_NPI_X121          byte = 0x03
	GSM_NPI_TELEX         byte = 0x04
	GSM_NPI_LAND_MOBILE   byte = 0x06
	GSM_NPI_NATIONAL      byte = 0x08
	GSM_NPI_PRIVATE       byte = 0x09
	GSM_NPI_ERMES         byte = 0x0A
	GSM_NPI_INTERNET      byte = 0x0E
	GSM_NPI_WAP_CLIENT_ID byte = 0x12
	GSM_NPI_RESERVED_EXTN byte = 0x0F

	//Service_Type
	SERVICE_NULL string = ""
	SERVICE_CMT  string = "CMT"
	SERVICE_CPT  string = "CPT"
	SERVICE_VMN  string = "VMN"
	SERVICE_VMA  string = "VMA"
	SERVICE_WAP  string = "WAP"
	SERVICE_USSD string = "USSD"

	SMPP_PROTOCOL                 byte = 1
	SMPPP_PROTOCOL                byte = 2
	SM_SERVICE_MOBILE_TERMINATED  byte = 0
	SM_SERVICE_MOBILE_ORIGINATED  byte = 1
	SM_SERVICE_MOBILE_TRANSCEIVER byte = 2

	// State of message at SMSC
	SM_STATE_EN_ROUTE      uint32 = 1 // default state for messages in transit
	SM_STATE_DELIVERED     uint32 = 2 // message is delivered
	SM_STATE_EXPIRED       uint32 = 3 // validity period expired
	SM_STATE_DELETED       uint32 = 4 // message has been deleted
	SM_STATE_UNDELIVERABLE uint32 = 5 // undeliverable
	SM_STATE_ACCEPTED      uint32 = 6 // message is in accepted state
	SM_STATE_INVALID       uint32 = 7 // message is in invalid state
	SM_STATE_REJECTED      uint32 = 8 // message is in rejected state

	//******************
	// ESMClass Defines
	//******************

	// Messaging Mode
	SM_ESM_DEFAULT        uint32 = 0x00 //Default SMSC Mode or Message Type
	SM_DATAGRAM_MODE      uint32 = 0x01 // Use one-shot express mode
	SM_FORWARD_MODE       uint32 = 0x02 // Do not use
	SM_STORE_FORWARD_MODE uint32 = 0x03 // Use store & forward

	// Send/Receive TDMA & CDMA Message Type
	SM_SMSC_DLV_RCPT_TYPE     uint32 = 0x04 // Recv Msg contains SMSC delivery receipt
	SM_ESME_DLV_ACK_TYPE      uint32 = 0x08 // Send/Recv Msg contains ESME delivery acknowledgement
	SM_ESME_MAN_USER_ACK_TYPE uint32 = 0x10 // Send/Recv Msg contains manual/user acknowledgment
	SM_CONV_ABORT_TYPE        uint32 = 0x18 // Recv Msg contains conversation abort (Korean CDMA)
	SM_INTMD_DLV_NOTIFY_TYPE  uint32 = 0x20 // Recv Msg contains intermediate notification

	// GSM Network features
	SM_NONE_GSM           uint32 = 0x00 // No specific features selected
	SM_UDH_GSM            uint32 = 0x40 // User Data Header indicator set
	SM_REPLY_PATH_GSM     uint32 = 0x80 // Reply path set
	SM_UDH_REPLY_PATH_GSM uint32 = 0xC0 // Both UDH & Reply path

	// Optional Parameter Tags, Min and Max Lengths
	// Following are the 2 byte tag and min/max lengths for
	// supported optional parameter (declann)

	OPT_PAR_MSG_WAIT uint16 = 2

	// Privacy Indicator
	OPT_PAR_PRIV_IND uint16 = 0x0201

	// Source Subaddress
	OPT_PAR_SRC_SUBADDR     uint16 = 0x0202
	OPT_PAR_SRC_SUBADDR_MIN uint32 = 2
	OPT_PAR_SRC_SUBADDR_MAX uint32 = 23

	// Destination Subaddress
	OPT_PAR_DEST_SUBADDR     uint16 = 0x0203
	OPT_PAR_DEST_SUBADDR_MIN uint32 = 2
	OPT_PAR_DEST_SUBADDR_MAX uint32 = 23

	// User Message Reference
	OPT_PAR_USER_MSG_REF uint16 = 0x0204

	// User Response Code
	OPT_PAR_USER_RESP_CODE uint16 = 0x0205

	// Language Indicator
	OPT_PAR_LANG_IND uint16 = 0x020D

	// Source Port
	OPT_PAR_SRC_PORT uint16 = 0x020A

	// Destination Port
	OPT_PAR_DST_PORT uint16 = 0x020B

	// Concat Msg Ref Num
	OPT_PAR_SAR_MSG_REF_NUM uint16 = 0x020C

	// Concat Total Segments
	OPT_PAR_SAR_TOT_SEG uint16 = 0x020E

	// Concat Segment Seqnums
	OPT_PAR_SAR_SEG_SNUM uint16 = 0x020F

	// SC Interface Version
	OPT_PAR_SC_IF_VER uint16 = 0x0210

	// Display Time
	OPT_PAR_DISPLAY_TIME uint16 = 0x1201

	// Validity Information
	OPT_PAR_MS_VALIDITY uint16 = 0x1204

	// DPF Result
	OPT_PAR_DPF_RES uint16 = 0x0420

	// Set DPF
	OPT_PAR_SET_DPF uint16 = 0x0421

	// MS Availability Status
	OPT_PAR_MS_AVAIL_STAT uint16 = 0x0422

	// Network Error Code
	OPT_PAR_NW_ERR_CODE     uint16 = 0x0423
	OPT_PAR_NW_ERR_CODE_MIN uint32 = 3
	OPT_PAR_NW_ERR_CODE_MAX uint32 = 3

	// Extended int16 Message has no size limit

	// Delivery Failure Reason
	OPT_PAR_DEL_FAIL_RSN uint16 = 0x0425

	// More Messages to Follow
	OPT_PAR_MORE_MSGS uint16 = 0x0426

	// Message State
	OPT_PAR_MSG_STATE uint16 = 0x0427

	// Callback Number
	OPT_PAR_CALLBACK_NUM     uint16 = 0x0381
	OPT_PAR_CALLBACK_NUM_MIN uint32 = 4
	OPT_PAR_CALLBACK_NUM_MAX uint32 = 19

	// Callback Number Presentation  Indicator
	OPT_PAR_CALLBACK_NUM_PRES_IND uint16 = 0x0302

	// Callback Number Alphanumeric Tag
	OPT_PAR_CALLBACK_NUM_ATAG     uint16 = 0x0303
	OPT_PAR_CALLBACK_NUM_ATAG_MIN uint32 = 1
	OPT_PAR_CALLBACK_NUM_ATAG_MAX uint32 = 65

	// Number of messages in Mailbox
	OPT_PAR_NUM_MSGS uint16 = 0x0304

	// SMS Received Alert
	OPT_PAR_SMS_SIGNAL uint16 = 0x1203

	// Message Delivery Alert
	OPT_PAR_ALERT_ON_MSG_DELIVERY uint16 = 0x130C

	// ITS Reply Type
	OPT_PAR_ITS_REPLY_TYPE uint16 = 0x1380

	// ITS Session Info
	OPT_PAR_ITS_SESSION_INFO uint16 = 0x1383

	// USSD Service Op
	OPT_PAR_USSD_SER_OP uint16 = 0x0501

	// Priority
	SM_NOPRIORITY uint32 = 0
	SM_PRIORITY   uint32 = 1

	// Registered delivery
	//   SMSC Delivery Receipt (bits 1 & 0)
	SM_SMSC_RECEIPT_MASK          byte = 0x03
	SM_SMSC_RECEIPT_NOT_REQUESTED byte = 0x00
	SM_SMSC_RECEIPT_REQUESTED     byte = 0x01
	SM_SMSC_RECEIPT_ON_FAILURE    byte = 0x02
	//   SME originated acknowledgement (bits 3 & 2)
	SM_SME_ACK_MASK               byte = 0x0c
	SM_SME_ACK_NOT_REQUESTED      byte = 0x00
	SM_SME_ACK_DELIVERY_REQUESTED byte = 0x04
	SM_SME_ACK_MANUAL_REQUESTED   byte = 0x08
	SM_SME_ACK_BOTH_REQUESTED     byte = 0x0c
	//   Intermediate notification (bit 5)
	SM_NOTIF_MASK          byte = 0x010
	SM_NOTIF_NOT_REQUESTED byte = 0x000
	SM_NOTIF_REQUESTED     byte = 0x010

	// Replace if Present flag
	SM_NOREPLACE uint32 = 0
	SM_REPLACE   uint32 = 1

	// Destination flag
	SM_DEST_SME_ADDRESS uint32 = 1
	SM_DEST_DL_NAME     uint32 = 2

	// Higher Layer Message Type
	SM_LAYER_WDP  uint32 = 0
	SM_LAYER_WCMP uint32 = 1

	// Operation Class
	SM_OPCLASS_DATAGRAM    uint32 = 0
	SM_OPCLASS_TRANSACTION uint32 = 3

	// Originating MSC Address
	OPT_PAR_ORIG_MSC_ADDR     uint16 = 0x8081 // int16(0x8081)
	OPT_PAR_ORIG_MSC_ADDR_MIN uint32 = 1
	OPT_PAR_ORIG_MSC_ADDR_MAX uint32 = 24

	// Destination MSC Address
	OPT_PAR_DEST_MSC_ADDR     uint16 = 0x8082 // int16(0x8082)
	OPT_PAR_DEST_MSC_ADDR_MIN uint32 = 1
	OPT_PAR_DEST_MSC_ADDR_MAX uint32 = 24

	// Unused Tag
	OPT_PAR_UNUSED uint32 = 0xffff

	// Destination Address Subunit
	OPT_PAR_DST_ADDR_SUBUNIT uint16 = 0x0005

	// Destination Network Type
	OPT_PAR_DST_NW_TYPE uint16 = 0x0006

	// Destination Bearer Type
	OPT_PAR_DST_BEAR_TYPE uint16 = 0x0007

	// Destination Telematics ID
	OPT_PAR_DST_TELE_ID uint16 = 0x0008

	// Source Address Subunit
	OPT_PAR_SRC_ADDR_SUBUNIT uint16 = 0x000D

	// Source Network Type
	OPT_PAR_SRC_NW_TYPE uint16 = 0x000E

	// Source Bearer Type
	OPT_PAR_SRC_BEAR_TYPE uint16 = 0x000F

	// Source Telematics ID
	OPT_PAR_SRC_TELE_ID uint16 = 0x0010

	// QOS Time to Live
	OPT_PAR_QOS_TIME_TO_LIVE     uint16 = 0x0017
	OPT_PAR_QOS_TIME_TO_LIVE_MIN uint32 = 1
	OPT_PAR_QOS_TIME_TO_LIVE_MAX uint32 = 4

	// Payload Type
	OPT_PAR_PAYLOAD_TYPE uint16 = 0x0019

	// Additional Status Info Text
	OPT_PAR_ADD_STAT_INFO     uint16 = 0x001D
	OPT_PAR_ADD_STAT_INFO_MIN uint32 = 1
	OPT_PAR_ADD_STAT_INFO_MAX uint32 = 256

	// Receipted Message ID
	OPT_PAR_RECP_MSG_ID     uint16 = 0x001E
	OPT_PAR_RECP_MSG_ID_MIN uint32 = 1
	OPT_PAR_RECP_MSG_ID_MAX uint32 = 65

	// Message Payload
	OPT_PAR_MSG_PAYLOAD     uint16 = 0x0424
	OPT_PAR_MSG_PAYLOAD_MIN uint32 = 1
	OPT_PAR_MSG_PAYLOAD_MAX uint32 = 1500

	/**
	 * @deprecated As of version 1.3 of the library there are defined
	 * new encoding constants for base set of encoding supported by Java Runtime.
	 * The <code>CHAR_ENC</code> is replaced by <code>ENC_ASCII</code>
	 * and redefined in this respect.
	 */

	DFLT_MSGID         string = ""
	DFLT_MSG           string = ""
	DFLT_SRVTYPE       string = ""
	DFLT_SYSID         string = ""
	DFLT_PASS          string = ""
	DFLT_SYSTYPE       string = ""
	DFLT_ADDR_RANGE    string = ""
	DFLT_DATE          string = ""
	DFLT_ADDR          string = ""
	DFLT_MSG_STATE     byte   = 0
	DFLT_ERR           byte   = 0
	DFLT_SCHEDULE      string = ""
	DFLT_VALIDITY      string = ""
	DFLT_REG_DELIVERY         = SM_SMSC_RECEIPT_NOT_REQUESTED | SM_SME_ACK_NOT_REQUESTED | SM_NOTIF_NOT_REQUESTED
	DFLT_DFLTMSGID     byte   = 0
	DFLT_MSG_LEN       uint16 = 0
	DFLT_ESM_CLASS     byte   = 0
	DFLT_DATA_CODING   byte   = 0
	DFLT_PROTOCOLID    byte   = 0
	DFLT_PRIORITY_FLAG byte   = 0
	DFTL_REPLACE_IFP   byte   = 0
	DFLT_DL_NAME       string = ""
	DFLT_GSM_TON              = GSM_TON_UNKNOWN
	DFLT_GSM_NPI              = GSM_NPI_UNKNOWN
	DFLT_DEST_FLAG     byte   = 0 // not set
	MAX_PDU_LEN        uint32 = 5000

	PDU_HEADER_SIZE uint32 = 16 // 4 integers
	TLV_HEADER_SIZE uint32 = 4  // 2 int16s: tag & length

	// all times in milliseconds
	RECEIVER_TIMEOUT           int64 = 60000
	CONNECTION_RECEIVE_TIMEOUT int64 = 10000
	UNBIND_RECEIVE_TIMEOUT     int64 = 5000
	CONNECTION_SEND_TIMEOUT    int64 = 20000
	COMMS_TIMEOUT              int64 = 60000
	QUEUE_TIMEOUT              int64 = 10000
	ACCEPT_TIMEOUT             int64 = 60000

	RECEIVE_BLOCKING int64 = -1

	MAX_VALUE_PORT     uint32 = 65535
	MIN_VALUE_PORT     uint32 = 100
	MIN_LENGTH_ADDRESS uint32 = 7
)

var defaultTon = DFLT_GSM_TON
var defaultNpi = DFLT_GSM_NPI

var l sync.RWMutex

// SetDefaultTon ..
func SetDefaultTon(dfltTon byte) {
	l.Lock()
	defaultTon = dfltTon
	l.Unlock()
}

// GetDefaultTon ...
func GetDefaultTon() (r byte) {
	l.RLock()
	r = defaultTon
	l.RUnlock()
	return
}

// SetDefaultNpi ...
func SetDefaultNpi(dfltNpi byte) {
	l.Lock()
	defaultNpi = dfltNpi
	l.Unlock()
}

// GetDefaultNpi ...
func GetDefaultNpi() (r byte) {
	l.RLock()
	r = defaultNpi
	l.RUnlock()
	return
}

// GetCurrentTime ...
func GetCurrentTime() time.Time {
	return time.Now()
}
