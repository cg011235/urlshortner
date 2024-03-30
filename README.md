URL Shortener Service
This project is a URL shortener service, designed to convert long URLs into more manageable, shorter versions, making it easier to share links. It includes a RESTful API to create shortened URLs and redirect users from short URLs to the original long URLs. The service also provides an endpoint to retrieve the top 3 most shortened domains and their counts.


Features
Shorten URLs: Convert long URLs into shorter, more shareable versions.
Redirection: Users accessing the shortened URL are redirected to the original URL.
Analytics: Retrieve the top 3 most frequently shortened domains and their corresponding counts.


Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.


Prerequisites
Go (version 1.15 or higher)


Build the Service
Run: ./build.sh


Run using docker
Run: ./run_docker.sh


Usage of REST APIs

1. Shorten a URL
Send a POST request to 127.0.0.1:8080/shorten with a JSON body containing the original URL:
{
  "originalUrl": "https://www.example.com/very/long/url"
}

In response you will get short URL as JSON response.
{
    "shortUrl": "aaaaaab"
}

2. Click / redirect to short URL
Use bowser or send get request to short URL: 127.0.0.1:8080/aaaaaab

In response, the API will lookup corresponding long URL and redirect to original URL.

3. Get Top 3 Shortened Domains
Send a GET request to /metrics to retrieve the top 3 most shortened domains.

In response, the API will produce following output.
{
    "topDomains": [
        {
            "domain": "docs.google.com",
            "count": 7
        },
        {
            "domain": "www.audible.in",
            "count": 1
        }
    ]
}

Running the Tests
Run the unit tests for this system with the following command:
go test ./...

Authors
Chaitanya Gotkindikar (cg011235@gmail.com)