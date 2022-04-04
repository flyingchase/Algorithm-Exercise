package channelimplemetsemaphore

type Seamphore chan struct{}

func (s Seamphore) Acquire() {
	s <- struct{}{}
}
func (s Seamphore) TryAcquire() bool {
	select {
	case s <- struct{}{}:
		// 还有
		return true
	default:
		return false
	}
}
func (s Seamphore) Release() {
	<-s
}
