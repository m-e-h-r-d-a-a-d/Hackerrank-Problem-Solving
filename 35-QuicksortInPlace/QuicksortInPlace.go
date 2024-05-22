    count := make([][]string, 100)
    size := len(arr)/2
    for i, v := range arr {
        if i < size{
            count[i] = append(count[i], "-")
        } else {
            count[i] = append(count[i], v...)
        }
        
        count[i] = append(count[i], " ")
    }
    sum := int32(0)
    for i := 0; i < 100; i++ {
        sum += count[i]
        count[i] = sum
    }