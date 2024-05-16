package common

//at compile time, fill this var
var versionInfo string

func GetVersion() string {
	return versionInfo
}
