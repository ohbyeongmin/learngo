# Learn Go

### 1. Packages and Imports

### 2. Variables and Constants

- Constants 선언

```go
const name string = "ohbyeongmin"
```

- Variables 선언

기본 :

```go
var name string = "ohbyeongmin
```

Func 안에서는 축약형이 사용 가능

```go
name := "byeongmin
```

### 3. Functions

- 함수 선언 시 arguments 의 타입과 return 타입을 같이 선언

```go
func add(a int, b int) int {
    return a + b
}

Or

func add(a, b int) int {
    return a + b
}
```

- 여러개의 return 값을 가질 수 있다.

```go
func lenAndUpper(name string) (int, string) {
    return len(name), strings.ToUpper(name)
}

func main(){
    totalLength, upperName := lenAndUpper("byeongmin")

    // go는 무언가를 만들고 사용하지 않으면 Error 를 발생시킨다.
    fmt.Println(totalLength, upperName)
}

// 무시되는 value 는 _ 로 처리 할 수 있다.
totalLength, _ := lenAndUpper("byeongmin")
```

- argument 를 여러개 받을 수 있다.

```go
func repeatMe(words ...string){
    fmt.Println(words)
}

func main(){
    repeatMe("a", "b", "c", "d")
}

// [a b c d]
```

- naked return

```go
func lenAndUpper(name string) (length int, uppercase string){
    length = len(name)
    uppercase = strings.ToUpper(name)
    return
}
```

- defer : 함수가 끝나고 뭔가를 더 실행시켜준다

```go
func lenAndUpper(name string) (length int, uppercase string){
    defer fmt.Println("done")
    length = len(name)
    uppercase = strings.ToUpper(name)
    return
}
```

### 4. for, range, ...args

- only for

```go
func superAdd(numbers ...int) int {
    // range 는 array에 loop를 적용할 수 있도록 해준다.
    // 기본적으로 index 를 준다.
    total := 0
    for index, number := range numbers {
        total += number
    }
    return total
}

func main() {
    total := superAdd(1,2,3,4,5,6)
}
```

### 5. if with a Twist

```go
func canIDrink(age int) bool {
    // 조건에 () 가 필요없다.
    if age < 18 {
        return false
    }
    return true
}
```

- variable experssion : if를 쓰는 순간에 variable 을 생성 할 수 있다.

```go
func canIDrink(age int) bool {
    if koreanAge := age + 2; koreanAge < 18 {
        return false
    }
    return true
}
```

### 6. switch

```go
func canIDrink(age int) bool {
    switch age {
        case 10:
            return false
        case 18:
            return true
    }
    return true
}
```

### 7. Pointers

- C와 비슷
- & 주소값
- - 주소로 접근

```go
func main(){
    a := 2
    b := &a
    *b = 6
    fmt.Println(a)
    // 6 출력
}
```

### 8. Arrays and Slices

```go
numbers := [5]int{1,2,3,4}
numbers[4] = 5

// Slice : length 없이 Array 사용
// append 는 직접 값을 수정 하는 것이 아니라 새로운 값을 return 해준다.
numbers := []int{1,2,3,4}
numbers = append(numbers, 5)
```

### 9. Maps

map[key]value{...}
map 도 iterator 하게 사용 할 수 있다.

```go
me := map[string]string{"name":"byeongmin", "age":"29"}
// print : map[age:29, name:byeongmin]

for key, value := range me {
    fmt.Println(key, value)
}
// print : name byeongmin
//         age 29
```

### 10. Structs

```go
type person struct {
    name string
    age int
    favFood []string
}

func main(){
    // 1. 순서대로 값을 넣어줌
    favFood := []string{"a", "b"}
    me := person{"obm", 29, favFood}

    // 2. key 값을 지정하여 작성
    me := person{name: "obm", age: 29, favFood: favFood}
}
```
