package simulator

import "context"

var CmdIpsecIkeModeCfgMethod = &CommandSpec{
	[]Token{TIpsec, TIke, TModeCfg, TMethod}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeModeCfgMethod = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TModeCfg, TMethod}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

