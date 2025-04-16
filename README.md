# Podcast Index

## About

A fully typed idiomatic go library to work with  [Podcast Index](https://podcastindex.org/). Documented, Tested, including integration tests. Requires Podcast Index API key and secret.

### Example Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/jjgmckenzie/podcastindex"
)

func main() {
	key := "<YOUR KEY>"
	secret := "<YOUR SECRET>"
	userAgent := "<YOUR USER AGENT>"
	client := podcastindex.NewClient(podcastindex.NewClientOptions{
		UserAgent: userAgent,
		APIKey:    key,
		APISecret: secret,
	})
	ctx := context.Background()
	podcasts := client.SearchPodcastsByTitle(ctx, "test", nil)
	for _, podcast := range podcasts {
		episodes := client.GetEpisodes(ctx, podcast, nil)
		for _, episode := range episodes {
			fmt.Printf("%s : %s\n", podcast.Title, episode.Title)
		}
	}
}
```


### API Coverage

See [API-COVERAGE.md](.API-COVERAGE.md) to see current API coverage by this library. Right now, the library is mostly limited to search, podcasts, and episodes.

If you'd like to contribute to this project by implementing more of the API endpoints such as Value4Value, please refer to the existing implementations as a guide for how to structure your code.

### Contributing

To run integration tests; create a .env file with the following variables.
```
PODCASTINDEX_API_KEY=<YOUR KEY>
PODCASTINDEX_API_SECRET=<YOUR KEY>
```
Please ensure all existing tests pass and that test coverage is full for any new feature/endpoint added.

## License
**License: MPL-2.0**  

You can use this library in both open-source and commercial projects. If you modify any of the source files, you must share those changes under the same license. You don’t need to open source your entire project.

_This is a human-friendly summary, not a substitute for the full license text. See [LICENSE](./LICENSE.md) for details._
