package simulator

import "context"

var CmdIpInterfaceRipAuthKey = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TAuth, TKey}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipAuthKey = &CommandSpec{
	[]Token{TIp, TPp, TRip, TAuth, TKey}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipAuthKey = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TAuth, TKey}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpInterfaceRipAuthKeyText = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TAuth, TKey, TText}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipAuthKeyText = &CommandSpec{
	[]Token{TIp, TPp, TRip, TAuth, TKey, TText}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipAuthKeyText = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TAuth, TKey, TText}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipAuthKey = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TAuth, TKey}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipAuthKey = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TAuth, TKey}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipAuthKey = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TAuth, TKey}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipAuthKeyText = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TAuth, TKey, TText}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipAuthKeyText = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TAuth, TKey, TText}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipAuthKeyText = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TAuth, TKey, TText}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

