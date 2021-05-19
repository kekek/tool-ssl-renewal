package version

// 打印当前的编译信息

import (
	"fmt"
	"strings"
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

func (r PkgVersion) LogVersion() string {
	return r.BuildDateTime + "-" + r.BuildGitCommit
}

func (m PkgVersion) PrintVersion() {
	fmt.Printf("%s", m.TplVersion())
	fmt.Println()
}

func (m PkgVersion) TplVersion() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildTime", BuildDateTime))
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildGitBranch", BuildGitBranch))
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildGitCommit", BuildGitCommit))
	str.WriteString(fmt.Sprintf("%15s: %s \n", "BuildPkgModule", BuildPackageModule))
	return str.String()
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
