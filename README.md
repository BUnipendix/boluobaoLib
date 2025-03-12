# BoluobaoLib

BoluobaoLib is a Go library for interacting with the Boluobao API. It provides a convenient way to access various features and data related to novels, user accounts, and more.

## Installation

To use BoluobaoLib in your Go project, you can install it using the following command:

```
go get github.com/BUnipendix/boluobaoLib
```

## Usage

### Creating a Client

To start using BoluobaoLib, you need to create a new client instance. You can do this by calling the `NewClient` function and optionally providing configuration options:

```go
client := boluobaoLib.NewClient(
    boluobaoLib.WithCookie("your_cookie"),
    boluobaoLib.WithDeviceId("your_device_id"),
    boluobaoLib.WithDebug(),
    // other options...
)
```

The available options are:

- `WithCookie(cookie string)`: Set the cookie value for authentication.
- `WithDeviceId(deviceId string)`: Set the device ID for identification.
- `WithDebug()`: Enable debug mode for detailed logging.
- `WithOutputDebug()`: Enable output debugging.
- `WithProxyURLArray(proxyURLArray []string)`: Set an array of proxy URLs.
- `WithProxyURL(proxyURL string)`: Set a single proxy URL.
- `WithAPIBaseURL(apiBaseURL string)`: Set the base URL for the API endpoints.
- `WithUserAgent(userAgent string)`: Set the user agent for HTTP requests.
- `WithAuthorization(authorization string)`: Set the authorization token.
- `WithAndroidApiKey(androidApiKey string)`: Set the Android API key.
- `WithSFCommunity(sfCommunity string)`: Set the SFCommunity value in the cookie.
- `WithSessionAPP(sessionApp string)`: Set the session_APP value in the cookie.

### Making API Requests

Once you have created a client, you can use it to make API requests. The `API` struct provides various methods for interacting with the Boluobao API. Here are some examples:

```go
api := client.API()

// Get book shelf information
shelfData, err := api.GetBookShelfInfo()
if err != nil {
    log.Fatal(err)
}

// Get book information
bookID := 123
bookInfo, err := api.GetBookInfo(bookID)
if err != nil {
    log.Fatal(err)
}

// Get chapter content
chapterID := 456
content, err := api.GetChapterContent(chapterID)
if err != nil {
    log.Fatal(err)
}
```

For a complete list of available methods, please refer to the `API` struct in the `api.go` file.

### Models

BoluobaoLib provides several model structs to represent the data returned by the API. These models are defined in the `boluobaomodel` package. Some of the key models include:

- `BookInfoData`: Represents the information of a book.
- `ChapterList`: Represents a list of chapters.
- `ContentData`: Represents the content of a chapter.
- `AccountData`: Represents the information of a user account.

For more details on the available models, please refer to the files in the `boluobaomodel` package.

## Examples

Here are a few examples to help you get started with BoluobaoLib:

### Example 1: Get Book Information

```go
client := boluobaoLib.NewClient()
api := client.API()

bookID := 123
bookInfo, err := api.GetBookInfo(bookID)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Book Title: %s\n", bookInfo.BookName)
fmt.Printf("Author: %s\n", bookInfo.AuthorName)
fmt.Printf("Introduction: %s\n", bookInfo.Expand.Intro)
```

### Example 2: Get Chapter Content

```go
client := boluobaoLib.NewClient()
api := client.API()

chapterID := 456
content, err := api.GetChapterContent(chapterID)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Chapter Title: %s\n", content.ChapName)
fmt.Printf("Content: %s\n", content.Expand.Content)
```

### Example 3: Login and Get User Information

```go
client := boluobaoLib.NewClient()
api := client.API()

username := "your_username"
password := "your_password"
cookie, err := api.Login(username, password)
if err != nil {
    log.Fatal(err)
}

// Set the obtained cookie in the client
client = boluobaoLib.NewClient(boluobaoLib.WithCookie(cookie))
api = client.API()

userInfo, err := api.GetUserInfo()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Username: %s\n", userInfo.UserName)
fmt.Printf("Nickname: %s\n", userInfo.NickName)
fmt.Printf("Email: %s\n", userInfo.Email)
```

## Contributing

Contributions to BoluobaoLib are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.
 