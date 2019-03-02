package mlog

import (
	"fmt"
	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

// load_console 将日志输出到console(Debug)
func load_console() {
	appConfig := `
<seelog minlevel="debug">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time %EscM(36)[%LEV]%EscM(49)%EscM(0) [%File:%Line] %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_console_trace 将日志输出到console (Trace)
func load_console_trace() {

	appConfig := `
<seelog minlevel="trace">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time %EscM(36)[%LEV]%EscM(49)%EscM(0) [%File:%Line] %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_console_info 将日志输出到console (Info)
func load_console_info() {

	appConfig := `
<seelog minlevel="info">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time %EscM(36)[%LEV]%EscM(49)%EscM(0) %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_console_err 将日志输出到console(Error)
func load_console_err() {
	appConfig := `
<seelog minlevel="error">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time %EscM(36)[%LEV]%EscM(49)%EscM(0) %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_trace_nofile 将日志输出到 console, 不显示文件名和行号(trace).
func load_trace_nofile() {
	appConfig := `
<seelog minlevel="trace">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time [%LEV] %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_debug_nofile 将日志输出到 console, 不显示文件名和行号(debug).
func load_debug_nofile() {
	appConfig := `
<seelog minlevel="debug">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time [%LEV] %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_info_nofile 将日志输出到 console, 不显示文件名和行号(Info).
func load_info_nofile() {
	appConfig := `
<seelog minlevel="info">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time [%LEV] %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// load_err_nofile 将日志输出到 console, 不显示文件名和行号(Error).
func load_err_nofile() {
	appConfig := `
<seelog minlevel="error">
    <outputs formatid="colored">
        <console  />
    </outputs>
    <formats>
        <format id="colored" format="%Date %Time [%LEV] %Msg%n%EscM(0)" />
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// SetLogWriter 设置日志配置参数，输入为一个文件地址.
func SetLogWriter(conf_file string) {
	SetLogConf(conf_file)
}

// SetLogString 设置日志显示，配置从字符串读入.
func SetLogString(config_str string) {
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(config_str))
	if err != nil {
		load_trace_nofile()
		return
	}
	UseLogger(logger)
}

// SetLogConf 设置日志记录方式，配置从传入的文件路径读入.
func SetLogConf(conf_file string) {
	logger, err := seelog.LoggerFromConfigAsFile(conf_file)
	if err != nil {
		load_console()
		return
	}
	UseLogger(logger)
}

/*
设置Console显示的日志格式，level支持:
debug, info, err, trace, debug_nofile, info_nofile, err_nofile, trace_nofile
*/
func SetLogConsole(level string) {
	if level == "debug" {
		load_console()
	} else if level == "info" {
		load_console_info()
	} else if level == "err" {
		load_console_err()
	} else if level == "trace" {
		load_console_trace()
	} else if level == "trace_nofile" {
		load_trace_nofile()
	} else if level == "debug_nofile" {
		load_debug_nofile()
	} else if level == "info_nofile" {
		load_info_nofile()
	} else if level == "err_nofile" {
		load_err_nofile()
	}
}

// DisableLog 禁止日志使用.
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger 设置日志使用对象
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

// FlushLog 刷新日志.
func FlushLog() {
	Logger.Flush()
}

func init() {
	DisableLog()
	load_console()
}
