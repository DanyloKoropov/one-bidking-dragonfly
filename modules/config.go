package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.5" )

func wApKl0mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRk3wCEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjEeIvVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f01vJQ9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9YZUbI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rth7mYYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5GWtASOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYvyK8CpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MF74dJ7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEjX2ehwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChPZKkctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32N6gbb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXJ1YrhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdgZVsHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNDGUlu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLxRbL7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXet27XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdgA5MEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtfDKaUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoNhD2oEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PxDouEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCPvlBBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqzEorAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLV3cVoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKcDOZzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gosuGNFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OmANC7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0g6rJadeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvG2Xe12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjneV6swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iznOYkteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kvlaaW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXw7JNpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkqNhYz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 442qThUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tNwjfsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ9Ao1QiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OD383H4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTrLwFoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xksQqUJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbQNH8MIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZUxqcEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sl1dGIE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ho65P9gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqxp1JQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxdgpByFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsUA3ConWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZAODP1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jecc1YMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zv9CtU6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 880D9sERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8oyHNA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AO4YlXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RzWjvtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZmbwIUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUFc9FasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgpXuPSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func un3fyswoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWxC5VxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7skS6vktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtPLah1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGRspYxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yW2nOVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLNDbr0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToRiNN6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YI3rDXF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ot3LIVgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3suFWb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUFFkr42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEJPtI7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNqUpHY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4MaN7sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdsebiYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHyKedSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNCL9tNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHkK4ByNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kG61cpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vvBYa7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xr6VFu6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlSCTmOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fgweNSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSr9QjsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9mcntHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YK90BW43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oo9bie90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhL9l2M2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oj8c3zkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3neG289Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIuDrPVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nYkKyGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4jHYVpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGWjbtNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Esc6pMbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DF74G6vrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3hAEOV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVLQyiFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7NsbzBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pseISVEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nO3gVWh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5l5jSEmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVO6tO6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKgvZdUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRQOTpSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BlkFKXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cez1THgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVgWzUJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BArTq24JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQOvv3TpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEsf4C10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yQKmsU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tILxP8jcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWcczIyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4CmxXGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GN5wllFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsP3IF85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRo3T6RYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fn0UMAX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOXchwZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzXROjDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ll2YISbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQNABe0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SC9TsLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebS5aUddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0twQhVwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upSHKW3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8yPEGhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sq48lG8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkPBYJTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbtASKpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfoK5aNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWJMpBD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AatXHscvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3EPHCvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jwuC583Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31ILWXV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfZFgW9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKzjDKpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXZUiETqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHuXpdyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qImEzCwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpI9MdpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjTpnSFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBPF0br8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8cffUZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXU2RpIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRMxgVjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUexc4a9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ya9Re3skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrOFHbiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ru463L9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMOdLUt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEXrK54uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wETlzsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3gQNtSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgY87QMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpFqVChlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0AgvzaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEbL23MaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHaSraT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rM8worCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDmad1LQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRRL37XiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afPIkHtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDtA3DuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvFUHauAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hgN1qfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBJkOgPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCxkQVQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUFqVCzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuM4CqdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCgcKDOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuB7CipIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfWMFecoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CONgRuXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mB8yxiZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2gF4GWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePQytSqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZU5xhoC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsXtaMQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8f9mhJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBhkeKW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUbQ090GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8m0zQrLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTznAzUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7tPfFBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wdqaf7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNBnA6SaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9epTLLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8FybwDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uEdxSOQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ux1scqfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vm7GnOs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSrGUaUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cADjtNhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMBzYKHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nieaoDDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NweTWAOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmWW1nJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8NvKjGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEWRIKcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEDzQoPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RAD7TOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDCmXYNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPgbhxJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvxL0MLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euKFMIKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhqKYGbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ys998livWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsBQlUweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7j0I9OX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRG7oqHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDIKMhjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJtHVfUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfEHL4BMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVaf5OxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4dStYHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyVOU3yUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4r3RKNbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvK9sM41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdmoCoDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dc07wczTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MM0qsTGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iq92KmgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eqp55sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph6ZR4fgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhnhKNSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jeyruazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxN6eWdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVJGRu7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPjIUIVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kopnnAvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgYEkpE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJsqdI1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9akisiuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1DwL1W6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFY9uO8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1hvHGZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cGAHTJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6RPJBtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGLzdoaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXiM4yr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oRabMvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtjmtlyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05FS44f6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8E5tYguYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcvz1kOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXI4QExTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBXMVaFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNy3NfI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nNv4eU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pf1B75yMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wExQfhmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5CDLzeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGQP2WYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pcbhv0RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPJTObAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0su7WmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVWcH76kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKmiI6hHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgT4nCOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdnqhyVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGLfcoX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cEhVpSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEUdGztSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AolzikIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugveNJaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbQYglXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K490NL5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zha4U8fXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DE2OUg5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0HGo78ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkv7W5s9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rszZgKTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DFZVZWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5yoPP4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nClFIkrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwLS1yTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7huO4sAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z97SJZXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLVjr1rQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6avx3eFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYzf3jNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVpltpXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uIEaa9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMrGW4dzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzbiGyTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwgkAwzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPvtaq7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcOuwuZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8LVbzUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LImdr3a0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWmvZP6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8v0cIPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMg2DfTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJkgMoDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IoszLW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiRraTk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPcWNenVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVdGC961Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKYSDGqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7Gv2GJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kj1952rRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKlWhIu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yq7zItMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYP4EpceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpdyqTQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnmwCohAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYrZO4QCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZfwzjGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THX2kxloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKielrRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2nFCAIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6Od3EB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5FrneUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9itfsZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXXSEVz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdpP9f4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IBBprqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nxa1wScoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6lzY8nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGaYyaDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func So1SLmjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YkoK8oXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Fx7weerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyDIxkjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCam5xHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUROeK1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKLYtSLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6QDMtjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4T4SHHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwDPFJWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4NyDeRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJ2j6wUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0hkEgZJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPOt5DBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poIvifZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ae9hcCJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpyZ8SfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuUmA4GwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbgvXLlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0WNlrpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgSwj2KdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkILQCPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpR10cDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D97qIY0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBJB1oChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmxTMRuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxBuq0ByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNzKAVSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aFfQYXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuZlW68fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NooDWlu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYefAR8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUew1j60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRi80Nx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZONahPRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Shrgf0HpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cW2qNyMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ufp6quUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzqlPrygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoY4Lc0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LO4fuM6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ztl129HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUYyVDgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqdzn7CiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03I4r5aiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8XgpO3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcUf6oHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6Ia0pszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5arFG3goWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXqPlms7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3hAeR6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaNMF3WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7em9AhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROro6XRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leW0L3PrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQPQK8QHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epAyGTYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5UTeLinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5UWWj6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXrp0E1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qphnpiVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjAGCv02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjRnERvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Pcra8cWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQhDFSMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzI9RUgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i54umeQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T65UnGcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoYSAxiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdUzv8OSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GANA3cCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJxENzPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJkjhEE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XACKejeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNnmHZtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlmQlvi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHEHMw47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vl3yFU0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLH8sRw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjzSZiIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rClWH7LTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xtfl5c09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCWtQeEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRWD0wV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpfRUinWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7inKbj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DchiNJ1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoxD65l9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dKtGVUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C25XfIJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZsTQeUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5HPDNIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2N7uQwtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJel6fYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCYv8zdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbrx6ExnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0ury3wyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAnxpLdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msUjBjTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nbsxXtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI9UEouNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AleVYurBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1XeYv56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1U9OKUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func do7ety5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMqvGncBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoDEvALqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Dbs3xJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7ItbXJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9wD9awQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWwXuBb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wx3qXJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mab9ibDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLnu6Mt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5H0dIsKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtSiJKR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtkJkH4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZyldfs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QI0cY839Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcL8lUEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlK9EsMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woaBvxALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKplvLsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwcpX1gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEwJJiXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgr7rUn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqeBaexKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlZmaU1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaWfKjZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpdXb6vZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4BrdCHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9zWgoQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghlG7nILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpRAZdjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9reHEasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3yrfmLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhtUUb5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hj8wyUvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSlNNFVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ha9fY1D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vymc4G13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlK0nLrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rg5V99PrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogvYlilsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBMKylywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fk8nTPbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhEZcoKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPK2MSViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEiW61ZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xpRyqDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6EwMBHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2l68kbQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tb3iRC7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBKo8cHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GV5MnFUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EavmqZFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTj3B7fRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h66y3cDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYOoFxteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAZ4vnUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hREK8qH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FtYly3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNEPuuUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TupjOTCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4YmziXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GnpvVfTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1LXX1FiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qk0M9goBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CQMpFB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt50nWM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1UedOBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnCBYPFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7TaQyxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zr04UFAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueC0OomYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BHnVdEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8LragqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDYvOdFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kX1OEMwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhcPVTAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fALl0SyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6ddFW1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdTKEz5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tX5TTerNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9Udshp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioQPPqs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zBBc6WUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPJODen1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1McCtiL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BfqdR6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoM6eC6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBKbQDvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r89hbJLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZQmGoRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEAn0Zt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjQnYZaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Y9VwDOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiahGK4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7t3N1zIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRr6FNxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVxAwrOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkAH9kq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzNCcrrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5DpUc63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTGGT0yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZG3Aj5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUEXfk76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func in6jE1fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kw1qQRwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlVTAGmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzcbafydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXdmtNXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfTGQzlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77dOL5RPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKRuQ0llWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFRqU5G4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulJnEdaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmbDJXT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiPpiWL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNFBF5EAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVe0qajMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TplOxlJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l90iHmmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRGcK2IuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcDtMQuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRr8xjZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtBEoCGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoA5cEipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGL07P7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhae7i6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXTSkh7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2UwFzQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPlEoDmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYYGbPQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CenbSuNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bG75MoZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfcVImVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6dF3XpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKG6J71OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lV2tBryoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkQTOt8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79CadOnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owJg8PQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UL3kB8Z1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTjGsSuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzCTzqhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUZsgslGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ju9hFGUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gr9DS1u5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7JqcAReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtuOle23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zz92sIEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NednLz98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYsPM7A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPTaYkX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5yOSRmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3LMyRl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhRJ6WF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsuP9UZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkDI5ZLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnzRLjFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UB77fu3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqC4DpKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS5qI6AAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4XZKSSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpxJrrbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRndYwBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3MZYDXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdlPjmyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4urPl7CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtofJCkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvhjG0hzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOxHJTt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1pbFqxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX3RRU99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfNk0QvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23hYBYmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQ41u86xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwliqNbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crQTvCjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzaGuHLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEpw7BFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gi1Z1bZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5b5gcvCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiUcyMy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IX65SfKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9io8FEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIxPjtmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYc5eeiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77azvXmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mrUSE8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRORXXGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QW2lHm14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPHGEXkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9mUPbXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfgyBxGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qi3if8MSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxs1leCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsZ81hUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKUdSgsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61KJieIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anZEXbEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiEpzRtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZKEYGGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjhEqYg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92hQzL4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwtb0dqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4ZR8O4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ct65VHLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UjWTdTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXcFGJTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cenyMfWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFHyOqieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDY9CfR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3ZC9BSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIzaaPFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqnefvYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxHml0pDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func editfj7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXNNL5vWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mr38xBxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfOyfKzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCWNJpXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xU5CSatUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func En3BeQ1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4a2kyHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIdEYRkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhyrVlXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y67tqcnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cset5MSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2smaW2HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzGDno9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKcg5o5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOZ6Hit0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQPpaRPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aH6JZ05EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJvoJCfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjYkIIjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSnTQQdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxUso2zQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiwWIbLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AirspmQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26C992s0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJhdYO9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tiwWp7oxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBtwpzLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77wQp99LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkpV6rrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKMQr6ekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRsPPmxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTVEMdo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvsmFes7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qxdz3T4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yd6p5CzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsX1pYuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lc5SscWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8sXCxbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iJjTkh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EX3EhnAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLxOWnYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ykl9r45vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DbPbnfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6jrkAlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Y7sLipzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZG0etBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7AGdn5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz4aSExxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9nAFErcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5Y0xJlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crBoDoOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOFt8BJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9owJPBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wS0pbgV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgk4DkFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyYdB7bQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSn70WheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfkgREEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uUkKLsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KpX4jtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhSDQ8YTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffulZUJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy6IsVZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWTglPqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6kocde3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZiZ9DpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1ZeFhOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nn6nVoiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCrY6M20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGM973xWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHHIIEn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y6H7TcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vmjo2mORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8SAdJE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0rXS1MJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPFdWw21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouLhf1MhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3J4M9QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlO2WX4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0vJAThPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ur1ElMcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qDKutNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOA0OiPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atoiJwoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEDh84z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vnoehAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGoMK5YFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XklsIfuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C28vXpDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 254F6KWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeJW6DMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObRnid5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkZ6MRl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RV6vV3zxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMx2AC6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyWJ9RoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhHEI2HXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGsaNtPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioHByTDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJmVCarbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCJZ3hV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sY0Dhc34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvSldF3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLVLQMcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5miE0rDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shNGO3RCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxO93BmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSqlxTUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5jO65ADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1LB58X9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8M1WkEn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3VCqrmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46oxas70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qn08qURGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbRpFVo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjTCVQywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rW2EYZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H02V60rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzMofa2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeUJLFfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KVEXwr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zErRqUUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYI6NPmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7c8xWTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSX3ssECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQEA3sV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

