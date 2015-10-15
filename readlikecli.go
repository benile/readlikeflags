package readlikeflags
import (
"github.com/benile/cli"
"github.com/chzyer/readline"
	"strings"
)
var on bool=true
func Exit(){
	on=false
}
type Options struct{
	Commands []cli.Command
	ReadlineConfig *readline.Config
	Usage string
	Version string
	ErrorHandler func(err error)
	AppName string
}
func StartSession(options Options){
	commands:=options.Commands
	readlineConfig:=options.ReadlineConfig
	usage:=options.Usage
	version:=options.Version
	items:=make([] *readline.PrefixCompleter,0)
	items=append(items,readline.PcItem("help"));
	for _,cmd:=range commands{
		items=append(items,readline.PcItem(cmd.Name))
	}
	var completer = readline.NewPrefixCompleter()
	completer.Children=items
	if readlineConfig==nil{
		readlineConfig=&readline.Config{
			Prompt:">",
		}
	}
	readlineConfig.AutoComplete=completer
	rl, err := readline.NewEx(readlineConfig)
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	app:=cli.NewApp()
	app.Usage=usage
	app.Version=version
	app.Commands=commands
	app.BuildinApp=true
	app.Name=options.AppName
	startLoop(app,rl, options.ErrorHandler)

}
func startLoop(app *cli.App,rl *readline.Instance ,errorHandler func(err error)){
	for on==true{
		defer func() {
			if r := recover(); r != nil {
				if err,ok:=r.(error);ok{
					errorHandler(err)
					startLoop(app,rl,errorHandler)
				}

			}
		}()
		line, err := rl.Readline()
		if err != nil {
			break
		}
		params:=strings.Split(line," ")
		args:=[]string{app.Name}
		for _,param:=range params{
			trimed:=strings.TrimSpace(param)
			if trimed!=""{
				args=append(args,trimed)
			}
		}
		if len(args) >1 {
			app.RunInside(args)
		}
	}
}