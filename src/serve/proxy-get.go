package serve

import (
    "comic-hero/store"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "strconv"
)

func getProxyImage(response http.ResponseWriter, request *http.Request) {
    var params = mux.Vars(request)
    var reqIdStr = params["id"]
    var reqHashStr = params["hash"]
    log.Info("HTTP Get for proxy image: id=[", reqIdStr, "], hash=[", reqHashStr, "]: ", request.RequestURI)

    if reqIdStr == "" {
        log.Warn("HTTP request for proxy image, without any comic ID")
        response.WriteHeader(http.StatusNotFound)
        return
    }

    var reqId, err = strconv.Atoi(reqIdStr)
    if err != nil {
        log.Warn("HTTP request for proxy image, with invalid comic ID")
        response.WriteHeader(http.StatusNotFound)
        return
    }

    if reqHashStr == "" {
        log.Warn("HTTP request for proxy image, without any issue hash")
        response.WriteHeader(http.StatusNotFound)
        return
    }

    var issueLink = store.GetIssueLinkByComicIdAndHash(reqId, reqHashStr)
    if issueLink == nil {
        log.Warn("HTTP request for proxy image, but cannot find the comic with that ID or the issue with that hash")
        response.WriteHeader(http.StatusNotFound)
        return
    }

    log.Info("Image proxy: Retrieving: ", issueLink.Issue.ImageUrl)
    httpResp, err := http.Get(issueLink.Issue.ImageUrl)
    if err != nil {
        log.Warn("Image proxy: failed to retrieve the content of the issue image: ", err)
        response.WriteHeader(http.StatusNotFound)
        return
    }

    defer httpResp.Body.Close()

    if httpResp.StatusCode != 200 {
        log.Warn("Image proxy: Got bad status code: ", httpResp.StatusCode)
        response.WriteHeader(http.StatusNotFound)
        return
    }

    var headerContentType = httpResp.Header.Get("Content-Type");
    var headerContentLength = httpResp.Header.Get("Content-Length");
    var headerAcceptEncoding = httpResp.Header.Get("Accept-Encoding");
    htmlContent, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
        log.Warn("Image proxy: failed to read the issue image")
        response.WriteHeader(http.StatusNotFound)
        return
    }

    setContentTypeHeader(response, headerContentType)
    setContentLengthHeader(response, headerContentLength)
    if headerAcceptEncoding != "" {
        response.Header().Set("Accept-Encoding", headerAcceptEncoding)
    }

    response.WriteHeader(http.StatusOK)
    _, err = response.Write(htmlContent)
    if err != nil {
        log.Warn("Image proxy: Failed to write issue image to HTTP response: ", err)
        return
    }
}
