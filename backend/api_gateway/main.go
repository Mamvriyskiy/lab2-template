package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	DTO "github.com/lapayka/rsoi-2/Common"
	http_utils "github.com/lapayka/rsoi-2/Common/HTTP_Utils"
	"github.com/lapayka/rsoi-2/Common/Logger"
	FS_structs "github.com/lapayka/rsoi-2/flight_service/Structs"
	TS_structs "github.com/lapayka/rsoi-2/ticket_service/structs"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/manage/health", http_utils.HealthCkeck).Methods("GET")
	router.HandleFunc("/api/v1/flights", flight_proxy).Methods("GET")
	router.HandleFunc("/api/v1/tickets/{ticketUid}", ticket_proxy).Methods("GET")
	router.HandleFunc("/api/v1/tickets", ticket_proxy).Methods("GET")
	router.HandleFunc("/api/v1/tickets/{ticketUid}", ticket_proxy).Methods("DELETE")
	// router.HandleFunc("/api/v1/tickets", ticket_proxy).Methods("Post")
	router.HandleFunc("/api/v1/privilege", bonus_proxy).Methods("GET")
	router.HandleFunc("/api/v1/me", meHandler).Methods("GET")

	router.HandleFunc("/api/v1/tickets", buy_ticket).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		Logger.GetLogger().Print(err)
	}
}

func GetDefaultClient() *http.Client {
	client := http.DefaultClient
	client.Timeout = 2 * time.Second

	return client
}

func check_flght_number(flight_number string) bool {
	req, _ := http.NewRequest("GET", "http://flight_service:8060/api/v1/flights", nil)
	r, err := GetDefaultClient().Do(req)

	if err != nil {
		Logger.GetLogger().Print(err)
		return false
	}

	flights := FS_structs.Flights{}
	err = http_utils.ReadSerializableFromResponse(r, &flights)

	if err != nil {
		Logger.GetLogger().Print(err)
		return false
	}

	for i := range flights {
		if flights[i].FlightNumber == flight_number {
			return true
		}
	}

	return false
}

func buy_ticket_in_ticket_service(username string, buy_ticket_info DTO.BuyTicketDTO) (TS_structs.Ticket, error) {
	body, _ := json.Marshal(buy_ticket_info)
	reader := strings.NewReader(string(body))

	req, err := http.NewRequest("POST", "http://ticket_service:8070/api/v1/tickets", reader)
	req.Header.Add("X-User-Name", username)

	if err != nil {
		Logger.GetLogger().Print(err)
		return TS_structs.Ticket{}, err
	}

	var r *http.Response
	r, err = GetDefaultClient().Do(req)

	if err != nil {
		Logger.GetLogger().Print(err)
		return TS_structs.Ticket{}, err
	}

	ticket := TS_structs.Ticket{}
	err = http_utils.ReadSerializableFromResponse(r, &ticket)

	if err != nil {
		Logger.GetLogger().Print(err)
		return TS_structs.Ticket{}, err
	}

	return ticket, nil
}

func buy_ticket_in_privilege_service(username string, buy_ticket_info DTO.BuyTicketDTO) error {
	body, _ := json.Marshal(buy_ticket_info)
	reader := strings.NewReader(string(body))

	req, err := http.NewRequest("POST", "http://privilege-service:8050/api/v1/tickets", reader)
	req.Header.Add("X-User-Name", username)

	if err != nil {
		Logger.GetLogger().Print(err)
		return err
	}

	var r *http.Response
	r, err = GetDefaultClient().Do(req)

	if err != nil {
		Logger.GetLogger().Print(err)
		return err
	}

	if r.StatusCode == http.StatusCreated {
		return nil
	}

	return fmt.Errorf("status code was: %d\n", r.StatusCode)
}

