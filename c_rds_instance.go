/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type RDSInstance struct {
}

func (n *RDSInstance) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "rds_instances.create":
		lines = append(lines, Message{Body: "Creating rds instances:", Level: "INFO"})
	case "rds_instances.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances created", Level: "INFO"})
	case "rds_instances.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances creation failed", Level: "INFO"})
	case "rds_instances.update":
		lines = append(lines, Message{Body: "Updating rds instances:", Level: "INFO"})
	case "rds_instances.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances modified", Level: "INFO"})
	case "rds_instances.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances modification failed", Level: "INFO"})
	case "rds_instances.delete":
		return append(lines, Message{Body: "Deleting rds instances:", Level: "INFO"})
	case "rds_instances.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances deleted", Level: "INFO"})
	case "rds_instances.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *RDSInstance) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name, _ := r["name"].(string)
		engine, _ := r["engine"].(string)
		cluster, _ := r["cluster"].(string)
		endpoint, _ := r["endpoint"].(string)
		status, _ := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   Engine    : " + engine, Level: ""})
		lines = append(lines, Message{Body: "   Cluster   : " + cluster, Level: ""})
		lines = append(lines, Message{Body: "   Endpoint  : " + endpoint, Level: ""})
		if status == "errored" {
			err, _ := r["error"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}