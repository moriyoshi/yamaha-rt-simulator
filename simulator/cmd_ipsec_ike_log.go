package simulator

import "context"

var CmdIpsecIkeLog = &CommandSpec{
	[]Token{TIpsec, TIke, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeLog = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

