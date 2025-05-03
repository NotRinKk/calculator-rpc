// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/rpc"

// 	"github.com/gorilla/mux"
// )

// type Operation struct {
// 	Type    string  `json:"type"`
// 	A       float64 `json:"a,omitempty"`
// 	B       float64 `json:"b,omitempty"`
// 	Value   float64 `json:"value,omitempty"`
// 	Places  int     `json:"places,omitempty"`
// 	Percent float64 `json:"percent,omitempty"`
// 	Power   float64 `json:"power,omitempty"`
// }

// type Request struct {
// 	Operations []Operation `json:"operations"`
// }

// type Response struct {
// 	Results []float64 `json:"results"`
// 	Error   string    `json:"error,omitempty"`
// }

// type Calculator struct{}

// func (c *Calculator) PerformOperation(ctx context.Context, op Operation) (float64, error) {
// 	var result float64
// 	var err error

// 	switch op.Type {
// 	case "addition":
// 		client, err := rpc.DialHTTP("tcp", "addition-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to addition service: %v", err)
// 		}
// 		err = client.Call("AdditionService.Add", []float64{op.A, op.B}, &result)
// 	case "subtraction":
// 		client, err := rpc.DialHTTP("tcp", "subtraction-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to subtraction service: %v", err)
// 		}
// 		err = client.Call("SubtractionService.Subtract", []float64{op.A, op.B}, &result)
// 	case "multiplication":
// 		client, err := rpc.DialHTTP("tcp", "multiplication-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to multiplication service: %v", err)
// 		}
// 		err = client.Call("MultiplicationService.Multiply", []float64{op.A, op.B}, &result)
// 	case "division":
// 		client, err := rpc.DialHTTP("tcp", "division-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to division service: %v", err)
// 		}
// 		err = client.Call("DivisionService.Divide", []float64{op.A, op.B}, &result)
// 	case "square_root":
// 		client, err := rpc.DialHTTP("tcp", "square-root-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to square root service: %v", err)
// 		}
// 		err = client.Call("SquareRootService.Sqrt", op.Value, &result)
// 	case "percentage":
// 		client, err := rpc.DialHTTP("tcp", "percentage-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to percentage service: %v", err)
// 		}
// 		err = client.Call("PercentageService.Percentage", map[string]float64{"value": op.Value, "percent": op.Percent}, &result)
// 	case "rounding":
// 		client, err := rpc.DialHTTP("tcp", "rounding-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to rounding service: %v", err)
// 		}
// 		err = client.Call("RoundingService.Round", map[string]interface{}{"value": op.Value, "places": op.Places}, &result)
// 	case "power":
// 		client, err := rpc.DialHTTP("tcp", "power-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to power service: %v", err)
// 		}
// 		err = client.Call("PowerService.Power", map[string]float64{"base": op.Value, "exponent": op.Power}, &result)
// 	default:
// 		return 0, fmt.Errorf("unknown operation type: %s", op.Type)
// 	}

// 	if err != nil {
// 		return 0, fmt.Errorf("RPC call failed: %v", err)
// 	}

// 	return result, nil
// }

// func (c *Calculator) Calculate(ctx context.Context, req Request) (Response, error) {
// 	var resp Response
// 	var results []float64

// 	for _, op := range req.Operations {
// 		result, err := c.PerformOperation(ctx, op)
// 		if err != nil {
// 			return Response{}, fmt.Errorf("operation failed: %v", err)
// 		}
// 		results = append(results, result)
// 	}

// 	resp.Results = results
// 	return resp, nil
// }

// func main() {
// 	calculator := &Calculator{}
// 	rpc.Register(calculator)
// 	rpc.HandleHTTP()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
// 		var req Request
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		resp, err := calculator.Calculate(r.Context(), req)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	})

// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal("listen error:", err)
// 	}

//		log.Println("Calculator service running on :8080")
//		if err := http.Serve(listener, router); err != nil {
//			log.Fatal("serve error:", err)
//		}
//	}
//////////////////////////////////////////////////////////////////////////////////////////TODO
// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/rpc"

