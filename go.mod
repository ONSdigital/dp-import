module github.com/ONSdigital/dp-import

go 1.23.2

// to avoid the following vulnerabilities:
//     - CVE-2022-29153 # pkg:golang/github.com/hashicorp/consul/api@v1.1.0
replace github.com/spf13/cobra => github.com/spf13/cobra v1.4.0


// [CVE-2023-39325] CWE-770: Allocation of Resources Without Limits or Throttling
replace golang.org/x/net => golang.org/x/net v0.23.0

require (
	github.com/ONSdigital/dp-kafka/v2 v2.8.0
	github.com/ONSdigital/dp-kafka/v3 v3.10.0
)

require github.com/go-avro/avro v0.0.0-20171219232920-444163702c11 // indirect
