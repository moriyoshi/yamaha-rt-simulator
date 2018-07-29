package simulator

import "context"

var CmdIpHost = &CommandSpec{
	[]Token{TIp, THost}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdDnsStaticTypeNameValueTtl = &CommandSpec{
	[]Token{TDns, TStatic}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpHostFqdnValue = &CommandSpec{
	[]Token{TNo, TIp, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsStaticTypeNameValue = &CommandSpec{
	[]Token{TNo, TDns, TStatic}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

