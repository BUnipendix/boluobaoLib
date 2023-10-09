# boluobaoLib

`boluobaoLib` is a Go package that provides a client for interacting with the Boluobao API. It allows you to perform various operations such as retrieving book information, getting user information, searching for books, and more.

## Usage

To use the `boluobaoLib` package, you need to import it in your Go code:

```go
import "github.com/AlexiaVeronica/boluobaoLib"
```

### Creating a Client

To interact with the Boluobao API, you first need to create a client. The client provides methods for performing API requests. You can create a client using the `NewClient` function:

```go
client := boluobaoLib.NewClient()
```

The `NewClient` function returns a new instance of the `Client` struct.

### Options

The `NewClient` function accepts optional configuration options that allow you to customize the client's behavior. The available options are:

- `WithCookie`: Sets the cookie value for the API requests.
- `WithDeviceId`: Sets the device ID for the API requests.
- `WithDebug`: Enables or disables debug mode for the HTTP client.
- `WithOutputDebug`: Enables or disables outputting debug information to a file.
- `WithProxyURL`: Sets the proxy URL for the HTTP client.

You can pass these options to the `NewClient` function as variadic arguments:

```go
client := boluobaoLib.NewClient(
    boluobaoLib.WithCookie("your-cookie-value"),
    boluobaoLib.WithDeviceId("your-device-id"),
    boluobaoLib.WithDebug(),
    boluobaoLib.WithOutputDebug(),
    boluobaoLib.WithProxyURL("your-proxy-url"),
)
```

### API Methods

The `Client` struct provides methods for interacting with the Boluobao API. Here are some of the available methods:

- `GetBookShelfInfo`: Retrieves the bookshelf information for the authenticated user.
- `GetBookInfo`: Retrieves information about a specific book.
- `GetCatalogue`: Retrieves the catalogue of chapters for a specific book.
- `GetChapterContent`: Retrieves the content of a specific chapter.
- `GetUserInfo`: Retrieves information about the authenticated user.
- `GetUserMoney`: Retrieves the user's money information.
- `GetCurreyIp`: Retrieves the current IP address of the user.
- `GetRankWeekArray`: Retrieves the weekly ranking list of books.
- `GetRankMonthArray`: Retrieves the monthly ranking list of books.
- `GetOtherUserInfo`: Retrieves information about another user.
- `GetUserWorks`: Retrieves the works (books) of a specific user.
- `Login`: Performs a login request with the specified username and password.
- `GetSearch`: Performs a search request with the specified keyword.

Each method returns the corresponding data or an error if the request fails.

## Example

Here's an example that demonstrates how to use the `boluobaoLib` package:

```go
package main

import (
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib"
)

func main() {
	// Create a client
	client := boluobaoLib.NewClient()

	// Retrieve bookshelf information
	bookshelf, err := client.GetBookShelfInfo()
	if err != nil {
		fmt.Println("Failed to retrieve bookshelf information:", err)
		return
	}

	// Print the bookshelf information
	for _, book := range bookshelf {
		fmt.Println("Book ID:", book.ID)
		fmt.Println("Book Title:", book.Title)
		fmt.Println("Book Author:", book.Author)
		fmt.Println("Book Status:", book.Status)
		fmt.Println("--------")
	}
}
```

In this example, we create a client using the `NewClient` function and then use the `GetBookShelfInfo` method to retrieve the bookshelf information. We iterate over the books in the bookshelf and print their details.

## Conclusion

The `boluobaoLib` package provides a convenient way to interact with the Boluobao API in Go. It simplifies the process of making API requests and handling the responses. With the `boluobaoLib` package, you can easily perform various operations related to books and users in your Go applications.