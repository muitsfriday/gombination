# gombination


## Usage

### Generate PermutateType
- with number `PermutateNum(5)`
- with slice `PermutateSlice([]interface{}{"a", "b", "c"})`

### Count the permutation
```Go
c := PermutateNum(5).Count() 
// 120
```

### Generate permutation from slice
```Go
p := PermutateSlice([]interface{}{"a", "b", "c"}).List()
// p is [a,b,c b,a,c c,a,b a,c,b b,c,a c,b,a]
```

