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
	fmt.Println("4.physicalCnt:", physicalCnt) //物理核心數
	fmt.Println("5.logicalCnt:", logicalCnt)   //邏輯核心數
	CPUName := model.CPUName()
	CPUPhysicalID := model.CPUPhysicalID()
	CPUVendorID := model.CPUVendorID()
	CPUmhz := model.CPUmhz()
	fmt.Println("6.CPU Name:", CPUName)             //cpu名稱
	fmt.Println("7.CPU PhysicalID:", CPUPhysicalID) //cpu id
	fmt.Println("8.CPU VendorID:", CPUVendorID)     //cpu VendorID
	fmt.Println("9.CPU mhz:", CPUmhz)               //cpu 赫茲

	GOARCH := model.GOARCH()
	fmt.Println("10.system GOARCH:", GOARCH) //cpu 赫茲

	Mac1Address := model.Mac1Address()
	fmt.Println("11  Mac1Address:", Mac1Address)

	Mac1AddressMask := model.Mac1AddressMask()
	fmt.Println("12  Mac1AddressMask:", Mac1AddressMask)

	Mac2AddressMask := model.Mac2AddressMask()
	fmt.Println("13  Mac2AddressMask:", Mac2AddressMask)

	Mac3AddressMask := model.Mac3AddressMask()
	fmt.Println("14  Mac3AddressMask:", Mac3AddressMask)

	Mac4AddressMask := model.Mac4AddressMask()
	fmt.Println("15  Mac4AddressMask:", Mac4AddressMask)

	Mac5AddressMask := model.Mac5AddressMask()
	fmt.Println("16  Mac5AddressMask:", Mac5AddressMask)

}
