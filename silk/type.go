package silk

type Type uint8

const (
	Close Type = iota
	Heartbeat
	Connect //json{..info...}
	ConnectAck
	Subscribe //topic
	SubscribeAck
	Unsubscribe //topic
	UnsubscribeAck
	Publish //topic, message
	PublishAck
	Message //topic, message

)

const (
	TunnelCreate       Type = iota + 16 //net,addr 例 tcp,127.0.0.1:8080
	TunnelCreateAck                     //id(uint16)
	TunnelClose                         //id
	TunnelCloseAck                      //id
	TunnelTransferData                  //id,data
	//TunnelTransferEnd

)

const (
	SystemShell      Type = iota + 32 // /bin/sh
	SystemShellAck                    //tunnel id(uint16)
	SystemExecute                     //command string
	SystemExecuteAck                  //stdout
	SystemConfig
	SystemConfigAck  //yaml
	SystemDbQuery    //sql
	SystemDbQueryAck //json
	SystemDbExec     //sql
	SystemDbExecAck  //text

)

const (
	StatsHost    Type = iota + 48
	StatsHostAck      //json
	StatsCpu
	StatsCpuAck //json
	StatsMem
	StatsMemAck //json
	StatsDisk
	StatsDiskAck //json
	StatsNet
	StatsNetAck //json
	StatsUser
	StatsUserAck //json
)

const (
	FsList      Type = iota + 64 //path
	FsListAck                    //json
	FsMkDir                      //path
	FsMkDirAck                   //
	FsRemove                     //path
	FsRemoveAck                  //
	FsRename                     //path,path
	FsRenameAck                  //
	FsStats                      //path
	FsStatsAck                   //json

	FsDownload           //path
	FsDownloadContent    //id,data
	FsDownloadContentAck //id
	FsDownloadEnd        //id

	FsUpload           //path
	FsUploadAck        //id
	FsUploadContent    //id,data
	FsUploadContentAck //id
	FsUploadEnd        //id
	FsUploadEndAck     //id

)

const (
	ex2 Type = iota + 96
)
const (
	ex3 Type = iota + 112
)

func no() {

}
