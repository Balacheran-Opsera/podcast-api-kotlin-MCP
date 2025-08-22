package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// EpisodeFull represents the EpisodeFull schema from the OpenAPI specification
type EpisodeFull struct {
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Title string `json:"title,omitempty"` // Episode name.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Transcript string `json:"transcript,omitempty"` // The transcript of this episode, in plain text (with the newline character \n). If there's not transcript, it is null. This field is available only in the PRO/ENTERPRISE plan.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Podcast PodcastSimple `json:"podcast,omitempty"`
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Link string `json:"link,omitempty"` // Web link of this episode.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
}

// SubmitPodcastForm represents the SubmitPodcastForm schema from the OpenAPI specification
type SubmitPodcastForm struct {
	Email string `json:"email,omitempty"` // A valid email address. If **email** is specified, then we'll notify this email address once the podcast is accepted.
	Rss string `json:"rss"` // A valid podcast rss url.
}

// EpisodeSearchResult represents the EpisodeSearchResult schema from the OpenAPI specification
type EpisodeSearchResult struct {
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Transcripts_highlighted []string `json:"transcripts_highlighted,omitempty"` // Up to 2 highlighted segments of the audio transcript of this episode.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Title_original string `json:"title_original,omitempty"` // Plain text of this episode' title
	Description_original string `json:"description_original,omitempty"` // Plain text of this episode's description
	Podcast map[string]interface{} `json:"podcast,omitempty"` // The podcast that this episode belongs to.
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of this episode's title
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of this episode's description
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Link string `json:"link,omitempty"` // Web link of this episode.
}

// Genre represents the Genre schema from the OpenAPI specification
type Genre struct {
	Id int `json:"id,omitempty"` // Genre id
	Name string `json:"name,omitempty"` // Genre name.
	Parent_id int `json:"parent_id,omitempty"` // Parent genre id.
}

// DeletedItem represents the DeletedItem schema from the OpenAPI specification
type DeletedItem struct {
	Status string `json:"status,omitempty"` // The status of this episode or podcast. For now, the only possible value is **deleted**.
	Title string `json:"title,omitempty"` // Episode title or podcast title.
	ErrorField string `json:"error,omitempty"` // Why this episode or podcast is deleted?
	Id string `json:"id,omitempty"` // Episode id or podcast id.
}

// GetPodcastRecommendationsResponse represents the GetPodcastRecommendationsResponse schema from the OpenAPI specification
type GetPodcastRecommendationsResponse struct {
	Recommendations []PodcastSimple `json:"recommendations"`
}

// TypeaheadResponse represents the TypeaheadResponse schema from the OpenAPI specification
type TypeaheadResponse struct {
	Genres []Genre `json:"genres,omitempty"` // Genre suggestions. It'll show up when the **show_genres** parameter is *1*.
	Podcasts []PodcastTypeaheadResult `json:"podcasts,omitempty"` // Podcast suggestions. It'll show up when the **show_podcasts** parameter is 1.
	Terms []string `json:"terms"` // Search term suggestions.
}

// EpisodeMinimum represents the EpisodeMinimum schema from the OpenAPI specification
type EpisodeMinimum struct {
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Title string `json:"title,omitempty"` // Episode name.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Link string `json:"link,omitempty"` // Web link of this episode.
}

// EpisodeSimple represents the EpisodeSimple schema from the OpenAPI specification
type EpisodeSimple struct {
	Title string `json:"title,omitempty"` // Episode name.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Podcast PodcastMinimum `json:"podcast,omitempty"`
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
}

// PlaylistItem represents the PlaylistItem schema from the OpenAPI specification
type PlaylistItem struct {
	Notes string `json:"notes,omitempty"` // Notes for this item.
	TypeField string `json:"type,omitempty"` // The type of this playlist item. If a playlist is **episode_list**, then an item could be either **episode** or **custom_audio**. If it's **podcast_list**, then an item can only be **podcast**.
	Added_at_ms int `json:"added_at_ms,omitempty"` // Timestamp (in milliseconds) when this item is added.
	Data interface{} `json:"data,omitempty"`
	Id int `json:"id,omitempty"` // Playlist item id.
}

