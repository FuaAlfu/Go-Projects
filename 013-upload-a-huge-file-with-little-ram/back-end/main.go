package main

import(
	"fmt"
	"log"
	"time"
	"strings"
)

func main() {
	targetUploadLocation := filepath.Join(uploadPath, filename)
if !fs.FileExists(uploadPath) {
    os.MkdirAll(uploadPath, 0755)
}

//Generate an UUID for this upload
uploadUUID := uuid.NewV4().String()
uploadFolder := filepath.Join(*tmp_directory, "uploads", uploadUUID)
if isHugeFile {
    //Upload to the same directory as the target location.
    uploadFolder = filepath.Join(uploadPath, ".metadata/.upload", uploadUUID)
}
os.MkdirAll(uploadFolder, 0700)

//Start websocket connection
var upgrader = websocket.Upgrader{}
upgrader.CheckOrigin = func(r *http.Request) bool { return true }
c, err := upgrader.Upgrade(w, r, nil)
if err != nil {
    log.Println("Failed to upgrade websocket connection: ", err.Error())
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 WebSocket upgrade failed"))
    return
}
defer c.Close()

//Handle WebSocket upload
blockCounter := 0
chunkName := []string{}
lastChunkArrivalTime := time.Now().Unix()

//Setup a timeout listener, check if connection still active every 1 minute
ticker := time.NewTicker(60 * time.Second)
done := make(chan bool)
go func() {
    for {
        select {
        case <-done:
            return
        case <-ticker.C:
            if time.Now().Unix()-lastChunkArrivalTime > 300 {
                //Already 5 minutes without new data arraival. Stop connection
                log.Println("Upload WebSocket connection timeout. Disconnecting.")
                c.WriteControl(8, []byte{}, time.Now().Add(time.Second))
                time.Sleep(1 * time.Second)
                c.Close()
                return
            }
        }
    }
}()

totalFileSize := int64(0)
for {
    mt, message, err := c.ReadMessage()
    if err != nil {
        //Connection closed by client. Clear the tmp folder and exit
        log.Println("Upload terminated by client. Cleaning tmp folder.")
        //Clear the tmp folder
        time.Sleep(1 * time.Second)
        os.RemoveAll(uploadFolder)
        return
    }
    //The mt should be 2 = binary for file upload and 1 for control syntax
    if mt == 1 {
        msg := strings.TrimSpace(string(message))
        if msg == "done" {
            //Start the merging process
            break
        } else {
            //Unknown operations

        }
    } else if mt == 2 {
        //File block. Save it to tmp folder
        chunkFilepath := filepath.Join(uploadFolder, "upld_"+strconv.Itoa(blockCounter))
        chunkName = append(chunkName, chunkFilepath)
        writeErr := ioutil.WriteFile(chunkFilepath, message, 0700)

        if writeErr != nil {
            //Unable to write block. Is the tmp folder fulled?
            log.Println("[Upload] Upload chunk write failed: " + err.Error())
            c.WriteMessage(1, []byte(`{\"error\":\"Write file chunk to disk failed\"}`))

            //Close the connection
            c.WriteControl(8, []byte{}, time.Now().Add(time.Second))
            time.Sleep(1 * time.Second)
            c.Close()

            //Clear the tmp files
            os.RemoveAll(uploadFolder)
            return
        }

        //Update the last upload chunk time
        lastChunkArrivalTime = time.Now().Unix()

        //Check if the file size is too big
        totalFileSize += fs.GetFileSize(chunkFilepath)
        if totalFileSize > max_upload_size {
            //File too big
            c.WriteMessage(1, []byte(`{\"error\":\"File size too large\"}`))

            //Close the connection
            c.WriteControl(8, []byte{}, time.Now().Add(time.Second))
            time.Sleep(1 * time.Second)
            c.Close()

            //Clear the tmp files
            os.RemoveAll(uploadFolder)
            return
        }
        blockCounter++

        //Request client to send the next chunk
        c.WriteMessage(1, []byte("next"))

    }
}

//Try to decode the location if possible
decodedUploadLocation, err := url.QueryUnescape(targetUploadLocation)
if err != nil {
    decodedUploadLocation = targetUploadLocation
}

//Do not allow % sign in filename. Replace all with underscore
decodedUploadLocation = strings.ReplaceAll(decodedUploadLocation, "%", "_")

//Merge the file
out, err := os.OpenFile(decodedUploadLocation, os.O_CREATE|os.O_WRONLY, 0755)
if err != nil {
    log.Println("Failed to open file:", err)
    c.WriteMessage(1, []byte(`{\"error\":\"Failed to open destination file\"}`))
    c.WriteControl(8, []byte{}, time.Now().Add(time.Second))
    time.Sleep(1 * time.Second)
    c.Close()
    return
}

for _, filesrc := range chunkName {
    srcChunkReader, err := os.Open(filesrc)
    if err != nil {
        log.Println("Failed to open Source Chunk", filesrc, " with error ", err.Error())
        c.WriteMessage(1, []byte(`{\"error\":\"Failed to open Source Chunk\"}`))
        return
    }
    io.Copy(out, srcChunkReader)
    srcChunkReader.Close()

    //Delete file immediately to save space
    os.Remove(filesrc)
}

out.Close()

//Return complete signal
c.WriteMessage(1, []byte("OK"))

//Stop the timeout listner
done <- true

//Clear the tmp folder
time.Sleep(300 * time.Millisecond)
err = os.RemoveAll(uploadFolder)
if err != nil {
    log.Println(err)
}

//Close WebSocket connection after finished
c.WriteControl(8, []byte{}, time.Now().Add(time.Second))
time.Sleep(300 * time.Second)
c.Close()

}