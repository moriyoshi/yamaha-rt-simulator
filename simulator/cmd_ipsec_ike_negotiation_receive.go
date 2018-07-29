package simulator

import "context"

var CmdIpsecIkeNegotiationReceive = &CommandSpec{
	[]Token{TIpsec, TIke, TNegotiation, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeNegotiationReceive = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TNegotiation, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

