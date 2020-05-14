好的沟通就像是一杯刺激的浓咖啡，然后就难以入睡

f, _ := os.Open("/etc/passwd");
defer f.Close()
r := bufio.NewReader(f)
s, ok := r.ReadString('\n')

f, e := os.Stat("name"); e != nil {
	os.Mkdir("name", 0755)
} else {
	// error
}

shell 
if [! -e name]; then
	mkdir name
else
	#error
fi

flag 
flag int bool string
intVar boolVar stringVAr

import "os/exec"
cmd = exec.Command("bin/ls", "-l")
err := cmd.Run()

net
conn, e := Dial("tcp", "192.0.32.10:80")
conn, e := Dial("udp", "192.0.32.10:80")
conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80") // 方括号是必须的

Read(b []byte)(n int, err error)

"net/http"
"io/ioutil"
r, err := http.Get("http://www.google.com/robots.txt")
ioutil.ReadAll(r.Body)
r.Body.Close()









