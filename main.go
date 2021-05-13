package main

import (
	"fmt"
	"golang/model"
)

func main() {
	platform := model.Platform() //c001
	fmt.Println("1.platform: ", platform)
	phys := model.GHWMemory()
	usable := model.GHWUsable()
	fmt.Println("2.phys Memory: ", phys)    //c023
	fmt.Println("3.usable Memory:", usable) //現在使用率

	physicalCnt := model.CPUPhy()
	logicalCnt := model.CPUlogical()
	fmt.Println("4.physicalCnt:", physicalCnt) //現在使用率
	fmt.Println("5.logicalCnt:", logicalCnt)   //現在使用率
	CPUName := model.CPUName()
	CPUPhysicalID := model.CPUPhysicalID()
	CPUVendorID := model.CPUVendorID()
	CPUmhz := model.CPUmhz()
	fmt.Println("6.CPU Name:", CPUName)             //現在使用率
	fmt.Println("7.CPU PhysicalID:", CPUPhysicalID) //現在使用率
	fmt.Println("8.CPU VendorID:", CPUVendorID)     //現在使用率
	fmt.Println("9.CPU mhz:", CPUmhz)               //現在使用率

	GOARCH := model.GOARCH()
	fmt.Println("9.system GOARCH:", GOARCH)

	AndroidCPUinfo := model.AndroidShellCPUInfo()
	fmt.Println("10 system shell cpuinfo:", AndroidCPUinfo)
	// model.PCIdb()
}
