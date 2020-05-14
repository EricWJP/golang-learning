byte(byte) 
byte(rune) 求余结果
rune(byte) 
byte(rune)
[]byte(string)
[]rune(string)
func EqualFold(s, t string) bool
	判断两个utf-8编码字符串是否相同，不区分大小写
	fmt.Println(strings.EqualFold("Go", "go"))
	true
func HasPrefix(s, prefix string) bool
	判断s是否有前缀prefix
	fmt.Println(strings.HasPrefix("Go", "G"))
	true
func HasSuffix(s, suffix string) bool
	判断s是否有后缀suffix
	类同 前缀
func Contains(s, substr string) bool
	判断s是否 包含 substr
	abz :     ---结果： false      ---结果(reverse)： false
    defef3f:  ---结果： false      ---结果(reverse)： false
    abc :     ---结果： true       ---结果(reverse)： false
    ab :      ---结果： true       ---结果(reverse)： false
func Compare(a, b string) int
	比较 两个字符串，逐个字符比较顺序(a..z)直到出现不同或者遍历完成
		 相同 -- 0  		< -- -1		> -- 1 
	"abc"
	   abd :     ---结果： -1      ---结果(reverse)： 1
	   def :     ---结果： -1      ---结果(reverse)： 1
	   abc :     ---结果： 0      ---结果(reverse)： 0
func Count(s, sep string) int
	返回s中有几个不重复的sep子串
	abcdeabcmoiiokkopkpook
   		abz :     ---结果： 0      ---结果(reverse)： 0
   		defef3 :  ---结果： 0      ---结果(reverse)： 0
   		abc :     ---结果： 2      ---结果(reverse)： 0
   		ab :      ---结果： 2      ---结果(reverse)： 0
func Index(s, sep string) int
	子串sep在s中第一次出现的位置
func IndexByte(s string, c byte) int
	字符c在s中第一次出现的位置	不存在返回-1
	abcdeabcm
	    98 :     ---结果： 1
	    100 :     ---结果： 3
	    102 :     ---结果： -1
	    115 :     ---结果： -1
func IndexRune(s string, r rune)
	unicode码值 r 在s中第一次出现的位置，不存在则返回-1
	abcdeabcm
		9898 :     ---结果： -1
		98 :     ---结果： 1
		100 :     ---结果： 3
		102 :     ---结果： -1
		115 :     ---结果： -1
func IndexAny(s, chars string) int
	字符串chars中的 任一 utf-8码值 在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1
	abcdehlnopm
		abz :     ---结果： 0
		defef :   ---结果： 3
		bc :      ---结果： 1
		os :      ---结果： 8
func IndexFunc(s string, f func(rune) bool) int
	s中第一个字符(rune)满足函数 f 的位置i，
		函数f: 这里的rune（类似其他语言的char类型只不过这里是数字utf-8码值）返回true
		不存在则返回-1。
func LastIndex(s, sep string) int
	类同Index，不同的是 最后一次出现
func LastIndexAny(s, chars string) int
	类同IndexAny，不同的是 最后一次出现
	abcdehlnopm
	   abz :    ---结果： 1
	   defef :  ---结果： 4
	   bc :     ---结果： 2
	   os :     ---结果： 8
func LastIndexFunc(s string, f func(rune) bool) int
	类同IndexFunc，不同的是 最后一次出现
func Title(s string) string
	返回s中每个单词的首字母改为标题格式的字符串副本
	fmt.Println(strings.Title("Hello, my world!"))
		=> Hello, My World!
func ToLower(s string) string
	所有字母小写
func ToLowerSpecial(_case unicode.SpecialCase, s string) string
	使用_case规定的字符映射，返回将所有字母都转为对应的小写版本的拷贝
func ToUpper(s string) string
	所有字母大写
func ToUpperSpecial(_case unicode.SpecialCase, s string) string
	类同ToLowerSpecial，不同的是 大写版本的拷贝
func ToTitle(s string) string
	返回将所有字母都转为 对应的 标题版本 的拷贝。 类似 ToUpper
func ToTitleSpecial(_case unicode.SpecialCase, s string) string
	使用_case规定的字符映射，返回将所有字母都转为对应的标题版本的拷贝
