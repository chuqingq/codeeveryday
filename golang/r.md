```
package main

import (
        "log"
        "time"
)

func main() {
        start := time.Now()
        for i := 0; i < 1000000000; i++ {
        }   
        log.Printf("%v", time.Since(start))
}
```

```
$ time ./my 
2017/08/28 14:09:17 333.116064ms

real	0m0.335s
user	0m0.332s
sys	0m0.000s
```



```
            long startTime=System.currentTimeMillis();
            for (long i = 0; i < 1000000000; i++) { // Long.MAX_VALUE
            }   
            long endTime=System.currentTimeMillis();
            System.out.println(endTime - startTime);
```

```
335
```

