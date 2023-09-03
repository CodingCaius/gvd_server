package flags

import (
	"gvd_server/global"
)

func Port(port int) {

	global.Config.System.Port = port
}
