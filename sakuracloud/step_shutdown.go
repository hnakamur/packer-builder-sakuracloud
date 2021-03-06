package sakuracloud

import (
	"fmt"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/sacloud/libsacloud/api"
)

type stepShutdown struct {
	Debug bool
}

func (s *stepShutdown) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*api.Client)
	ui := state.Get("ui").(packer.Ui)
	serverID := state.Get("server_id").(int64)

	ui.Say("Gracefully shutting down server...")

	_, err := client.Server.Shutdown(serverID)
	if err != nil {
		err := fmt.Errorf("Error shutting down server: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	client.Server.SleepUntilDown(serverID, 1*time.Minute)

	return multistep.ActionContinue
}

func (s *stepShutdown) Cleanup(state multistep.StateBag) {
	// no cleanup
}
