package simulator

import "context"

var CmdUsbhostModemInitialize = &CommandSpec{
	[]Token{TUsbhost, TModem, TInitialize}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostModemInitialize = &CommandSpec{
	[]Token{TNo, TUsbhost, TModem, TInitialize}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

