package main

type Affiliate struct {
	Name    string
	Website string
	Address string
}

type Product struct {
	Uuid       string
	Affiliates []Affiliate
}
