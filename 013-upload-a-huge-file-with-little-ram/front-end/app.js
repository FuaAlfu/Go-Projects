let hugeFileMode = "";
if (file.size > largeFileCutoffSize){
       //Filesize over cutoff line. Use huge file mode
       hugeFileMode = "&hugefile=true";
}

let socket = new WebSocket(
    protocol +
     window.location.hostname + ":" + port +
      "/system/file_system/lowmemUpload?filename=" + 
      encodeURIComponent(filename) + "&path=" + 
      encodeURIComponent(uploadDir) + hugeFileMode
      );