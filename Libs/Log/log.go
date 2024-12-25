package Log

import (
	"Libs/DateTime"
	"Libs/Enum"
	"Libs/Ext"

	"fmt"
	"runtime"
	"strconv"
	"time"
)

type logValue int

const (
	UsePath = true
	NoPath  = false
)

type logLv struct {
	v     logValue
	color string
}

func (f logValue) FlagMap() map[int]string {
	return map[int]string{
		1: "Info",
		2: "Debug",
		3: "Warn",
		4: "Error",
	}
}

var (
	lv_info = logLv{
		v:     1,
		color: "\033[97m",
	}
	lv_debug = logLv{
		v:     2,
		color: "\033[94m",
	}
	lv_warn = logLv{
		v:     3,
		color: "\033[33m",
	}
	lv_error = logLv{
		v:     4,
		color: "\033[91m",
	}
)

var (
	Info = struct {
		Print  func(...any)
		Printf func(string, ...any)
		Printp func(...any)
	}{
		Print: func(p ...any) {
			I(fmt.Sprint(p...))
		},
		Printf: func(temp string, p ...any) {
			I(fmt.Sprintf(temp, p...))
		},
		Printp: func(p ...any) {
			I(fmt.Sprint(p...), UsePath)
		},
	}

	Debug = struct {
		Print  func(...any)
		Printf func(string, ...any)
		Printp func(...any)
	}{
		Print: func(p ...any) {
			D(fmt.Sprint(p...))
		},
		Printf: func(temp string, p ...any) {
			D(fmt.Sprintf(temp, p...))
		},
		Printp: func(p ...any) {
			D(fmt.Sprint(p...), UsePath)
		},
	}

	Warn = struct {
		Print  func(...any)
		Printf func(string, ...any)
		Printp func(...any)
	}{
		Print: func(p ...any) {
			W(fmt.Sprint(p...))
		},
		Printf: func(temp string, p ...any) {
			W(fmt.Sprintf(temp, p...))
		},
		Printp: func(p ...any) {
			W(fmt.Sprint(p...), UsePath)
		},
	}

	Error = struct {
		Print  func(...any)
		Printf func(string, ...any)
		Printp func(...any)
	}{
		Print: func(p ...any) {
			E(fmt.Sprint(p...))
		},
		Printf: func(temp string, p ...any) {
			E(fmt.Sprintf(temp, p...))
		},
		Printp: func(p ...any) {
			E(fmt.Sprint(p...), UsePath)
		},
	}
)

/*
param[0]:bool *print runtime file path*
*/
func I(msg string, param ...any) {
	lv := lv_info
	msg = fmt.Sprintf("%s[%s]\033[0m %s", lv.color, Enum.ToString(lv.v), msg)
	if len(param) < 1 {
		logBuild2(msg)
		return
	}

	if path := param[0].(bool); path {
		_, filename, line, _ := runtime.Caller(1)
		logBuild2WithRuntime(msg, filename, line)
	} else {
		logBuild2(msg)
	}
}

/*
param[0]:bool *print runtime file path*
*/
func D(msg string, param ...any) {
	lv := lv_debug
	msg = fmt.Sprintf("%s[%s]\033[0m %s", lv.color, Enum.ToString(lv.v), msg)
	if len(param) < 1 {
		logBuild2(msg)
		return
	}

	if path := param[0].(bool); path {
		_, filename, line, _ := runtime.Caller(1)
		logBuild2WithRuntime(msg, filename, line)
	} else {
		logBuild2(msg)
	}
}

/*
param[0]:bool *print runtime file path*
*/
func W(msg string, param ...any) {
	lv := lv_warn
	msg = fmt.Sprintf("%s[%s]\033[0m %s", lv.color, Enum.ToString(lv.v), msg)
	if len(param) < 1 {
		logBuild2(msg)
		return
	}

	if path := param[0].(bool); path {
		_, filename, line, _ := runtime.Caller(1)
		logBuild2WithRuntime(msg, filename, line)
	} else {
		logBuild2(msg)
	}
}

/*
param[0]:bool *print runtime file path*
param[0]:error *catch error*
param[1]:bool *print runtime file path*
*/
func E(msg string, param ...any) {
	lv := lv_error
	msg = fmt.Sprintf("%s[%s]\033[0m %s", lv.color, Enum.ToString(lv.v), msg)
	if len(param) < 1 {
		logBuild2(msg)
		return
	}

	p1 := param[0]

	if b := Ext.TypeEqual[bool](p1); b {
		if path := p1.(bool); path {
			_, filename, line, _ := runtime.Caller(1)
			logBuild2WithRuntime(msg, filename, line)
		} else {
			logBuild2(msg)
		}
		return
	}

	e := p1.(error)
	msg = fmt.Sprintf("%s\033[91m Error:%s\033[0m", msg, e.Error())

	if v, e := Ext.TryGet(param, 1); e == nil {
		//fmt.Println("Error p2=", v)
		if v.(bool) {
			_, filename, line, _ := runtime.Caller(1)
			logBuild2WithRuntime(msg, filename, line)
		} else {
			logBuild2(msg)
		}
	} else {
		logBuild2(msg)
	}
}

func logBuild2(msg string) {
	timeMask := time.Now().Format(DateTime.Format)
	msg = fmt.Sprintf("%s %s", timeMask, msg)

	// 勿刪
	fmt.Println(msg)
}
func logBuild2WithRuntime(msg string, filename string, line int) {
	timeMask := time.Now().Format(DateTime.Format)

	msg = fmt.Sprintf("%s %s at \033[35m%s:%s\033[0m", timeMask, msg, filename, strconv.Itoa(line))
	//errMsg := timeMask + " " + level
	// if Config.ModeLog == Config.CML_Dev {
	//errMsg += " \033[35m" + filename + ":" + strconv.Itoa(line) + "\033[0m =>\n"
	// } else {
	// 	errMsg += " "
	// }
	// errMsg += string(msg)
	// errMsg += msgStr

	// 勿刪
	fmt.Println(msg)
}