// SpellCheckResponse represents the SpellCheckResponse schema from the OpenAPI specification
type SpellCheckResponse struct {
	Corrected_text_html string `json:"corrected_text_html"` // The corrected text for the entire search term (multiple words/tokens), where misspelled tokens are replaced with the correct texts and html tags <b><i>
	Tokens []map[string]interface{} `json:"tokens"` // The word in the text query string that is not spelled correctly
}

// GetPodcastsInBatchForm represents the GetPodcastsInBatchForm schema from the OpenAPI specification
type GetPodcastsInBatchForm struct {
	Itunes_ids string `json:"itunes_ids,omitempty"` // Comma-separated Apple Podcasts (iTunes) ids, e.g., 659155419
	Next_episode_pub_date int `json:"next_episode_pub_date,omitempty"` // For latest episodes pagination. It's the value of **next_episode_pub_date** from the response of last request. If not specified, just return latest 15 episodes.
	Rsses string `json:"rsses,omitempty"` // Comma-separated rss urls.
	Show_latest_episodes int `json:"show_latest_episodes,omitempty"` // Whether or not to fetch up to 15 latest episodes from these podcasts, sorted by pub_date. 1 is yes, and 0 is no.
	Spotify_ids string `json:"spotify_ids,omitempty"` // Comma-separated Spotify ids, e.g., 3DDfEsKDIDrTlnPOiG4ZF4
	Ids string `json:"ids,omitempty"` // Comma-separated list of podcast ids.
}

// GetRegionsResponse represents the GetRegionsResponse schema from the OpenAPI specification
type GetRegionsResponse struct {
	Regions map[string]interface{} `json:"regions"`
}

// TrendingSearchesResponse represents the TrendingSearchesResponse schema from the OpenAPI specification
type TrendingSearchesResponse struct {
	Terms []string `json:"terms"` // Trending search terms
}

// PodcastSimple represents the PodcastSimple schema from the OpenAPI specification
type PodcastSimple struct {
	Looking_for PodcastLookingForField `json:"looking_for,omitempty"`
	Title string `json:"title,omitempty"` // Podcast name.
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Genre_ids []int `json:"genre_ids,omitempty"`
	Is_claimed bool `json:"is_claimed,omitempty"` // Whether this podcast is claimed by its producer on [ListenNotes.com](https://www.ListenNotes.com).
	Country string `json:"country,omitempty"` // The country where this podcast is produced.
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Language string `json:"language,omitempty"` // The language of this podcast. You can get all supported languages from `GET /languages`.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	TypeField string `json:"type,omitempty"` // The type of this podcast - episodic or serial.
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Extra PodcastExtraField `json:"extra,omitempty"`
}

// RelatedSearchesResponse represents the RelatedSearchesResponse schema from the OpenAPI specification
type RelatedSearchesResponse struct {
	Terms []string `json:"terms"` // Related search terms
}

// PodcastMinimum represents the PodcastMinimum schema from the OpenAPI specification
type PodcastMinimum struct {
	Title string `json:"title,omitempty"` // Podcast name.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
}

// BestPodcastsResponse represents the BestPodcastsResponse schema from the OpenAPI specification
type BestPodcastsResponse struct {
	Listennotes_url string `json:"listennotes_url"` // Url of the list of best podcasts on [ListenNotes.com](https://www.ListenNotes.com).
	Previous_page_number int `json:"previous_page_number"`
	Has_next bool `json:"has_next"`
	Has_previous bool `json:"has_previous"`
	Id int `json:"id"` // The id of this genre
	Total int `json:"total"`
	Next_page_number int `json:"next_page_number"`
	Parent_id int `json:"parent_id"` // The id of parent genre.
	Name string `json:"name"` // This genre's name.
	Page_number int `json:"page_number"`
	Podcasts []PodcastSimple `json:"podcasts"`
}

