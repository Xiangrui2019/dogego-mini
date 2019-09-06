package global

import (
	"dogego-mini/models"
	"encoding/json"
	"fmt"
)

func AsyncTaskData(taskname string, data models.TaskData) string {
	v, _ := json.Marshal(data)
	return fmt.Sprintf("%s#$#%s", taskname, v)
}
