package handler

import (
	"log"
	"time"
)

func (h handler) TrnConfirmer(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			h.service.Repository().TrnConfirmer()
			log.Println("Обработчик тразнакций завершил свою работу")
		}
	}
}
