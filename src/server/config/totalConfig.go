package config

import (
	"fmt"
)

var StringInfoMgr ConfigMap
var WorldLevelInfoMgr ConfigMap

func init() {
	fmt.Println("totalConfig init")
	var basePath = getCurrentDirectory() + "/../../config/excel/"
	StringInfoMgr.LoadConfigByFilePath(basePath + "String.txt")
	WorldLevelInfoMgr.LoadConfigByFilePath(basePath + "WorldLevelInfo.txt")
}
