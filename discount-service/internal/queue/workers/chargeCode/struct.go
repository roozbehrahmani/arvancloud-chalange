package charge_code_worker

import (
	charge_code_job "github.com/roozbehrahmani/abrarvan_test/internal/queue/jobs/chargeCode"
)

///////

func NewChargeWalletWorker(id int, workerPool chan chan charge_code_job.ChargeWalletJob) ChargeWalletWorker {
	return ChargeWalletWorker{
		id:         id,
		jobQueue:   make(chan charge_code_job.ChargeWalletJob),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

type ChargeWalletWorker struct {
	id         int
	jobQueue   chan charge_code_job.ChargeWalletJob
	workerPool chan chan charge_code_job.ChargeWalletJob
	quitChan   chan bool
}
