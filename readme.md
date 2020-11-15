# Vanish

<br />
<p align="center">
  <a href="https://github.com/kurojs/vanish">
    <img src="images/vanish.jpg" alt="Logo" width="125px">
  </a>

  <p  align="center">Vanish is an Golang package to remove JSON fields without struct</p>
</p>
<br />

## Table of Contents

- [About the Project](#about-the-project)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [License](#license)
- [Contact](#contact)

## About The Project

Vanish is what you need when wanting remove some unwanted fields in JSON strings with Golang.

Here's why:

- No need of define a struct to marshal/unmarshal JSON and remove theme. For example:
  To remove `metadata` field, we need to define an struct and do Unmarshal/Marshal in Go like

```go
struct Person {
    Name string `json:"name,omitempty"`
}

person := Person{}
err := json.Unmarshal([]byte(`
    {
        "name": "Here my name",
        "metadata" : {
            "nested_string": "abc"
        },
    }
`), &person)
if err != nil {
    // handle error
}

wantedJSON, err := json.Marshal(&person)
...
```

- You don't know exactly what type of the field is holding
- You tell me

An this is not the silver bullet for your problem, use with carefull.

## Getting Started

Getting it via Go `get`

```sh
go get github.com/kurojs/vanish
```

## Usage

To using vanish, just import it an define a list of fields need to be removed

```go
jsonData := `
    {
        "name": "Here my name",
        "metadata" : {
            "nested_string": "abc"
        },
        "number": 100
    }
`
tobeRemoved := []string{"metadata.nested_string", "number"}

wantedJSON, err := vanish.RemoveFields([]byte(jsonData), tobeRemoved)
if err != nil {
    // Handle error
}

fmt.Println(string(wantedJSON))
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

kuro (Chi Pham) - [via Email](thanhchi.fit.hcmus@gmail.com)

Project Link: [https://github.com/kurojs/vanish](https://github.com/kurojs/vanish)
