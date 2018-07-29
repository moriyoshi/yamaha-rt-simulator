package simulator

import "context"

var CmdNgnRenumberingLinkRefresh = &CommandSpec{
	[]Token{TNgn, TRenumbering, TLinkRefresh}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNgnRenumberingLinkRefresh = &CommandSpec{
	[]Token{TNo, TNgn, TRenumbering, TLinkRefresh}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

