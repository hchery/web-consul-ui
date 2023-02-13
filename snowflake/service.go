package snowflake

import (
	"fmt"
	"sync"
	"time"
)

/**********************************
 * Date: 2023/2/2
 * Author: hchery
 * Home: https://github.com/hchery
 *********************************/

const (
	machineBits  = int64(5)  //机器id位数
	serviceBits  = int64(5)  //服务id位数
	sequenceBits = int64(12) //序列id位数

	maxMachineID  = int64(-1) ^ (int64(-1) << machineBits)  //最大机器id
	maxServiceID  = int64(-1) ^ (int64(-1) << serviceBits)  //最大服务id
	maxSequenceID = int64(-1) ^ (int64(-1) << sequenceBits) //最大序列id

	timeLeft    = uint8(22) //时间id向左移位的量
	machineLeft = uint8(17) //机器id向左移位的量
	serviceLeft = uint8(12) //服务id向左移位的量

	initMs = int64(1640966400000) //初始毫秒,时间是: Wed Nov  9 13:40:27 CST 2022
)

type Worker struct {
	sync.Mutex
	lastTimestamp int64
	machineId     int64
	serviceId     int64
	sequenceId    int64
}

type SFError struct {
	desc string
}

func (e *SFError) Error() string {
	return e.desc
}

func NewWorker(serviceId int64) (*Worker, *SFError) {
	if serviceId < 0 || serviceId > maxServiceID {
		return nil, &SFError{
			desc: fmt.Sprintf("Invalid service id: %d", serviceId),
		}
	}
	return &Worker{
		lastTimestamp: 0,
		machineId:     props.Machine.Id,
		serviceId:     serviceId,
		sequenceId:    0,
	}, nil
}

func (w *Worker) Next() int64 {
	w.Lock()
	defer w.Unlock()
	mill := time.Now().UnixMilli()
	if mill == w.lastTimestamp {
		w.sequenceId = (w.sequenceId + 1) & maxSequenceID
		if w.sequenceId == 0 {
			for mill > w.lastTimestamp {
				mill = time.Now().UnixMilli()
			}
		}
	} else {
		w.sequenceId = 0
	}
	w.lastTimestamp = mill
	return (w.lastTimestamp-initMs)<<timeLeft |
		w.machineId<<machineLeft |
		w.serviceId<<serviceLeft |
		w.sequenceId
}
