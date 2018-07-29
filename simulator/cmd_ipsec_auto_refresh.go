package simulator

import "context"

var CmdIpsecAutoRefresh = &CommandSpec{
	[]Token{TIpsec, TAuto, TRefresh}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecAutoRefresh = &CommandSpec{
	[]Token{TNo, TIpsec, TAuto, TRefresh}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

