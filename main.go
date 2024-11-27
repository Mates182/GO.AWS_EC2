package main

import (
	"html/template"
	"net"
	"net/http"
)

type PageData struct {
	LocalIP string
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}
	for _, addr := range addrs {
		// Verifica si la dirección es una IP y no un localhost
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "Unknown"
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Obtén la IP local
	ip := getLocalIP()

	// Define la estructura de datos para la plantilla
	data := PageData{
		LocalIP: ip,
	}

	// Carga y ejecuta la plantilla
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
