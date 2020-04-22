# snmpsim-check
[![Go Report Card](https://goreportcard.com/badge/github.com/inexio/snmpsim-check)](https://goreportcard.com/report/github.com/inexio/snmpsim-check)
[![GitHub license](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/inexio/check_eve_ng/blob/master/LICENSE)

## Description
monitoring check plugin for the [snmpsim](https://github.com/etingof/snmpsim) [REST API](https://github.com/etingof/snmpsim-control-plane) (written in golang). The plugin complies with the [Monitoring Plugins Development Guidelines](https://www.monitoring-plugins.org/doc/guidelines.html) and should therefore be compatible with [nagios](https://www.nagios.org/), [icinga2](https://icinga.com/), [zabbix](https://www.zabbix.com/), [checkmk](https://checkmk.com/), etc.

## Usage of the check-management-client

    Usage:
      main [OPTIONS]

    Application Options:
      -U, --url=      The base URL of the SNMPSIM server
      -u, --username= The username for the server if set
      -p, --password= The username for the server if set
      -P, --path=     The data path to a agent file on the server

    Help Options:
      -h, --help      Show this help message

    exit status 3
    
## Usage of the check-metrics-client

    Usage:
      main [OPTIONS]

    Application Options:
      -U, --url=      The base URL of the SNMPSIM server
      -u, --username= The username for the server if set
      -p, --password= The username for the server if set

    Help Options:
      -h, --help      Show this help message

    exit status 3
    
## Installation

To install, use `go get` or `git clone`:

    go get https://github.com/inexio/snmpsim-check
    
or
    
    git clone https://github.com/inexio/snmpsim-check.git

You can also directly download the binaries under the [releases](https://github.com/inexio/snmpsim-check/releases) tab.
