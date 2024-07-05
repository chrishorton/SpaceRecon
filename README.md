![SpaceRecon Logo](/assets/spacerecon-logo.png)

SpaceRecon is an Open-Source Intelligence (OSINT) tool for satellite reconnaissance, built with Golang. The tool can retrieve and parse TLE (Two-Line Element) data, and get the latest satellite conjunctions.

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

    ```yaml
    username: "<Space-Track username>"
    password: "<Space-Track password>"
    ```

    You can also configure the amount of CDM's to fetch by adding
    ```yaml
    cdm_limit: <Max CDM's to be loaded>
    ```


3. **Build the project**:

    ```sh
    go build -o spacerecon
    ./spacerecon
    ```
    or
   ```sh
   go run main.go
   ```

## API Integration

SpaceRecon uses the following APIs:

- **[Space-Track](https://www.space-track.org/)**: For retrieving satellite catalogs and TLE information.

- **[Space Data Standards](https://spacedatastandards.org/#/standards)**: For serializing data into binary for efficient reading and communications
## Development

### Prerequisites

- **Golang**: Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

### Contributing

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Space-Track](https://www.space-track.org/)
- [Space Data Standards](https://spacedatastandards.org/#/standards)
