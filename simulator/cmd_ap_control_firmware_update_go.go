package simulator

import "context"

var CmdApControlFirmwareUpdateGo = &CommandSpec{
	[]Token{TAp, TControl, TFirmware, TUpdate, TGo}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

