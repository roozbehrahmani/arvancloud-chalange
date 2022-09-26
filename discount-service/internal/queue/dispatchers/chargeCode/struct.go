package charge_code_dispatcher

import (
	charge_code_job "github.com/roozbehrahmani/abrarvan_test/internal/queue/jobs/chargeCode"
)

///////

////////

func NewChargeWalletDispatcher(jobQueue chan charge_code_job.ChargeWalletJob, maxWorkers int) *ChargeWalletDispatcher {
	workerPool := make(chan chan charge_code_job.ChargeWalletJob, maxWorkers)
	return &ChargeWalletDispatcher{
		JobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

type ChargeWalletDispatcher struct {
	workerPool chan chan charge_code_job.ChargeWalletJob
	maxWorkers int
	JobQueue   chan charge_code_job.ChargeWalletJob
}
