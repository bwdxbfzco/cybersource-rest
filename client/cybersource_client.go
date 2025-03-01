// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/bwdxbfzco/cybersource-rest/client/asymmetric_key_management"
	"github.com/bwdxbfzco/cybersource-rest/client/batches"
	"github.com/bwdxbfzco/cybersource-rest/client/billing_agreements"
	"github.com/bwdxbfzco/cybersource-rest/client/capture"
	"github.com/bwdxbfzco/cybersource-rest/client/chargeback_details"
	"github.com/bwdxbfzco/cybersource-rest/client/chargeback_summaries"
	"github.com/bwdxbfzco/cybersource-rest/client/conversion_details"
	"github.com/bwdxbfzco/cybersource-rest/client/create_new_webhooks"
	"github.com/bwdxbfzco/cybersource-rest/client/credit"
	"github.com/bwdxbfzco/cybersource-rest/client/customer"
	"github.com/bwdxbfzco/cybersource-rest/client/customer_payment_instrument"
	"github.com/bwdxbfzco/cybersource-rest/client/customer_shipping_address"
	"github.com/bwdxbfzco/cybersource-rest/client/decision_manager"
	"github.com/bwdxbfzco/cybersource-rest/client/download_d_t_d"
	"github.com/bwdxbfzco/cybersource-rest/client/download_x_s_d"
	"github.com/bwdxbfzco/cybersource-rest/client/e_m_v_tag_details"
	"github.com/bwdxbfzco/cybersource-rest/client/instrument_identifier"
	"github.com/bwdxbfzco/cybersource-rest/client/interchange_clearing_level_details"
	"github.com/bwdxbfzco/cybersource-rest/client/invoice_settings"
	"github.com/bwdxbfzco/cybersource-rest/client/invoices"
	"github.com/bwdxbfzco/cybersource-rest/client/key_management"
	"github.com/bwdxbfzco/cybersource-rest/client/key_management_password"
	"github.com/bwdxbfzco/cybersource-rest/client/key_management_pgp"
	"github.com/bwdxbfzco/cybersource-rest/client/key_management_scmp"
	"github.com/bwdxbfzco/cybersource-rest/client/manage_webhooks"
	"github.com/bwdxbfzco/cybersource-rest/client/microform_integration"
	"github.com/bwdxbfzco/cybersource-rest/client/net_fundings"
	"github.com/bwdxbfzco/cybersource-rest/client/notification_of_changes"
	"github.com/bwdxbfzco/cybersource-rest/client/payer_authentication"
	"github.com/bwdxbfzco/cybersource-rest/client/payment_batch_summaries"
	"github.com/bwdxbfzco/cybersource-rest/client/payment_instrument"
	"github.com/bwdxbfzco/cybersource-rest/client/payments"
	"github.com/bwdxbfzco/cybersource-rest/client/payouts"
	"github.com/bwdxbfzco/cybersource-rest/client/plans"
	"github.com/bwdxbfzco/cybersource-rest/client/purchase_and_refund_details"
	"github.com/bwdxbfzco/cybersource-rest/client/push_funds"
	"github.com/bwdxbfzco/cybersource-rest/client/refund"
	"github.com/bwdxbfzco/cybersource-rest/client/replay_webhooks"
	"github.com/bwdxbfzco/cybersource-rest/client/report_definitions"
	"github.com/bwdxbfzco/cybersource-rest/client/report_downloads"
	"github.com/bwdxbfzco/cybersource-rest/client/report_subscriptions"
	"github.com/bwdxbfzco/cybersource-rest/client/reports"
	"github.com/bwdxbfzco/cybersource-rest/client/retrieval_details"
	"github.com/bwdxbfzco/cybersource-rest/client/retrieval_summaries"
	"github.com/bwdxbfzco/cybersource-rest/client/reversal"
	"github.com/bwdxbfzco/cybersource-rest/client/search_transactions"
	"github.com/bwdxbfzco/cybersource-rest/client/secure_file_share"
	"github.com/bwdxbfzco/cybersource-rest/client/subscriptions"
	"github.com/bwdxbfzco/cybersource-rest/client/symmetric_key_management"
	"github.com/bwdxbfzco/cybersource-rest/client/taxes"
	"github.com/bwdxbfzco/cybersource-rest/client/token"
	"github.com/bwdxbfzco/cybersource-rest/client/transaction_batches"
	"github.com/bwdxbfzco/cybersource-rest/client/transaction_details"
	"github.com/bwdxbfzco/cybersource-rest/client/transient_token_data"
	"github.com/bwdxbfzco/cybersource-rest/client/unified_checkout_capture_context"
	"github.com/bwdxbfzco/cybersource-rest/client/user_management"
	"github.com/bwdxbfzco/cybersource-rest/client/user_management_search"
	"github.com/bwdxbfzco/cybersource-rest/client/verification"
	"github.com/bwdxbfzco/cybersource-rest/client/void"
)

