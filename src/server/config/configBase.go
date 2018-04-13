package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type EConfigContextType int

const (
	_ EConfigContextType = iota
	EConfigContextType_Int
	EConfigContextType_String
)

type ConfigCell interface {
	GetAsInt() int
	GetAsString() string
	ToString() string
	ReadByString(string) bool
}
type ConfigCellString struct {
	context string
}

type ConfigCellInt struct {
	context int
}

func (self *ConfigCellInt) GetAsInt() int {
	return self.context
}
func (self *ConfigCellInt) GetAsString() string {
	panic("can not get int field as string")
	return self.ToString()
}
func (self *ConfigCellInt) ToString() string {
	return strconv.Itoa(self.context)
}
func (self *ConfigCellInt) ReadByString(val string) bool {
	num, err := strconv.Atoi(val)
	if err != nil {
		panic("error: config string read error." + val + "is not a valid number")
		return false
	}
	self.context = num
	return true
}
func (self *ConfigCellString) GetAsInt() int {
	panic("can not get string field as int")
	return 0
}
func (self *ConfigCellString) GetAsString() string {
	return self.context
}
func (self *ConfigCellString) ToString() string {
	return self.context
}
func (self *ConfigCellString) ReadByString(val string) bool {
	self.context = val
	return true
}

type ConfigMap struct {
	infos map[int](map[string]ConfigCell)
}

func (self *ConfigMap) GetInfo(id int) map[string]ConfigCell {
	return self.infos[id]
}
func (self *ConfigMap) LoadConfigByFilePath(filePath string) bool {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("config string read error" + err.Error())
		return false
	}

	return self.LoadConfig(string(file))
}
func (self *ConfigMap) LoadConfig(context string) bool {
	infos := make(map[int]map[string](ConfigCell))
	lines := strings.Split(context, "\r\n")
	count := 0
	var fields []string
	for _, line := range lines {
		if line == "" { // 空行
			continue
		}
		count++
		vals := strings.Split(line, "\t")
		if count == 1 {
			// 字段行
			fields = vals
			continue
		}
		info := make(map[string](ConfigCell))
		id, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("error: id[" + vals[0] + "] is not a valid number")
			return false
		}
		for index, fieldName := range fields {
			tp := fieldName[1:1]
			var cell ConfigCell
			if tp == "i" {
				cell = new(ConfigCellInt)
			} else {
				cell = new(ConfigCellString)
			}
			if !cell.ReadByString(vals[index]) {
				return false
			}
			info[fieldName] = cell
		}
		infos[id] = info
	}
	self.infos = infos
	return true
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir //strings.Replace(dir, "\\", "/", -1)
}
