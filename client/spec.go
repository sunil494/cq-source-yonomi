package client

type Spec struct {
	Devices       []YonomiDeviceConfigBlock `json:"devices"`
	Authorization string                    `json:"authorization"`
}

type YonomiDeviceConfigBlock struct {
	DeviceId string `json:"deviceId"`
	Name     string `json:"name"`
}
