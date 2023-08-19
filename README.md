##isup - Domain Reachability Verifier

The *isup* application is a simple command-line tool that allows you to verify whether a domain is reachable or not using the isup command. It performs DNS record checks, retrieves the IP address associated with the domain, and then determines if the domain is reachable from your network.

This tool is particularly useful for network administrators, developers, and anyone who needs to quickly check if a specific domain is accessible. Whether you're troubleshooting connectivity issues, testing your own domains, or monitoring network health, isup provides a straightforward solution.

## Installation

Follow the steps below to install the isup application from the source repository:

1. Clone the Repository: Open your terminal and run the following command to clone the isup repository:
```bash
git clone git@github.com:denysovkos/isup.git
```

2. Install Using Go: After cloning the repository, navigate to the project directory using the terminal:
```bash
cd isup
```

3. Then, install the application using the Go programming language:
```bash
go install
```

Note: Make sure you have Go (Golang) installed on your system before proceeding with the installation.

## Usage

Once isup is installed, you can use it to check the reachability of a domain by providing the isup command followed by the domain you want to verify. For example:
```bash
isup www.google.com
```

The application will perform the following steps to determine domain reachability:

* DNS Record Check: The application will query the DNS records for the provided domain to retrieve its associated IP address.
* IP Address Retrieval: The retrieved IP address will be displayed along with additional information.
* Reachability Check: The application will then attempt to establish a connection to the domain's IP address to determine if it is reachable from your network.
* Result: Based on the reachability test, the application will provide you with a clear message indicating whether the domain is reachable or not.

# Example output:
```bash
isup google.com

2023/08/19 11:00:09 Processing:  google.com
google.com. IN A 2a00:1450:4001:80b::200e
google.com. IN A 172.217.18.14
IPv4:  172.217.18.14
Domain [172.217.18.14] http://google.com reposned with status 200 OK REACHEBLE ✅
```

## Contributions

Contributions to the isup application are welcome! If you encounter issues, have ideas for improvements, or want to add new features, feel free to submit a pull request to the GitHub repository. Your contributions will help make the tool even more valuable to the community.

## License

The isup application is distributed under the MIT License. Feel free to use, modify, and distribute it according to the terms of the license.

## Contact

If you have any questions, suggestions, or feedback, you can reach out to the project maintainer at k@denysov.me.


# Thank you for using the isup application! We hope it proves to be a valuable addition to your toolkit for domain reachability verification.