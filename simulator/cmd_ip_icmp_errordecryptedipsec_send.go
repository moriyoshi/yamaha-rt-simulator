package simulator

import "context"

var CmdIpIcmpErrorDecryptedIpsecSend = &CommandSpec{
	[]Token{TIp, TIcmp, TErrorDecryptedIpsec, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpErrorDecryptedIpsecSend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TErrorDecryptedIpsec, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