// PodcastAudienceResponse represents the PodcastAudienceResponse schema from the OpenAPI specification
type PodcastAudienceResponse struct {
	By_regions []map[string]interface{} `json:"by_regions,omitempty"`
}

// CuratedListSearchResult represents the CuratedListSearchResult schema from the OpenAPI specification
type CuratedListSearchResult struct {
	Podcasts []PodcastMinimum `json:"podcasts,omitempty"` // Up to 5 podcasts in this curated list.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of this curated list's title
	Title_original string `json:"title_original,omitempty"` // Plain text of this curated list's title
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of this curated list's description
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Description_original string `json:"description_original,omitempty"` // Plain text of this curated list's description
}

// CuratedListFull represents the CuratedListFull schema from the OpenAPI specification
type CuratedListFull struct {
	Podcasts []PodcastSimple `json:"podcasts,omitempty"` // Complete meta data of all podcasts in this curated list.
	Title string `json:"title,omitempty"` // Curated list name.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Description string `json:"description,omitempty"` // This curated list's description.
}

// PodcastMinimumRss represents the PodcastMinimumRss schema from the OpenAPI specification
type PodcastMinimumRss struct {
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title string `json:"title,omitempty"` // Podcast name.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
}

// PodcastLookingForField represents the PodcastLookingForField schema from the OpenAPI specification
type PodcastLookingForField struct {
	Cohosts bool `json:"cohosts,omitempty"` // Whether this podcast is looking for cohosts.
	Cross_promotion bool `json:"cross_promotion,omitempty"` // Whether this podcast is looking for cross promotion opportunities with other podcasts.
	Guests bool `json:"guests,omitempty"` // Whether this podcast is looking for guests.
	Sponsors bool `json:"sponsors,omitempty"` // Whether this podcast is looking for sponsors.
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	Next_offset int `json:"next_offset,omitempty"` // Pass this value to the **offset** parameter to do pagination of search results.
	Results []interface{} `json:"results,omitempty"` // A list of search results.
	Took float64 `json:"took,omitempty"` // The time it took to fetch these search results. In seconds.
	Total int `json:"total,omitempty"` // The total number of search results.
	Count int `json:"count,omitempty"` // The number of search results in this page.
}

// GetEpisodesInBatchResponse represents the GetEpisodesInBatchResponse schema from the OpenAPI specification
type GetEpisodesInBatchResponse struct {
	Episodes []EpisodeSimple `json:"episodes"`
}

// GetLanguagesResponse represents the GetLanguagesResponse schema from the OpenAPI specification
type GetLanguagesResponse struct {
	Languages []string `json:"languages"`
}

// CustomAudio represents the CustomAudio schema from the OpenAPI specification
type CustomAudio struct {
	Title string `json:"title,omitempty"` // Custom audio title.
	Audio string `json:"audio,omitempty"` // Audio url, which can be played directly.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length in seconds.
	Image string `json:"image,omitempty"` // High resolution image url of this custom audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date (in milliseconds) of this custom audio. For now, it's the same as **added_at_ms** of this playlist item.
	Thumbnail string `json:"thumbnail,omitempty"` // Low resolution image url of this custom audio.
}

