/*
 * talkkonnect headless mumble client/gateway with lcd screen and channel control
 * Copyright (C) 2018-2019, Suvir Kumar <suvir@talkkonnect.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Software distributed under the License is distributed on an "AS IS" basis,
 * WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
 * for the specific language governing rights and limitations under the
 * License.
 *
 * talkkonnect is the based on talkiepi and barnard by Daniel Chote and Tim Cooper
 *
 * The Initial Developer of the Original Code is
 * Suvir Kumar <suvir@talkkonnect.com>
 * Portions created by the Initial Developer are Copyright (C) Suvir Kumar. All Rights Reserved.
 *
 * Contributor(s):
 *
 * Suvir Kumar <suvir@talkkonnect.com>
 *
 * My Blog is at www.talkkonnect.com
 * The source code is hosted at github.com/talkkonnect
 *
 *
 */

package talkkonnect

import (
	"fmt"
	"log"
	"net/http"
)

func (b *Talkkonnect) httpAPI(w http.ResponseWriter, r *http.Request) {
	commands, ok := r.URL.Query()["command"]
	if !ok || len(commands[0]) < 1 {
		log.Println("error: URL Param 'command' is missing example http api commands should be of the format http://a.b.c.d/?command=StartTransmitting")
		fmt.Fprintf(w, "error: API should be of the format http://a.b.c.d:"+APIListenPort+"/?command=StartTransmitting\n")
		return
	}

	command := commands[0]
	log.Println("debug: http command " + string(command))
	b.BackLightTimer()

	switch string(command) {
	case "DisplayMenu":
		if APIDisplayMenu {
			b.cmdDisplayMenu()
			fmt.Fprintf(w, "API Display Menu Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Display Menu Request Denied\n")
		}
	case "ChannelUp":
		if APIChannelUp {
			b.cmdChannelUp()
			fmt.Fprintf(w, "API Channel Up Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Channel Up Request Denied\n")
		}
	case "ChannelDown":
		if APIChannelDown {
			b.cmdChannelDown()
			fmt.Fprintf(w, "API Channel Down Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Channel Down Request Denied\n")
		}
	case "Mute-Toggle":
		if APIMute {
			b.cmdMuteUnmute("toggle")
			fmt.Fprintf(w, "API Mute/UnMute Speaker Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Mute/Unmute Speaker Request Denied\n")
		}
	case "Mute":
		if APIMute {
			b.cmdMuteUnmute("mute")
			fmt.Fprintf(w, "API Mute/UnMute Speaker Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Mute/Unmute Speaker Request Denied\n")
		}
	case "Unmute":
		if APIMute {
			b.cmdMuteUnmute("unmute")
			fmt.Fprintf(w, "API Mute/UnMute Speaker Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Mute/Unmute Speaker Request Denied\n")
		}
	case "CurrentVolume":
		if APICurrentVolumeLevel {
			b.cmdCurrentVolume()
			fmt.Fprintf(w, "API Current Volume Level Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Current Volume Level Request Denied\n")
		}
	case "VolumeUp":
		if APIDigitalVolumeUp {
			b.cmdVolumeUp()
			fmt.Fprintf(w, "API Digital Volume Up Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Digital Volume Up Request Denied\n")
		}
	case "VolumeDown":
		if APIDigitalVolumeDown {
			b.cmdVolumeDown()
			fmt.Fprintf(w, "API Digital Volume Down Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Digital Volume Down Request Denied\n")
		}
	case "ListChannels":
		if APIListServerChannels {
			b.cmdListServerChannels()
			fmt.Fprintf(w, "API List Server Channels Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API List Server Channels Request Denied\n")
		}
	case "StartTransmitting":
		if APIStartTransmitting {
			b.cmdStartTransmitting()
			fmt.Fprintf(w, "API Start Transmitting Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Start Transmitting Request Denied\n")
		}
	case "StopTransmitting":
		if APIStopTransmitting {
			b.cmdStopTransmitting()
			fmt.Fprintf(w, "API Stop Transmitting Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Stop Transmitting Request Denied\n")
		}
	case "ListOnlineUsers":
		if APIListOnlineUsers {
			b.cmdListOnlineUsers()
			fmt.Fprintf(w, "API List Online Users Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API List Online Users Request Denied\n")
		}
	case "Stream-Toggle":
		if APIPlayStream {
			b.cmdPlayback()
			fmt.Fprintf(w, "API Play/Stop Stream Request Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Play/Stop Stream Request Denied\n")
		}
	case "GPSPosition":
		if APIRequestGpsPosition {
			b.cmdGPSPosition()
			fmt.Fprintf(w, "API Request GPS Position Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Request GPS Position Denied\n")
		}

	case "SendEmail":
		if APIEmailEnabled {
			b.cmdSendEmail()
			fmt.Fprintf(w, "API Send Email Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Send Email Config Denied\n")
		}
	case "ConnPreviousServer":
		if APINextServer {
			b.cmdConnPreviousServer()
			fmt.Fprintf(w, "API Previous Server Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Previous Server Denied\n")
		}
	case "ConnNextServer":
		if APINextServer {
			b.cmdConnNextServer()
			fmt.Fprintf(w, "API Next Server Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Next Server Denied\n")
		}

	case "ClearScreen":
		if APIClearScreen {
			b.cmdClearScreen()
			fmt.Fprintf(w, "API Clear Screen Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Clear Screen Denied\n")
		}
	case "PingServers":
		if APIEmailEnabled {
			b.cmdPingServers()
			fmt.Fprintf(w, "API Ping Servers Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Ping Servers Denied\n")
		}
	case "PanicSimulation":
		if APIPanicSimulation {
			b.cmdPanicSimulation()
			fmt.Fprintf(w, "API Request Panic Simulation Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Request Panic Simulation Denied\n")
		}
	case "RepeatTxLoop":
		if APIRepeatTxLoopTest {
			b.cmdRepeatTxLoop()
			fmt.Fprintf(w, "API Request Repeat Tx Loop Test Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Request Repeat Tx Loop Test Denied\n")
		}
	case "ScanChannels":
		if APIScanChannels {
			b.cmdScanChannels()
			fmt.Fprintf(w, "API Request Scan Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Request Scan Denied\n")
		}
	case "Thanks":
		if true {
			b.cmdThanks()
			fmt.Fprintf(w, "API Request Show Acknowledgements Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Request Show Acknowledgements Denied\n")
		}
	case "ShowUptime":
		if APIDisplayVersion {
			b.cmdShowUptime()
			fmt.Fprintf(w, "API Request Current Version Successfully\n")
		} else {
			fmt.Fprintf(w, "API Request Current Version Denied\n")
		}
	case "DumpXMLConfig":
		if APIPrintXmlConfig {
			b.cmdDumpXMLConfig()
			fmt.Fprintf(w, "API Print XML Config Processed Successfully\n")
		} else {
			fmt.Fprintf(w, "API Print XML Congfig Denied\n")
		}
	case "PlayRepeaterTone":
		b.cmdPlayRepeaterTone()
		fmt.Fprintf(w, "API Play Repeater Tone Processed Successfully\n")
	default:
		fmt.Fprintf(w, "API Command Not Defined\n")
	}
}
