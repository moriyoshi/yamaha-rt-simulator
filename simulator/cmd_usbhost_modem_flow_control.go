package simulator

import "context"

var CmdUsbhostModemFlowControl = &CommandSpec{
	[]Token{TUsbhost, TModem, TFlow, TControl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostModemFlowControl = &CommandSpec{
	[]Token{TNo, TUsbhost, TModem, TFlow, TControl}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

