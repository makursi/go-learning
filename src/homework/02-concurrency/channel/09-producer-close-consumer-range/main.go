// 题目：发送方 close、接收方 range——把「发」与「收」分给两个协程，靠 close 交接收尾。
// 生产者在这头往无缓冲通道依次发 1、2、3、4，发完就 close；消费者在那头 for range，
// 收满 4 个后循环自己结束，打印个数与总和（4 个、和 10）。
// 再 range 一个刚建立就 close 的空通道，确认一次迭代都不进（0 次）——
// 验证 range 的出口只有「已关闭且已排空」，而 close 必须落在发送方手上。

package main

import "fmt"

func main() {
	ch := make(chan int) // 无缓冲：发一个就得收一个，range 与发送同步推进

	go func() {
		for i := 1; i <= 4; i++ {
			ch <- i
		}
		close(ch) // 必须由发送方关闭：没有这一句，下面的 range 会永久等下去
	}()

	sum, count := 0, 0
	for v := range ch {
		fmt.Println("收到", v)
		sum += v
		count++
	}
	fmt.Printf("共收到 %d 个，求和 %d\n", count, sum)

	// 空通道一建立就关闭：range 立刻结束，循环体一次都不执行
	empty := make(chan int)
	close(empty)

	n := 0
	for range empty {
		n++
	}
	fmt.Println("空通道被 range 的次数：", n)
}
