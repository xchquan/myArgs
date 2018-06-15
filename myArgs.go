/**
 * Created by huQg on 2018/6/14,014.
 */

package myArgs

import (
	"fmt"
	"os"
	"strings"
)

const (
	VERSION = "1.0.0"
)

type tgf_set_args struct {
	sKey      string
	sSign     string
	sTag      string
	bOnlyflag bool
}

type tgf_val_args struct {
	sKey      string
	setArgs   *tgf_set_args
	sVal      string
	bOnlyflag bool
}

type Fc_tgc_args struct {
	lstSetArgs []*tgf_set_args
	mapSetArgs map[string]*tgf_set_args
	mapCfgArgs map[string]*tgf_val_args
}

func CreateArgsParser() *Fc_tgc_args {
	res := Fc_tgc_args{}
	res.mapSetArgs = make(map[string]*tgf_set_args)
	res.mapCfgArgs = make(map[string]*tgf_val_args)
	return &res
}

func (s *Fc_tgc_args) findKey(skey string) (*tgf_set_args, error) {
	if args, ok := s.mapSetArgs[skey]; ok {
		return args, nil
	} else {
		return nil, fmt.Errorf("key[%s] isn't exist", skey)
	}
}

func (s *Fc_tgc_args) findVal(sKey string) (*tgf_val_args, error) {
	if args, ok := s.mapCfgArgs[sKey]; ok {
		return args, nil
	} else {
		return nil, fmt.Errorf("key[%s] isn't exist", sKey)
	}
}

func (s *Fc_tgc_args) saveSet(sKey string, args *tgf_set_args) error {
	s.mapSetArgs[sKey] = args
	return nil
}

func (s *Fc_tgc_args) saveCfg(skey string, args *tgf_val_args) error {
	s.mapCfgArgs[skey] = args
	return nil
}

/// print Help
func (s *Fc_tgc_args) printfHelp(skey string) {

	if !strings.EqualFold(skey, "-h") && !strings.EqualFold(skey, "--help") {
		return
	}

	fmt.Printf("parse args, version[%s]\n", VERSION)
	fmt.Println("parse args [options] ...")
	fmt.Println("Options:")

	var iTol = 30
	for _, v := range s.lstSetArgs {
		if v.bOnlyflag {
			str := fmt.Sprintf(" %s,%s", v.sKey, v.sSign)
			for len(str) > iTol {
				iTol = iTol + 5
			}
			fmt.Printf("%s%s%s\n", str, strings.Repeat(" ", iTol-len(str)), v.sTag)
		} else {
			str := fmt.Sprintf(" %s,%s    <argv>    ", v.sKey, v.sSign)
			for len(str) > iTol {
				iTol = iTol + 5
			}
			fmt.Printf("%s%s%s\n", str, strings.Repeat(" ", iTol-len(str)), v.sTag)
		}
	}
	fmt.Println("Options for ArgsParse finished")
	os.Exit(0)
}

/// print version
func (s *Fc_tgc_args) printfVersion(skey string) {
	if !strings.EqualFold(skey, "-v") && !strings.EqualFold(skey, "--version") {
		return
	}

	fmt.Printf("version[%s]\n", VERSION)
	os.Exit(0)
}

/// args set
func (s *Fc_tgc_args) ConfigArgs(sKey string, sSign string, sTag string, bOnlyFlag ...bool) {
	var bFlag bool = false
	for _, v := range bOnlyFlag {
		bFlag = v
		break
	}

	if len(sKey) == 0 {
		return
	}
	strRes := strings.Split(sKey, ",")

	if len(sSign) > 0 {
		sRes := strings.Split(sSign, ",")
		strRes = append(strRes, sRes...)
	}
	sargs := &tgf_set_args{
		sKey:      sKey,
		sSign:     sSign,
		sTag:      sTag,
		bOnlyflag: bFlag,
	}

	for _, v := range strRes {
		s.saveSet(v, sargs)
	}
	s.lstSetArgs = append(s.lstSetArgs, sargs)
}

/// parse args
func (s *Fc_tgc_args) ParseArgs(sArgs []string) {
	if s == nil {
		return
	}

	if len(sArgs) <= 1 {
		return
	}

	args := sArgs[1:len(sArgs)]
	var bfindv = false
	var vargs *tgf_val_args
	for _, v := range args {
		res, err := s.findKey(v)
		if err != nil {
			if bfindv {
				vargs.sVal = v
				s.saveCfg(vargs.sKey, vargs)
				bfindv = false
			} else {
				fmt.Println(err)
			}
		} else {
			vargs = &tgf_val_args{
				sKey:      v,
				setArgs:   res,
				bOnlyflag: res.bOnlyflag,
			}
			if res.bOnlyflag {
				s.saveCfg(v, vargs)
				bfindv = false
			} else {
				bfindv = true
			}
		}
	}
}

type dealargs func(sVal string, bOnlyFlag bool) error

/// get cfg and deal by call custom func dealargs
func (s *Fc_tgc_args) DealArgs(sKey string, dargs dealargs) error {

	s.printfHelp(sKey)
	s.printfVersion(sKey)

	cfg, err := s.findVal(sKey)
	if err != nil {
		return err
	}
	return dargs(cfg.sVal, cfg.bOnlyflag)
}