// 	"github.com/gorilla/mux"
// )

// type Operation struct {
// 	Type    string  `json:"type"`
// 	A       float64 `json:"a,omitempty"`
// 	B       float64 `json:"b,omitempty"`
// 	Value   float64 `json:"value,omitempty"`
// 	Places  int     `json:"places,omitempty"`
// 	Percent float64 `json:"percent,omitempty"`
// 	Power   float64 `json:"power,omitempty"`
// }

// type Request struct {
// 	Operations []Operation `json:"operations"`
// }

// type Response struct {
// 	Results []float64 `json:"results"`
// 	Error   string    `json:"error,omitempty"`
// }

// type Calculator struct{}

// func (c *Calculator) PerformOperation(ctx context.Context, op Operation) (float64, error) {
// 	var result float64
// 	var err error

// 	switch op.Type {
// 	case "addition":
// 		client, err := rpc.DialHTTP("tcp", "addition-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to addition service: %v", err)
// 		}
// 		err = client.Call("AdditionService.Add", []float64{op.A, op.B}, &result)
// 	case "subtraction":
// 		client, err := rpc.DialHTTP("tcp", "subtraction-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to subtraction service: %v", err)
// 		}
// 		err = client.Call("SubtractionService.Subtract", []float64{op.A, op.B}, &result)
// 	case "multiplication":
// 		client, err := rpc.DialHTTP("tcp", "multiplication-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to multiplication service: %v", err)
// 		}
// 		err = client.Call("MultiplicationService.Multiply", []float64{op.A, op.B}, &result)
// 	case "division":
// 		client, err := rpc.DialHTTP("tcp", "division-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to division service: %v", err)
// 		}
// 		err = client.Call("DivisionService.Divide", []float64{op.A, op.B}, &result)
// 	case "square_root":
// 		client, err := rpc.DialHTTP("tcp", "square-root-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to square root service: %v", err)
// 		}
// 		err = client.Call("SquareRootService.Sqrt", op.Value, &result)
// 	case "percentage":
// 		client, err := rpc.DialHTTP("tcp", "percentage-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to percentage service: %v", err)
// 		}
// 		err = client.Call("PercentageService.Percentage", map[string]float64{"value": op.Value, "percent": op.Percent}, &result)
// 	case "rounding":
// 		client, err := rpc.DialHTTP("tcp", "rounding-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to rounding service: %v", err)
// 		}
// 		err = client.Call("RoundingService.Round", map[string]interface{}{"value": op.Value, "places": op.Places}, &result)
// 	case "power":
// 		client, err := rpc.DialHTTP("tcp", "power-service:8080")
// 		if err != nil {
// 			return 0, fmt.Errorf("could not connect to power service: %v", err)
// 		}
// 		err = client.Call("PowerService.Power", map[string]float64{"base": op.Value, "exponent": op.Power}, &result)
// 	default:
// 		return 0, fmt.Errorf("unknown operation type: %s", op.Type)
// 	}

// 	if err != nil {
// 		return 0, fmt.Errorf("RPC call failed: %v", err)
// 	}

// 	return result, nil
// }

// func (c *Calculator) Calculate(ctx context.Context, req Request) (Response, error) {
// 	var resp Response
// 	var results []float64

// 	for _, op := range req.Operations {
// 		result, err := c.PerformOperation(ctx, op)
// 		if err != nil {
// 			return Response{}, fmt.Errorf("operation failed: %v", err)
// 		}
// 		results = append(results, result)
// 	}

// 	resp.Results = results
// 	return resp, nil
// }

// func main() {
// 	calculator := &Calculator{}
// 	rpc.Register(calculator)
// 	rpc.HandleHTTP()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
// 		var req Request
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		resp, err := calculator.Calculate(r.Context(), req)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(resp)
// 	})

// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal("listen error:", err)
// 	}

//		log.Println("Calculator service running on :8080")
//		if err := http.Serve(listener, router); err != nil {
//			log.Fatal("serve error:", err)
//		}
//	}
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/gorilla/mux"
)

