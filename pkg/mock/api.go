package mock

import (
	"github.com/fastly/go-fastly/fastly"
)

// API is a mock implementation of api.Interface that's used for testing.
// The zero value is useful, but will panic on all methods. Provide function
// implementations for the method(s) your test will call.
type API struct {
	GetTokenSelfFn func() (*fastly.Token, error)

	CreateServiceFn     func(*fastly.CreateServiceInput) (*fastly.Service, error)
	ListServicesFn      func(*fastly.ListServicesInput) ([]*fastly.Service, error)
	GetServiceFn        func(*fastly.GetServiceInput) (*fastly.Service, error)
	GetServiceDetailsFn func(*fastly.GetServiceInput) (*fastly.ServiceDetail, error)
	UpdateServiceFn     func(*fastly.UpdateServiceInput) (*fastly.Service, error)
	DeleteServiceFn     func(*fastly.DeleteServiceInput) error

	CloneVersionFn      func(*fastly.CloneVersionInput) (*fastly.Version, error)
	ListVersionsFn      func(*fastly.ListVersionsInput) ([]*fastly.Version, error)
	UpdateVersionFn     func(*fastly.UpdateVersionInput) (*fastly.Version, error)
	ActivateVersionFn   func(*fastly.ActivateVersionInput) (*fastly.Version, error)
	DeactivateVersionFn func(*fastly.DeactivateVersionInput) (*fastly.Version, error)
	LockVersionFn       func(*fastly.LockVersionInput) (*fastly.Version, error)
	LatestVersionFn     func(*fastly.LatestVersionInput) (*fastly.Version, error)

	CreateDomainFn func(*fastly.CreateDomainInput) (*fastly.Domain, error)
	ListDomainsFn  func(*fastly.ListDomainsInput) ([]*fastly.Domain, error)
	GetDomainFn    func(*fastly.GetDomainInput) (*fastly.Domain, error)
	UpdateDomainFn func(*fastly.UpdateDomainInput) (*fastly.Domain, error)
	DeleteDomainFn func(*fastly.DeleteDomainInput) error

	CreateBackendFn func(*fastly.CreateBackendInput) (*fastly.Backend, error)
	ListBackendsFn  func(*fastly.ListBackendsInput) ([]*fastly.Backend, error)
	GetBackendFn    func(*fastly.GetBackendInput) (*fastly.Backend, error)
	UpdateBackendFn func(*fastly.UpdateBackendInput) (*fastly.Backend, error)
	DeleteBackendFn func(*fastly.DeleteBackendInput) error

	CreateHealthCheckFn func(*fastly.CreateHealthCheckInput) (*fastly.HealthCheck, error)
	ListHealthChecksFn  func(*fastly.ListHealthChecksInput) ([]*fastly.HealthCheck, error)
	GetHealthCheckFn    func(*fastly.GetHealthCheckInput) (*fastly.HealthCheck, error)
	UpdateHealthCheckFn func(*fastly.UpdateHealthCheckInput) (*fastly.HealthCheck, error)
	DeleteHealthCheckFn func(*fastly.DeleteHealthCheckInput) error

	CreateBigQueryFn func(*fastly.CreateBigQueryInput) (*fastly.BigQuery, error)
	ListBigQueriesFn func(*fastly.ListBigQueriesInput) ([]*fastly.BigQuery, error)
	GetBigQueryFn    func(*fastly.GetBigQueryInput) (*fastly.BigQuery, error)
	UpdateBigQueryFn func(*fastly.UpdateBigQueryInput) (*fastly.BigQuery, error)
	DeleteBigQueryFn func(*fastly.DeleteBigQueryInput) error

	CreateS3Fn func(*fastly.CreateS3Input) (*fastly.S3, error)
	ListS3sFn  func(*fastly.ListS3sInput) ([]*fastly.S3, error)
	GetS3Fn    func(*fastly.GetS3Input) (*fastly.S3, error)
	UpdateS3Fn func(*fastly.UpdateS3Input) (*fastly.S3, error)
	DeleteS3Fn func(*fastly.DeleteS3Input) error

	CreateSyslogFn func(*fastly.CreateSyslogInput) (*fastly.Syslog, error)
	ListSyslogsFn  func(*fastly.ListSyslogsInput) ([]*fastly.Syslog, error)
	GetSyslogFn    func(*fastly.GetSyslogInput) (*fastly.Syslog, error)
	UpdateSyslogFn func(*fastly.UpdateSyslogInput) (*fastly.Syslog, error)
	DeleteSyslogFn func(*fastly.DeleteSyslogInput) error

	CreateLogentriesFn func(*fastly.CreateLogentriesInput) (*fastly.Logentries, error)
	ListLogentriesFn   func(*fastly.ListLogentriesInput) ([]*fastly.Logentries, error)
	GetLogentriesFn    func(*fastly.GetLogentriesInput) (*fastly.Logentries, error)
	UpdateLogentriesFn func(*fastly.UpdateLogentriesInput) (*fastly.Logentries, error)
	DeleteLogentriesFn func(*fastly.DeleteLogentriesInput) error

	CreatePapertrailFn func(*fastly.CreatePapertrailInput) (*fastly.Papertrail, error)
	ListPapertrailsFn  func(*fastly.ListPapertrailsInput) ([]*fastly.Papertrail, error)
	GetPapertrailFn    func(*fastly.GetPapertrailInput) (*fastly.Papertrail, error)
	UpdatePapertrailFn func(*fastly.UpdatePapertrailInput) (*fastly.Papertrail, error)
	DeletePapertrailFn func(*fastly.DeletePapertrailInput) error

	CreateSumologicFn func(*fastly.CreateSumologicInput) (*fastly.Sumologic, error)
	ListSumologicsFn  func(*fastly.ListSumologicsInput) ([]*fastly.Sumologic, error)
	GetSumologicFn    func(*fastly.GetSumologicInput) (*fastly.Sumologic, error)
	UpdateSumologicFn func(*fastly.UpdateSumologicInput) (*fastly.Sumologic, error)
	DeleteSumologicFn func(*fastly.DeleteSumologicInput) error

	GetUserFn func(*fastly.GetUserInput) (*fastly.User, error)
}

