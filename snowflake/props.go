package snowflake

import (
	"fmt"
	"os"
	"web-consul-ui/app"
)

/**********************************
 * Date: 2023/2/2
 * Author: hchery
 * Home: https://github.com/hchery
 *********************************/

type Props struct {
	Machine Machine `yaml:"machine"`
}

type Machine struct {
	Id int64 `yaml:"id"`
}

var props = new(Props)

func init() {
	app.YamlUnmarshal(app.SnowflakeConf, props)
	machineId := props.Machine.Id
	if machineId < 0 || machineId > maxMachineID {
		fmt.Printf("Invalid snowflake machine id: %d", machineId)
		os.Exit(app.ExitWithSnowflakeMachineIdError)
	}
}
