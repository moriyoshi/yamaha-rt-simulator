package simulator

import "context"

var CmdIpInterfaceIntrusionDetectionRepeatControl = &CommandSpec{
	[]Token{TIp, TL2Interface, TIntrusion, TDetection, TRepeatControl}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIntrusionDetectionRepeatControl = &CommandSpec{
	[]Token{TIp, TPp, TIntrusion, TDetection, TRepeatControl}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelIntrusionDetectionRepeatControl = &CommandSpec{
	[]Token{TIp, TTunnel, TIntrusion, TDetection, TRepeatControl}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceIntrusionDetectionRepeatControl = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIntrusion, TDetection, TRepeatControl}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIntrusionDetectionRepeatControl = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIntrusion, TDetection, TRepeatControl}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelIntrusionDetectionRepeatControl = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIntrusion, TDetection, TRepeatControl}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