// GetTokenSelf implements Interface.
func (m API) GetTokenSelf() (*fastly.Token, error) {
	return m.GetTokenSelfFn()
}

// CreateService implements Interface.
func (m API) CreateService(i *fastly.CreateServiceInput) (*fastly.Service, error) {
	return m.CreateServiceFn(i)
}

// ListServices implements Interface.
func (m API) ListServices(i *fastly.ListServicesInput) ([]*fastly.Service, error) {
	return m.ListServicesFn(i)
}

// GetService implements Interface.
func (m API) GetService(i *fastly.GetServiceInput) (*fastly.Service, error) {
	return m.GetServiceFn(i)
}

// GetServiceDetails implements Interface.
func (m API) GetServiceDetails(i *fastly.GetServiceInput) (*fastly.ServiceDetail, error) {
	return m.GetServiceDetailsFn(i)
}

// UpdateService implements Interface.
func (m API) UpdateService(i *fastly.UpdateServiceInput) (*fastly.Service, error) {
	return m.UpdateServiceFn(i)
}

// DeleteService implements Interface.
func (m API) DeleteService(i *fastly.DeleteServiceInput) error {
	return m.DeleteServiceFn(i)
}

// CloneVersion implements Interface.
func (m API) CloneVersion(i *fastly.CloneVersionInput) (*fastly.Version, error) {
	return m.CloneVersionFn(i)
}

// ListVersions implements Interface.
func (m API) ListVersions(i *fastly.ListVersionsInput) ([]*fastly.Version, error) {
	return m.ListVersionsFn(i)
}

// UpdateVersion implements Interface.
func (m API) UpdateVersion(i *fastly.UpdateVersionInput) (*fastly.Version, error) {
	return m.UpdateVersionFn(i)
}

// ActivateVersion implements Interface.
func (m API) ActivateVersion(i *fastly.ActivateVersionInput) (*fastly.Version, error) {
	return m.ActivateVersionFn(i)
}

// DeactivateVersion implements Interface.
func (m API) DeactivateVersion(i *fastly.DeactivateVersionInput) (*fastly.Version, error) {
	return m.DeactivateVersionFn(i)
}

// LockVersion implements Interface.
func (m API) LockVersion(i *fastly.LockVersionInput) (*fastly.Version, error) {
	return m.LockVersionFn(i)
}

// LatestVersion implements Interface.
func (m API) LatestVersion(i *fastly.LatestVersionInput) (*fastly.Version, error) {
	return m.LatestVersionFn(i)
}

