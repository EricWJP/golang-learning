package main
import ("fmt"; "strings"; "strconv")

// func main() {
// 	ktype := "user_id:fk:users(id)"
// 	ka := strings.SplitN(ktype, ":", 2)

// 	fmt.Println(ka[1])
// 	fmt.Printf("%T\n%v", ka, len(ka))

// }
func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	mp := map[int]int{}
	fmt.Println(mp[1])
	if mp[1] == 0 {
		fmt.Println(1110000000)
		fmt.Println(mp)
	}
	iEnd := []int{10}
	for i:=0; i<iEnd[0];i++ {
		fmt.Println(i)
		// fmt.Println(iEnd)
		if i == 3 {
			iEnd[0] = 5
		}
	}
	switch iEnd[0] {
	case 1:
		fmt.Println(111111)
	case 5:
		fmt.Println(55555)
	default:
		fmt.Println("bucunzai")
	}
	fmt.Println(26/3)

	fmt.Println(strconv.Atoi("33"))
	sSlice := []string{"aa", "bb", "cc"}
	rStrings(sSlice)
	fmt.Println(sSlice)
	fmt.Println(strings.SplitN("abc", "", 5))
	fmt.Println(string('a'+1))
}
func rStrings(sSlice []string) []string {
	// sSlice = []string{"a", "b"}
	sSlice[0] = "aaaaaa"
	return []string{}
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}