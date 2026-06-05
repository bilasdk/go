# Shared Response Types

- <a href="https://pkg.go.dev/github.com/bilasdk/go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go/shared#PaginationMetaDto">PaginationMetaDto</a>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountDetailsDto">AccountDetailsDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountResponseDto">AccountResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountGetResponse">AccountGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountListResponse">AccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountGetBalanceResponse">AccountGetBalanceResponse</a>

Methods:

- <code title="get /api/v1/bila/accounts/{id}">client.Accounts.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/accounts/{id}/balance">client.Accounts.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountService.GetBalance">GetBalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#AccountGetBalanceResponse">AccountGetBalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TransferRecipients

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#RecipientResponseDto">RecipientResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientGetResponse">TransferRecipientGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientListResponse">TransferRecipientListResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientNewBankAccountResponse">TransferRecipientNewBankAccountResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientNewMobileMoneyResponse">TransferRecipientNewMobileMoneyResponse</a>

Methods:

- <code title="get /api/v1/bila/transfer-recipients/{id}">client.TransferRecipients.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientGetResponse">TransferRecipientGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/transfer-recipients">client.TransferRecipients.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientListParams">TransferRecipientListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientListResponse">TransferRecipientListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/transfer-recipients/bank-account">client.TransferRecipients.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientService.NewBankAccount">NewBankAccount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientNewBankAccountParams">TransferRecipientNewBankAccountParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientNewBankAccountResponse">TransferRecipientNewBankAccountResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/transfer-recipients/mobile-money">client.TransferRecipients.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientService.NewMobileMoney">NewMobileMoney</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientNewMobileMoneyParams">TransferRecipientNewMobileMoneyParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientNewMobileMoneyResponse">TransferRecipientNewMobileMoneyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transfers

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferRecipientDto">TransferRecipientDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferResponseDto">TransferResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferGetResponse">TransferGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferListResponse">TransferListResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferGetStatusByReferenceResponse">TransferGetStatusByReferenceResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferInitiateBankTransferResponse">TransferInitiateBankTransferResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferInitiateMobileMoneyTransferResponse">TransferInitiateMobileMoneyTransferResponse</a>

Methods:

- <code title="get /api/v1/bila/transfers/{id}">client.Transfers.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferGetResponse">TransferGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/transfers">client.Transfers.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferListParams">TransferListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferListResponse">TransferListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/transfers/status/{reference}">client.Transfers.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferService.GetStatusByReference">GetStatusByReference</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reference <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferGetStatusByReferenceResponse">TransferGetStatusByReferenceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/transfers/bank-account">client.Transfers.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferService.InitiateBankTransfer">InitiateBankTransfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferInitiateBankTransferParams">TransferInitiateBankTransferParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferInitiateBankTransferResponse">TransferInitiateBankTransferResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/transfers/mobile-money">client.Transfers.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferService.InitiateMobileMoneyTransfer">InitiateMobileMoneyTransfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferInitiateMobileMoneyTransferParams">TransferInitiateMobileMoneyTransferParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransferInitiateMobileMoneyTransferResponse">TransferInitiateMobileMoneyTransferResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Collections

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#BilaCollectionCustomerDto">BilaCollectionCustomerDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#BilaCollectionResponseDto">BilaCollectionResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionGetResponse">CollectionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionListResponse">CollectionListResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionGetStatusByReferenceResponse">CollectionGetStatusByReferenceResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionInitiateMobileMoneyCollectionResponse">CollectionInitiateMobileMoneyCollectionResponse</a>

Methods:

- <code title="get /api/v1/bila/collections/{id}">client.Collections.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionGetResponse">CollectionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/collections">client.Collections.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionListParams">CollectionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionListResponse">CollectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/collections/status/{reference}">client.Collections.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionService.GetStatusByReference">GetStatusByReference</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reference <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionGetStatusByReferenceResponse">CollectionGetStatusByReferenceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/collections/mobile-money">client.Collections.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionService.InitiateMobileMoneyCollection">InitiateMobileMoneyCollection</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionInitiateMobileMoneyCollectionParams">CollectionInitiateMobileMoneyCollectionParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#CollectionInitiateMobileMoneyCollectionResponse">CollectionInitiateMobileMoneyCollectionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionResponseDto">TransactionResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionGetResponse">TransactionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionListResponse">TransactionListResponse</a>

Methods:

- <code title="get /api/v1/bila/transactions/{id}">client.Transactions.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionGetResponse">TransactionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookConfigResponseDto">WebhookConfigResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookUpdateResponse">WebhookUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookDeactivateResponse">WebhookDeactivateResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookGetDeliveriesResponse">WebhookGetDeliveriesResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookListEventsResponse">WebhookListEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookRotateSecretResponse">WebhookRotateSecretResponse</a>

Methods:

- <code title="post /api/v1/bila/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/v1/bila/webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookUpdateResponse">WebhookUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/bila/webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.Deactivate">Deactivate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookDeactivateResponse">WebhookDeactivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/webhooks/{id}/deliveries">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.GetDeliveries">GetDeliveries</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookGetDeliveriesParams">WebhookGetDeliveriesParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookGetDeliveriesResponse">WebhookGetDeliveriesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/bila/webhooks/events">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookListEventsResponse">WebhookListEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/webhooks/{id}/rotate-secret">client.Webhooks.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookService.RotateSecret">RotateSecret</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#WebhookRotateSecretResponse">WebhookRotateSecretResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Banks

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#BankListResponse">BankListResponse</a>

Methods:

- <code title="get /api/v1/bila/banks">client.Banks.<a href="https://pkg.go.dev/github.com/bilasdk/go#BankService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#BankListParams">BankListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#BankListResponse">BankListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Resolve

Response Types:

- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolvedAccountResponseDto">ResolvedAccountResponseDto</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveBankAccountResponse">ResolveBankAccountResponse</a>
- <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveMobileMoneyResponse">ResolveMobileMoneyResponse</a>

Methods:

- <code title="post /api/v1/bila/resolve/bank-account">client.Resolve.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveService.BankAccount">BankAccount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveBankAccountParams">ResolveBankAccountParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveBankAccountResponse">ResolveBankAccountResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/bila/resolve/mobile-money">client.Resolve.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveService.MobileMoney">MobileMoney</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveMobileMoneyParams">ResolveMobileMoneyParams</a>) (\*<a href="https://pkg.go.dev/github.com/bilasdk/go">bila</a>.<a href="https://pkg.go.dev/github.com/bilasdk/go#ResolveMobileMoneyResponse">ResolveMobileMoneyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
