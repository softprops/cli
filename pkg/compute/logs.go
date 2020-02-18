package compute

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/fastly/cli/pkg/common"
	"github.com/fastly/cli/pkg/compute/manifest"
	"github.com/fastly/cli/pkg/config"
	"github.com/fastly/cli/pkg/text"

	"github.com/r3labs/sse"
)

const endpointBaseURL = "https://log-bin.glitch.me"

// LogEvent models an individual log event from the log stream endpoint.
type LogEvent struct {
	Timestamp int64  `json:"time"`
	Raw       string `json:"raw"`
}

// LogsCommand streams logs from a given Compute@Edge service.
type LogsCommand struct {
	common.Base
	manifest manifest.Data
}

// NewLogsCommand returns a usable command registered under the parent.
func NewLogsCommand(parent common.Registerer, globals *config.Data) *LogsCommand {
	var c LogsCommand
	c.Globals = globals
	c.manifest.File.Read(manifest.Filename)
	c.CmdClause = parent.Command("logs", "Stream logs emitted from a Fastly Compute@Edge service")
	c.CmdClause.Flag("service-id", "Service ID").Short('s').StringVar(&c.manifest.Flag.ServiceID)
	return &c
}

// Exec implements the command interface.
func (c *LogsCommand) Exec(in io.Reader, out io.Writer) error {
	serviceID, source := c.manifest.ServiceID()
	if source == manifest.SourceUndefined {
		return fmt.Errorf("error reading service: no service ID found. Please provide one via the --service-id flag or within your package manifest")
	}

	endpoint := fmt.Sprintf("%s/%s", endpointBaseURL, serviceID)
	client := sse.NewClient(endpoint)

	client.Subscribe("log", func(msg *sse.Event) {
		var log LogEvent
		if string(msg.Event) == "log" {
			err := json.Unmarshal(msg.Data, &log)
			if err == nil {
				time := time.Unix(log.Timestamp/1000, 0)
				fmt.Fprintf(out, "%s: %s\n", text.Bold(time), log.Raw)
			}
		}
	})

	return nil
}
