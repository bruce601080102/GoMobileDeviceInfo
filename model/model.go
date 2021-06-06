package model

//https://onedrive.live.com/edit.aspx?resid=F45737980BFDD13B!55244&ithint=file%2cxlsx&authkey=!ALh0SdqmkpNi7u4
import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"github.com/distatus/battery"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
	// This is required to use H264 video encoder
)

var (
	sizeInMB float64 = 999 // This is in megabytes
	suffixes [5]string
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

//c001
func Platform() string {
	os := runtime.GOOS
	return fmt.Sprintf("%s", string(os))
}

//c023
func GHWMemory() string {
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
		return "nan"
	}
	phys := memory.TotalPhysicalBytes
	return fmt.Sprintf("%d", phys)
}

func GHWUsable() string {
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
		return fmt.Sprintf("%d", "nan")
	}

	usable := memory.TotalUsableBytes
	return fmt.Sprintf("%d ", usable)
}

//c019
func CPUPhy() string {
	physicalCnt, err := cpu.Counts(false)
	if err != nil {
		fmt.Printf("Error getting physicalCnt info: %v", err)
		return "nan"
	}
	return fmt.Sprintf("%d ", physicalCnt)
}

func CPUlogical() string {
	logicalCnt, err := cpu.Counts(true)
	if err != nil {
		fmt.Printf("Error getting CPU logical info: %v", err)
		return "nan"
	}
	return fmt.Sprintf("%d ", logicalCnt)
}

func CPUName() string {

	infos, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting CPU Name info: %v", err)
		return fmt.Sprintf("%d", "nan")
	}
	return infos[0].ModelName
}

func CPUPhysicalID() string {

	infos, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting PhysicalID info: %v", err)
		return fmt.Sprintf("%d", "nan")
	}
	return infos[0].PhysicalID
}

func CPUVendorID() string {

	infos, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting VendorID info: %v", err)
		return fmt.Sprintf("%d", "nan")
	}
	return infos[0].VendorID
}

func CPUmhz() string {

	infos, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting hz info: %v", err)
		return fmt.Sprintf("%d", "nan")
	}
	hz := infos[0].Mhz
	return fmt.Sprintf("%f", hz)
}

func Battery() {

	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}
	// for i, battery := range batteries {
	// 	fmt.Printf("Bat%d: ", i)
	// 	fmt.Printf("state: %s, ", battery.State.String())
	// 	fmt.Printf("current capacity: %f mWh, ", battery.Current)
	// 	fmt.Printf("last full capacity: %f mWh, ", battery.Full)
	// 	fmt.Printf("design capacity: %f mWh, ", battery.Design)
	// 	fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
	// 	fmt.Printf("voltage: %f V, ", battery.Voltage)
	// 	fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)
	// }

	fmt.Printf("current capacity: %f mWh", batteries[0].Current)
	fmt.Printf("last full capacity: %f mWh : ", batteries[0].Full)
	// dictBattery := make(map[string]string)
	// dictBattery["Brand"] =
	// dictBattery["Brand"] =

}

// =================================1 mac address end=================================
func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		for i := 0; i < 5; i++ {
			macAddrs = append(macAddrs, "nan")
		}
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}

	for i := 0; i < 5; i++ {
		macAddrs = append(macAddrs, "nan")
	}

	return macAddrs
}

func getIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func Mac1AddressMask() string {
	return getMacAddrs()[0]
}

func Mac1Address() string {
	if len(getIPs()) == 0 {
		return "nan"
	}

	return getIPs()[0]
}

func Mac2AddressMask() string {
	return getMacAddrs()[1]
}

func Mac3AddressMask() string {

	return getMacAddrs()[2]
}

func Mac4AddressMask() string {

	return getMacAddrs()[3]
}

func Mac5AddressMask() string {

	return getMacAddrs()[4]
}

// =================================1 mac address end=================================
func GOARCH() string {

	return runtime.GOARCH
}

