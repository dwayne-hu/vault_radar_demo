package main

import "fmt"

// Customer test data - should be removed before production
var testCustomers = []Customer{
    {
        Name:        "John Smith",
        Email:       "john.smith@email.com",
        Phone:       "+1-555-123-4567",
        SSN:         "123-45-6789",                    // 🚨 PII: Social Security Number
        CreditCard:  "4532-1234-5678-9012",           // 🚨 PII: Credit Card Number
        Address:     "123 Main St, New York, NY 10001",
        DateOfBirth: "1985-03-15",                     // 🚨 PII: Date of Birth
    },
    {
        Name:        "Sarah Johnson", 
        Email:       "sarah.johnson@company.com",
        Phone:       "+1-555-987-6543",
        SSN:         "987-65-4321",                    // 🚨 PII: Social Security Number
        CreditCard:  "5105-1051-0510-5100",           // 🚨 PII: Credit Card Number
        Address:     "456 Oak Ave, Los Angeles, CA 90210",
        DateOfBirth: "1990-07-22",                     // 🚨 PII: Date of Birth
    },
}

func printCustomerInfo() {
    for _, customer := range testCustomers {
        fmt.Printf("Customer: %s, SSN: %s, Card: %s\n", 
                   customer.Name, customer.SSN, customer.CreditCard)
    }
}