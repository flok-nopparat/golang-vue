
# Line TH Interview

Implement a website that can upload the list of websites as a CSV file.


## Acknowledgements

 - [Vue CLI](https://cli.vuejs.org/guide/)
 - [Golang Routine and Channel](https://www.somkiat.cc/golang-goroutine-and-channel)
 - [Unit Testing on Golang](https://medium.com/the-existing/unit-testing-in-golang-2077ad8ae215)
 - [Creating A Simple Web Server With Golang](https://tutorialedge.net/golang/creating-simple-web-server-with-golang/)

## Tech Stack

**Client:** VueJS 2

**Server:** Golang 1.18

**Run and Deploy:** Docker , Docker-Compose


## API Reference

#### Upload File With Multicore

```http
  POST /uploadMulticore
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `File` | `file` | **Required**. Your CSV File |

Return { \
    total-->(int) \
    success-->(int) \
    fail-->(int) \
    time_use-->(float64)  
} 
#### Upload File

```http
  POST /upload
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `File` | `file` | **Required**. Your CSV File |

Return { \
    total-->(int) \
    success-->(int) \
    fail-->(int) \
    time_use-->(float64)  
} 


## Installation
1.Run Code

Select one choice.

1.1 By Docker-Compose (recommend)

```bash
  docker-compose up -d --build
```

1.2. By DockerFile (Optional)

```bash
  docker build -t line-interview .
  docker run -p 8080:8080 -it --rm line-interview
```
    
1.3. Run a separate system (Optional)

 - With yarn
```bash
  cd ./client/line-interview/ && yarn install
  yarn build
  cd ../.. 
  go build -v -o /docker-golang
  /docker-golang
```

2.Open http://localhost:8080/
## Running Tests

To run tests, run the following command

```bash
  go test
```


## Screenshots

![App Screenshot](https://www.img.in.th/images/db1cd9de5f4744d783e2cc2306bd5232.png)

#
![App Screenshot2](https://www.img.in.th/images/e20b9f1fe5736f7f8f9befd762df8458.png)


## Authors

- [@nopparat-yuyen](https://github.com/nopparat-yuyen)


![Logo](https://seeklogo.com/images/G/go-logo-046185B647-seeklogo.com.png)
![Logo](https://miro.medium.com/max/666/1*yGrOUQyqX3MBekvP5d-pCA.png)

