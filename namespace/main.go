package main
import (
	"log"
	"os"
	"os/exec"
	"syscall"
	"fmt"
)
// 注: 运行时需要 root 权限。
func main() {
	fmt.Println("Start!")
	// 创建bash shell命令对象
	cmd := exec.Command("bash")
	fmt.Println("1")
	// syscall.SysProcAttr 是 Go 中定义的一个结构体
	// Cloneflags 是 SysProcAttr 中的一个字段，它用来指定 Linux 内核中 clone 系统调用的标志位
	// syscall.CLONE_NEWUTS 是 Cloneflags 的一个具体标志，表示该进程将会进入一个新的 UTS 命名空间
	// clone：创建一个新的进程并把他放到新的 namespace 中
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags: syscall.CLONE_NEWUTS,
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | 
		syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER |
		syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 启动进程
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Command ran successfully")
	}
	fmt.Println("END!")
}
