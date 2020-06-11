package transmission

import (
	"encoding/json"
	"reflect"
)

type Response struct {
	Result    string      `json:"result"`
	Arguments interface{} `json:"arguments,omitempty"`
	Tag       int         `json:"tag,omitempty"`
}

type SessionStats struct {
	ActiveTorrentCount int64  `json:"activeTorrentCount"`
	DownloadSpeed      int64  `json:"downloadSpeed"`
	PausedTorrentCount int64  `json:"pausedTorrentCount"`
	TorrentCount       int64  `json:"torrentCount"`
	UploadSpeed        int64  `json:"uploadSpeed"`
	CumulativeStats    *Stats `json:"cumulative-stats"`
	CurrentStats       *Stats `json:"current-stats"`
}

type Stats struct {
	UploadedBytes   int64 `json:"uploadedBytes"`
	DownloadedBytes int64 `json:"downloadedBytes"`
	FilesAdded      int64 `json:"filesAdded"`
	SessionCount    int64 `json:"sessionCount"`
	SecondsActive   int64 `json:"secondsActive"`
}

type TorrentData struct {
	Torrents []Torrent `json:"torrents"`
}

type Torrent struct {
	ActivityDate            int64         `json:"activityDate"`
	AddedDate               int64         `json:"addedDate"`
	BandwidthPriority       int64         `json:"bandwidthPriority"`
	Comment                 string        `json:"comment"`
	CorruptEver             int64         `json:"corruptEver"`
	Creator                 string        `json:"creator"`
	DateCreated             int64         `json:"dateCreated"`
	DesiredAvailable        int64         `json:"desiredAvailable"`
	DoneDate                int64         `json:"doneDate"`
	DownloadDir             string        `json:"downloadDir"`
	DownloadedEver          int64         `json:"downloadedEver"`
	DownloadLimit           int64         `json:"downloadLimit"`
	DownloadLimited         bool          `json:"downloadLimited"`
	EditDate                int64         `json:"editDate"`
	Error                   int64         `json:"error"`
	ErrorString             string        `json:"errorString"`
	ETA                     int64         `json:"eta"`
	ETAIdle                 int64         `json:"etaIdle"`
	Files                   []File        `json:"files"`
	FileStats               []FileStat    `json:"fileStats"`
	HashString              string        `json:"hashString"`
	HaveUnchecked           int64         `json:"haveUnchecked"`
	HaveValid               int64         `json:"haveValid"`
	HonorsSessionLimits     bool          `json:"honorsSessionLimits"`
	ID                      int64         `json:"id"`
	IsFinished              bool          `json:"isFinished"`
	IsPrivate               bool          `json:"isPrivate"`
	IsStalled               bool          `json:"isStalled"`
	Labels                  []Label       `json:"labels"`
	LeftUntilDone           int64         `json:"leftUntilDone"`
	MagnetLink              string        `json:"magnetLink"`
	ManualAnnounceTime      int64         `json:"manualAnnounceTime"`
	MaxConnectedPeers       int64         `json:"maxConnectedPeers"`
	MetadataPercentComplete float64       `json:"metadataPercentComplete"`
	Name                    string        `json:"name"`
	PeerLimit               int64         `json:"peer-limit"`
	Peers                   []Peer        `json:"peers"`
	PeersConnected          int64         `json:"peersConnected"`
	PeersFrom               *PeersFrom    `json:"peersFrom"`
	PeersGettingFromUs      int64         `json:"peersGettingFromUs"`
	PeersSendingToUs        int64         `json:"peersSendingToUs"`
	PercentDone             float64       `json:"percentDone"`
	Pieces                  *Pieces       `json:"pieces"`
	PieceCount              int64         `json:"pieceCount"`
	PieceSize               int64         `json:"pieceSize"`
	Priorities              []Priority    `json:"priorities"`
	QueuePosition           int64         `json:"queuePosition"`
	RateDownload            int64         `json:"rateDownload"`
	RateUpload              int64         `json:"rateUpload"`
	RecheckProgress         float64       `json:"recheckProgress"`
	SecondsDownloading      int64         `json:"secondsDownloading"`
	SecondsSeeding          int64         `json:"secondsSeeding"`
	SeedIdleLimit           int64         `json:"seedIdleLimit"`
	SeedIdleMode            int64         `json:"seedIdleMode"`
	SeedRatioLimit          float64       `json:"seedRatioLimit"`
	SeedRatioMode           int64         `json:"seedRatioMode"`
	SizeWhenDone            int64         `json:"sizeWhenDone"`
	StartDate               int64         `json:"startDate"`
	Status                  int64         `json:"status"`
	Trackers                []Tracker     `json:"trackers"`
	TrackerStats            []TrackerStat `json:"trackerStats"`
	TotalSize               int64         `json:"totalSize"`
	TorrentFile             string        `json:"torrentFile"`
	UploadedEver            int64         `json:"uploadedEver"`
	UploadLimit             int64         `json:"uploadLimit"`
	UploadLimited           bool          `json:"uploadLimited"`
	UploadRatio             float64       `json:"uploadRatio"`
	Wanted                  []Wanted      `json:"wanted"`
	Webseeds                []Webseed     `json:"webseeds"`
	WebseedsSendingToUs     int64         `json:"webseedsSendingToUs"`
}

