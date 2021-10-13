func f(n:Int) -> Int{
    var sum = 0
    for i in 0..<n{
        sum += i
    }
    return sum
}

print(f(n:2<<25))


// 2251799780130816

// real    0m19.554s
// user    0m19.540s
// sys     0m0.007s
