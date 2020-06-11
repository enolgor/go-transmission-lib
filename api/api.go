package api

import (
	"net/url"

	"github.com/enolgor/go-transmission-lib/transmission"
)

type Api struct {
	Client *transmission.Client
	tag    int
}

func NewApi(rpcURL string) (*Api, error) {
	parsedURL, err := url.Parse(rpcURL)
	if err != nil {
		return nil, err
	}
	return &Api{Client: &transmission.Client{URL: parsedURL}}, nil
}

func (api *Api) newTag() int {
	api.tag++
	return api.tag
}

func (api *Api) exec(resp interface{}, method string, args map[string]interface{}) error {
	response := &transmission.Response{Arguments: resp}
	request := transmission.NewRequest(method, args, api.newTag())
	return api.Client.Do(request, response)
}

func (api *Api) GetTorrents(ids ...int) ([]transmission.Torrent, error) {
	var torrentData transmission.TorrentData
	reqArgs := make(map[string]interface{})
	if len(ids) > 0 {
		reqArgs["ids"] = ids
	}
	reqArgs["fields"] = torrentGetFields
	if err := api.exec(&torrentData, "torrent-get", reqArgs); err != nil {
		return nil, err
	}
	return torrentData.Torrents, nil
}

func (api *Api) GetSession() (*transmission.Session, error) {
	session := &transmission.Session{}
	if err := api.exec(session, "session-get", nil); err != nil {
		return nil, err
	}
	return session, nil
}

func (api *Api) SessionStats() (*transmission.SessionStats, error) {
	sessionStats := &transmission.SessionStats{}
	if err := api.exec(sessionStats, "session-stats", nil); err != nil {
		return nil, err
	}
	return sessionStats, nil
}

var torrentGetFields []string = []string{
	"activityDate",
	"addedDate",
	"bandwidthPriority",
	"comment",
	"corruptEver",
	"creator",
	"dateCreated",
	"desiredAvailable",
	"doneDate",
	"downloadDir",
	"downloadedEver",
	"downloadLimit",
	"downloadLimited",
	"editDate",
	"error",
	"errorString",
	"eta",
	"etaIdle",
	"files",
	"fileStats",
	"hashString",
	"haveUnchecked",
	"haveValid",
	"honorsSessionLimits",
	"id",
	"isFinished",
	"isPrivate",
	"isStalled",
	"labels",
	"leftUntilDone",
	"magnetLink",
	"manualAnnounceTime",
	"maxConnectedPeers",
	"metadataPercentComplete",
	"name",
	"peer-limit",
	"peers",
	"peersConnected",
	"peersFrom",
	"peersGettingFromUs",
	"peersSendingToUs",
	"percentDone",
	"pieces",
	"pieceCount",
	"pieceSize",
	"priorities",
	"queuePosition",
	"rateDownload",
	"rateUpload",
	"recheckProgress",
	"secondsDownloading",
	"secondsSeeding",
	"seedIdleLimit",
	"seedIdleMode",
	"seedRatioLimit",
	"seedRatioMode",
	"sizeWhenDone",
	"startDate",
	"status",
	"trackers",
	"trackerStats",
	"totalSize",
	"torrentFile",
	"uploadedEver",
	"uploadLimit",
	"uploadLimited",
	"uploadRatio",
	"wanted",
	"webseeds",
	"webseedsSendingToUs",
}
