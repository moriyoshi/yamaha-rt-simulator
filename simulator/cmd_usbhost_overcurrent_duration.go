package simulator

import "context"

var CmdUsbhostOvercurrentDuration = &CommandSpec{
	[]Token{TUsbhost, TOvercurrent, TDuration}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostOvercurrentDuration = &CommandSpec{
	[]Token{TNo, TUsbhost, TOvercurrent, TDuration}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

