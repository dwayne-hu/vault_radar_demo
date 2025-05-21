package main

import "fmt"

func main() {
	// Klartext-Secrets
	githubToken := "ghp_1234567890abcdefghijklmnopqrstuvwx"
	slackToken := "xoxb-123456789012-1234567890123-abcdefghijklmnopqrstuvwxyz"

	// Base64-codierte Secrets 
	githubTokenBase64 := "Z2hwXzEyMzQ1Njc4OTBhYmNkZWZnaGlqa2xtbm9wcHFyc3R1dnd4"
	azureClientSecretBase64 := "YWJjMTIzLWF6dXJlLWNsaWVudC1zZWNyZXQta2V5"

	// Ausgabe (nur zur Vermeidung von unused-Warnungen)
	fmt.Println("Vault Radar Test Secrets – Klartext:")
	fmt.Println(awsSecret)
	fmt.Println(githubToken)
	fmt.Println(slackToken)
	fmt.Println(stripeSecret)
	fmt.Println(sendgridKey)
	fmt.Println(azureClientSecret)
	fmt.Println(postmanKey)
	fmt.Println(googleKey)

	fmt.Println("\nVault Radar Test Secrets – Base64-codiert:")
	fmt.Println(awsSecretBase64)
	fmt.Println(githubTokenBase64)
	fmt.Println(slackTokenBase64)
	fmt.Println(stripeSecretBase64)
	fmt.Println(sendgridKeyBase64)
	fmt.Println(azureClientSecretBase64)
	fmt.Println(postmanKeyBase64)
	fmt.Println(googleKeyBase64)
}