// PodcastFull represents the PodcastFull schema from the OpenAPI specification
type PodcastFull struct {
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Is_claimed bool `json:"is_claimed,omitempty"` // Whether this podcast is claimed by its producer on [ListenNotes.com](https://www.ListenNotes.com).
	Language string `json:"language,omitempty"` // The language of this podcast. You can get all supported languages from `GET /languages`.
	Country string `json:"country,omitempty"` // The country where this podcast is produced.
	Looking_for PodcastLookingForField `json:"looking_for,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Extra PodcastExtraField `json:"extra,omitempty"`
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Next_episode_pub_date int `json:"next_episode_pub_date,omitempty"` // Passed to the **next_episode_pub_date** parameter of `GET /podcasts/{id}` to paginate through episodes of that podcast.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Title string `json:"title,omitempty"` // Podcast name.
	Episodes []EpisodeMinimum `json:"episodes,omitempty"`
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	TypeField string `json:"type,omitempty"` // The type of this podcast - episodic or serial.
	Genre_ids []int `json:"genre_ids,omitempty"`
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
}

// CuratedListSimple represents the CuratedListSimple schema from the OpenAPI specification
type CuratedListSimple struct {
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Title string `json:"title,omitempty"` // Curated list name.
	Description string `json:"description,omitempty"` // This curated list's description.
	Podcasts []PodcastMinimum `json:"podcasts,omitempty"` // Minimum meta data of up to 5 podcasts in this curated list.
}

// PodcastExtraField represents the PodcastExtraField schema from the OpenAPI specification
type PodcastExtraField struct {
	Amazon_music_url string `json:"amazon_music_url,omitempty"` // Amazon Music url for this podcast
	Facebook_handle string `json:"facebook_handle,omitempty"` // Facebook username affiliated with this podcast
	Patreon_handle string `json:"patreon_handle,omitempty"` // Patreon username affiliated with this podcast
	Twitter_handle string `json:"twitter_handle,omitempty"` // Twitter username affiliated with this podcast
	Linkedin_url string `json:"linkedin_url,omitempty"` // LinkedIn url affiliated with this podcast
	Spotify_url string `json:"spotify_url,omitempty"` // Spotify url for this podcast
	Url1 string `json:"url1,omitempty"` // Url affiliated with this podcast
	Url2 string `json:"url2,omitempty"` // Url affiliated with this podcast
	Wechat_handle string `json:"wechat_handle,omitempty"` // WeChat username affiliated with this podcast
	Youtube_url string `json:"youtube_url,omitempty"` // YouTube url affiliated with this podcast
	Google_url string `json:"google_url,omitempty"` // Google Podcasts url for this podcast
	Instagram_handle string `json:"instagram_handle,omitempty"` // Instagram username affiliated with this podcast
	Url3 string `json:"url3,omitempty"` // Url affiliated with this podcast
}

// GetGenresResponse represents the GetGenresResponse schema from the OpenAPI specification
type GetGenresResponse struct {
	Genres []Genre `json:"genres"`
}

// DeletePodcastResponse represents the DeletePodcastResponse schema from the OpenAPI specification
type DeletePodcastResponse struct {
	Status string `json:"status"` // The status of this podcast deletion request.
}

// PlaylistResponse represents the PlaylistResponse schema from the OpenAPI specification
type PlaylistResponse struct {
	Items []PlaylistItem `json:"items,omitempty"` // A list of playlist items.
	Last_timestamp_ms int `json:"last_timestamp_ms,omitempty"` // Passed to the **last_timestamp_ms** parameter of `GET /playlists/{id}` to paginate through items of that playlist.
	TypeField string `json:"type,omitempty"` // The type of this playlist, which should be either **episode_list** or **podcast_list**.
	Description string `json:"description,omitempty"` // Playlist description.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this playlist on ListenNotes.com.
	Name string `json:"name,omitempty"` // Playlist name.
	Thumbnail string `json:"thumbnail,omitempty"` // Low resolution image url of the playlist.
	Total int `json:"total,omitempty"` // Total number of items in this playlist.
	Visibility string `json:"visibility,omitempty"` // Visibility of this playlist.
	Id string `json:"id,omitempty"` // A 11-character playlist id, which can be used to further fetch detailed playlist metadata via `GET /playlists/{id}`.
	Total_audio_length_sec int `json:"total_audio_length_sec,omitempty"` // Total audio length of all episodes in this playlist, in seconds. It will have a valid value only when type is **episode_list**. In other words, it will be 0 if type is **podcast_list**.
	Image string `json:"image,omitempty"` // High resolution image url of the playlist.
}

// PodcastTypeaheadResult represents the PodcastTypeaheadResult schema from the OpenAPI specification
type PodcastTypeaheadResult struct {
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Publisher_highlighted string `json:"publisher_highlighted,omitempty"` // Highlighted segment of this podcast's publisher name.
	Publisher_original string `json:"publisher_original,omitempty"` // Plain text of this podcast's publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of podcast name.
	Title_original string `json:"title_original,omitempty"` // Plain text of podcast name.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
}

// SubmitPodcastResponse represents the SubmitPodcastResponse schema from the OpenAPI specification
type SubmitPodcastResponse struct {
	Podcast PodcastMinimum `json:"podcast"`
	Status string `json:"status"` // The status of this submission.
}

// GetEpisodeRecommendationsResponse represents the GetEpisodeRecommendationsResponse schema from the OpenAPI specification
type GetEpisodeRecommendationsResponse struct {
	Recommendations []EpisodeSimple `json:"recommendations"`
}

// PlaylistsResponse represents the PlaylistsResponse schema from the OpenAPI specification
type PlaylistsResponse struct {
	Page_number int `json:"page_number,omitempty"`
	Playlists []map[string]interface{} `json:"playlists,omitempty"`
	Previous_page_number int `json:"previous_page_number,omitempty"`
	Total int `json:"total,omitempty"`
	Has_next bool `json:"has_next,omitempty"`
	Has_previous bool `json:"has_previous,omitempty"`
	Next_page_number int `json:"next_page_number,omitempty"`
}

// PodcastSearchResult represents the PodcastSearchResult schema from the OpenAPI specification
type PodcastSearchResult struct {
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of podcast description
	Description_original string `json:"description_original,omitempty"` // Plain text of podcast description
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of podcast name.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Title_original string `json:"title_original,omitempty"` // Plain text of podcast name.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Genre_ids []int `json:"genre_ids,omitempty"`
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher_original string `json:"publisher_original,omitempty"` // Plain text of this podcast's publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Publisher_highlighted string `json:"publisher_highlighted,omitempty"` // Highlighted segment of this podcast's publisher name.
}

// GetPodcastsInBatchResponse represents the GetPodcastsInBatchResponse schema from the OpenAPI specification
type GetPodcastsInBatchResponse struct {
	Latest_episodes []EpisodeSimple `json:"latest_episodes,omitempty"` // Up to 10 latest episodes from these podcasts, sorted by **pub_date**. This field shows up only when **show_latest_episodes** is 1.
	Podcasts []PodcastSimple `json:"podcasts"`
}

// PodcastDomainResponse represents the PodcastDomainResponse schema from the OpenAPI specification
type PodcastDomainResponse struct {
	Previous_page_number int `json:"previous_page_number,omitempty"`
	Has_next bool `json:"has_next,omitempty"`
	Has_previous bool `json:"has_previous,omitempty"`
	Next_page_number int `json:"next_page_number,omitempty"`
	Page_number int `json:"page_number,omitempty"`
	Podcasts []PodcastSimple `json:"podcasts,omitempty"`
}

// GetEpisodesInBatchForm represents the GetEpisodesInBatchForm schema from the OpenAPI specification
type GetEpisodesInBatchForm struct {
	Ids string `json:"ids"` // Comma-separated list of episode ids.
}

// GetCuratedPodcastsResponse represents the GetCuratedPodcastsResponse schema from the OpenAPI specification
type GetCuratedPodcastsResponse struct {
	Has_previous bool `json:"has_previous"`
	Next_page_number int `json:"next_page_number"`
	Page_number int `json:"page_number"`
	Previous_page_number int `json:"previous_page_number"`
	Total int `json:"total"`
	Curated_lists []CuratedListSimple `json:"curated_lists"`
	Has_next bool `json:"has_next"`
}
