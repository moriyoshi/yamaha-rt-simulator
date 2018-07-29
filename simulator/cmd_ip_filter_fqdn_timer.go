package simulator

import "context"

var CmdIpFilterFqdnTimer = &CommandSpec{
	[]Token{TIp, TFilter, TFqdn, TTimer, TSomething /* time */, TAuto}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterFqdnTimerTime = &CommandSpec{
	[]Token{TNo, TIp, TFilter, TFqdn, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

