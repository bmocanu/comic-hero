package serve

import (
    "encoding/base64"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
)

const favIconBase64 = `iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAABG2lUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4KPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNS41LjAiPgogPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIi8+CiA8L3JkZjpSREY+CjwveDp4bXBtZXRhPgo8P3hwYWNrZXQgZW5kPSJyIj8+Gkqr6gAAAYNpQ0NQc1JHQiBJRUM2MTk2Ni0yLjEAACiRdZG7SwNBEIe/JGpEIwqmSGERJGoTxQcEbSwSNApqkUQwapOceQh5HHcJEmwF24CCaOOr0L9AW8FaEBRFEBsba0UblXPOCBExs+zst7+dGXZnwRrJKFm9rh+yuYIWCvrdc9F5t/0RCw04cdETU3R1Ojweoaa93Ui02FWvWat23L/WvJTQFbA0Co8qqlYQnhCeWimoJm8KO5V0bEn4WNiryQWFr009XuEnk1MV/jBZi4QCYG0Tdqd+cfwXK2ktKywvx5PNFJWf+5gvcSRys2FZO2V2oBMiiB83k4wRwMcAI+J99DJIn+yokd//nT9DXnIV8SolNJZJkaaAV9SiVE/ImhQ9ISNDyez/377qyaHBSnWHH+ofDOOlC+wb8Fk2jPd9w/g8ANs9nOWq+fk9GH4VvVzVPLvQugYn51UtvgWn6+C6U2Na7FuyybQmk/B8BC1RaL+EpoVKz37OObyFyKp81QVs70C3xLcufgFTEGfds4KRawAAAFFQTFRFymsArlwAw2cA3XUAmVEAuGEA5HkAdD0ARSUAsF0AXzIAjUsAbToAfkIAbDkA/4cAUywA8H8A9IEAAAAAuWIAsV4AxWgAjksAplgA+oQAyUUmrzuymgAAABt0Uk5T//////////////////////////////////8AJzQLNQAAAAlwSFlzAAALEwAACxMBAJqcGAAAAJFJREFUOI2l0zsSgzAMBFAoGBgIcbGd7n9QUBI5tixLCajzzis8+gwU1PAHSGdR8xQAjJxM6ABkABgAKAEaAAVgAkmy+AKJPokNqAZQn6QSkAVIA9wCiwe4VgZb6oMHg90BTwZzB8jk3E+G4EIf6lm0rcYvQI1bDStaGAlfibly4dK+hbf2VBxOvhyn1T1gVwgOJ1dXYvT1DU8AAAAASUVORK5CYII=`

func getFavIcon(response http.ResponseWriter, request *http.Request) {
    log.Info("HTTP Get for favicon: ", request.RequestURI)
    pngData, err := base64.StdEncoding.DecodeString(favIconBase64)
    if err != nil {
        log.Error("Failed to decode base64 PNG content: ", err)
        response.WriteHeader(http.StatusInternalServerError)
        return
    }

    setContentTypeHeader(response, "image/png")
    setContentLengthHeader(response, strconv.Itoa(len(pngData)))
    setCacheControlHeader(response, "public, max-age=2592000")
    _, err = response.Write(pngData)
    if err != nil {
        log.Error("Failed to write CSS content to HTTP response: ", err)
        response.WriteHeader(http.StatusInternalServerError)
    }
}