type File struct {
	BytesCompleted int64  `json:"bytesCompleted"`
	Length         int64  `json:"length"`
	Name           string `json:"name"`
}

type FileStat struct {
	BytesCompleted int64 `json:"bytesCompleted"`
	Wanted         bool  `json:"wanted"`
	Priority       int64 `json:"priority"`
}

type Label string

type Peer struct {
	Address            string  `json:"address"`
	ClientName         string  `json:"clientName"`
	ClientIsChoked     bool    `json:"clientIsChoked"`
	ClientIsInterested bool    `json:"clientIsInterested"`
	FlagStr            string  `json:"flagStr"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom"`
	IsEncrypted        bool    `json:"isEncrypted"`
	IsIncoming         bool    `json:"isIncoming"`
	IsUploadingTo      bool    `json:"isUploadingTo"`
	IsUTP              bool    `json:"isUTP"`
	PeerIsChoked       bool    `json:"peerIsChoked"`
	PeerIsInterested   bool    `json:"peerIsInterested"`
	Port               int64   `json:"port"`
	Progress           float64 `json:"progress"`
	RateToClient       int64   `json:"rateToClient"`
	RateToPeer         int64   `json:"rateToPeer"`
}

type PeersFrom struct {
	FromCache    int64 `json:"fromCache"`
	FromDHT      int64 `json:"fromDht"`
	FromIncoming int64 `json:"fromIncoming"`
	FromLPD      int64 `json:"fromLpd"`
	FromLTEP     int64 `json:"fromLtep"`
	FromPEX      int64 `json:"fromPex"`
	FromTracker  int64 `json:"fromTracker"`
}

type Pieces string

type Priority int64

type Tracker struct {
	Announce string `json:"announce"`
	ID       int64  `json:"id"`
	Scrape   string `json:"scrape"`
	Tier     int64  `json:"tier"`
}

type TrackerStat struct {
	Announce              string `json:"announce"`
	AnnounceState         int64  `json:"announceState"`
	DownloadCount         int64  `json:"downloadCount"`
	HasAnnounced          bool   `json:"hasAnnounced"`
	HasScraped            bool   `json:"hasScraped"`
	Host                  string `json:"host"`
	ID                    int64  `json:"id"`
	IsBackup              bool   `json:"isBackup"`
	LastAnnouncePeerCount int64  `json:"lastAnnouncePeerCount"`
	LastAnnounceResult    string `json:"lastAnnounceResult"`
	LastAnnounceStartTime int64  `json:"lastAnnounceStartTime"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded"`
	LastAnnounceTime      int64  `json:"lastAnnounceTime"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut"`
	LastScrapeResult      string `json:"lastScrapeResult"`
	LastScrapeStartTime   int64  `json:"lastScrapeStartTime"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded"`
	LastScrapeTime        int64  `json:"lastScrapeTime"`
	LastScrapeTimedOut    bool   `json:"lastScrapeTimedOut"` // (bug) it should be bool but the output is number. Thats why the Unmarshal implementation
	LeecherCount          int64  `json:"leecherCount"`
	NextAnnounceTime      int64  `json:"nextAnnounceTime"`
	NextScrapeTime        int64  `json:"nextScrapeTime"`
	Scrape                string `json:"scrape"`
	ScrapeState           int64  `json:"scrapeState"`
	SeederCount           int64  `json:"seederCount"`
	Tier                  int64  `json:"tier"`
}

