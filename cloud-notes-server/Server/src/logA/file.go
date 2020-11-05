package logA

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写信息

// FileLogger 文件日志结构体
type FileLogger struct {
	level    Level
	FileName string
	FilePath string
	file     *os.File
	errFile  *os.File
	maxSize  int64
}

//构造函数
func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl :=
		&FileLogger{
			level:    logLevel,
			FileName: fileName,
			FilePath: filePath,
			maxSize:  1024,
		}
	fl.initFile()
	return fl
}

//初始化函数
func (f *FileLogger) initFile() {
	logName := path.Join(f.FilePath, f.FileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open logger_file: %s faile,%v", logName, err))
	}
	f.file = fileObj
	//错误日志文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open logger_file: %s faile,%v", errLogName, err))
	}
	f.errFile = errFileObj
}

//检测是否切分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize
}

//日志切分方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	//检测日志大学是否超过maxSize

	//切分文件
	fileName := file.Name()
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	//把源文件关闭
	file.Close()
	//备份原来的文件
	os.Rename(fileName, backupName)
	//创建一个新文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("open file faile"))
	}
	return fileObj

}

//日志输出方法（模板）
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	levelstr := getLevelStr(level)
	msg := fmt.Sprintf(format, args...)
	//日志格式：【时间】【文件:行号】【函数名】【日志级别】日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName, levelstr, msg)
	if f.checkSplit(f.file) {
		f.file = f.splitLogFile(f.file)
	}
	fmt.Fprintln(f.file, logMsg)
	if level >= ErrorLevel {
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

//日志Debug输出方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

//日志Info输出方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

//日志Warn输出方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarnLevel, format, args...)
}

//日志Error输出方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

//日志Fatal输出方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

//Close关闭日志文件
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
