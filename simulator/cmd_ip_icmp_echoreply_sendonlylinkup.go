package simulator

import "context"

var CmdIpIcmpEchoReplySendOnlyLinkup = &CommandSpec{
	[]Token{TIp, TIcmp, TEchoReply, TSendOnlyLinkup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpEchoReplySendOnlyLinkup = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TEchoReply, TSendOnlyLinkup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

