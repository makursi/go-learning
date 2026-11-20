package main

// 按题号顺序调度。07（先发送后起接收）与 08（缓冲已满再发送）会 deadlock，
// 故意不接进来；要观察就把 main 里换成对应的入口函数单独跑。
func main() {
	UnbufferedChannel()      // 01
	BufferedChannel()        // 02
	CloseAndRange()          // 03
	SelectFastest()          // 04
	UnbufferedDocking()      // 05
	UnbufferedReceiveFirst() // 06
}
