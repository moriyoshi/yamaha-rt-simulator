package simulator

import "context"

var CmdIpv6IcmpEchoReplySendOnlyLinkup = &CommandSpec{
	[]Token{TIpv6, TIcmp, TEchoReply, TSendOnlyLinkup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpEchoReplySendOnlyLinkup = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TEchoReply, TSendOnlyLinkup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

