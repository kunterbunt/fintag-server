package controller

import (
    "net/http"
    "encoding/json"
    "github.com/kunterbunt/fintag-server/model"
    "time"
    // "strings"
    // "fmt"
)

/**
* Cateogry Controller.
*/
type CoordinatesController struct {
    ControllerBase
}

func NewCoordinatesController(model model.Model) *CoordinatesController {
    return &CoordinatesController{ControllerBase{model}}
}

/**
* For HTTP request parameters.
*/
type httpParametersCategory struct {
    Id string
    Long, Lat int
    Timestamp time.Time
}

func (this *CoordinatesController) parseHttpParameters(request *http.Request) (parameters httpParametersCategory, err error) {
  // Decode JSON parameters from HTTP body.
  decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&parameters)
	return parameters, err
}

func (this *CoordinatesController) Get(writer http.ResponseWriter, request *http.Request) {
    // Get parameter values from URL.
    values := request.URL.Query()
    // These two variables will propagate to the end.
    var jsonResult []byte
    var err error
    // Return all categories.
    if (len(values) == 0) {
        var coordinates []*model.Coordinates
        coordinates, err = this.model.GetAllCoordinates()
        jsonResult, err = json.Marshal(coordinates)
    // Return specific category.
    } else {
        id := values.Get("id")
        if (len(id) != 0) {
            var coordinates *model.Coordinates
            coordinates, err = this.model.GetCoordinates(id)
            jsonResult, err = json.Marshal(coordinates)
        } else {
            http.Error(writer, "Invalid request - 'id' field missing.", http.StatusBadRequest)
        }
    }
    // Finally send out the result.
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    } else {
        writer.Header().Set("Content-Type", "application/json")
        writer.Header().Set("Access-Control-Allow-Origin", "*")
        writer.Write(jsonResult)
    }
}

func (this *CoordinatesController) Post(writer http.ResponseWriter, request *http.Request) {
    parameters, err := this.parseHttpParameters(request)
    if err != nil {
        http.Error(writer, "Error parsing HTTP request body: " + err.Error(), http.StatusBadRequest)
        return
    }
    // Check if all needed parameters are present.
    if parameters.Long == 0 {
        http.Error(writer, "Invalid request: 'long' missing.", http.StatusBadRequest)
        return
    }
    if parameters.Lat == 0 {
        http.Error(writer, "Invalid request: 'lat' missing.", http.StatusBadRequest)
        return
    }
    coordinates := model.Coordinates{"", parameters.Long, parameters.Lat, time.Now()}
    // Send to model.
    err = this.model.AddCoordinates(&coordinates)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    } else {
        writer.Header().Set("Access-Control-Allow-Origin", "*")
        writer.Write([]byte("Soordinates successfully saved."))
    }
}

// func (this *CoordinatesController) Put(writer http.ResponseWriter, request *http.Request) {
//     parameters, err := this.parseHttpParameters(request)
//     if err != nil {
//         http.Error(writer, "Error parsing HTTP request body: " + err.Error(), http.StatusBadRequest)
//         return
//     }
//     // Check if identification email is provided.
//     if len(parameters.Id) == 0 {
//         http.Error(writer, "Invalid request: 'id' missing.", http.StatusBadRequest)
//         return
//     }
//     // Get currently saved category from database.
//     category, err := this.model.GetCategory(parameters.Id)
//     if err != nil {
//         http.Error(writer, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     // Update values if they're provided.
//     updatedString := ""
//     if len(parameters.Titles) != 0 {
//         category.Titles = parameters.Titles
//         updatedString += " titles"
//     }
//     if len(parameters.Parent) != 0 {
//         category.Parent = parameters.Parent
//         updatedString += " parent"
//     }
//     if updatedString == "" {
//         http.Error(writer, "Invalid request: nothing to update.", http.StatusBadRequest)
//         return
//     }
//     // Make sure all provided languages are known.
//     for lang, title := range category.Titles {
//         // Enforce uppercase.
//         delete(category.Titles, lang)
//         category.Titles[strings.ToUpper(lang)] = title
//         _, err := this.model.GetLanguage(strings.ToUpper(lang))
//         if err != nil {
//             http.Error(writer, "Error checking for provided language '" + strings.ToUpper(lang) + "': " + err.Error(), http.StatusInternalServerError)
//             return
//         }
//     }
//     err = this.model.UpdateCategory(category)
//     if err != nil {
//         http.Error(writer, err.Error(), http.StatusInternalServerError)
//     } else {
//         writer.Header().Set("Access-Control-Allow-Origin", "*")
//         writer.Write([]byte("Successfully updated:" + updatedString))
//     }
// }
//
// func (this *CoordinatesController) Delete(writer http.ResponseWriter, request *http.Request) {
//     parameters, err := this.parseHttpParameters(request)
//     if err != nil {
//         http.Error(writer, "Error parsing HTTP request body: " + err.Error(), http.StatusBadRequest)
//         return
//     }
//     // Check for ID field.
//     if len(parameters.Id) == 0 {
//         http.Error(writer, "Invalid request: 'id' missing.", http.StatusBadRequest)
//         return
//     }
//     // Tell model to remove category.
//     err = this.model.RemoveCategory(parameters.Id)
//     if err != nil {
//         http.Error(writer, err.Error(), http.StatusInternalServerError)
//     } else {
//         writer.Header().Set("Access-Control-Allow-Origin", "*")
//         writer.Write([]byte("Category successfully removed."))
//     }
// }