// Default cybersource HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "apitest.cybersource.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new cybersource HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cybersource {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cybersource HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cybersource {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cybersource client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cybersource {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cybersource)
	cli.Transport = transport
	cli.AsymmetricKeyManagement = asymmetric_key_management.New(transport, formats)
	cli.Batches = batches.New(transport, formats)
	cli.BillingAgreements = billing_agreements.New(transport, formats)
	cli.Capture = capture.New(transport, formats)
	cli.ChargebackDetails = chargeback_details.New(transport, formats)
	cli.ChargebackSummaries = chargeback_summaries.New(transport, formats)
	cli.ConversionDetails = conversion_details.New(transport, formats)
	cli.CreateNewWebhooks = create_new_webhooks.New(transport, formats)
	cli.Credit = credit.New(transport, formats)
	cli.Customer = customer.New(transport, formats)
	cli.CustomerPaymentInstrument = customer_payment_instrument.New(transport, formats)
	cli.CustomerShippingAddress = customer_shipping_address.New(transport, formats)
	cli.DecisionManager = decision_manager.New(transport, formats)
	cli.Downloaddtd = download_d_t_d.New(transport, formats)
	cli.Downloadxsd = download_x_s_d.New(transport, formats)
	cli.EmvTagDetails = e_m_v_tag_details.New(transport, formats)
	cli.InstrumentIdentifier = instrument_identifier.New(transport, formats)
	cli.InterchangeClearingLevelDetails = interchange_clearing_level_details.New(transport, formats)
	cli.InvoiceSettings = invoice_settings.New(transport, formats)
	cli.Invoices = invoices.New(transport, formats)
	cli.KeyManagement = key_management.New(transport, formats)
	cli.KeyManagementPassword = key_management_password.New(transport, formats)
	cli.KeyManagementPgp = key_management_pgp.New(transport, formats)
	cli.KeyManagementScmp = key_management_scmp.New(transport, formats)
	cli.ManageWebhooks = manage_webhooks.New(transport, formats)
	cli.MicroformIntegration = microform_integration.New(transport, formats)
	cli.NetFundings = net_fundings.New(transport, formats)
	cli.NotificationOfChanges = notification_of_changes.New(transport, formats)
	cli.PayerAuthentication = payer_authentication.New(transport, formats)
	cli.PaymentBatchSummaries = payment_batch_summaries.New(transport, formats)
	cli.PaymentInstrument = payment_instrument.New(transport, formats)
	cli.Payments = payments.New(transport, formats)
	cli.Payouts = payouts.New(transport, formats)
	cli.Plans = plans.New(transport, formats)
	cli.PurchaseAndRefundDetails = purchase_and_refund_details.New(transport, formats)
	cli.PushFunds = push_funds.New(transport, formats)
	cli.Refund = refund.New(transport, formats)
	cli.ReplayWebhooks = replay_webhooks.New(transport, formats)
	cli.ReportDefinitions = report_definitions.New(transport, formats)
	cli.ReportDownloads = report_downloads.New(transport, formats)
	cli.ReportSubscriptions = report_subscriptions.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.RetrievalDetails = retrieval_details.New(transport, formats)
	cli.RetrievalSummaries = retrieval_summaries.New(transport, formats)
	cli.Reversal = reversal.New(transport, formats)
	cli.SearchTransactions = search_transactions.New(transport, formats)
	cli.SecureFileShare = secure_file_share.New(transport, formats)
	cli.Subscriptions = subscriptions.New(transport, formats)
	cli.SymmetricKeyManagement = symmetric_key_management.New(transport, formats)
	cli.Taxes = taxes.New(transport, formats)
	cli.Token = token.New(transport, formats)
	cli.TransactionBatches = transaction_batches.New(transport, formats)
	cli.TransactionDetails = transaction_details.New(transport, formats)
	cli.TransientTokenData = transient_token_data.New(transport, formats)
	cli.UnifiedCheckoutCaptureContext = unified_checkout_capture_context.New(transport, formats)
	cli.UserManagement = user_management.New(transport, formats)
	cli.UserManagementSearch = user_management_search.New(transport, formats)
	cli.Verification = verification.New(transport, formats)
	cli.Void = void.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Cybersource is a client for cybersource
