package utils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"line/interview/models"
	"net/http"
	"runtime"
	"time"
)

// Function to find number split in half data depend on number of core
func RoundInt(x int) int {
	cpuNum := runtime.NumCPU()
	if x%cpuNum == 0 {
		return x / cpuNum
	} else {
		return (x / cpuNum) + 1
	}
}

// Function to split data to array
func SplitData(data []string, index int) [][]string {
	var result [][]string
	for i := 0; i < len(data); i += index {
		result = append(result, data[i:i+index])
	}
	return result
}

// Function to update file & return data
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Decalre a new variable to store the result
	countTime := time.Now()
	// result := make(chan bool)
	sumData := []string{}

	//File Upload Endpoint
	file, handler, err := r.FormFile("File")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error Retrieving the File"))
		return
	}
	defer file.Close()

	//File Content
	// Max file 5MB  1 MB = 1048576 Byte
	if handler.Size > 5*1048576 {
		fmt.Fprintf(w, "File size exceeds the maximum size \n")
		return
	}

	//config file type
	fileTypeSupport := []string{"text/csv"}
	for _, v := range fileTypeSupport {
		if v != handler.Header.Get("Content-Type") {
			fmt.Fprintf(w, "File Type Not Supported\n")
			return
		}
	}

	// Read the file to byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// CSV Read the byte array
	reader := csv.NewReader(bytes.NewBuffer(fileBytes))
	for {
		line, err := reader.Read()
		if err != nil {
			// End of file is expected at this point
			if err == io.EOF {
				fmt.Println(err)
				break
			} else {
				fmt.Fprintf(w, "Error Reading CSV File\n")
				return
			}
		}

		sumData = append(sumData, line[0])
	}
	// round index
	total := IsUpOrDown(sumData)
	dataReturn := models.ReturnData{
		Total:   total.CountSite,
		Sucess:  total.CountSuccess,
		Fail:    total.CountFail,
		TimeUse: time.Since(countTime).Seconds(),
	}
	// Set Header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")
	w.Header().Set("Content-Type", "application/json")
	// Convert the data to JSON
	json.NewEncoder(w).Encode(dataReturn)

}

// Function to upload file multicore
func UploadFileMulticore(w http.ResponseWriter, r *http.Request) {

	// Decalre a new variable to store the result
	countTime := time.Now()
	// result := make(chan bool)
	sumData := []string{}

	//File Upload Endpoint
	file, handler, err := r.FormFile("File")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error Retrieving the File"))
		return
	}
	defer file.Close()

	//File Content
	// Max file 5MB  1 MB = 1048576 Byte
	if handler.Size > 5*1048576 {
		fmt.Fprintf(w, "File size exceeds the maximum size \n")
		return
	}

	//config file type
	fileTypeSupport := []string{"text/csv"}
	for _, v := range fileTypeSupport {
		if v != handler.Header.Get("Content-Type") {
			fmt.Fprintf(w, "File Type Not Supported\n")
			return
		}
	}

	// Read the file to byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// CSV Read the byte array
	reader := csv.NewReader(bytes.NewBuffer(fileBytes))
	for {
		line, err := reader.Read()
		if err != nil {
			// End of file is expected at this point
			if err == io.EOF {
				break
			} else {
				fmt.Fprintf(w, "Error Reading CSV File\n")
				return
			}
		}
		// Collect data to Array
		sumData = append(sumData, line[0])
	}
	// round index
	index := RoundInt(len(sumData))
	//split data
	dataSplit := SplitData(sumData, index)
	//multi core
	runtime.GOMAXPROCS(runtime.NumCPU())
	//create channel
	result := make(chan models.Total)
	//create goroutine
	for _, v := range dataSplit {
		go func(data []string) {
			total := IsUpOrDown(data)
			result <- total
		}(v)
	}
	//wait for goroutine and summarize data
	var total models.Total
	for i := 0; i < len(dataSplit); i++ {
		if i == 0 {
			total = <-result
			continue
		}
		roundOther := <-result
		total.CountSuccess += roundOther.CountSuccess
		total.CountFail += roundOther.CountFail
		total.CountSite += roundOther.CountSite
	}
	// set data return
	dataReturn := models.ReturnData{
		Total:   total.CountSite,
		Sucess:  total.CountSuccess,
		Fail:    total.CountFail,
		TimeUse: time.Since(countTime).Seconds(),
	}
	// Set header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")
	w.Header().Set("Content-Type", "application/json")
	// Encode data to json
	json.NewEncoder(w).Encode(dataReturn)

}

// Function to check if the Website is up or down
func IsUpOrDown(url []string) models.Total {
	// Declare a new variable to store the result
	var result models.Total

	// Create a new client
	client := &http.Client{}

	for _, index := range url {
		// Check if the url is valid
		if index == "" {
			continue
		}
		// Create a new request
		req, err := http.NewRequest("GET", index, nil)
		if err != nil {
			fmt.Println(err)
		}

		// Make the request
		resp, err := client.Do(req)
		// Check if the error is valid
		if err != nil {
			result.CountFail++
			result.CountSite++
			continue
		}

		// Check the status code
		if resp.StatusCode == 200 {
			result.CountSuccess++
		}
		result.CountSite++

	}
	return result
}
