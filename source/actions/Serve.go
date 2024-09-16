package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "net/http"
import "os"
import "path/filepath"
import "strconv"
import "strings"

func serveFileRange(response http.ResponseWriter, file string, start int64, end int64) {

	stat, err0 := os.Stat(file)

	if err0 == nil && !stat.IsDir() {

		last_modified := stat.ModTime().Format(http.TimeFormat)

		buffer, err1 := os.ReadFile(file)

		if err1 == nil {

			content_range := "bytes " + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10) + "/" + strconv.FormatInt(stat.Size(), 10)
			partial_buffer := buffer[start:end+1]

			console.Log("Serve " + filepath.Base(file) + " " + content_range)

			header := response.Header()
			header.Set("Content-Encoding", "identity")
			header.Set("Content-Length", strconv.Itoa(len(partial_buffer)))
			header.Set("Content-Range", content_range)
			header.Set("Content-Type", "application/octet-stream")
			header.Set("Date", last_modified)
			header.Set("Last-Modified", last_modified)

			response.WriteHeader(http.StatusPartialContent)
			response.Write(partial_buffer)

		} else {

			console.Warn("Serve " + filepath.Base(file) + " failed")

			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte{})

		}

	} else {

		console.Warn("Serve " + filepath.Base(file) + " failed")

		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte{})

	}

}

func serveFile(response http.ResponseWriter, file string) {

	stat, err0 := os.Stat(file)

	if err0 == nil && !stat.IsDir() {

		last_modified := stat.ModTime().Format(http.TimeFormat)

		buffer, err1 := os.ReadFile(file)

		if err1 == nil {

			console.Log("Serve " + filepath.Base(file))

			header := response.Header()
			header.Set("Accept-Ranges", "bytes")
			header.Set("Content-Encoding", "identity")
			header.Set("Content-Length", strconv.FormatInt(stat.Size(), 10))
			header.Set("Content-Type", "application/octet-stream")
			header.Set("Date", last_modified)
			header.Set("Last-Modified", last_modified)

			response.WriteHeader(http.StatusOK)
			response.Write(buffer)

		} else {

			console.Warn("Serve " + filepath.Base(file) + " failed")

			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte{})

		}

	} else {

		console.Warn("Serve " + filepath.Base(file) + " failed")

		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte{})

	}

}

func serveFileHeader(response http.ResponseWriter, file string) {

	stat, err0 := os.Stat(file)

	if err0 == nil && !stat.IsDir() {

		last_modified := stat.ModTime().Format(http.TimeFormat)

		header := response.Header()
		header.Set("Accept-Ranges", "bytes")
		header.Set("Content-Encoding", "identity")
		header.Set("Content-Length", strconv.FormatInt(stat.Size(), 10))
		header.Set("Content-Type", "application/octet-stream")
		header.Set("Date", last_modified)
		header.Set("Last-Modified", last_modified)

		response.WriteHeader(http.StatusOK)
		response.Write([]byte{})

	} else {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte{})
	}

}

func Serve(sync_folder string, pkgs_folder string) bool {

	console.Group("Serve")

	var result bool

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		path := request.URL.Path
		file := filepath.Base(path)

		if request.Method == "GET" {

			range_header := strings.TrimSpace(request.Header.Get("Range"))

			if strings.HasPrefix(range_header, "bytes=") {

				tmp := strings.Split(range_header[6:], "-")

				if !strings.Contains(range_header, ",") && len(tmp) == 2 {

					// 0 is special end value to serve the rest of the file
					if tmp[1] == "" {
						tmp[1] = "0"
					}

					start, err1 := strconv.ParseInt(tmp[0], 10, 64)
					end,   err2 := strconv.ParseInt(tmp[1], 10, 64)

					if err1 == nil && err2 == nil {

						if pacman.IsDatabaseFilename(file) {
							serveFileRange(response, sync_folder + "/" + file, start, end)
						} else if pacman.IsPackageFilename(file) {
							serveFileRange(response, pkgs_folder + "/" + file, start, end)
						} else {
							response.WriteHeader(http.StatusNotFound)
							response.Write([]byte{})
						}

					} else {
						response.WriteHeader(http.StatusBadRequest)
						response.Write([]byte{})
					}

				} else {
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte{})
				}

			} else {

				if_modified_since := strings.TrimSpace(request.Header.Get("If-Modified-Since"))

				if if_modified_since != "" {

					time, err0 := http.ParseTime(if_modified_since)

					if err0 == nil {

						if pacman.IsDatabaseFilename(file) {

							stat, err1 := os.Stat(sync_folder + "/" + file)

							if err1 == nil {

								if stat.ModTime().After(time) {
									serveFile(response, sync_folder + "/" + file)
								} else {
									response.WriteHeader(http.StatusNotModified)
									response.Write([]byte{})
								}

							} else {
								response.WriteHeader(http.StatusNotFound)
								response.Write([]byte{})
							}

						} else if pacman.IsPackageFilename(file) {

							stat, err1 := os.Stat(pkgs_folder + "/" + file)

							if err1 == nil {

								if stat.ModTime().After(time) {
									serveFile(response, pkgs_folder + "/" + file)
								} else {
									response.WriteHeader(http.StatusNotModified)
									response.Write([]byte{})
								}

							} else {
								response.WriteHeader(http.StatusNotFound)
								response.Write([]byte{})
							}

						} else {

							response.WriteHeader(http.StatusNotFound)
							response.Write([]byte{})

						}

					} else {

						if pacman.IsDatabaseFilename(file) {
							serveFile(response, sync_folder + "/" + file)
						} else if pacman.IsPackageFilename(file) {
							serveFile(response, pkgs_folder + "/" + file)
						} else {
							response.WriteHeader(http.StatusNotFound)
							response.Write([]byte{})
						}

					}

				} else {

					if pacman.IsDatabaseFilename(file) {
						serveFile(response, sync_folder + "/" + file)
					} else if pacman.IsPackageFilename(file) {
						serveFile(response, pkgs_folder + "/" + file)
					} else {
						response.WriteHeader(http.StatusNotFound)
						response.Write([]byte{})
					}

				}

			}

		} else if request.Method == "HEAD" {

			if pacman.IsDatabaseFilename(file) {
				serveFileHeader(response, sync_folder + "/" + file)
			} else if pacman.IsPackageFilename(file) {
				serveFileHeader(response, pkgs_folder + "/" + file)
			} else {
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte{})
			}

		} else {
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte{})
		}

	})

	err := http.ListenAndServe(":15678", nil)

	if err == nil {
		result = true
	}

	console.GroupEndResult(result, "Serve")

	return result

}
