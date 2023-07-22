package utils

import (
	"flag"
	"fmt"
	"os"
)

type argMap struct {
	version *bool
	force   *bool
	extra   []string
}

func getArgs() argMap {
	args := argMap{}

	args.version = flag.Bool("v", false, "Display mkpath version")
	args.force = flag.Bool("f", false, "Force mkpath operation")
	flag.Parse()

	args.extra = flag.Args()

	return args
}

func ApplyFlags() {
	args := getArgs()
	if *args.version {
		fmt.Printf("mkpath: Ver. %s\n", MKPATH_VERSION)
		os.Exit(OS_CODES["NO_ERROR"])
	}

	if len(args.extra) >= 1 {
		filePath := args.extra[0]
		if pathExists(filePath) && !*args.force {
			fmt.Printf("mkpath-ERR: Creating: %s. File already exists.\n", filePath)
			fmt.Println("mkpath-ERR: To force overwrite, please run using -f flag.")
			os.Exit(OS_CODES["FILE_ALREADY_EXISTS"])
		}
		createFile(filePath)
	}
}
