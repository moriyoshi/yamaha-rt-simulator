package simulator

import "context"

var CmdBridgeMember = &CommandSpec{
	[]Token{TBridge, TMember}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBridgeMember = &CommandSpec{
	[]Token{TNo, TBridge, TMember}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

