package webrtc

import(
	"log"
	"sync"
	"github.com/gofiber/websocket"
	"github.com/pion/webrtc/v3"
)

func RoomConn(c *websocket.Conn, p *Peers){
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil{
		log.Print(err)
		return
	}

	newPeer := PeerconnectionState{
		PeerConnection: peerConnection,
		webSocket: &ThreadSafeWriter{
		Conn: c,
		Mutex: sync.Mutex{}	
		},
	}

}

