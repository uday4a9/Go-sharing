package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

var sentMsg = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "message",
		Help: "The total number of sent messages",
	},
	[]string{"user"},
)

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(sentMsg)

	http.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))

	var chatConn *websocket.Conn
	http.HandleFunc("/chatws", func(w http.ResponseWriter, r *http.Request) {
		var err error
		chatConn, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Chat connection established")
		msg := struct {
			Username string `json:"username"`
			Message  string `json:"message"`
		}{}
		for {
			err = chatConn.ReadJSON(&msg)
			if err != nil {
				return
			}

			// fmt.Printf("%s sent: %s\n", chatConn.RemoteAddr(), string(msg))
			// fmt.Printf("msgType: %d\n", msgType)

		}
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if chatConn == nil {
			fmt.Println("Chat connection not established")
			http.Error(w, "Chat connection not established", http.StatusServiceUnavailable)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		msg := struct {
			Username string `json:"username"`
			Message  string `json:"message"`
		}{}
		for {
			err := conn.ReadJSON(&msg)
			if err != nil {
				return
			}

			fmt.Printf("msg: %v\n", msg)
			sentMsg.WithLabelValues(msg.Username).Set(float64(len(msg.Message)))
			// sentMsg.WithLabelValues(conn.RemoteAddr().String()).Set(float64(len(msg)))

			// msgString := strings.ToUpper(string(msg))

			// fmt.Println("msgString: ", msgString, msgType)

			// if err = chatConn.WriteMessage(msg); err != nil {
			// 	return
			// }

			// FRom here i need to send to other wesocket client \chat

			if err = chatConn.WriteJSON(msg); err != nil {

				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "chat.html")
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
