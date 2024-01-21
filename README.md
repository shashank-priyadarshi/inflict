# inflict

Inflation calculator, written in Go, calculates the current equivalent of an amount from a given year,
taking into account inflation.
It allows users to input an amount and the year it was earned, and returns the current value of that amount.

## Features

* Calculates the current equivalent of an amount from a given year, accounting for inflation
* Refers to the official inflation data of a country available on World Bank's [inflation database](https://www.worldbank.org/en/research/brief/inflation-database) 
* Provides a simple and easy-to-use interface
* It can be used to compare the value of money over time
* It can be used to make informed financial decisions

## Usage

To use the application, run the following command:

```bash
 inflict [inflation_type] [country] [value] [year]
```

where:

* `inflation_type` is one of the types available in the [inflation database](https://www.worldbank.org/en/research/brief/inflation-database)
* `country` is the country amount has been earned in
* `value` is the amount you want to calculate the current equivalent of
* `year` is the year the amount was earned

For example, to calculate the current equivalent of $100 earned in 2000, you would run the following command:

```bash
inflict "ccpi_a_e" "ICELAND" "2011" "100000"
```

The application would then output the current equivalent of $100 earned in 2000.

For more information on the available command line options, run the following command:

```bash
inflict -h
```

## Contributing

Contributions are welcome!

## License

This application is licensed under the MIT License.
