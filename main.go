package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

type sheet_type string

const (
	// Headline consumer price inflation: This measures the overall change in the prices of a basket of goods and services that represents the average consumer's expenditure. It includes all items, such as food, fuel, housing, and other goods and services.
	// hcpi_m sheet_type = "hcpi_m"
	// hcpi_q sheet_type = "hcpi_q"
	// hcpi_a sheet_type = "hcpi_a"

	// Energy price inflation: This specifically measures the change in the prices of energy-related goods and services, such as gasoline, electricity, and heating oil.
	// fcpi_m sheet_type = "fcpi_m"
	// fcpi_q sheet_type = "fcpi_q"
	// fcpi_a sheet_type = "fcpi_a"

	// Food price inflation: This measures the change in the prices of food items, including groceries and dining out.
	// ecpi_m sheet_type = "ecpi_m"
	// ecpi_q sheet_type = "ecpi_q"
	// ecpi_a sheet_type = "ecpi_q"

	// Producer price inflation, annual: This measures the change in the prices received by domestic producers for their output. It reflects the prices of goods and services at the wholesale level before they reach the consumer.
	// ccpi_m sheet_type = "ccpi_m"
	// ccpi_q sheet_type = "ccpi_q"
	// ccpi_a sheet_type = "ccpi_a"

	// Official core consumer price inflation: This measures the change in the prices of consumer goods and services, excluding volatile items such as food and energy. It provides a more stable measure of inflation.
	// ppi_m sheet_type = "ppi_m"
	// ppi_q sheet_type = "ppi_q"
	// ppi_a sheet_type = "ppi_a"

	// GDP deflator growth rate: This measures the change in the prices of all new, domestically produced, final goods and services in an economy, providing an indication of the overall inflation within an economy.
	// def_q sheet_type = "def_q"
	// def_a sheet_type = "def_a"

	// Estimated core consumer price inflation: Similar to official core consumer price inflation, this measure excludes volatile items such as food and energy, but it may use different estimation methods.
	// ccpi_m_e sheet_type = "ccpi_m_e"
	// ccpi_q_e sheet_type = "ccpi_q_e"
	ccpi_a_e sheet_type = "ccpi_a_e"

	// Estimated transitory (cyclical) component of headline CPI inflation: This measures the temporary or cyclical factors affecting the headline consumer price inflation, such as changes in energy prices due to supply disruptions or changes in demand.
	// hcpi_q_t sheet_type = "hcpi_q_t"
	// hcpi_q_c sheet_type = "hcpi_q_c"

	// Aggregate annual average inflation, by inflation measures, country groups: This provides an average of the annual inflation rates across different measures for specific country groups, such as advanced economies or emerging markets, offering a comprehensive view of inflation across different economic contexts.
	// aggregate sheet_type = "aggregate"
)

type row_data struct {
	country_code, imf_country_code, indicator_type, series_name string
	time_series_data                                            map[string]string
}

var (
	excel_file_path = "./assets/data/Inflation-data.xlsx"      // os.Getenv("EXCEL_FILE_PATH")
	row_map         = make(map[sheet_type]map[string]row_data) // make(map[sheet_type]map[string]row_data): represents sheet type mapping that contains a country wise map
)

func main() {
	if err := parse(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data parsing completed successfully: ", row_map)

	var country string
	var amount int64

	fmt.Println(calculator(country, amount))
}

func calculator(country string, amount int64) (inflation_adjusted_equivalent int64) { return }

func parse() (err error) {
	f, err := excelize.OpenFile(excel_file_path)
	if err != nil {
		return fmt.Errorf("error while opening excel file from path %s: %w", excel_file_path, err)
	}

	defer func() (err error) {
		if err := f.Close(); err != nil {
			return fmt.Errorf("error while closing excel file form path %s: %w", excel_file_path, err)
		}
		return
	}()

	sheets := f.GetSheetList()

	for _, sheet := range sheets {

		// Only annual estimated core consumer price inflation data is parsed as of now, all other data is discarded
		if !strings.EqualFold(sheet, string(ccpi_a_e)) {
			continue
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Printf("error while reading sheet %s: \n%v\n", sheet, err)
			continue
		}

		country_wise_data := make(map[string]row_data)

		// Each row represents a country
		for i := 0; i < len(rows); i++ {

			row := rows[i]

			if len(row) < 4 {
				continue
			}

			country := string(row[2])

			country_data := row_data{
				country_code:     string(row[0]),
				imf_country_code: string(row[1]),
				indicator_type:   string(row[3]),
				series_name:      string(row[4]),
				time_series_data: make(map[string]string),
			}

			// Each column represents the fields like country code, imf country code, country name, indicator type, series name and time series inflation data for a country
			for key, colCell := range row {

				// The first 5 columns of a row contain column names, from where the years are picked, all other column names are discarded
				if key < 5 {
					continue
				}

				if i == 0 {
					country_data.time_series_data[colCell] = ""
					continue
				}
				country_data.time_series_data[rows[0][key]] = colCell
			}

			country_wise_data[country] = country_data
		}

		// Sheet type is added to map
		row_map[sheet_type(sheet)] = country_wise_data
	}

	return
}
