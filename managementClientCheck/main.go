package main

import (
	"github.com/inexio/go-monitoringplugin"
	"github.com/inexio/snmpsim-restapi-go-client"
	"github.com/jessevdk/go-flags"
	"os"
)

var opts struct {
	URL      string `short:"U" long:"url" description:"The base URL of the SNMPSIM server" required:"true"`
	Username string `short:"u" long:"username" description:"The username for the server if set" required:"false"`
	Password string `short:"p" long:"password" description:"The username for the server if set" required:"false"`
	Path     string `short:"P" long:"path" description:"The data path to a agent file on the server" required:"true"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(3) //parseArgs() prints errors to stdout
	}
	response := monitoringplugin.NewResponse("checked")
	defer response.OutputAndExit()

	client, err := snmpsimclient.NewManagementClient(opts.URL)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't connect to server")
		return
	}

	if opts.Username != "" && opts.Password != "" {
		err = client.SetUsernameAndPassword(opts.Username, opts.Password)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't login client")
			return
		}
	}

	//Create and delete requests

	agent, err := client.CreateAgent("testAgent", opts.Path)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create agent")
		return
	}
	defer func() {
		err = client.DeleteAgent(agent.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete agent")
			return
		}
	}()

	endpoint, err := client.CreateEndpoint("testEndpoint", "192.168.100.203:6666", "udpv4")
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create endpoint")
		return
	}
	defer func() {
		err = client.DeleteEndpoint(endpoint.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete endpoint")
			return
		}
	}()

	engine, err := client.CreateEngine("testEngine", "0")
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create engine")
		return
	}
	defer func() {
		err = client.DeleteEngine(engine.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete engine")
			return
		}
	}()

	lab, err := client.CreateLab("testLab")
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create lab")
		return
	}
	defer func() {
		err = client.DeleteLab(lab.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete lab")
			return
		}
	}()

	//TODO CreateSelector/DeleteSelector

	tag, err := client.CreateTag("testTag", "Tag for testing")
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create tag")
		return
	}
	defer func() {
		err = client.DeleteTag(tag.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete tag")
			return
		}
	}()

	user, err := client.CreateUser("testUser", "testUser", "0x50dd4d3ec79a1cf4dfa5fee9f76b0847647fcf74", "sha", "0x50dd4d3ec79a1cf4dfa5fee9f76b0847", "des")
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create user")
		return
	}
	defer func() {
		err = client.DeleteUser(user.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete user")
			return
		}
	}()

	//Create with tag requests

	agentWithTag, err := client.CreateAgentWithTag("testAgentWithTag", opts.Path, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create agent with tag")
		return
	}
	defer func() {
		err = client.DeleteAgent(agentWithTag.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete agent with tag")
			return
		}
	}()

	endpointWithTag, err := client.CreateEndpointWithTag("testEndpointWithTag", "192.168.100.203:6667", "udpv4", tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create endpoint with tag")
		return
	}
	defer func() {
		err = client.DeleteEndpoint(endpointWithTag.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete endpoint with tag")
			return
		}
	}()

	engineWithTag, err := client.CreateEngineWithTag("testEngineWithTag", "testEngineWithTag", tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create engine with tag")
		return
	}
	defer func() {
		err = client.DeleteEngine(engineWithTag.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete engine with tag")
			return
		}
	}()

	labWithTag, err := client.CreateLabWithTag("testLabWithTag", tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create lab with tag")
		return
	}
	defer func() {
		err = client.DeleteLab(labWithTag.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete lab with tag")
			return
		}
	}()

	userWithTag, err := client.CreateUserWithTag("testUserWithTag", "testUserWithTag", "0x50dd4d3ec79a1cf4dfa5fee9f76b0847647fcf74", "sha", "0x50dd4d3ec79a1cf4dfa5fee9f76b0847", "des", tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't create user with tag")
		return
	}
	defer func() {
		err = client.DeleteUser(userWithTag.ID)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't delete user with tag")
			return
		}
	}()

	//Add requests

	err = client.AddAgentToLab(lab.ID, agent.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add agent to lab")
		return
	}

	err = client.AddEndpointToEngine(engine.ID, endpoint.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add endpoint to engine")
		return
	}

	err = client.AddEngineToAgent(agent.ID, engine.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add engine to agent")
		return
	}

	//TODO AddSelectorToAgent

	err = client.AddTagToAgent(agent.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add tag to agent")
		return
	}

	err = client.AddTagToEndpoint(endpoint.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add tag to endpoint")
		return
	}

	err = client.AddTagToEngine(endpoint.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add tag to engine")
		return
	}

	err = client.AddTagToLab(lab.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add tag to lab")
		return
	}

	err = client.AddTagToUser(user.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add tag to user")
		return
	}

	err = client.AddUserToEngine(engine.ID, user.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't add user to engine")
		return
	}

	//Remove requests

	err = client.RemoveAgentFromLab(lab.ID, agent.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove agent from lab")
		return
	}

	err = client.RemoveEndpointFromEngine(engine.ID, endpoint.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove endpoint from engine")
		return
	}

	err = client.RemoveEngineFromAgent(agent.ID, engine.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove engine from agent")
		return
	}

	//TODO AddSelectorToAgent

	err = client.RemoveTagFromAgent(agent.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove tag from agent")
		return
	}

	err = client.RemoveTagFromEndpoint(endpoint.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove tag from endpoint")
		return
	}

	err = client.RemoveTagFromEngine(endpoint.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove tag from engine")
		return
	}

	err = client.RemoveTagFromLab(lab.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove tag from lab")
		return
	}

	err = client.RemoveTagFromUser(user.ID, tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove tag from user")
		return
	}

	err = client.RemoveUserFromEngine(engine.ID, user.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't remove user from engine")
		return
	}

	//Get requests

	_, err = client.GetAgent(agent.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get agent")
		return
	}

	_, err = client.GetAgents(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get agents")
		return
	}

	_, err = client.GetEndpoint(endpoint.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get endpoint")
		return
	}

	_, err = client.GetEndpoints(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get endpoints")
		return
	}

	_, err = client.GetEngine(engine.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get engine")
		return
	}

	_, err = client.GetEngines(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get engines")
		return
	}

	_, err = client.GetLab(lab.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get lab")
		return
	}

	_, err = client.GetLabs(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get labs")
		return
	}

	files, err := client.GetRecordFiles()
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get record files")
		return
	}

	if files != nil {
		_, err = client.GetRecordFile(files[0].Path)
		if err != nil {
			response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get record file")
			return
		}
	}

	//TODO GetSelector

	//TODO GetSelectors

	_, err = client.GetTag(tag.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get tag")
		return
	}

	_, err = client.GetTags(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get tags")
		return
	}

	_, err = client.GetUser(user.ID)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get user")
		return
	}

	_, err = client.GetUsers(nil)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't get users")
		return
	}

	//Set requests

	err = client.SetLabPower(lab.ID, false)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't set lab power")
		return
	}

	err = client.SetUsernameAndPassword(opts.Username, opts.Password)
	if err != nil {
		response.UpdateStatus(monitoringplugin.CRITICAL, "Can't set username and password")
		return
	}
}
