package config

// Config struct.
type Config struct {
	version string
	appid   string
	key     string
}

// ConfigIns is ins of config.
var configIns Config

// init package config.
func init() {
	configIns.version = "0.01"
	configIns.appid = "20151203000007279"
	configIns.key = "D5hPGkfdOipB12SK5ljx"
}

// GetVersion of this software.
func GetVersion() (version string) {
	version = configIns.version
	return
}

// GetAppid : get appid .
func GetAppid() (appid string) {
	appid = configIns.appid
	return
}

// GetKey : get key.
func GetKey() (key string) {
	key = configIns.key
	return
}
