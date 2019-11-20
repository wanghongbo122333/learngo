 package cpu

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type cpu struct {
	//每个cpu读取的数据
	time   string
	name   string  //“all” 或者cpu序号
	user   float64 //用户级别运行使用CPU总时间的百分比
	nice   float64 //在用户级别，用于nice操作，所占用CPu总时间百分比
	system float64 //内核态运行占用百分比
	iowait float64 //等待IO操作占用时间百分比
	steal  float64 //管理程序为另一个虚拟进程提供服务而等待虚拟CPU的百分比
	idle   float64 //CPU空闲时间百分比

}
type cpuNetwork struct {
	time     string  //时间
	IFACE    string  //网络接口名称
	rxpckps  float64 //每秒收包数量
	txpckps  float64 //每秒发包的数量
	rxkBps   float64 //每秒收的数据量（kb）
	txkBps   float64 //每秒发送数据量（kb）
	rxcmpps  float64 //每秒钟接收的压缩数据包
	txcmpps  float64 //每秒钟发送的压缩数据包
	rxmcstps float64 //每秒中接收的多播数据包
	//检测计算的数据

}
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

//ConvertByte2String 格式转换 返回输出为字符串
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

//TestTerminal 从命令行执行命令arg 并获取输出,返回string
func TestTerminal(arg0 string, arg1 string) (string, error) {
	cmd := exec.Command(arg0, arg1)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return "", err
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return "", err
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return "", err
	}
	fmt.Printf("stdout:\n %s", ConvertByte2String(bytes, GB18030)) //打印输出
	return ConvertByte2String(bytes, GB18030), nil
}

