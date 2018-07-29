package simulator

import "context"

var CmdSshdHostKeyGenerate = &CommandSpec{
	[]Token{TSshd, THost, TKey, TGenerate}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdHostKeyGenerate = &CommandSpec{
	[]Token{TNo, TSshd, THost, TKey, TGenerate}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