type Operation struct {
	Type    string  `json:"type"`
	A       float64 `json:"a,omitempty"`
	B       float64 `json:"b,omitempty"`
	Value   float64 `json:"value,omitempty"`
	Places  int     `json:"places,omitempty"`
	Percent float64 `json:"percent,omitempty"`
	Power   float64 `json:"power,omitempty"`
}

type Request struct {
	Operations []Operation `json:"operations"`
}

type Response struct {
	Results     []float64 `json:"results"`     // Результаты каждой операции
	FinalResult float64   `json:"finalResult"` // Общий итоговый результат
	Error       string    `json:"error,omitempty"`
}

type Calculator struct{}

func (c *Calculator) PerformOperation(ctx context.Context, op Operation, currentResult float64) (float64, error) {
	var result float64
	var err error

	switch op.Type {
	case "addition":
		client, err := rpc.DialHTTP("tcp", "addition-service:8080")
		if err != nil {
			return 0, fmt.Errorf("could not connect to addition service: %v", err)
		}
		// Если это первая операция, используем op.A и op.B, иначе currentResult и op.B
		if currentResult == 0 {
			err = client.Call("AdditionService.Add", []float64{op.A, op.B}, &result)
		} else {
			err = client.Call("AdditionService.Add", []float64{currentResult, op.B}, &result)
		}
	case "subtraction":
		client, err := rpc.DialHTTP("tcp", "subtraction-service:8080")
		if err != nil {
			return 0, fmt.Errorf("could not connect to subtraction service: %v", err)
		}
		if currentResult == 0 {
			err = client.Call("SubtractionService.Subtract", []float64{op.A, op.B}, &result)
		} else {
			err = client.Call("SubtractionService.Subtract", []float64{currentResult, op.B}, &result)
		}
	case "multiplication":
		client, err := rpc.DialHTTP("tcp", "multiplication-service:8080")
		if err != nil {
			return 0, fmt.Errorf("could not connect to multiplication service: %v", err)
		}
		if currentResult == 0 {
			err = client.Call("MultiplicationService.Multiply", []float64{op.A, op.B}, &result)
		} else {
			err = client.Call("MultiplicationService.Multiply", []float64{currentResult, op.B}, &result)
		}
	case "division":
		client, err := rpc.DialHTTP("tcp", "division-service:8080")
		if err != nil {
			return 0, fmt.Errorf("could not connect to division service: %v", err)
		}
		if currentResult == 0 {
			err = client.Call("DivisionService.Divide", []float64{op.A, op.B}, &result)
		} else {
			err = client.Call("DivisionService.Divide", []float64{currentResult, op.B}, &result)
		}
	default:
		return 0, fmt.Errorf("unknown operation type: %s", op.Type)
	}

	if err != nil {
		return 0, fmt.Errorf("RPC call failed: %v", err)
	}

	return result, nil
}

func (c *Calculator) Calculate(ctx context.Context, req Request) (Response, error) {
	var resp Response
	var results []float64
	var finalResult float64

	for i, op := range req.Operations {
		result, err := c.PerformOperation(ctx, op, finalResult)
		if err != nil {
			return Response{}, fmt.Errorf("operation failed: %v", err)
		}
		results = append(results, result)
		finalResult = result // Обновляем итоговый результат

		// Для первой операции используем op.A, если это бинарная операция
		if i == 0 && op.Type != "addition" && op.Type != "subtraction" &&
			op.Type != "multiplication" && op.Type != "division" {
			finalResult = result
		}
	}

	resp.Results = results
	resp.FinalResult = finalResult
	return resp, nil
}

func main() {
	calculator := &Calculator{}
	rpc.Register(calculator)
	rpc.HandleHTTP()

	router := mux.NewRouter()
	router.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := calculator.Calculate(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Calculator service running on :8080")
	if err := http.Serve(listener, router); err != nil {
		log.Fatal("serve error:", err)
	}
}
