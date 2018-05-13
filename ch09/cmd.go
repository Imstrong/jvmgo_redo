package main

import (
	"flag"
	"fmt"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	jreOption   string
	class       string
	args        []string
	verboseClassFlag bool
	verboseInstFlag bool
}

func ParseCmd() (cmd *Cmd) {
	cmd = &Cmd{}
	flag.Usage=PrintUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version message")
	flag.BoolVar(&cmd.verboseClassFlag, "verboseClass", false, "print version message")
	flag.BoolVar(&cmd.verboseInstFlag, "verboseInst", false, "print version message")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version message")
	flag.StringVar(&cmd.jreOption,"jre","","jrepath")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if (len(args) > 0) {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
func PrintUsage() {
	fmt.Print("usage:ch01.exe [-v|version|help|?|classpath] <params>")
}
