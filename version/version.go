package version

// 打印当前的编译信息

import (
	"fmt"
)

var (
	BuildDateTime      = ""
	BuildGitBranch     = ""
	BuildGitCommit     = ""
	BuildPackageModule = ""
)

var BuildInfo = PkgVersion{}

func init() {
	BuildInfo = PkgVersion{
		BuildGitBranch:     BuildGitBranch,
		BuildDateTime:      BuildDateTime,
		BuildGitCommit:     BuildGitCommit,
		BuildPackageModule: BuildPackageModule,
	}
}

type PkgVersion struct {
	BuildDateTime      string
	BuildGitBranch     string
	BuildGitCommit     string
	BuildPackageModule string
}

func (r PkgVersion) String() string {
	return r.BuildDateTime + "-" + r.BuildGitCommit
}

func (m PkgVersion) PrintVersion() {
	fmt.Printf("%15s: %s \n", "BuildTime", BuildDateTime)
	fmt.Printf("%15s: %s \n", "BuildGitBranch", BuildGitBranch)
	fmt.Printf("%15s: %s \n", "BuildGitCommit", BuildGitCommit)
	fmt.Printf("%15s: %s \n", "BuildPkgModule", BuildPackageModule)
	fmt.Println()
}

func PrintVersion() {
	BuildInfo.PrintVersion()
}

// -------------------------------- main 调用代码---------------------------------------
//package main
//
//import (
//"flag"
//"fmt"
//
//"wps.ktkt.com/debug/debug/version"
//)
//
//var printVersion = flag.Bool("v", false, "show build version for the program")
//
//func main() {
//	flag.Parse()
//
//	if *printVersion {
//		version.PrintVersion()
//	}
//
//	fmt.Println("test ------")
//}
// ** -------------------------------- main 调用代码 END---------------------------------------

// -------------------------------- build bash 代码---------------------------------------
//#!/usr/bin/env bash
////
//moduleName=$(go list -m)
//commit=$(git rev-parse --short HEAD)
//branch=$(git rev-parse --abbrev-ref HEAD)
//buildTime=$(date +%Y%m%d%H%M)
//
////versionPkgName="version"
//
//flags="-X '${moduleName}/${versionPkgName}.BuildGitBranch=${branch}' -X '${moduleName}/${versionPkgName}.BuildGitCommit=${commit}'  -X '${moduleName}/${versionPkgName}.BuildDateTime=${buildTime}' -X '${moduleName}/${versionPkgName}.BuildPackageModule=${moduleName}' "
//
//program=$(basename ${moduleName})
//
//#echo "$program"
//#
//#echo "${flags}"
//#
//#echo "\$(pwd)"
//#
//
////go build -o bin/$program -ldflags "$flags" -v -mod=vendor  ./cmd/version/main.go
//
// ** -------------------------------- build bash 代码 END---------------------------------------
