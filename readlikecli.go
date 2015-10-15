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
func StartSession(commands []cli.Command,readlineConfig *readline.Config,usage string,version string){
	items:=make([] *readline.PrefixCompleter,10)
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
	for on==true{
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
