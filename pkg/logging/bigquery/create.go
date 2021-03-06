package bigquery

import (
	"io"

	"github.com/fastly/cli/pkg/common"
	"github.com/fastly/cli/pkg/compute/manifest"
	"github.com/fastly/cli/pkg/config"
	"github.com/fastly/cli/pkg/errors"
	"github.com/fastly/cli/pkg/text"
	"github.com/fastly/go-fastly/fastly"
)

// CreateCommand calls the Fastly API to create BigQuery logging endpoints.
type CreateCommand struct {
	common.Base
	manifest manifest.Data
	Input    fastly.CreateBigQueryInput
}

// NewCreateCommand returns a usable command registered under the parent.
func NewCreateCommand(parent common.Registerer, globals *config.Data) *CreateCommand {
	var c CreateCommand
	c.Globals = globals
	c.manifest.File.Read(manifest.Filename)
	c.CmdClause = parent.Command("create", "Create a BigQuery logging endpoint on a Fastly service version").Alias("add")
	c.CmdClause.Flag("name", "The name of the BigQuery logging object. Used as a primary key for API access").Short('n').Required().StringVar(&c.Input.Name)
	c.CmdClause.Flag("service-id", "Service ID").Short('s').StringVar(&c.manifest.Flag.ServiceID)
	c.CmdClause.Flag("version", "Number of service version").Required().IntVar(&c.Input.Version)
	c.CmdClause.Flag("project-id", "Your Google Cloud Platform project ID").Required().StringVar(&c.Input.ProjectID)
	c.CmdClause.Flag("dataset", "Your BigQuery dataset").Required().StringVar(&c.Input.Dataset)
	c.CmdClause.Flag("table", "Your BigQuery table").Required().StringVar(&c.Input.Table)
	c.CmdClause.Flag("user", "Your Google Cloud Platform service account email address. The client_email field in your service account authentication JSON.").Required().StringVar(&c.Input.User)
	c.CmdClause.Flag("secret-key", "Your Google Cloud Platform account secret key. The private_key field in your service account authentication JSON.").Required().StringVar(&c.Input.SecretKey)
	c.CmdClause.Flag("template-suffix", "BigQuery table name suffix template").StringVar(&c.Input.Template)
	c.CmdClause.Flag("format", "Apache style log formatting. Must produce JSON that matches the schema of your BigQuery table").StringVar(&c.Input.Format)
	c.CmdClause.Flag("format-version", "The version of the custom logging format used for the configured endpoint. Can be either 2 (the default, version 2 log format) or 1 (the version 1 log format). The logging call gets placed by default in vcl_log if format_version is set to 2 and in vcl_deliver if format_version is set to 1").UintVar(&c.Input.FormatVersion)
	// TODO(phamann): It seems `placement` and `response_condition` aren't
	// exposed via the go-fastly input struct. They should be added here when
	// they do.
	return &c
}

// Exec invokes the application logic for the command.
func (c *CreateCommand) Exec(in io.Reader, out io.Writer) error {
	serviceID, source := c.manifest.ServiceID()
	if source == manifest.SourceUndefined {
		return errors.ErrNoServiceID
	}
	c.Input.Service = serviceID

	d, err := c.Globals.Client.CreateBigQuery(&c.Input)
	if err != nil {
		return err
	}

	text.Success(out, "Created BigQuery logging endpoint %s (service %s version %d)", d.Name, d.ServiceID, d.Version)
	return nil
}
