package task

// Status - Tasks status
type Status struct {
	Worker   int  `json:"worker"`
	Waiting  int  `json:"waiting"`
	Count    int  `json:"count"`
	Free     int  `json:"free"`
	Running  int  `json:"running"`
	IsClosed bool `json:"is_closed"`
}

// GetStatus get tasks status
//
//	@receiver t *Tasks
//	@return Status
func (t *Tasks) GetStatus() Status {
	// t.mu.Lock()
	// defer t.mu.Unlock()
	st := Status{
		Worker:   t.Cap(),
		Waiting:  t.Waiting(),
		Free:     t.Free(),
		Running:  t.Running(),
		IsClosed: t.IsClosed(),
	}
	taskCount := t.GetCount()
	if s := (taskCount - st.Running); st.Waiting < s {
		st.Waiting = s
	}
	return st
}
