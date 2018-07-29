package simulator

import "context"

var CmdShowStatusIpKeepalive = &CommandSpec{
	[]Token{TShow, TStatus, TIp, TKeepalive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

