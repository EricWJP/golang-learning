package main
import ( "fmt"; "os/exec"; "sort"; "strconv"; "strings" )

func main() {
	ps := exec.Command("ps", "-e", "-o pid,ppid,comm")
	output, _ := ps.Output()
	fmt.Println("完整", ps.Run())
	fmt.Printf("fileds Type：%T", output)
	child := make(map[int][]int)
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 { continue } // kill first line
		// fmt.Println("长度", len(s))
		if len(s) == 0 { continue } // kill empty line
		// fmt.Println("完整行", s)
		f := strings.Fields(s)
		// fmt.Println("fileds：", f)
		// fmt.Printf("fileds Type：%T", f)
		fpp, _ := strconv.Atoi(f[1]) // parent pid
		fp, _ := strconv.Atoi(f[0]) // child pid
		child[fpp] = append(child[fpp], fp)
	}
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child { schild[i] = k; i++ }
	sort.Ints(schild)
	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
		if len(child[ppid]) == 1 {
			fmt.Printf(": %v\n", child[ppid])
			continue
		}
		fmt.Printf("ren: %v\n", child[ppid])
	}
}