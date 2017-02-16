package middleware

import (
	"errors"
	"fmt"

	"sync"

	"github.com/chapin/spider/base"
)

var defaultChanLen uint = 10
var statusNameMap = map[ChannelManagerStatus]string{
	CHANNEL_MANAGER_STATUS_UNINITIALIZED: "uninitialized",
	CHANNEL_MANAGER_STATUS_INITIALIZED:   "initialzed",
	CHANNEL_MANAGER_STATUS_CLOSED:        "closed",
}
var chanmanSummaryTemplate = "status: %s, " +
	"requestChannel: %d/%d, " +
	"responseChannel: %d/%d, " +
	"ItemChannel: %d/%d, " +
	"errorChannel: %d/%d"

// myChannelManager struct .
type myChannelManager struct {
	channelLen uint                 // requestChnnal Length
	reqCh      chan base.Request    // request channel
	respCh     chan base.Response   // response channel
	itemCh     chan base.Item       // item channel
	errorCh    chan error           // error channel
	status     ChannelManagerStatus // ChannelManager status
	rwmutex    sync.RWMutex         // RWMutex
}

func NewChannelManager(channLen uint) ChannelManager {
	if channLen == 0 {
		channLen = defaultChanLen
	}

	chanman := &myChannelManager{}
	chanman.Init(channLen, true)

	return chanman
}

// Init method.
func (chanman *myChannelManager) Init(channelLen uint, reset bool) bool {
	if channelLen == 0 {
		panic(errors.New("The channel length is invalid!"))
	}

	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED && !reset {
		return false
	}

	chanman.channelLen = channelLen
	chanman.reqCh = make(chan base.Request, channelLen)
	chanman.respCh = make(chan base.Response, channelLen)
	chanman.itemCh = make(chan base.Item, channelLen)
	chanman.errorCh = make(chan error, channelLen)
	chanman.status = CHANNEL_MANAGER_STATUS_INITIALIZED
	return true
}

// Close method.
func (chanman *myChannelManager) Close() bool {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()

	if chanman.status != CHANNEL_MANAGER_STATUS_INITIALIZED {
		return false
	}

	close(chanman.reqCh)
	close(chanman.respCh)
	close(chanman.itemCh)
	close(chanman.errorCh)
	chanman.status = CHANNEL_MANAGER_STATUS_CLOSED

	return true
}

// ReqChan
func (chanman *myChannelManager) ReqChan() (chan base.Request, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.reqCh, nil
}

// RespChan
func (chanman *myChannelManager) RespChan() (chan base.Response, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.respCh, nil
}

// ItemChan
func (chanman *myChannelManager) ItemChan() (chan base.Item, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.itemCh, nil
}

// errorCh
func (chanman *myChannelManager) ErrorChan() (chan error, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.errorCh, nil
}

// Summary
func (chanman *myChannelManager) Summary() string {
	summary := fmt.Sprintf(chanmanSummaryTemplate,
		statusNameMap[chanman.status],
		len(chanman.reqCh), cap(chanman.reqCh),
		len(chanman.respCh), cap(chanman.respCh),
		len(chanman.itemCh), cap(chanman.itemCh),
		len(chanman.errorCh), cap(chanman.errorCh))

	return summary
}

// ChannelLen
func (chanman *myChannelManager) ChannelLen() uint {
	return chanman.channelLen
}

// Status
func (chanman *myChannelManager) Status() ChannelManagerStatus {
	return chanman.status
}

// status check
func (chanman *myChannelManager) checkStatus() error {
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED {
		return nil
	}
	statusName, ok := statusNameMap[chanman.status]
	if !ok {
		statusName = fmt.Sprintf("%d", chanman.status)
	}
	errMsg := fmt.Sprintf("The undesirable status of channel manager: %s!\n", statusName)
	return errors.New(errMsg)
}
