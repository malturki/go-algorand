package agreement

import (
	"sort"
	"time"

	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/protocol"
	"github.com/algorand/go-algorand/util/timers"
)

// clockManager managers multiple clocks used by different pipelined rounds.
// XXX garbage-collect old rounds
type clockManager struct {
	m  map[round]timers.Clock
	t0 timers.Clock
}

func makeClockManager(t0 timers.Clock) *clockManager {
	return &clockManager{m: make(map[round]timers.Clock), t0: t0}
}

func (cm *clockManager) setZero(r round) {
	cm.m[r] = cm.t0.Zero()
}

// nextDeadlineCh returns a timeout channel that will fire when the earliest Deadline among all of
// the rounds described in externalDemuxSignals has occurred. It also returns the corresponding
// round (including speculative branch) this timeout channel corresponds to.
func (cm *clockManager) nextDeadlineCh(es []externalDemuxSignals) (<-chan time.Time, round) {
	if len(es) == 0 {
		return nil, roundZero
	}
	sort.Slice(es, func(i, j int) bool {
		ti := cm.m[es[i].CurrentRound].GetTimeout(es[i].Deadline)
		tj := cm.m[es[j].CurrentRound].GetTimeout(es[j].Deadline)
		return ti.Before(tj)
	})
	r := es[0].CurrentRound
	c, ok := cm.m[r]
	if !ok {
		// no rezeroAction has set up this clock yet
		// XXX this should probably only happen at service bootstrap or in tests
		logging.Base().Warnf("clockManager.nextDeadlineCh making new clock for round %+v", r)
		c = cm.t0.Zero()
		cm.m[r] = c
	}
	return c.TimeoutAt(es[0].Deadline), r
}

// nextFastDeadlineCh returns a timeout channel that will fire when the earliest FastRecoveryDeadline among all of
// the rounds described in externalDemuxSignals has occurred. It also returns the corresponding
// round (including speculative branch) this timeout channel corresponds to.
func (cm clockManager) nextFastDeadlineCh(es []externalDemuxSignals) (<-chan time.Time, round) {
	if len(es) == 0 {
		return nil, roundZero
	}
	sort.Slice(es, func(i, j int) bool {
		ti := cm.m[es[i].CurrentRound].GetTimeout(es[i].FastRecoveryDeadline)
		tj := cm.m[es[j].CurrentRound].GetTimeout(es[j].FastRecoveryDeadline)
		return ti.Before(tj)
	})
	return cm.m[es[0].CurrentRound].TimeoutAt(es[0].FastRecoveryDeadline), es[0].CurrentRound
}

type clockManagerSerialized struct {
	Clocks []struct{ R, C []byte }
}

func (cm *clockManager) Decode(data []byte) (*clockManager, error) {
	var s clockManagerSerialized
	err := protocol.DecodeReflect(data, &s)
	if err != nil {
		logging.Base().Errorf("clockManager decode error: %v", err)
		return nil, err
	}
	m := make(map[round]timers.Clock)
	for _, rc := range s.Clocks {
		var r round
		err := protocol.DecodeReflect(rc.R, &r)
		if err != nil {
			return nil, err
		}
		clk, err := cm.t0.Decode(rc.C)
		if err != nil {
			return nil, err
		}
		m[r] = clk
	}
	return &clockManager{m: m, t0: cm.t0}, err
}

func (cm *clockManager) Encode() []byte {
	var s clockManagerSerialized
	for r, c := range cm.m {
		s.Clocks = append(s.Clocks, struct{ R, C []byte }{
			R: protocol.EncodeReflect(r),
			C: c.Encode(),
		})
	}
	return protocol.EncodeReflect(&s)
}