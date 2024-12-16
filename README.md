# RooLink SDK

RooLink SDK is a Golang library designed for seamless interaction with the RooLink API. It provides utilities for API request limits, parsing scripts, generating sensor data, and more.

## Features

- Fetch API request limits
- Parse script data
- Generate sensor data for validation
- Create SBSD body
- Generate pixel data
- Solve sec-cpt challenges

## Usage

### Installing
```shell
go get github.com/RooLinkIO/roolink-sdk-go/roolink
```

### Initialize RooLink SDK Session
```go
ctx := context.Background()
API_KEY = "your_api_key"
protectedURL = "https://protected.example.net"
userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

session := roolink.NewSession(API_KEY, protectedURL, userAgent)
```

### Fetch API Request Limit
```go
limit, err := session.RequestLimit(ctx)
if err != nil {
    fmt.Println("Error fetching request limit:", err)
    panic(err)
}
fmt.Println("Request Limit:", limit.Requests)
```

### Parse Script Data
```go
scriptBody := []byte("function example() { return 'sample'; }")
parsedData, err := session.ParseScriptData(ctx, scriptBody)
if err != nil {
    fmt.Println("Error parsing script data:", err)
    panic(err)
}
fmt.Println("Parsed Data:", parsedData)
```

### Generate Sensor Data

```go

payload := &roolink.SensorPayload{
    URL:        protectedURL,
    UserAgent:  session.UserAgent,
    Abck:       "replace _abck cookie",
    BmSz:       "replace bm_sz cookie",
    ScriptData: parsedData,
    SecCpt:     false,
    Stepper:    false,
    Index:      0,
    Flags:      "sample flags",
}
sensorData, err := session.GenerateSensorData(ctx, payload)
if err != nil {
    fmt.Println("Error generating sensor data:", err)
    panic(err)
}
fmt.Println("Sensor Data:", sensorData.Sensor)
```

### Generate SBSD Body
```go
sbsd_payload := &roolink.SbsdPayload{
    UserAgent: session.UserAgent,
    Vid:       "sample_vid",
    Cookie:    "sample_cookie",
    Static:    true,
}
sbsdBody, err := session.GenerateSbsdBody(ctx, sbsd_payload)
if err != nil {
    fmt.Println("Error generating SBSD body:", err)
    panic(err)
}
fmt.Println("SBSD Body:", sbsdBody.Body)
```

### Generate Pixel Data
- get `Bazadebezolkohpepadr` value with `roolink.GetBazadebezolkohpepadr(protectedPageHtml)`

```go
pixelPayload := &roolink.PixelPayload{
    UserAgent:            session.UserAgent,
    Bazadebezolkohpepadr: 12345,
    Hash:                 "sample_hash",
}
pixelData, err := session.GeneratePixelData(ctx, pixelPayload)
if err != nil {
    fmt.Println("Error generating pixel data:", err)
    panic(err)
}
fmt.Println("Pixel Data:", pixelData)
```

### Solve sec-cpt Challenges
```go
cptPayload := &roolink.CptChallenge{
    Token:      "sample_token",
    Timestamp:  1234567890,
    Nonce:      "sample_nonce",
    Difficulty: 3,
    Cookie:     "sample_cookie",
}
secCptAnswers, err := session.GenerateSecCptAnswers(ctx, cptPayload)
if err != nil {
    fmt.Println("Error solving sec-cpt challenge:", err)
    panic(err)
}
fmt.Println("Sec-CPT Answers:", secCptAnswers)
```