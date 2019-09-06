package executers

import (
	"dogego-mini/global"
	"dogego-mini/modules"
	"errors"
	"log"
	"time"
)

func TimeExecuter() {
	modules.RedisMQModule.Custome(
		global.TimeTaskQueueKey(),
		executeTask,
	)
}

func executeTask(message string) error {
	for _, item := range modules.TasksModule {
		if item.Taskname == message {
			defer func() {
				err := recover()
				if err != nil {
					modules.LockerModule.Free(item.Taskname)
					panic(err)
				}
			}()

			if item.Type != modules.TimeJob {
				return errors.New("Job type error.")
			}

			if !modules.LockerModule.Lock(item.Taskname, 0) {
				return errors.New("Lock error.")
			}

			job := item.Job.(modules.TimeTask)

			from := time.Now().UnixNano()
			err := job()
			to := time.Now().UnixNano()

			if err != nil {
				log.Printf("%s Execute Error: %dms\n", item.Taskname, (to-from)/int64(time.Millisecond))
			} else {
				log.Printf("%s Execute Success: %dms\n", item.Taskname, (to-from)/int64(time.Millisecond))
			}

			modules.LockerModule.Free(item.Taskname)
		}
	}

	return nil
}
