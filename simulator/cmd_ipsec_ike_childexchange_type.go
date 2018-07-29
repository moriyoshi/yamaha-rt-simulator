package simulator

import "context"

var CmdIpsecIkeChildExchangeType = &CommandSpec{
	[]Token{TIpsec, TIke, TChildExchange, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeChildExchangeType = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TChildExchange, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