func buy_ticket(w http.ResponseWriter, r *http.Request) {
	print("aaaaaaaaaaa")
	username := r.Header.Get("X-User-Name")

	var buyTicketInfo DTO.BuyTicketDTO
	if err := http_utils.ReadSerializable(r, &buyTicketInfo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if !check_flght_number(buyTicketInfo.FlightNumber) {
		http.Error(w, "flight not found", http.StatusNotFound)
		return
	}

	ticket, err := buy_ticket_in_ticket_service(username, buyTicketInfo)
	if err != nil {
		http.Error(w, "failed to buy ticket", http.StatusInternalServerError)
		return
	}
	buyTicketInfo.TicketUid = ticket.TicketUid

	err = buy_ticket_in_privilege_service(username, buyTicketInfo)
	if err != nil {
		http.Error(w, "failed to update privilege", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	bonusReq, _ := http.NewRequest("GET", "http://privilege-service:8050/api/v1/privilege", nil)
	bonusReq.Header = r.Header

	bonusResp, err := client.Do(bonusReq)
	if err != nil {
		http.Error(w, "Error fetching privilege", http.StatusInternalServerError)
		return
	}
	defer bonusResp.Body.Close()

	if bonusResp.StatusCode != http.StatusOK {
		http.Error(w, "failed to get privilege", bonusResp.StatusCode)
		return
	}

	var privilege struct {
		Balance int    `json:"balance"`
		Status  string `json:"status"`
	}
	if err := json.NewDecoder(bonusResp.Body).Decode(&privilege); err != nil {
		http.Error(w, "invalid privilege response", http.StatusInternalServerError)
		return
	}

	flight, err := getFlightByNumber(buyTicketInfo.FlightNumber)
	if err != nil {
		http.Error(w, "flight not found", http.StatusNotFound)
		return
	}

	// loc, _ := time.LoadLocation("Europe/Moscow")
	print("bbbbbbbbb")
	response := map[string]interface{}{
		"ticketUid":     ticket.TicketUid,
		"flightNumber":  flight.FlightNumber,
		"fromAirport":   flight.FromAirport,
		"toAirport":     flight.ToAirport,
		"date":          flight.Date,
		"price":         flight.Price,
		"paidByMoney":   flight.Price,
		"paidByBonuses": 0,
		"status":        "PAID",
		"privilege": map[string]interface{}{
			"balance": privilege.Balance,
			"status":  privilege.Status,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func echo_request(w http.ResponseWriter, r *http.Request, service_url string) {
	Logger.GetLogger().Printf("Proxying to: %s%s", service_url, r.URL.String())
	
	req, err := http.NewRequestWithContext(r.Context(), r.Method, service_url+r.URL.String(), r.Body)
	if err != nil {
		Logger.GetLogger().Printf("Error creating request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	req.Header = r.Header.Clone()
	response, err := GetDefaultClient().Do(req)

	if err != nil {
		Logger.GetLogger().Print(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		defer response.Body.Close()
		
		// Копируем все заголовки КРОМЕ Content-Type
		for key, values := range response.Header {
			if key != "Content-Type" { // Не копируем исходный Content-Type
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
		}

		// Форсируем application/json для API endpoints
		w.Header().Set("Content-Type", "application/json")
		
		w.WriteHeader(response.StatusCode)
		io.Copy(w, response.Body)
	}
}

func bonus_proxy(w http.ResponseWriter, r *http.Request) {
    echo_request(w, r, "http://privilege-service:8050")
}

func flight_proxy(w http.ResponseWriter, r *http.Request) {
    echo_request(w, r, "http://flight_service:8060")
}

func ticket_proxy(w http.ResponseWriter, r *http.Request) {
    echo_request(w, r, "http://ticket_service:8070")
}

func meHandler(w http.ResponseWriter, r *http.Request) {
    client := &http.Client{}

    // --- Получаем билеты ---
    ticketsReq, _ := http.NewRequest("GET", "http://ticket_service:8070/api/v1/tickets", nil)
    ticketsReq.Header = r.Header // проксируем X-User-Name и другие заголовки
    ticketsResp, err := client.Do(ticketsReq)
    if err != nil {
        http.Error(w, "Error fetching tickets", http.StatusInternalServerError)
        return
    }
    defer ticketsResp.Body.Close()

    ticketsBody, _ := io.ReadAll(ticketsResp.Body)
    var tickets []map[string]interface{}
    if err := json.Unmarshal(ticketsBody, &tickets); err != nil {
        http.Error(w, "Error parsing tickets response", http.StatusInternalServerError)
        return
    }

    // --- Получаем привилегии ---
    bonusReq, _ := http.NewRequest("GET", "http://privilege-service:8050/api/v1/privilege", nil)
    bonusReq.Header = r.Header
    bonusResp, err := client.Do(bonusReq)
    if err != nil {
        http.Error(w, "Error fetching privilege", http.StatusInternalServerError)
        return
    }
    defer bonusResp.Body.Close()

    bonusBody, _ := io.ReadAll(bonusResp.Body)
    var privilege map[string]interface{}
    if err := json.Unmarshal(bonusBody, &privilege); err != nil {
        http.Error(w, "Error parsing privilege response", http.StatusInternalServerError)
        return
    }

    // --- Формируем объединённый ответ ---
    response := map[string]interface{}{
        "tickets":   tickets,
        "privilege": map[string]interface{}{"balance": privilege["balance"], "status": privilege["status"]},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

type FlightShort struct {
	FlightNumber string `json:"flightNumber"`
	FromAirport  string `json:"fromAirport"`
	ToAirport    string `json:"toAirport"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
}

func getFlightByNumber(flightNumber string) (FlightShort, error) {
	// Формируем запрос к flight_service
	req, _ := http.NewRequest("GET", "http://flight_service:8060/api/v1/flights", nil)
	q := req.URL.Query()
	q.Add("flightNumber", flightNumber)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return FlightShort{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return FlightShort{}, fmt.Errorf("flight not found")
	}

	// Структура, соответствующая JSON ответа flight_service
	var result struct {
		Items []FlightShort `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return FlightShort{}, err
	}

	if len(result.Items) == 0 {
		return FlightShort{}, fmt.Errorf("flight not found")
	}

	return result.Items[0], nil
}

