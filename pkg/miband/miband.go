// Package miband provides all the requried protocols and functions to
// interact and work with MiBand 2 and MiBand 3
package miband

// UUIDBase is to get UUID Base for x
func uuidBase(x string) string {
	return "0000" + x + "-0000-3512-2118-0009af100700"
}

// The UUID data for MiBand is fetched using the nRF Connect band on android
const (
	/********************** PRIMARY SERVICE *************************/
	// UUIDServiceGenericAccess is the uint16 UUID to access service
	UUIDServiceGenericAccess = 0x1800
	// UUIDServiceGenericAttribute is the uint16 UUID to access attribute service
	UUIDServiceGenericAttribute = 0x1801
	// UUIDServiceDeviceInformation is the uint16 UUID to get device information
	UUIDServiceDeviceInformation = 0x180a
	// UUIDServiceFirmware to access device firmware
	UUIDServiceFirmware = "00001530-0000-3512-2118-0009af100700"
	// UUIDServiceAlertNotification is the uint16 UUID to send notification
	UUIDServiceAlertNotification = 0x1811
	// UUIDServiceImmediateAlert is the uint16 UUID to send immediate alert
	UUIDServiceImmediateAlert = 0x1802
	// UUIDServiceHeartRate is the uint16 UUID to get heart rate
	UUIDServiceHeartRate = 0x180d
	// UUIDServiceMiband1 is the unknown service UUID
	UUIDServiceMiband1 = 0xfee0
	// UUIDServiceband2 is the unknown service UUID
	UUIDServiceMiband2 = 0xfee1

	/************************* SECONDARY SERVICE *****************************/
	// UUIDServiceGenericDeviceName is the uint16 UUID to READ device name
	UUIDServiceGenericDeviceName = 0x2a00
	// UUIDServiceGenericAppearance is the uint16 UUID to READ appearance of band
	UUIDServiceGenericAppearance = 0x2a01
	// UUIDServiceGenericPeripheralPreferredConnectionParameter is the uint16 to READ peripheral preferred connection parameter
	UUIDServiceGenericPeripheralPreferredConnectionParameter = 0x2a04

	// UUIDServiceGenericAttributeChanged is the uint16 UUID to INDICATE, READ service change
	UUIDServiceGenericAttributeChanged = 0x2a05
	// UUIDServiceGenericAttributeChangedClient is the uint16 UUID for client characteristic configuration
	UUIDServiceGenericAttributeChangedClient = 0x2902

	// UUIDServiceDeviceInformationSerial is the uint16 UUID to READ serial number
	UUIDServiceDeviceInformationSerial = 0x2a25
	// UUIDServiceDeviceInformationHardwareRevision is the uint16 UUID to READ hardware revision
	UUIDServiceDeviceInformationHardwareRevision = 0x2a27
	// UUIDServiceDeviceInformationSoftwareRevision is the uint16 UUID to READ software revision
	UUIDServiceDeviceInformationSoftwareRevision = 0x2a28
	// UUIDServiceDeviceInformationSystemID is the uint16 UUID to READ system id
	UUIDServiceDeviceInformationSystemID = 0x2a23
	// UUIDServiceDeviceInformationPnpID is the uint16 UUID to READ PnP ID
	UUIDServiceDeviceInformationPnpID = 0x2a50

	// UUIDServiceAlertNotificationNew is the uint16 UUID to WRITE notification
	UUIDServiceAlertNotificationNew = 0x2a46
	// UUIDServiceAlertNotificationNewClient is the uint16 UUID for WRITE client characteristic configuration
	UUIDServiceAlertNotificationNewClient = 0x2901
	// UUIDServiceAlertNotificationControlPoint is the uint16 UUID to WRITE, READ, NOTIFY notification control point
	UUIDServiceAlertNotificationControlPoint = 0x2a44
	// UUIDServiceAlertNotificationControlPointClient is the uint16 UUID configure client characteristic
	UUIDServiceAlertNotificationControlPointClient = 0x2902

	// UUIDServiceImmediateAlertLevel is the uint16 UUID to WRITE NO RESPONSE but alert
	UUIDServiceImmediateAlertLevel = 0x2a06

	// UUIDServiceHeartRateMeasurement is the uint16 UUID to NOTIFY heart rate
	UUIDServiceHeartRateMeasurement = 0x2a37
	// UUIDServiceHeartRateMeasurementClient is the uint16 UUID to NOTIFY heart rate client
	UUIDServiceHeartRateMeasurementClient = 0x2902
	// UUIDServiceHeartRateControlPoint is the uint16 UUID to READ, WRITE heart rate control point
	UUIDServiceHeartRateControlPoint = 0x2a39
)