// =================================1 Android Shell CPUInfo start=================================
func JsonShellCPUInfo(shellText string) string {
	listText := strings.Split(shellText, "\n")
	dictText := make(map[string]string)

	for _, c := range listText {
		if len(c) == 1 {
			break
		}

		listSplit := strings.Split(c, ":")
		if len(listSplit) > 1 {
			tabKeyReplace := strings.ReplaceAll(listSplit[0], "	", "")
			dictText[strings.ReplaceAll(tabKeyReplace, " ", "")] = listSplit[1]
		}
	}

	jsonBytes, err := json.MarshalIndent(dictText, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	jsonText := string(jsonBytes)
	// fmt.Println(jsonText)
	return jsonText
}
func AndroidShellCPUInfo() string {

	cmd1 := exec.Command("/system/bin/cat", "/proc/cpuinfo")
	out1, err := cmd1.Output()
	if err != nil {
		log.Println("cmd CPUInfo.Run() failed with ", err)
		return "nan"
	}
	return fmt.Sprintf("%s", JsonShellCPUInfo(string(out1)))
}

// =================================1 Android Shell CPUInfo end=================================

// =================================2 Android Shell brand start=================================

func AndroidShellgetprop() string {

	var (
		outputBrand        string
		outputSdk          string
		outputDevice       string
		outputModel        string
		outputRelease      string
		outputTimezone     string
		outputSerialNumber string
		outputGPSVersion   string
		outputISPName      string
		outputDeviceid     string
		outputWifimac      string
		outputWifimac2     string
		outputWigigmac     string
	)
	cmd0 := exec.Command("/system/bin/getprop", "ro.product.brand")

	cmd1 := exec.Command("/system/bin/getprop", "ro.build.version.sdk")
	cmd2 := exec.Command("/system/bin/getprop", "ro.product.device")
	cmd3 := exec.Command("/system/bin/getprop", "ro.product.model")
	cmd4 := exec.Command("/system/bin/getprop", "ro.build.version.release")
	cmd5 := exec.Command("/system/bin/getprop", "persist.sys.timezone")
	cmd6 := exec.Command("/system/bin/getprop", "ro.serialno")
	cmd7 := exec.Command("/system/bin/getprop", "gps.version.driver")
	cmd8 := exec.Command("/system/bin/getprop", "gsm.operator.alpha")
	cmd9 := exec.Command("/system/bin/getprop", "ro.deviceid")
	cmd10 := exec.Command("/system/bin/getprop", "ro.wifimac")
	cmd11 := exec.Command("/system/bin/getprop", "ro.wifimac_2")
	cmd12 := exec.Command("/system/bin/getprop", "ro.wigigmac")

	out0, err := cmd0.Output()
	if err != nil {
		log.Println("cmd brand.Run() failed with ", err)
		outputBrand = "nan"
	} else {
		outputBrand = string(out0)
	}

	out1, err1 := cmd1.Output()
	if err1 != nil {
		log.Println("cmd sdk.Run() failed with ", err1)
		outputSdk = "nan"
	} else {
		outputSdk = string(out1)
	}

	out2, err2 := cmd2.Output()
	if err2 != nil {
		log.Println("cmd device.Run() failed with ", err2)
		outputDevice = "nan"
	} else {
		outputDevice = string(out2)
	}

	out3, err3 := cmd3.Output()
	if err3 != nil {
		log.Println("cmd model.Run() failed with ", err3)
		outputModel = "nan"
	} else {
		outputModel = string(out3)
	}

	out4, err4 := cmd4.Output()
	if err4 != nil {
		log.Println("cmd release.Run() failed with ", err4)
		outputRelease = "nan"
	} else {
		outputRelease = string(out4)
	}

	out5, err5 := cmd5.Output()
	if err5 != nil {
		log.Println("cmd release.Run() failed with ", err5)
		outputTimezone = "nan"
	} else {
		outputTimezone = string(out5)
	}

	out6, err6 := cmd6.Output()
	if err6 != nil {
		log.Println("cmd release.Run() failed with ", err6)
		outputSerialNumber = "nan"
	} else {
		outputSerialNumber = string(out6)
	}

	out7, err7 := cmd7.Output()
	if err7 != nil {
		log.Println("cmd gps version.Run() failed with ", err7)
		outputGPSVersion = "nan"
	} else {
		outputGPSVersion = string(out7)
	}

	out8, err8 := cmd8.Output()
	if err8 != nil {
		log.Println("cmd ISP Name.Run() failed with ", err8)
		outputISPName = "nan"
	} else {
		outputISPName = string(out8)
	}

	out9, err9 := cmd9.Output()
	if err9 != nil {
		log.Println("cmd Deviceid.Run() failed with ", err9)
		outputDeviceid = "nan"
	} else {
		outputDeviceid = string(out9)
	}
	out10, err10 := cmd10.Output()
	if err10 != nil {
		log.Println("cmd Wifimac.Run() failed with ", err10)
		outputWifimac = "nan"
	} else {
		outputWifimac = string(out10)
	}
	out11, err11 := cmd11.Output()
	if err11 != nil {
		log.Println("cmd Wifimac2.Run() failed with ", err11)
		outputWifimac2 = "nan"
	} else {
		outputWifimac2 = string(out11)
	}
	out12, err12 := cmd12.Output()
	if err12 != nil {
		log.Println("cmd Wigigmac.Run() failed with ", err12)
		outputWigigmac = "nan"
	} else {
		outputWigigmac = string(out12)
	}

	dictGetprop := make(map[string]string)
	dictGetprop["Brand"] = strings.ReplaceAll(outputBrand, "\n", "")
	dictGetprop["sdkVersion"] = strings.ReplaceAll(outputSdk, "\n", "")
	dictGetprop["Device"] = strings.ReplaceAll(outputDevice, "\n", "")
	dictGetprop["Model"] = strings.ReplaceAll(outputModel, "\n", "")
	dictGetprop["ReleaseVersion"] = strings.ReplaceAll(outputRelease, "\n", "")
	dictGetprop["Timezone"] = strings.ReplaceAll(outputTimezone, "\n", "")
	dictGetprop["DeviceSerialNumber"] = strings.ReplaceAll(outputSerialNumber, "\n", "")
	dictGetprop["DeviceGPSVersion"] = strings.ReplaceAll(outputGPSVersion, "\n", "")
	dictGetprop["ISPName"] = strings.ReplaceAll(outputISPName, "\n", "")
	dictGetprop["Deviceid"] = strings.ReplaceAll(outputDeviceid, "\n", "")
	dictGetprop["Wifimac"] = strings.ReplaceAll(outputWifimac, "\n", "")
	dictGetprop["Wifimac2"] = strings.ReplaceAll(outputWifimac2, "\n", "")
	dictGetprop["Wigigmac"] = strings.ReplaceAll(outputWigigmac, "\n", "")

	jsonBytes, err := json.MarshalIndent(dictGetprop, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	jsonText := string(jsonBytes)

	return fmt.Sprintf("%s", jsonText)
}

// =================================2 Android Shell brand end=================================

// =================================3 Android Shell screen siz start=================================
func AndroidShellSreenSize() string {

	var outputSreenSize string
	cmd0 := exec.Command("/system/bin/dumpsys", "battery")
	out0, err := cmd0.Output()
	log.Println("cmd Sreen Size.Run() out0 ", out0)
	if err != nil {
		log.Println("cmd Sreen Size.Run() failed with ", err)
		outputSreenSize = "nan"
	} else {
		// outputSreenSize = strings.ReplaceAll(strings.Split(string(out0), ":")[1], "	", "")
		outputSreenSize = fmt.Sprintf("combined out:\n%s\n", string(out0))
	}
	return outputSreenSize
}

// =================================3 Android Shell screen siz end=================================

// =================================4 Android Shell device IMEI start=================================
func AndroidShellIMEI() string {

	var outputIMEI string
	cmd0 := exec.Command("/system/bin/service", "call", "iphonesubinfo", "1")
	out0, err := cmd0.Output()
	if err != nil {
		log.Println("cmd IMEI.Run() failed with ", err)
		outputIMEI = "nan"
	} else {
		outputIMEI = string(out0)
	}
	return outputIMEI
}

// =================================4 Android Shell screen siz end=================================
