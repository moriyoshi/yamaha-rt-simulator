package simulator

import "context"

var CmdIpsecRefreshSa = &CommandSpec{
	[]Token{TIpsec, TRefresh, TSa}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

