package simulator

import "context"

var CmdIsdnCallProhibitTime = &CommandSpec{
	[]Token{TIsdn, TCall, TProhibit, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallProhibitTime = &CommandSpec{
	[]Token{TNo, TIsdn, TCall, TProhibit, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

