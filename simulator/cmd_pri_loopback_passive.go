package simulator

import "context"

var CmdPriLoopbackPassive = &CommandSpec{
	[]Token{TPri, TLoopback, TPassive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdPriLoopbackPassivePriOff = &CommandSpec{
	[]Token{TPri, TLoopback, TPassive, TSomething /* pri */, TOff}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
