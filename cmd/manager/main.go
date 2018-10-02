package manager

import (
	"flag"
	"github.com/kubermatic/cluster-api-provider-digitalocean/pkg/cloud/digitalocean/actuators/machine/machinesetup"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	flag.Parse()

	// Get configuration for talking to the API server
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create new command and start components
	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Initializing dependencies.
	log.Print("Initializing dependencies...")
	err = initStaticDeps(mgr)
	if err != nil {
		log.Fatal(err)
	}

	// Registering components.
	log.Print("Registering components...")

}

// initStaticDeps initializes and set ups static dependencies.
// TODO: Find a way to improve this function.
func initStaticDeps(mgr manager.Manager) error {
	// TODO: This should not be static.
	machineSetupConfigPath := "/etc/machineconfig/config.yaml"

	configWatch, err := machinesetup.NewConfigWatch(machineSetupConfigPath)
	if err != nil {
		return err
	}

	// TODO: Finish this.
}