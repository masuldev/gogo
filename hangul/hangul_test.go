package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}

	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)

	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_array() {
	fruits := [3]string{"사과", "바나나", "토마토"}

	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
}

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

	// Output:
	// [사과 바나나 토마토]
	// [포도 딸기]
	// [사과 바나나 토마토 포도 딸기]
	// [사과 바나나 포도 딸기]
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
	// Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// [1 2 3]
	// len: 3
	// cap: 5
	//
	// [3 4 5]
	// len: 3
	// cap: 3
	//
	// [1 2 3 4]
	// len: 4
	// cap: 5
	//
	// [1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}

func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	anotherDest := append([]int(nil), src...)

	copy(dest, src)

	if n := copy(dest, src); n != len(src) {
		fmt.Println("복사가 덜 됐습니다.")
	}

	fmt.Println(dest)
	fmt.Println(anotherDest)
	// Output:
	// [30 20 50 10 40]
	// [30 20 50 10 40]
}
