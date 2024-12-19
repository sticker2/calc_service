package application

import (
    "calc_service/pkg/calculation"
    "encoding/json"
    "errors"
    "net/http"
    "strconv"
)

type CalcRequest struct {
    Expression string `json:"expression"`
}

type CalcResponse struct {
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CalcRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(CalcResponse{Error: "Internal server error"})
        return
    }

    result, err := calculation.Calc(req.Expression)
    if err != nil {
        if errors.Is(err, calculation.ErrInvalidExpression) {
            w.WriteHeader(http.StatusUnprocessableEntity)
            json.NewEncoder(w).Encode(CalcResponse{Error: "Expression is not valid"})
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(CalcResponse{Error: "Internal server error"})
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(CalcResponse{Result: strconv.FormatFloat(result, 'f', -1, 64)})
}
