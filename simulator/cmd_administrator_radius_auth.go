package simulator

import "context"

var CmdAdministratorRadiusAuth = &CommandSpec{
	[]Token{TAdministrator, TRadius, TAuth}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAdministratorRadiusAuth = &CommandSpec{
	[]Token{TNo, TAdministrator, TRadius, TAuth}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

