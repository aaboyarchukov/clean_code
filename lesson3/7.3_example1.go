package lesson3

import (
	"fmt"
	"math/rand"
	"time"
)

func (restaurant *Restaurant) run(duration time.Duration) {
	stopTimer := time.After(duration)

	// 7.3
	for i := 0; i < restaurant.waiters; i++ {
		restaurant.wg.Add(1)
		go restaurant.waiter(i + 1)
	}

	// 7.3
	for i := 0; i < restaurant.chefs; i++ {
		restaurant.wg.Add(1)
		go restaurant.chef(i + 1)
	}

	go func() {
		fmt.Println("заказ создается")
		orderID := 1
		for {
			select {
			case <-stopTimer:
				close(restaurant.orders)
				close(restaurant.preparedOrders)
				restaurant.wg.Done()
				return
			default:
				order := Order{
					id:    orderID,
					dish:  dishes[rand.Intn(len(dishes))],
					table: rand.Intn(10) + 1,
				}
				fmt.Printf("Новый заказ #%d: %s (стол %d)\n", order.id, order.dish, order.table)
				restaurant.orders <- order
				orderID++
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+500))
			}
		}
	}()
}

func (restaurant *Restaurant) printWaiterStats() {
	fmt.Println("\nСтатистика работы официантов:")

	// 7.3
	for waiterID, stats := range restaurant.waiterStats {
		stats.mu.Lock()
		tablesCount := len(stats.tablesServed)
		ordersCount := stats.ordersCount
		stats.mu.Unlock()

		fmt.Printf("Официант #%d: обслужил %d столов, принял %d заказов\n",
			waiterID, tablesCount, ordersCount)
	}
}
