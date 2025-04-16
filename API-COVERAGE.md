# Podcast Index API Implementation Coverage

This document tracks which endpoints from the [Podcast Index API](https://podcastindex-org.github.io/docs-api/) are implemented in this Go library.

## API Endpoints Coverage

### Search
Search the index

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/search/byterm` | Search for podcasts by term | ✅ | `SearchPodcastsByTerm()` |
| `/search/bytitle` | Search for podcasts by title | ✅ | `SearchPodcastsByTitle()` |
| `/search/byperson` | Search for podcasts by person | ✅ | `SearchPodcastsByPerson()` |
| `/search/music/byterm` | Search for music podcasts by term | ✅ | `SearchMusicPodcastsByTerm()` |

### Podcasts
Find details about a Podcast and its feed.

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-----------|-----------------|
| `/podcasts/byfeedid` | Get podcast by feed ID | ✅ | `GetPodcastByFeedID()` |
| `/podcasts/byfeedurl` | Get podcast by feed URL | ✅ | `GetPodcastByURL()` |
| `/podcasts/byitunesid` | Get podcast by iTunes ID | ✅ | `GetPodcastByITunesID()` |
| `/podcasts/byguid` | Get podcast by GUID | ✅ | `GetPodcastByGUID()` |
| `/podcasts/bytag` | Get podcast by Tag | ❌ | - |
| `/podcasts/bymedium` | Get podcast by Medium | ❌ | - |
| `/podcasts/trending` | Get trending podcasts | ❌ | - |
| `/podcasts/dead` | Get feeds that have been marked dead | ❌ | - |
| `/podcasts/batch/byguid` | Get feed info from GUIDs provided in JSON array | ❌ | - |

### Episodes
Find details about one or more episodes of a podcast or podcasts.

| Endpoint                   | Description                                 | Implemented | Client Function |
|----------------------------|---------------------------------------------|-------------|----------------|
| **N/A  - Helper Function** | Get episode by using a podcastindex Podcast |✅ | `GetEpisodes()`  |
| `/episodes/byfeedid`       | Get episodes by podcast feed ID             | ❌ | -              |
| `/episodes/byfeedurl`      | Get episodes by podcast feed URL            | ❌ | -              |
| `/episodes/bypodcastguid`  | Get episodes by podcast feed GUID           | ❌ | -              |
| `/episodes/byitunesid`     | Get episodes by podcast feed iTunes ID      | ❌ | -              |
| `/episodes/byid`           | Get episode metadata by ID                  | ❌ | -              |
| `/episodes/byguid`         | Get episode metadata by GUID                | ❌ | -              |
| `/episodes/live`           | Get episodes with podcast:liveitem tag      | ❌ | -              |
| `/episodes/random`         | Get random batch of episodes                | ❌ | -              |

### Recent
Find recent additions to the index

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/recent/episodes` | Get recent episodes | ❌ | - |
| `/recent/feeds` | Get recent feeds | ❌ | - |
| `/recent/newfeeds` | Get recent new feeds | ❌ | - |
| `/recent/newvaluefeeds` | Get recent new feeds with a value tag | ❌ | - |
| `/recent/data` | This call returns every new feed and episode added to the index over the past 24 hours in reverse chronological order. | ❌ | - |
| `/recent/soundbites` | Get recent soundbites | ❌ | - |

### Value
The podcast's "Value for Value" information

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/value/byfeedid` | Get value4value by feed ID | ❌ | - |
| `/value/byitunesid` | Get value4value by iTunes ID | ❌ | - |
| `/value/bypodcastguid` | Get value4value by podcast GUID | ❌ | - |
| `/value/byepisodeguid` | Get value4value by episode GUID | ❌ | - |
| `/value/batch/byepisodeguid` | This call returns the information for supporting the podcast episode via one of the "Value for Value" methods from a JSON object containing one or more podcast GUID and one or more episode GUID for the podcast. | ❌ | - |

### Stats
Statistics for items in the Podcast Index

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/stats/current` | Get current stats | ❌ | - |

### Categories
Categories used by the Podcast Index

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/categories/list` | Get list of podcast categories | ✅ | `Categories()` |


### Hub
Notify the index that a feed has changed

| Endpoint | Description | Implemented | Library Function |
|----------|-------------|-------------|-----------------|
| `/hub/pubnotify` | Notify the index that a feed has changed | ❌ | - |

### Add
Add new podcast feeds to the index.

**NOTE: To add to the index, the API Key must have write or publisher permissions.**

| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/add/byfeedurl` | Add podcast by feed URL | ❌ | - |
| `/add/byitunesid` | Add podcast by iTunes ID | ❌ | - |

### Apple Replacement
| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| `/search` | Replaces the Apple search API but returns data from the Podcast Index database. | ❌ | - |
| `/lookup` | Replaces the Apple podcast lookup API but returns data from the Podcast Index database. | ❌ | - |

### Static Data
| Endpoint | Description | Implemented | Client Function |
|----------|-------------|-------------|-----------------|
| ` /static/stats/daily_counts.json` | Report a number of statistics about the feeds in Podcast Index's database. Updated daily. | ❌ | - |

## Contributing

If you'd like to contribute to this project by implementing more API endpoints, please refer to the existing implementations as a guide for how to structure your code.
