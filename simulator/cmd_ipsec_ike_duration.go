package simulator

import "context"

var CmdIpsecIkeDuration = &CommandSpec{
	[]Token{TIpsec, TIke, TDuration}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeDuration = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TDuration}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

