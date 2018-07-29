package simulator

import "context"

var CmdIpv6Stealth = &CommandSpec{
	[]Token{TIpv6, TStealth}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6StealthInterfaceInterface = &CommandSpec{
	[]Token{TIpv6, TStealth}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6Stealth = &CommandSpec{
	[]Token{TNo, TIpv6, TStealth}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

