package simulator

import "context"

var CmdIpArpTimer = &CommandSpec{
	[]Token{TIp, TArp, TTimer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpArpTimerTimerRetry = &CommandSpec{
	[]Token{TNo, TIp, TArp, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

