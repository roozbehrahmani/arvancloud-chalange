package charge_code_worker

import "fmt"

func (w ChargeWalletWorker) Start() {
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to my jobQueue.
				reponse, err := job.DoWork(job.ChargeCode, job.User)
				if err != nil {
					fmt.Printf("job fialed :(  err:%v", err)
				} else {
					fmt.Printf("job done :)     response:%v", reponse)
				}

			case <-w.quitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.id)
				return
			}
		}
	}()
}

func (w ChargeWalletWorker) stop() {
	go func() {
		w.quitChan <- true
	}()
}
