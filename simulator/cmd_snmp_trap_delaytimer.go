package simulator

import "context"

var CmdSnmpTrapDelayTimer = &CommandSpec{
	[]Token{TSnmp, TTrap, TDelayTimer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdSnmpTrapDelayTimerOff = &CommandSpec{
	[]Token{TSnmp, TTrap, TDelayTimer, TOff}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapDelayTimer = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TDelayTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

