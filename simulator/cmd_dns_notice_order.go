package simulator

import "context"

var CmdDnsNoticeOrder = &CommandSpec{
	[]Token{TDns, TNotice, TOrder}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsNoticeOrder = &CommandSpec{
	[]Token{TNo, TDns, TNotice, TOrder}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