// CreateDomain implements Interface.
func (m API) CreateDomain(i *fastly.CreateDomainInput) (*fastly.Domain, error) {
	return m.CreateDomainFn(i)
}

// ListDomains implements Interface.
func (m API) ListDomains(i *fastly.ListDomainsInput) ([]*fastly.Domain, error) {
	return m.ListDomainsFn(i)
}

// GetDomain implements Interface.
func (m API) GetDomain(i *fastly.GetDomainInput) (*fastly.Domain, error) {
	return m.GetDomainFn(i)
}

// UpdateDomain implements Interface.
func (m API) UpdateDomain(i *fastly.UpdateDomainInput) (*fastly.Domain, error) {
	return m.UpdateDomainFn(i)
}

// DeleteDomain implements Interface.
func (m API) DeleteDomain(i *fastly.DeleteDomainInput) error {
	return m.DeleteDomainFn(i)
}

// CreateBackend implements Interface.
func (m API) CreateBackend(i *fastly.CreateBackendInput) (*fastly.Backend, error) {
	return m.CreateBackendFn(i)
}

// ListBackends implements Interface.
func (m API) ListBackends(i *fastly.ListBackendsInput) ([]*fastly.Backend, error) {
	return m.ListBackendsFn(i)
}

// GetBackend implements Interface.
func (m API) GetBackend(i *fastly.GetBackendInput) (*fastly.Backend, error) {
	return m.GetBackendFn(i)
}

// UpdateBackend implements Interface.
func (m API) UpdateBackend(i *fastly.UpdateBackendInput) (*fastly.Backend, error) {
	return m.UpdateBackendFn(i)
}

// DeleteBackend implements Interface.
func (m API) DeleteBackend(i *fastly.DeleteBackendInput) error {
	return m.DeleteBackendFn(i)
}

// CreateHealthCheck implements Interface.
func (m API) CreateHealthCheck(i *fastly.CreateHealthCheckInput) (*fastly.HealthCheck, error) {
	return m.CreateHealthCheckFn(i)
}

// ListHealthChecks implements Interface.
func (m API) ListHealthChecks(i *fastly.ListHealthChecksInput) ([]*fastly.HealthCheck, error) {
	return m.ListHealthChecksFn(i)
}

// GetHealthCheck implements Interface.
func (m API) GetHealthCheck(i *fastly.GetHealthCheckInput) (*fastly.HealthCheck, error) {
	return m.GetHealthCheckFn(i)
}

// UpdateHealthCheck implements Interface.
func (m API) UpdateHealthCheck(i *fastly.UpdateHealthCheckInput) (*fastly.HealthCheck, error) {
	return m.UpdateHealthCheckFn(i)
}

// DeleteHealthCheck implements Interface.
func (m API) DeleteHealthCheck(i *fastly.DeleteHealthCheckInput) error {
	return m.DeleteHealthCheckFn(i)
}

// CreateBigQuery implements Interface.
func (m API) CreateBigQuery(i *fastly.CreateBigQueryInput) (*fastly.BigQuery, error) {
	return m.CreateBigQueryFn(i)
}

// ListBigQueries implements Interface.
func (m API) ListBigQueries(i *fastly.ListBigQueriesInput) ([]*fastly.BigQuery, error) {
	return m.ListBigQueriesFn(i)
}

// GetBigQuery implements Interface.
func (m API) GetBigQuery(i *fastly.GetBigQueryInput) (*fastly.BigQuery, error) {
	return m.GetBigQueryFn(i)
}

// UpdateBigQuery implements Interface.
func (m API) UpdateBigQuery(i *fastly.UpdateBigQueryInput) (*fastly.BigQuery, error) {
	return m.UpdateBigQueryFn(i)
}

// DeleteBigQuery implements Interface.
func (m API) DeleteBigQuery(i *fastly.DeleteBigQueryInput) error {
	return m.DeleteBigQueryFn(i)
}

// CreateS3 implements Interface.
func (m API) CreateS3(i *fastly.CreateS3Input) (*fastly.S3, error) {
	return m.CreateS3Fn(i)
}

// ListS3s implements Interface.
func (m API) ListS3s(i *fastly.ListS3sInput) ([]*fastly.S3, error) {
	return m.ListS3sFn(i)
}

// GetS3 implements Interface.
func (m API) GetS3(i *fastly.GetS3Input) (*fastly.S3, error) {
	return m.GetS3Fn(i)
}

