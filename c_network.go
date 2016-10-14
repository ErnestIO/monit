/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Network struct {
}

func (n *Network) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "networks.create":
		return append(lines, Message{Body: "Creating networks:", Level: "INFO"})
	case "networks.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks successfully created", Level: "INFO"})
	case "networks.delete":
		return append(lines, Message{Body: "Deleting networks:", Level: "INFO"})
	case "networks.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks deleted", Level: "INFO"})
	}
	return lines
}

func (n *Network) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		ip := r["range"].(string)
		name := r["name"].(string)
		status := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   IP     : " + ip, Level: ""})
		id := r["network_aws_id"].(string)
		if id != "" {
			lines = append(lines, Message{Body: "   AWS ID : " + id, Level: ""})
		}
		lines = append(lines, Message{Body: "   Status : " + status, Level: ""})
		if status == "errored" {
			err := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}
	return lines
}