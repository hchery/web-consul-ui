package main

import "web-consul-ui/snowflake"

/**********************************
 * Date: 2023/2/2
 * Author: hchery
 * Home: https://github.com/hchery
 *********************************/

func main() {
	worker, _ := snowflake.NewWorker(31)
	println(worker.Next())
}
