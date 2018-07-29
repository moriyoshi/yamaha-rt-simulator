package simulator

import "context"

var CmdQacTmUnqualifiedClientAccessControl = &CommandSpec{
	[]Token{TQacTm, TUnqualified, TClient, TAccess, TControl}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmUnqualifiedClientAccessControl = &CommandSpec{
	[]Token{TNo, TQacTm, TUnqualified, TClient, TAccess, TControl}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

