package main

import (
	"github.com/inexio/go-monitoringplugin"
	"github.com/inexio/snmpsim-restapi-go-client"
	"github.com/jessevdk/go-flags"
	"os"
)

var opts struct {
	URL       string `short:"U" long:"url" description:"The base URL of the SNMPSIM server" required:"true"`
	Username  string `short:"u" long:"username" description:"The username for the server if set" required:"false"`
	Password  string `short:"p" long:"password" description:"The username for the server if set" required:"false"`
	FullCheck []bool `short:"F" long:"full" description:"Run a full check of the API" required:"false"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(3) //parseArgs() prints errors to stdout
	}
	response := monitoringplugin.NewResponse("checked")
	defer response.OutputAndExit()

	metricsClient, err := snmpsimclient.NewMetricsClient(opts.URL)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't connect to server")
		return
	}

	if opts.Username != "" && opts.Password != "" {
		err = metricsClient.SetUsernameAndPassword(opts.Username, opts.Password)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't login client")
			return
		}
	}

	messageFilters, err := metricsClient.GetMessageFilters()
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get message filters")
		return
	}

	_, err = metricsClient.GetMessages(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get messages")
		return
	}

	packetFilters, err := metricsClient.GetPacketFilters()
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get packet filters")
		return
	}

	_, err = metricsClient.GetPackets(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get packets")
		return
	}

	if messageFilters != nil {
		_, err := metricsClient.GetPossibleValuesForMessageFilter(messageFilters[0])
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get possible values for message filter")
			return
		}
	}

	if packetFilters != nil {
		_, err := metricsClient.GetPossibleValuesForPacketFilter(packetFilters[0])
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get possible values for packet filter")
			return
		}
	}

	processes, err := metricsClient.GetProcesses(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get processes")
		return
	}

	if processes != nil {
		_, err := metricsClient.GetProcess(processes[0].ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get process")
			return
		}

		processConsolePages, err := metricsClient.GetProcessConsolePages(processes[0].ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get process console pages")
			return
		}

		if processConsolePages != nil {
			_, err = metricsClient.GetProcessConsolePage(processes[0].ID, processConsolePages[0].ID)
			if err != nil {
				response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get process console page")
				return
			}
		}

		processEndpoints, err := metricsClient.GetProcessEndpoints(processes[0].ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get process endpoints")
			return
		}

		if processEndpoints != nil {
			_, err = metricsClient.GetProcessEndpoint(processes[0].ID, processEndpoints[0].ID)
			if err != nil {
				response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get process endpoint")
				return
			}
		}
	}

	if opts.FullCheck[0] == true {
		err = metricsClient.SetUsernameAndPassword(opts.Username, opts.Password)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't set username and password")
			return
		}
	}
}