func Repeat(s string, count int) string
	返回count个s串连的字符串
func Replace(s, old, new string, n int) string
	返回 将 s 中前n个不重叠old子串都替换为new的新字符串
		n<0时，会替换所有old子串
	函数参数没有默认值，不能没有
func Map(mapping func(rune) rune, s string) string
	将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝
	如果mapping是负值，将会丢弃该码值而不会替换。
func Trim(s string, cutset string) string
	返回将s前后端 所有cutset包含的 字符 的连续字符 都去掉后的字符串
func TrimSpace(s string) string
	返回将s前后端所有空白（unicode.IsSpace指定）串 都去掉的 字符串
func TrimFunc(s string, f func(rune) bool) string
	返回将s前后端所有满足f的 字符串 都去掉的字符串
func TrimLeft(s string, cutset string) string
	类同 Trim，不同的是 前端
func TrimLeftFunc(s string, f func(rune) bool) string
	类同 TrimFunc，不同的是 前端
func TrimPrefix(s, prefix string) string
	返回去除s可能的前端 prefix完全字符串 后的字符串
func TrimRight(s string, cutset string) string
	类同 Trim，不同的是 后端
func TrimRightFunc(s string, f func(rune) bool) string
	类同 TrimFunc，不同的是 后端
func TrimSuffix(s, suffix string) string
	返回s去除可能的后缀suffix完全字符串 后的字符串
func Fields(s string) []string
	返回将字符串s按照空白分隔(unicode.IsSpace确定，可以是多个)分割成一个字符串切片
		全部都是空白就返回空切片
func FieldsFunc(s string, f func(rune) bool) []string
	返回 将字符串s按照 满足f的字符 组成的连续字符串 作为分隔符进行分割

func Split(s, sep string) []string
	字符串s 按照 完整字符串sep 进行分割，分割到结尾， 两个相同的字符串sep相邻也要进行分割
	sep为空字符串，一个字符作为一个元素进行分割
func SplitN(s, sep string, n int) []string
	类同Split
	n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	n == 0: 返回 空切片
	n < 0 : 返回所有的子字符串组成的切片
func SplitAfter(s, sep string) []string
	类同Split，不同的是 分割的时候是从字符串sep后面分割，sep串是保留的
	strings.SplitAfter("a,b,c", ",")
	 =>	["a," "b," "c"]
func SplitAfterN(s, sep string, n int) []string
	类同SplitAfter，不同的是 
		n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
		n == 0: 返回nil
		n < 0 : 返回所有的子字符串组成的切
func Join(a []string, sep string) string
	一系列字符串连接为一个字符串，之间用sep来分隔

type Reader struct {
	//内含隐藏或非导出字段
}
	Reader类型 通过从一个字符串读取数据，实现了 io.Reader、io.Seeker、io.ReaderAt、
		io.WriterTo、io.ByteScanner、io.RuneScanner接口
	func NewReader(s string) *Reader
		创建了一个从字符串s读取数据的Reader。更有效率，只读
	func (r *Reader) Len() int
		返回 r 包含的字符串还没有被读取的部分
	func (r *Reader) Read(b []byte) (n int, err error)
	func (r *Reader) ReadByte() (b byte, err error)
	func (r *Reader) UnreadByte() error
	func (r *Reader) ReadRune() (ch rune, size int, err error)
	func (r *Reader) UnreadRune() error
	func (r *Reader) Seek(offset int64, whence int) (int64, error)
		Seek 实现了 io.Seeker 接口
	func (r *Reader) Readat(b []byte, off int64) (n int, err error)
	func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
		实现了 io.WriterTo 接口
type Replacer struct {
	//内含隐藏或非导出字段
}
	Replacer 类型进行一系列字符串的替换
	func NewReplacer(oldnew ...string) *Replacer
		多组old、new字符串对 创建并返回一个 *Replacer。 替换依次进行，匹配时不会重叠

		r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
		fmt.Println(r.Replace("This is <b>HTML</b>!"))
			=> This is &lt;b&gt;HTML&lt;/b&gt;!
	func (r *Replacer) Replace(s string) string
		返回s所有替换进行完成后的副本
	func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
		向w中写入s的所有替换完成后的副本












