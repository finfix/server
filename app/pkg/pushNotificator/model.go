package pushNotificator

type PushReq struct {
	Notification      NotificationSettings
	NotificationToken string
	BundleID          string
}

type NotificationSettings struct {
	Title    *string
	Subtitle *string
	Message  *string
	Badge    *uint8
}
