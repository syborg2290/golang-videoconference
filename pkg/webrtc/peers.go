package webrtc

type Room struct{
	Peers *Peers
	Hub *chat.Hub
}
type Peers struct{
	ListLock sync.RWMutex
	Connections []peerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame(){

}