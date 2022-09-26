package charge_code_dispatcher

import (
	"fmt"
	charge_code_worker "github.com/roozbehrahmani/abrarvan_test/internal/queue/workers/chargeCode"
)

func (d *ChargeWalletDispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := charge_code_worker.NewChargeWalletWorker(i+1, d.workerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *ChargeWalletDispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				fmt.Printf("fetching workerJobQueue for: %s\n", job.Name)
				workerJobQueue := <-d.workerPool
				fmt.Printf("adding %s to workerJobQueue\n", job.Name)
				workerJobQueue <- job
			}()
		}
	}
}
