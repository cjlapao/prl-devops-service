package cmd

import (
	"fmt"

	"github.com/Parallels/prl-devops-service/constants"
)

func processHelp(command string) {
	fmt.Printf("%v\n", constants.Name)
	fmt.Printf("\n")
	fmt.Printf("  Find out more at: https://github.com/Parallesl/prl-devops-service\n")
	fmt.Printf("\n")
	switch command {
	case constants.API_COMMAND:
		processApiHelp()
	case constants.CATALOG_COMMAND:
		processCatalogHelp()
	case constants.TEST_COMMAND:
		processTestHelp()
	case constants.GENERATE_SECURITY_KEY_COMMAND:
		processGenerateSecurityKeyHelp()
	case constants.REVERSE_PROXY_COMMAND:
		processReverseProxyHelp()
	default:
		processDefaultHelp()
	}
	fmt.Printf("\n")
}

func processDefaultHelp() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\n")
	fmt.Printf("  %v [command] [flags]\n", constants.ExecutableName)
	fmt.Printf("\n")
	fmt.Printf("Available Commands:\n")
	fmt.Printf("\n")
	fmt.Printf("  %s\t\t Starts the API Service\n", constants.API_COMMAND)
	fmt.Printf("  %s\t Starts the Reverse Proxy Service\n", constants.REVERSE_PROXY_COMMAND)
	fmt.Printf("  %s\t Prints the API Catalog\n", constants.CATALOG_COMMAND)
	fmt.Printf("  %s\t Generates a new Security Key\n", constants.GENERATE_SECURITY_KEY_COMMAND)
	fmt.Printf("  %s\t Installs the API Service\n", constants.INSTALL_SERVICE_COMMAND)
	fmt.Printf("  %s\t Uninstalls the API Service\n", constants.UNINSTALL_SERVICE_COMMAND)
	fmt.Printf("  %s\t\t Tests the Remote providers\n", constants.TEST_COMMAND)
	fmt.Printf("  %s\t Prints the Version\n", constants.VERSION_COMMAND)
	fmt.Printf("  %s\t\t Prints this help\n", constants.HELP_COMMAND)
}
