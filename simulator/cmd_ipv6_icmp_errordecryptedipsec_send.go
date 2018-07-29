package simulator

import "context"

var CmdIpv6IcmpErrorDecryptedIpsecSend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TErrorDecryptedIpsec, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpErrorDecryptedIpsecSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TErrorDecryptedIpsec, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