// UpdateS3 implements Interface.
func (m API) UpdateS3(i *fastly.UpdateS3Input) (*fastly.S3, error) {
	return m.UpdateS3Fn(i)
}

// DeleteS3 implements Interface.
func (m API) DeleteS3(i *fastly.DeleteS3Input) error {
	return m.DeleteS3Fn(i)
}

// CreateSyslog implements Interface.
func (m API) CreateSyslog(i *fastly.CreateSyslogInput) (*fastly.Syslog, error) {
	return m.CreateSyslogFn(i)
}

// ListSyslogs implements Interface.
func (m API) ListSyslogs(i *fastly.ListSyslogsInput) ([]*fastly.Syslog, error) {
	return m.ListSyslogsFn(i)
}

// GetSyslog implements Interface.
func (m API) GetSyslog(i *fastly.GetSyslogInput) (*fastly.Syslog, error) {
	return m.GetSyslogFn(i)
}

// UpdateSyslog implements Interface.
func (m API) UpdateSyslog(i *fastly.UpdateSyslogInput) (*fastly.Syslog, error) {
	return m.UpdateSyslogFn(i)
}

// DeleteSyslog implements Interface.
func (m API) DeleteSyslog(i *fastly.DeleteSyslogInput) error {
	return m.DeleteSyslogFn(i)
}

// CreateLogentries implements Interface.
func (m API) CreateLogentries(i *fastly.CreateLogentriesInput) (*fastly.Logentries, error) {
	return m.CreateLogentriesFn(i)
}

// ListLogentries implements Interface.
func (m API) ListLogentries(i *fastly.ListLogentriesInput) ([]*fastly.Logentries, error) {
	return m.ListLogentriesFn(i)
}

// GetLogentries implements Interface.
func (m API) GetLogentries(i *fastly.GetLogentriesInput) (*fastly.Logentries, error) {
	return m.GetLogentriesFn(i)
}

// UpdateLogentries implements Interface.
func (m API) UpdateLogentries(i *fastly.UpdateLogentriesInput) (*fastly.Logentries, error) {
	return m.UpdateLogentriesFn(i)
}

// DeleteLogentries implements Interface.
func (m API) DeleteLogentries(i *fastly.DeleteLogentriesInput) error {
	return m.DeleteLogentriesFn(i)
}

// CreatePapertrail implements Interface.
func (m API) CreatePapertrail(i *fastly.CreatePapertrailInput) (*fastly.Papertrail, error) {
	return m.CreatePapertrailFn(i)
}

// ListPapertrails implements Interface.
func (m API) ListPapertrails(i *fastly.ListPapertrailsInput) ([]*fastly.Papertrail, error) {
	return m.ListPapertrailsFn(i)
}

// GetPapertrail implements Interface.
func (m API) GetPapertrail(i *fastly.GetPapertrailInput) (*fastly.Papertrail, error) {
	return m.GetPapertrailFn(i)
}

// UpdatePapertrail implements Interface.
func (m API) UpdatePapertrail(i *fastly.UpdatePapertrailInput) (*fastly.Papertrail, error) {
	return m.UpdatePapertrailFn(i)
}

// DeletePapertrail implements Interface.
func (m API) DeletePapertrail(i *fastly.DeletePapertrailInput) error {
	return m.DeletePapertrailFn(i)
}

// CreateSumologic implements Interface.
func (m API) CreateSumologic(i *fastly.CreateSumologicInput) (*fastly.Sumologic, error) {
	return m.CreateSumologicFn(i)
}

// ListSumologics implements Interface.
func (m API) ListSumologics(i *fastly.ListSumologicsInput) ([]*fastly.Sumologic, error) {
	return m.ListSumologicsFn(i)
}

// GetSumologic implements Interface.
func (m API) GetSumologic(i *fastly.GetSumologicInput) (*fastly.Sumologic, error) {
	return m.GetSumologicFn(i)
}

// UpdateSumologic implements Interface.
func (m API) UpdateSumologic(i *fastly.UpdateSumologicInput) (*fastly.Sumologic, error) {
	return m.UpdateSumologicFn(i)
}

// DeleteSumologic implements Interface.
func (m API) DeleteSumologic(i *fastly.DeleteSumologicInput) error {
	return m.DeleteSumologicFn(i)
}

// GetUser implements Interface.
func (m API) GetUser(i *fastly.GetUserInput) (*fastly.User, error) {
	return m.GetUserFn(i)
}
