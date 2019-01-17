/*
*Auth: JackColor
*Date: 2019/1/13 下午4:35
*/
package logging

type LogConsole struct {
	Level int
}

func NewLogConsole(conf map[string]string) (Logger LoggerInterface, err error) {

	return
}

func (f *LogConsole) Init() {

}

func (f *LogConsole) SetLevel(level int) {

}

func (f *LogConsole) LogDebug(format string,args... interface{}){

}

func (f *LogConsole) LogTrace(format string,args... interface{}) {

}
func (f *LogConsole) LogInfo(format string,args... interface{}){

}
func (f *LogConsole) LogWarn(format string,args... interface{}) {

}
func (f *LogConsole) LogErrors(format string,args... interface{}){

}

func (f *LogConsole) LogFatal(format string,args... interface{}) {

}
