package simulator

import "context"

var CmdIpsecIkeKeepaliveLog = &CommandSpec{
	[]Token{TIpsec, TIke, TKeepalive, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeKeepaliveLog = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TKeepalive, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

