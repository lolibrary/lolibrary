package cmd

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/cmd/flower/rpc"
)

var (
	// EnvProd is where we'll spin up a port-forward for service/edge-proxy-internal.
	// edge-proxy.internal is able to route requests to any defined internal service.
	EnvProd = "localhost:8080"

	// EnvProdAPI is where we'll connect to service/edge-proxy-external via our ingress.
	// edge-proxy.external can only route to service.api.* services.
	EnvProdAPI = "https://external-api.lolibrary.org"

	// local service/edge-proxy-internal
	EnvLocal = "localhost:18080"

	// local service/edge-proxy-external
	EnvLocalAPI = "http://localhost:18081"
)

func environment() error {
	switch env {
	case "local", "":
		env = "local"
		fmt.Printf("🌸  Using environment %v\n\n", aurora.Blue("local"))
		rpc.SetInternalEdgeProxy(EnvLocal)
		rpc.SetExternalEdgeProxy(EnvLocalAPI)
	case "prod", "production":
		env = "prod"
		fmt.Printf("🌸  Using environment %v\n\n", aurora.Blue("prod"))
		rpc.SetInternalEdgeProxy(EnvProd)
		rpc.SetExternalEdgeProxy(EnvProdAPI)

		// fmt.Printf("➡️  Enabling port-forward to %v...", aurora.Blue("service/edge-proxy-internal"))
		// if err := commands.Enable(); err != nil {
		// 	fmt.Printf(" ❌\n\n")
		// 	return err
		// } else {
		// 	fmt.Printf(" ✅\n\n")
		// }
	default:
		return fmt.Errorf("⚠️  Valid environments are %v and %v",
			aurora.Green("local"), aurora.Green("prod"))
	}

	return nil
}

func setEnvironment(input string) bool {
	env = input
	if err := environment(); err != nil {
		fmt.Printf("❌  Error setting environment:\n%v", err)
		return false
	}

	return true
}
