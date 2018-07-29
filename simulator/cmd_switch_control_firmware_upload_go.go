package simulator

import "context"

var CmdSwitchControlFirmwareUploadGo = &CommandSpec{
	[]Token{TSwitch, TControl, TFirmware, TUpload, TGo}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

