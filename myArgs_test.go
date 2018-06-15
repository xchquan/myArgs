/**
 * Created by huQg on 2018/6/15,015.
 */

package myArgs

import (
	"fmt"
	"strconv"
	"testing"
)

func TestArgs(t *testing.T) {
	argsParser := CreateArgsParser()
	argsParser.ConfigArgs("--help", "-h", "help", true)
	argsParser.ConfigArgs("--version", "-v", "version for args", true)
	argsParser.ConfigArgs("--port", "-p", "special port for config")
	argsParser.ConfigArgs("--ip", "-ip", "ip for config")

	//var sArgs = []string{"", "--a", "b", "--help", "--version"}
	var sArgs = []string{"", "--version", "-p", "6379", "-ip", "10.15.43.226"}

	argsParser.ParseArgs(sArgs)

	DealargsFunc := func(sKey string, doFunc dealargs) {
		err := argsParser.DealArgs(sKey, func(sVal string, bOnlyFlag bool) error {
			fmt.Printf("key[%s], val[%s], onlyflag[%v]\n", sKey, sVal, bOnlyFlag)

			if doFunc != nil {
				doFunc(sVal, bOnlyFlag)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("key[%s] error[%v]\n", sKey, err)
		}
	}

	DealargsFunc("-h", nil)

	var iPort int = 9001
	var sIp string = "127.0.0.1"
	DealargsFunc("-p", func(sVal string, bOnlyFlag bool) error {
		if !bOnlyFlag {
			iPort, _ = strconv.Atoi(sVal)
		}
		return nil
	})
	fmt.Printf("port[%d]\n", iPort)

	DealargsFunc("-ip", func(sVal string, bOnlyFlag bool) error {
		if !bOnlyFlag {
			sIp = sVal
		}
		return nil
	})
	fmt.Printf("ip[%s]\n", sIp)
}