// JudgeCPU 针对"sar -P ALL 1 2"命令解析JudgeCPU状态数据
func JudgeCPU(str string) {

	var badCPUs []cpu
	badCPUs = make([]cpu, 0)
	totalLines := 0
	countOfBadCPU := 0  //system>60的核数量
	countOfGoodCPU := 0 //<60
	var maxSystem *cpu  //存放最大cpu使用率的cpu信息
	maxSystem = new(cpu)
	//每行读取
	res := strings.Split(str, "\n")
	res = res[1:]
	for _, line := range res {
		totalLines++
		//获取该行信息
		if strings.Trim(line, " ") == "" { //空行跳过
			//补充判断条件
		} else {
			res := deleteSpace(strings.Split(line, " "))
			var c *cpu
			c = new(cpu)
			//各属性赋值
			c.time = res[0] + res[1]
			c.name = res[2]
			user, err := strconv.ParseFloat(strings.Trim(res[3], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			c.user = user
			nice, err := strconv.ParseFloat(strings.Trim(res[4], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			c.nice = nice
			system, err := strconv.ParseFloat(strings.Trim(res[5], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			c.system = system
			iowait, err := strconv.ParseFloat(strings.Trim(res[6], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			c.iowait = iowait
			steal, err := strconv.ParseFloat(strings.Trim(res[7], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			c.steal = steal
			idle, err := strconv.ParseFloat(strings.Trim(res[8], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			c.idle = idle

			//存最大cpu使用率的cpu
			if maxSystem == nil {
				maxSystem = c
			} else {
				if maxSystem.system < c.system {
					maxSystem = c
				}
			}
			//不同检测结果操作
			isSingleCPU, isBadCPU := judgeCPU(c, 60)
			if isSingleCPU {
				if isBadCPU {
					countOfBadCPU++
					fmt.Print("该cpu使用率较高：")
					badCPUs = append(badCPUs, *c)
					c.printInfo()
				} else {
					countOfGoodCPU++
				}
			} else {
				if isBadCPU {
					fmt.Println("总体平均CPU使用率过高")
					c.printInfo()
				} else {
					fmt.Println("总体平均CPU使用率正常")
				}
			}
		}
	}
	fmt.Println("共检测了", totalLines, "行记录.检测结果：cpu使用率小于60%的核的个数是", countOfGoodCPU, ",大于60%的核的个数是", countOfBadCPU)
	fmt.Println(badCPUs)
	fmt.Println("当前cpu使用率最大的CPU信息：")
	maxSystem.printInfo()
}

//JudgeNetWork 针对 "sar -n DEV 1 2" 命令解析 network状态数据
func JudgeNetWork(str string) {
	totalLines := 0
	countOfBadRecord := 0
	countOfGoodRecord := 0
	var maxPack *cpuNetwork //包量最大的
	maxPack = new(cpuNetwork)
	var maxCurrent *cpuNetwork //流量最大的
	maxCurrent = new(cpuNetwork)

	res := strings.Split(str, "\n")
	for _, line := range res {
		totalLines++
		if strings.Trim(line, " ") == "" {
			//补充判断条件
		} else {
			res := deleteSpace(strings.Split(line, " "))
			//存各项数据
			var n *cpuNetwork
			n = new(cpuNetwork)
			n.time = res[0] + res[1]
			n.IFACE = res[2]
			rxpckps, err := strconv.ParseFloat(strings.Trim(res[3], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.rxpckps = rxpckps
			txpckps, err := strconv.ParseFloat(strings.Trim(res[4], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.txpckps = txpckps
			rxkBps, err := strconv.ParseFloat(strings.Trim(res[5], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.rxkBps = rxkBps
			txkBps, err := strconv.ParseFloat(strings.Trim(res[6], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.txkBps = txkBps
			rxcmpps, err := strconv.ParseFloat(strings.Trim(res[7], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.rxcmpps = rxcmpps
			txcmpps, err := strconv.ParseFloat(strings.Trim(res[8], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.txcmpps = txcmpps
			rxmcstps, err := strconv.ParseFloat(strings.Trim(res[9], " "), 64)
			if err != nil {
				fmt.Println("strconvParseFloat error", err.Error())
				return
			}
			n.rxmcstps = rxmcstps
			//包量最大
			if maxPack == nil {
				maxPack = n
			} else {
				if maxCurrent.txpckps+maxPack.rxpckps < n.txpckps+n.rxpckps {
					maxPack = n
				}
			}

			//流量最大
			if maxCurrent == nil {
				maxCurrent = n
			} else {
				if maxCurrent.txkBps+maxCurrent.rxkBps < n.txkBps+n.rxkBps {
					maxCurrent = n
				}
			}

			flag := judgeDDOS(n, 500000, 4*1024*1024, 20)
			if flag {
				fmt.Print("这个接口有嫌疑，详细信息：")
				countOfBadRecord++
				n.printInfo()
			} else {
				countOfGoodRecord++
			}
		}
	}
	fmt.Println("共检测到", totalLines, "行记录", "可疑记录", countOfBadRecord, "条", "正常记录", countOfGoodRecord, "条")
	fmt.Print("当前包交换量最大的是：")
	maxPack.printInfo()
	fmt.Print("当前流量交换量最大的是：")
	maxCurrent.printInfo()

}

//处理每行字符串切割后的空格
func deleteSpace(str []string) []string {
	var res []string
	res = make([]string, 10)
	count := 0
	for _, data := range str {
		if data != "" {
			res[count] = data
			count++
		}
	}
	// fmt.Println(res)
	return res
}
func judgeCPU(c *cpu, CPUMax float64) (bool, bool) { //是否为单核信息，是否可疑
	return c.name != "all", c.system >= CPUMax
}

//判断是否有DDOS攻击嫌疑
func judgeDDOS(netInfo *cpuNetwork, MaxPack float64, MaxCurrent float64, MaxRatio float64) bool {
	flag := false //初始化没有进行ddos攻击
	//检测是否有ddos攻击
	//1.流量>4Gbps
	// fmt.Println("流量：",n.rxkBps,n.txkBps)
	if netInfo.rxkBps > MaxCurrent || netInfo.txkBps > MaxCurrent {
		//有嫌疑
		flag = true
	}
	//2.包量大于50000PPS
	// fmt.Println("包量:",n.rxpckps,n.txpckps)
	if netInfo.rxpckps > MaxPack || netInfo.txpckps > MaxPack {
		flag = true
	}
	//3.txpckps/rxpckps >20或者 txkBps/rxkBps>20 收发比例超过一定阈值
	if netInfo.rxpckps > 0 && netInfo.rxkBps > 0 { //分母不为0
		// fmt.Println("发送/接受包量比例：",n.txpckps/n.rxpckps,"发送/接收速率比例",n.txkBps/n.rxkBps)
		if netInfo.txpckps/netInfo.rxpckps >= MaxRatio || netInfo.txkBps/netInfo.rxkBps >= MaxRatio {
			flag = true
		}
	}
	return flag
}

type showInfo interface {
	printInfo()
}

func (netInfo *cpuNetwork) printInfo() {
	fmt.Println("Time:", netInfo.time, "Interface:", netInfo.IFACE, "rxpckps:", netInfo.rxpckps, "txpckps:", netInfo.txpckps, "rxkBps:", netInfo.rxkBps, "txkBps:",
		netInfo.txkBps, "rxcmpps:", netInfo.rxcmpps, "txcmpps:", netInfo.txcmpps, "rxmcstps:", netInfo.rxmcstps)
}
func (c *cpu) printInfo() {
	fmt.Println("time:", c.time, "name:", c.name, "user%:", c.user, "nice%:", c.nice, "iowait%:", c.iowait, "system%:", c.system, "idle%:", c.idle)
}
