package simulator

import "context"

var CmdIpsecIkePayloadType = &CommandSpec{
	[]Token{TIpsec, TIke, TPayload, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkePayloadType = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TPayload, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

