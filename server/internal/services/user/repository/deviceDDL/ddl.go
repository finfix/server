package deviceDDL

import "server/internal/ddl"

const (
	Table          = ddl.SchemaCoin + "." + "devices"
	TableWithAlias = Table + " " + alias
	alias          = "d"
)

const (
	ColumnID                  = "id"
	ColumnRefreshToken        = "refresh_token"
	ColumnDeviceID            = "device_id"
	ColumnUserID              = "user_id"
	ColumnNotificationToken   = "notification_token"
	ColumnApplicationBundleID = "application_bundle_id"
	ColumnDeviceIPAddress     = "device_ip_address"
	ColumnDeviceUserAgent     = "device_user_agent"
	ColumnDeviceOSName        = "device_os_name"
	ColumnDeviceOSVersion     = "device_os_version"
	ColumnDeviceName          = "device_name"
	ColumnDeviceModelName     = "device_model_name"
	ColumnApplicationVersion  = "application_version"
	ColumnApplicationBuild    = "application_build"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
