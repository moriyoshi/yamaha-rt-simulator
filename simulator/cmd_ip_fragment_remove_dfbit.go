package simulator

import "context"

var CmdIpFragmentRemoveDfBit = &CommandSpec{
	[]Token{TIp, TFragment, TRemove, TDfBit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFragmentRemove = &CommandSpec{
	[]Token{TNo, TIp, TFragment, TRemove, TDfBit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

