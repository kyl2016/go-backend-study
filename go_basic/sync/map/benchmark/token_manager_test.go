package token_manager

import (
	"sync"
	"testing"
)

// 适用 sync.Map 要比直接使用 map 快近 10 秒，sync.Map 适用于读多写少的场景

func TestAddToken(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				FindToken("alsdjlasdfjasldjf")
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				AddToken("alsdjlasdfjasldjf")
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestAddTokenToSyncMap(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				FindTokenFromSyncMap("alsdjlasdfjasldjf")
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				AddTokenToSyncMap("alsdjlasdfjasldjf")
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
