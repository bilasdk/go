//go:build ignore

// Webhooks examples
//
// Demonstrates how to configure webhooks and manage delivery history.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

const webhookID = "68f11209-451f-4a15-bfcd-d916eb8b09f4"

func main() {
	apiKey := os.Getenv("BILA_API_KEY")
	if apiKey == "" {
		apiKey = "sk_test_your_api_key_here"
	}

	client := bila.NewClient(
		option.WithAPIKey(apiKey),
		option.WithEnvironmentSandbox(),
	)
	ctx := context.Background()

	// Create webhook
	createParams := bila.WebhookNewParams{
		Events: []string{"payment.completed", "withdrawal.completed", "transfer.completed"},
		URL:    "https://example.com/webhooks",
	}
	created, err := client.Webhooks.New(ctx, createParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("create:", created)

	// Update webhook
	updateParams := bila.WebhookUpdateParams{
		Events:   []string{"payment.completed", "collection.completed", "transfer.failed"},
		URL:      bila.String("https://example.com/webhooks/v2"),
		IsActive: bila.Bool(true),
	}
	updated, err := client.Webhooks.Update(ctx, webhookID, updateParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("update:", updated)

	// List webhooks
	webhooks, err := client.Webhooks.List(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", webhooks)

	// Get webhook deliveries
	deliveriesParams := bila.WebhookGetDeliveriesParams{
		StartDate: bila.String("2026-04-01T00:00:00.000Z"),
		EndDate:   bila.String("2026-04-30T23:59:59.999Z"),
		EventType: bila.String("payment.completed"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(20),
		Status:    bila.String("DELIVERED"),
	}
	deliveries, err := client.Webhooks.GetDeliveries(ctx, webhookID, deliveriesParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("getDeliveries:", deliveries)

	// List webhook events
	events, err := client.Webhooks.ListEvents(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("listEvents:", events)

	// Rotate webhook secret
	rotated, err := client.Webhooks.RotateSecret(ctx, webhookID)
	if err != nil {
		panic(err)
	}
	fmt.Println("rotateSecret:", rotated)

	// Deactivate webhook
	deactivated, err := client.Webhooks.Deactivate(ctx, webhookID)
	if err != nil {
		panic(err)
	}
	fmt.Println("deactivate:", deactivated)
}