func (tr *TrackerStat) UnmarshalJSON(data []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	tsType := reflect.TypeOf(tr).Elem()
	tsValue := reflect.ValueOf(tr).Elem()
	for i := 0; i < tsType.NumField(); i++ {
		f := tsType.Field(i)
		if f.PkgPath != "" {
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "" {
			continue
		}
		v := tmp[tag]
		vv := reflect.ValueOf(v)
		if vv.Kind() == reflect.Float64 {
			if tag == "lastScrapeTimedOut" {
				f := vv.Float()
				tsValue.Field(i).SetBool(f != 0)
			} else {
				tsValue.Field(i).SetInt(int64(vv.Float()))
			}
		} else {
			tsValue.Field(i).Set(vv)
		}
	}
	return nil
}

type Wanted int64 // (bug) should be bool according to spec
type Webseed string

type Session struct {
	AltSpeedDown              int64   `json:"alt-speed-down"`
	AltSpeedEnabled           bool    `json:"alt-speed-enabled"`
	AltSpeedTimeBegin         int64   `json:"alt-speed-time-begin"`
	AltSpeedTimeEnabled       bool    `json:"alt-speed-time-enabled"`
	AltSpeedTimeEnd           int64   `json:"alt-speed-time-end"`
	AltSpeedTimeDay           int64   `json:"alt-speed-time-day"`
	AltSpeedUp                int64   `json:"alt-speed-up"`
	BlocklistURL              string  `json:"blocklist-url"`
	BlocklistEnabled          bool    `json:"blocklist-enabled"`
	BlocklistSize             int64   `json:"blocklist-size"`
	CacheSizeMb               int64   `json:"cache-size-mb"`
	ConfigDir                 string  `json:"config-dir"`
	DownloadDir               string  `json:"download-dir"`
	DownloadQueueSize         int64   `json:"download-queue-size"`
	DownloadQueueEnabled      bool    `json:"download-queue-enabled"`
	DhtEnabled                bool    `json:"dht-enabled"`
	Encryption                string  `json:"encryption"`
	IdleSeedingLimit          int64   `json:"idle-seeding-limit"`
	IdleSeedingLimitEnabled   bool    `json:"idle-seeding-limit-enabled"`
	IncompleteDir             string  `json:"incomplete-dir"`
	IncompleteDirEnabled      bool    `json:"incomplete-dir-enabled"`
	LPDEnabled                bool    `json:"lpd-enabled"`
	PeerLimitGlobal           int64   `json:"peer-limit-global"`
	PeerLimitPerTorrent       int64   `json:"peer-limit-per-torrent"`
	PEXEnabled                bool    `json:"pex-enabled"`
	PeerPort                  int64   `json:"peer-port"`
	PeerPortRandomOnStart     bool    `json:"peer-port-random-on-start"`
	PortForwardingEnabled     bool    `json:"port-forwarding-enabled"`
	QueueStalledEnabled       bool    `json:"queue-stalled-enabled"`
	QueueStalledMinutes       int64   `json:"queue-stalled-minutes"`
	RenamePartialFiles        bool    `json:"rename-partial-files"`
	RPCVersion                int64   `json:"rpc-version"`
	RPCVersionMinimum         int64   `json:"rpc-version-minimum"`
	ScriptTorrentDoneFilename string  `json:"script-torrent-done-filename"`
	ScriptTorrentDoneEnabled  bool    `json:"script-torrent-done-enabled"`
	SeedRatioLimit            float64 `json:"seedRatioLimit"`
	SeedRatioLimited          bool    `json:"seedRatioLimited"`
	SeedQueueSize             int64   `json:"seed-queue-size"`
	SeedQueueEnabled          bool    `json:"seed-queue-enabled"`
	SpeedLimitDown            int64   `json:"speed-limit-down"`
	SpeedLimitDownEnabled     bool    `json:"speed-limit-down-enabled"`
	SpeedLimitUp              int64   `json:"speed-limit-up"`
	SpeedLimitUpEnabled       bool    `json:"speed-limit-up-enabled"`
	StartAddedTorrents        bool    `json:"start-added-torrents"`
	TrashOriginalTorrentFiles bool    `json:"trash-original-torrent-files"`
	Units                     *Units  `json:"units"`
	UTPEnabled                bool    `json:"utp-enabled"`
	Version                   string  `json:"version"`
}

type Units struct {
	SpeedUnits  []string `json:"speed-units"`
	SpeedBytes  int64    `json:"speed-bytes"`
	SizeUnits   []string `json:"size-units"`
	SizeBytes   int64    `json:"size-bytes"`
	MemoryUnits []string `json:"memory-units"`
	MemoryBytes int64    `json:"memory-bytes"`
}
