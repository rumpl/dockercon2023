target "lint" {
    target = "lint"
    output = ["type=cacheonly"]
}

target "binary" {
    target = "binary"
    output = ["./bin"]
}
