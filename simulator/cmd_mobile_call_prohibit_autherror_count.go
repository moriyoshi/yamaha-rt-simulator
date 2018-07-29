package simulator

import "context"

var CmdMobileCallProhibitAuthErrorCount = &CommandSpec{
	[]Token{TMobile, TCall, TProhibit, TAuthError, TCount}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileCallProhibitAuthErrorCount = &CommandSpec{
	[]Token{TNo, TMobile, TCall, TProhibit, TAuthError, TCount}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

