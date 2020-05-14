os包提供了操作系统函数的不依赖平台的接口。Unix风格

func Hostname() (name string, err error)
获取内核提供的主机名
func Getpagesize() int
底层的系统内存页的尺寸
func Environ() []string
环境变量的格式为"key=value"的字符串的切片拷贝
func Getenv(key string) []string
检索并返回名为key的环境变量的值。如果不存在返回空字符串
func Setenv(key, value string) error
设置名为key的环境变量。 出错返回该错误
func Clearenv()
删除所有环境变量
func Exit(code int)
以给出的状态码code退出。通常 0表示成功

// TODO
func Expand(s string, mapping func(string) string) string
替换s中的${var}或$var为 mapping(var)
func ExpandEnv(s string) string
ExpandEnv函数替换s中的${var}或$var为名为var 的环境变量的值。引用未定义环境变量会被替换为空字符串

func Getuid() int
返回调用者的用户ID
func Geteuid() int
返回调用者的有效用户ID

func Getgid() int
返回调用者 组ID
func Getegid() int
返回调用者有效组ID
func Getgroups() ([]int, error)
返回调用者所属的所有用户组的组ID
func Getpid() int
返回调用者所在进程的进程ID
func Getppid() int
返回调用者所在进程的父进程的进程ID

type Signal interface {
	String() string
	Signal() //用于区分其他实现了Stringer接口的类型
}
代表一个操作系统信号，其底层实现是依赖于操作系统的

type PathError struct {
	Op   string
	Path string 
	Err  error
}
记录一个错误，以及导致错误的路径
func (e *PathError) Error() string

type linkError struct {
	Op  string
	Old string
	New string
	Err error
}
用来记录在Link、Symlink、Rename系统调用时出现的错误，以及导致错误的路径

func (e *LinkError) Error() string


type SyscallError struct {
	Syscall string
	Err 	error
}
记录某个系统调用出现的错误

func (e *SyscallError) Error() string
func NewSyscallError(syscall string, err error) error
返回一个指定系统调用名称和错误细节的SyscallError. err为nil，返回nil

type FileMode uint32
代表文件的模式和权限位

func (m FileMode) IsDir() bool
判断是否是一个目录
func (m FileMode) IsRegular() bool
报告m是否是一个普通文件
func (m FileMode) Perm() FileMode
返回m的Unix权限位
func (m FileMode) String() string





