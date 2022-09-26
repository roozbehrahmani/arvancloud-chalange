package charge_code_dispatcher

import (
	"github.com/roozbehrahmani/abrarvan_test/config"
	charge_code_job "github.com/roozbehrahmani/abrarvan_test/internal/queue/jobs/chargeCode"
)

func Initialize(cnf *config.Config) *ChargeWalletDispatcher {

	workerPool := make(chan chan charge_code_job.ChargeWalletJob, cnf.MaxWorkerForChargeWallet)

	dispatcher := &ChargeWalletDispatcher{
		JobQueue:   make(chan charge_code_job.ChargeWalletJob),
		maxWorkers: cnf.MaxWorkerForChargeWallet,
		workerPool: workerPool,
	}
	dispatcher.Run()
	return dispatcher
}
