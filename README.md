# SMFP

'smfp' is short for 'simple functional programming'.  
so we aim for a functional programming library that can be used easily and universally.

## Getting Started

```go
// sum is 15
sum := smfp.Reduce(func(acummulator int, item int, index int) int {
    return acummulator + item
}, func() int {
    return 0
}).Execute([]int{
    1, 2, 3, 4, 5
}).(int)
```

### Pipe Connect

```go
// results is []int{15, 20, 25}
results := smfp.Pipe[[]string]([]*smfp.Worker{
    smfp.Map(func(item string, index int) int {
        num, _ := strconv.Atoi(item)
        return num
    }),
    smfp.Filter(func(item int, index int) bool {
        return item > 14
    }),
}).Execute([]string{
    "1", "5", "10", "15", "20", "25"
}).([]int)
```