type Cybersource struct {
	AsymmetricKeyManagement asymmetric_key_management.ClientService

	Batches batches.ClientService

	BillingAgreements billing_agreements.ClientService

	Capture capture.ClientService

	ChargebackDetails chargeback_details.ClientService

	ChargebackSummaries chargeback_summaries.ClientService

	ConversionDetails conversion_details.ClientService

	CreateNewWebhooks create_new_webhooks.ClientService

	Credit credit.ClientService

	Customer customer.ClientService

	CustomerPaymentInstrument customer_payment_instrument.ClientService

	CustomerShippingAddress customer_shipping_address.ClientService

	DecisionManager decision_manager.ClientService

	Downloaddtd download_d_t_d.ClientService

	Downloadxsd download_x_s_d.ClientService

	EmvTagDetails e_m_v_tag_details.ClientService

	InstrumentIdentifier instrument_identifier.ClientService

	InterchangeClearingLevelDetails interchange_clearing_level_details.ClientService

	InvoiceSettings invoice_settings.ClientService

	Invoices invoices.ClientService

	KeyManagement key_management.ClientService

	KeyManagementPassword key_management_password.ClientService

	KeyManagementPgp key_management_pgp.ClientService

	KeyManagementScmp key_management_scmp.ClientService

	ManageWebhooks manage_webhooks.ClientService

	MicroformIntegration microform_integration.ClientService

	NetFundings net_fundings.ClientService

	NotificationOfChanges notification_of_changes.ClientService

	PayerAuthentication payer_authentication.ClientService

	PaymentBatchSummaries payment_batch_summaries.ClientService

	PaymentInstrument payment_instrument.ClientService

	Payments payments.ClientService

	Payouts payouts.ClientService

	Plans plans.ClientService

	PurchaseAndRefundDetails purchase_and_refund_details.ClientService

	PushFunds push_funds.ClientService

	Refund refund.ClientService

	ReplayWebhooks replay_webhooks.ClientService

	ReportDefinitions report_definitions.ClientService

	ReportDownloads report_downloads.ClientService

	ReportSubscriptions report_subscriptions.ClientService

	Reports reports.ClientService

	RetrievalDetails retrieval_details.ClientService

	RetrievalSummaries retrieval_summaries.ClientService

	Reversal reversal.ClientService

	SearchTransactions search_transactions.ClientService

	SecureFileShare secure_file_share.ClientService

	Subscriptions subscriptions.ClientService

	SymmetricKeyManagement symmetric_key_management.ClientService

	Taxes taxes.ClientService

	Token token.ClientService

	TransactionBatches transaction_batches.ClientService

	TransactionDetails transaction_details.ClientService

	TransientTokenData transient_token_data.ClientService

	UnifiedCheckoutCaptureContext unified_checkout_capture_context.ClientService

	UserManagement user_management.ClientService

	UserManagementSearch user_management_search.ClientService

	Verification verification.ClientService

	Void void.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cybersource) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AsymmetricKeyManagement.SetTransport(transport)
	c.Batches.SetTransport(transport)
	c.BillingAgreements.SetTransport(transport)
	c.Capture.SetTransport(transport)
	c.ChargebackDetails.SetTransport(transport)
	c.ChargebackSummaries.SetTransport(transport)
	c.ConversionDetails.SetTransport(transport)
	c.CreateNewWebhooks.SetTransport(transport)
	c.Credit.SetTransport(transport)
	c.Customer.SetTransport(transport)
	c.CustomerPaymentInstrument.SetTransport(transport)
	c.CustomerShippingAddress.SetTransport(transport)
	c.DecisionManager.SetTransport(transport)
	c.Downloaddtd.SetTransport(transport)
	c.Downloadxsd.SetTransport(transport)
	c.EmvTagDetails.SetTransport(transport)
	c.InstrumentIdentifier.SetTransport(transport)
	c.InterchangeClearingLevelDetails.SetTransport(transport)
	c.InvoiceSettings.SetTransport(transport)
	c.Invoices.SetTransport(transport)
	c.KeyManagement.SetTransport(transport)
	c.KeyManagementPassword.SetTransport(transport)
	c.KeyManagementPgp.SetTransport(transport)
	c.KeyManagementScmp.SetTransport(transport)
	c.ManageWebhooks.SetTransport(transport)
	c.MicroformIntegration.SetTransport(transport)
	c.NetFundings.SetTransport(transport)
	c.NotificationOfChanges.SetTransport(transport)
	c.PayerAuthentication.SetTransport(transport)
	c.PaymentBatchSummaries.SetTransport(transport)
	c.PaymentInstrument.SetTransport(transport)
	c.Payments.SetTransport(transport)
	c.Payouts.SetTransport(transport)
	c.Plans.SetTransport(transport)
	c.PurchaseAndRefundDetails.SetTransport(transport)
	c.PushFunds.SetTransport(transport)
	c.Refund.SetTransport(transport)
	c.ReplayWebhooks.SetTransport(transport)
	c.ReportDefinitions.SetTransport(transport)
	c.ReportDownloads.SetTransport(transport)
	c.ReportSubscriptions.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.RetrievalDetails.SetTransport(transport)
	c.RetrievalSummaries.SetTransport(transport)
	c.Reversal.SetTransport(transport)
	c.SearchTransactions.SetTransport(transport)
	c.SecureFileShare.SetTransport(transport)
	c.Subscriptions.SetTransport(transport)
	c.SymmetricKeyManagement.SetTransport(transport)
	c.Taxes.SetTransport(transport)
	c.Token.SetTransport(transport)
	c.TransactionBatches.SetTransport(transport)
	c.TransactionDetails.SetTransport(transport)
	c.TransientTokenData.SetTransport(transport)
	c.UnifiedCheckoutCaptureContext.SetTransport(transport)
	c.UserManagement.SetTransport(transport)
	c.UserManagementSearch.SetTransport(transport)
	c.Verification.SetTransport(transport)
	c.Void.SetTransport(transport)
}
