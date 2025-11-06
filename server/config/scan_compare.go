package config

type ScanCompare struct {
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 扫码枪配对过期时间
	EioIp       string `mapstructure:"eio-ip" json:"eio-ip" yaml:"eio-ip"`                   // io设备ip
	EioPort     int    `mapstructure:"eio-port" json:"eio-port" yaml:"eio-port"`             // io设备端口
}
