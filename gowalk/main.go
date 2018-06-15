package main

import (
	"fmt"
	"github.com/StevenZack/tools/fileToolkit"
	"os"
	"os/exec"
)

func main() {
	program := "main.go"
	if len(os.Args) > 1 {
		program = os.Args[1]
	}
	name := getFileWithoutFormat(program)
	if !fileToolkit.FileExists(name + ".exe.manifest") {
		f, e := os.OpenFile(name+".exe.manifest", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if e != nil {
			fmt.Println(e)
			return
		}
		f.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0" xmlns:asmv3="urn:schemas-microsoft-com:asm.v3">
	<assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
	<dependency>
		<dependentAssembly>
			<assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
		</dependentAssembly>
	</dependency>
	<asmv3:application>
		<asmv3:windowsSettings xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">
			<dpiAware>true</dpiAware>
		</asmv3:windowsSettings>
	</asmv3:application>
</assembly>
`)
		f.Close()
	}
	e := exec.Command("go", "build", program).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
	e = exec.Command(name + ".exe").Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
func getFileWithoutFormat(name string) string {
	for i := 0; i < len(name); i++ {
		if name[i:i+1] == "." {
			return name[:i]
		}
	}
	return name
}
