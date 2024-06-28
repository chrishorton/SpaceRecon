![SpaceRecon Logo](/assets/spacerecon-logo.png)

SpaceRecon is an Open-Source Intelligence (OSINT) tool for satellite reconnaissance, built with Golang. The tool can fetch satellite telemetry data, retrieve and parse TLE (Two-Line Element) data, and get the latest predicted satellite conjunctions.

## Features

- **Satellite Telemetry Retrieval**: Fetch the latest telemetry data for a given satellite.
- **TLE Data Retrieval**: Retrieve the latest TLE data for a specific satellite.
- **Orbital Predictions**: Generate orbital predictions based on TLE data.
- **Conjunction Data**: Fetch the latest satellite conjunctions to identify potential collisions.

## Installation

1. **Clone the repository**:

    ```sh
    git clone https://github.com/yourusername/spacerecon.git
    cd spacerecon
    ```

2. **Enter Credentials**:
    Create a `config.yaml` file and enter your credentials

    ```sh
    username: "<Space-Track username>"
    password: "<Space-Track password>"
    ```

3. **Build the project**:

    ```sh
    go build -o spacerecon
    ```

## Usage

### Fetch TLE Data

To retrieve the latest TLE data for a satellite:

```sh
./spacerecon --satellite 25544
```

This will fetch the TLE data for the satellite with NORAD ID 25544 (e.g., the International Space Station).

### Get Latest Conjunctions

To fetch the latest satellite conjunctions:

```sh
./spacerecon --conjunctions
```

This will retrieve the latest conjunction data, showing potential close approaches between satellites.

## Example Output

### TLE Data

```sh
./spacerecon --satellite 25544
```

Output:

```
TLE Data:
1 25544U 98067A   20334.54791667  .00000260  00000-0  12672-4 0  9993
2 25544  51.6442  21.0722 0001310 342.3151  44.0314 15.49197822253325
```

### Conjunction Data

```sh
./spacerecon --conjunctions
```

Output:

```
Latest Conjunctions:
Date: 2024-06-01 12:34:56, Satellite 1: ISS (ZARYA), Satellite 2: STARLINK-1500, Min Range: 0.05 km, Probability: 0.000123
Date: 2024-06-01 12:30:45, Satellite 1: NOAA 19, Satellite 2: METEOR 2-1, Min Range: 0.10 km, Probability: 0.000045
```

## API Integration

SpaceRecon uses the following APIs:

- **[Space-Track](https://www.space-track.org/)**: For retrieving satellite catalogs and TLE information.

- **[N2YO](https://www.n2yo.com/)**: For obtaining pass predictions (integration example included, but not detailed in this README).

## Development

### Prerequisites

- **Golang**: Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

### Project Structure

```
spacerecon/
│
├── main.go
├── fetch.go
├── tle_parser.go
├── conjunctions.go
└── go.mod
```

### Contributing

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Space-Track](https://www.space-track.org/)
