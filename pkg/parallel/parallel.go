package parallel

import (
	"fmt"
	"sync"
)

type Parallel struct {
	wg sync.WaitGroup
}
