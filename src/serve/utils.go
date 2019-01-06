package serve

import (
    "net/http"
    "strings"
)

func setNoCachingHeaders(response http.ResponseWriter) {
    response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    response.Header().Set("Pragma", "no-cache")
    response.Header().Set("Expires", "0")
}

func setCacheControlHeader(response http.ResponseWriter, value string) {
    response.Header().Set("Cache-Control", value)
}

func setContentTypeHeader(response http.ResponseWriter, value string) {
    response.Header().Set("Content-Type", value)
}

func setContentLengthHeader(response http.ResponseWriter, value string) {
    response.Header().Set("Content-Length", value)
}

func urlConcat(part1 string, part2 string) string {
    var trimmedPart1 = strings.TrimSpace(part1)
    var trimmedPart2 = strings.TrimSpace(part2)
    if trimmedPart1 == "" {
        return trimmedPart2
    } else if trimmedPart2 == "" {
        return trimmedPart1
    } else if !strings.HasSuffix(trimmedPart1, "/") && !strings.HasPrefix(trimmedPart2, "/") {
        return trimmedPart1 + "/" + trimmedPart2
    } else if strings.HasSuffix(trimmedPart1, "/") && strings.HasPrefix(trimmedPart2, "/") {
        return trimmedPart1 + trimmedPart2[1:]
    } else {
        return trimmedPart1 + trimmedPart2
    }
}
