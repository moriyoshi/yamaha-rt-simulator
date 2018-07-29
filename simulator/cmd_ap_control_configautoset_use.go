package simulator

import "context"

var CmdApControlConfigAutoSetUse = &CommandSpec{
	[]Token{TAp, TControl, TConfigAutoSet, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoApControlConfigAutoSetUse = &CommandSpec{
	[]Token{TNo, TAp, TControl, TConfigAutoSet, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

