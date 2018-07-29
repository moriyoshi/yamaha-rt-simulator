package simulator

import "context"

var CmdTunnelBackup = &CommandSpec{
	[]Token{TTunnel, TBackup, &EitherToken{TPp, TTunnel}}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoTunnelBackup = &CommandSpec{
	[]Token{TNo, TTunnel, TBackup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
