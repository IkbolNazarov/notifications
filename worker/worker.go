package worker

import (
	"fmt"
	"notifications/entities"
	"notifications/usecases"
	"time"
)

type Worker struct {
	EventUsecase usecases.EventUsecase
	Interval     time.Duration
	QuitChan     chan bool
}

func NewWorker(eventUsecase usecases.EventUsecase, interval time.Duration) *Worker {
	return &Worker{
		EventUsecase: eventUsecase,
		Interval:     interval,
		QuitChan:     make(chan bool),
	}
}

func (w *Worker) Start() {
	ticker := time.NewTicker(w.Interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.processEvents()
			case <-w.QuitChan:
				ticker.Stop()
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.QuitChan <- true
}

func (w *Worker) processEvents() {
	events, err := w.EventUsecase.GetPendingEvents()
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}
	for _, event := range events {
		w.notify(event)
		w.EventUsecase.RemoveEvent(event)
	}
}

func (w *Worker) notify(event *entities.Event) {
	fmt.Printf("notification: %+v\n", event)
}
