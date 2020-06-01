package msg

import (
	"time"
)

const (
	//pongWait = 60 * time.Second
	pongWait = 30 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	pingMsg = "ping"
	pongMsg = "pong"
)

func (h *hub_struct) set_ping() {
	timer := time.NewTimer(pingPeriod)
	ping_err_count := 0
	for {
		select {
		case <-timer.C:
			if ping_err_count < 3 {
				if h.State == 1 {
					return
				}
				ping_err_count++
				h.send_msg([]byte(pingMsg))
				timer.Reset(pingPeriod)
			}else {
				h.Death()
			}
		case <-h.pong:
			ping_err_count = 0
			timer.Reset(pingPeriod)
		}

	}
}
func check_msg_is_pong(msg []byte) bool {
	if len(msg) == len([]byte(pongMsg)) {
		if string(msg) == pongMsg {
			return true
		}
	}
	return false
}