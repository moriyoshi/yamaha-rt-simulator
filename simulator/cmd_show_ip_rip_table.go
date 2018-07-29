package simulator

import "context"

var CmdShowIpRipTable = &CommandSpec{
	[]Token{TShow, TIp, TRip, TTable}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

