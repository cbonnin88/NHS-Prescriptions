package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
)

type PrescriptionData struct {
	Year_Month	int `json:"YEAR_MONTH"`
	Regional_Office_Name	string `json:"REGIONAL_OFFICE_NAME"`
	Regional_Office_Code	string	`json:"REGIONAL_OFFICE_CODE"`
	Icb_Name	string `json:"ICB_NAME"`
	Icb_Code	string `json:"ICB_CODE"`
	Pco_Name	string `json:"PCO_NAME"`
	Pco_Code	string `json:"PCO_CODE"`
	Practice_Name	string `json:"PRACTICE_NAME"`
	Practice_Code	string `json:"PRACTICE_CODE"`
	Address_1	string `json:"ADDRESS_1"`
	Address_2	string `json:"ADDRESS_2"`
	Address_3	string `json:"aADDRESS_3"`
	Address_4	string `json:"ADDRESS_4"`
	Postcode	string	`json:"POSTCODE"`
	Bnf_Chemical_Substance string `json:"BNF_CHEMICAL_SUBSTANCE"`
	Chemical_Substance_Bnf_Descur string `json:"CHEMICAL_SUBSTANCE_BNF_DESCUR"`
	Bnf_Code	string `json:"BNF_CODE"`
	Bnf_Description	string `json:"BNF_DESCRIPTION"`
	Bnf_Chapter_plus_code	string `json:"BNF_CHAPTER_PLUS_CODE"`
	Quantity float64 `json:"QUANTITY"`
	Items 	 int64 `json:"ITEMS"`
	Total_Quantity float64 `json:"TOTAL_QUANTITY"`
	Adqusage float64 `json:"ADQUSAGE"`
	NIC float64 `json:"NIC"`
	Actual_Cost float64 `json:"ACTUAL_COST"`
	Unidentified bool `json:"UNIDENTIFIED"`
}

type CKANAPIResponse struct {
	Help string `json:"help"`
	Success bool `json:"success"`
	Result struct {
		Records []PrescriptionData `json:"records"`
	} `json:"result"`
}

type StructuredPrescriptionData struct {
	Year  int
	Region string
	PracticeCode string
	PracticeName string
	BNFCode string
	BNFName string
	Items int
	NIC float64
	ACTCost float64
}

type BigQueryRow struct {
	Year int `bigquery:"year"`
	Region string `bigquery:"region"`
	PracticeCode string `bigquery:"practice_code"`
	PracticeName string `bigquery:"practice_name"`
	BNFCode string `bigquery:"bnf_code"`
	BNFName string `bigquery:"bnf_name"`
	Items int `bigquery:"items"`
	NIC float64 `bigquery:"nic"`
	ACTCost float64 `bigquery:"act_cost"`
	Timestamp time.Time `bigquery:"timestamp"`
}

func fetchNHSData(apiURL, resourceID string, limit int)([]PrescriptionData,error){
	requestBody, err := json.Marshal(map[string]interface{}{
		"resource_id": resourceID,
		"limit": limit,
})
if err != nil {
	return nil, fmt.Errorf("failed to marshal request body: %w", err)
}

req, err := http.NewRequest("POST", apiURL+"datastore_search", bytes.NewBuffer(requestBody))
if err != nil {
	return nil, fmt.Errorf("failed to create HTTP request: %w", err)
}
req.Header.Set("Content-Type", "application/json")

client := &http.Client{Timeout: 30 * time.Second} // Increased timeout for larger fetches
resp, err := client.Do(req)
if err != nil {
	return nil, fmt.Errorf("failed to make HTTP request to CKAN API: %w", err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return nil, fmt.Errorf("CKAN API returned non-OK status: %d %s, Body: %s", resp.StatusCode, resp.Status, string(bodyBytes))
}

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	return nil, fmt.Errorf("failed to read response body: %w", err)
}

var ckanResponse CKANAPIResponse
err = json.Unmarshal(body, &ckanResponse)
if err != nil {
	return nil, fmt.Errorf("failed to unmarshal JSON CKAN API response: %w", err)
}

if !ckanResponse.Success {
	return nil, fmt.Errorf("CKAN API call was not successful: %s", ckanResponse.Help)
}

	return ckanResponse.Result.Records, nil
}

func structureData(rawData []PrescriptionData)[]StructuredPrescriptionData{
	var structuredData []StructuredPrescriptionData
	for _, record := range rawData {
		record.Regional_Office_Name = strings.TrimSpace(record.Regional_Office_Name)
		record.Regional_Office_Code = strings.TrimSpace(record.Regional_Office_Code)
		record.Icb_Name = strings.TrimSpace(record.Icb_Name)
		record.Icb_Code = strings.TrimSpace(record.Icb_Code)
		record.Pco_Name = strings.TrimSpace(record.Pco_Name)
		record.Pco_Code = strings.TrimSpace(record.Pco_Code)
		record.Practice_Name = strings.TrimSpace(record.Practice_Name)
		record.Practice_Code = strings.TrimSpace(record.Practice_Code)
		record.Address_1 = strings.TrimSpace(record.Address_1)
		record.Address_2 = strings.TrimSpace(record.Address_2)
		record.Address_3 = strings.TrimSpace(record.Address_3)
		record.Address_4 = strings.TrimSpace(record.Address_4)
		record.Postcode = strings.TrimSpace(record.Postcode)
		record.Bnf_Chemical_Substance = strings.TrimSpace(record.Bnf_Chemical_Substance)
		record.Chemical_Substance_Bnf_Descur = strings.TrimSpace(record.Chemical_Substance_Bnf_Descur)
		record.Bnf_Code = strings.TrimSpace(record.Bnf_Code)
		record.Bnf_Description = strings.TrimSpace(record.Bnf_Description)
		record.Bnf_Chapter_plus_code = strings.TrimSpace(record.Bnf_Chapter_plus_code)


		structuredData = append(structuredData, StructuredPrescriptionData{
			Region: record.Regional_Office_Name,
			PracticeCode: record.Practice_Code,
			PracticeName: record.Practice_Name,
			BNFCode: record.Bnf_Code,
			BNFName: record.Bnf_Description,
			Items: int(record.Items),
			NIC: record.NIC,
			ACTCost: record.Actual_Cost,
		})
	}
	return structuredData

	}

func exportToCSV(data []StructuredPrescriptionData, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	
// Suggested code may be subject to a license. Learn more: ~LicenseLog:2040907952.
	header := []string{"Year","Regional_Office_Name", "Practice_Code", "Practice_Name", "BNF_Code", "BNF_Description", "Items", "NIC", "ACTCost",}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	
	for _, record := range data {
		row := []string{
			strconv.Itoa(int(record.Year)),
			record.Region,
			record.PracticeCode,
			record.PracticeName,
			record.BNFCode,
			record.BNFName,
			strconv.Itoa(record.Items),
			strconv.FormatFloat(record.NIC, 'f', 2, 64),
			strconv.FormatFloat(record.ACTCost, 'f', 2, 64),
			
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}
	return nil
}

func loadDataToBigQuery(projectID, datasetID, tableID string, data []StructuredPrescriptionData) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	inserter := client.Dataset(datasetID).Table(tableID).Inserter()

	var bigqueryRows []BigQueryRow
	ingestionTime := time.Now()
	for _, record := range data {
		bigqueryRows = append(bigqueryRows, BigQueryRow{
			Year:         record.Year,
			Region:		  record.Region,
			PracticeCode: record.PracticeCode,
			PracticeName: record.PracticeName,
			BNFCode:      record.BNFCode,
			BNFName:      record.BNFName,
			Items:        record.Items,
			NIC:          record.NIC,
			ACTCost:      record.ACTCost,
			Timestamp:    ingestionTime,
		})
	}

	// BigQuery allows batch insertion. Define the batch size.
	batchSize := 500
	for i := 0; i < len(bigqueryRows); i += batchSize {
		end := i + batchSize
		if end > len(bigqueryRows) {
			end = len(bigqueryRows)
		}
		batch := bigqueryRows[i:end]

		if err := inserter.Put(ctx, batch); err != nil {
			return fmt.Errorf("failed to insert batch into BigQuery: %w", err)
		}
		log.Printf("Successfully inserted %d rows into BigQuery.", len(batch))
	}
	return nil
}

func main() {
	projectID := "nhs-data-analysis"
	datasetID := "epd_data"
	tableID := "march_2025"
	ckanAPIURL := "https://opendata.nhsbsa.net/api/3/action/"
	nhsbsaResourceID := "EPD_202503"
	dataLimit := 100000
	outputJSONPath := "epd_prescription_data.csv"


	log.Printf("Fetching raw data from NHSBSA CKAN API using resource ID: %s", nhsbsaResourceID)
	prescriptionRecords, err := fetchNHSData(ckanAPIURL, nhsbsaResourceID, dataLimit)
	if err != nil {
		log.Fatalf("Error fetching NHS data: %v", err)
	}
	log.Printf("Fetched %d raw records from NHSBSA CKAN API.", len(prescriptionRecords))

	// 2. Perform initial data structuring
	log.Println("Structuring data...")
	structuredData := structureData(prescriptionRecords)
	log.Printf("Structured %d records.", len(structuredData))

	// 3. Export structured data to CSV for Python
	log.Printf("Exporting data to %s...", outputJSONPath)
	err = exportToCSV(structuredData, outputJSONPath)
	if err != nil {
		log.Fatalf("Error exporting data to CSV: %v", err)
	}
	log.Println("Data exported to CSV successfully!")

	// 4. Load structured data into BigQuery
	log.Println("Loading data into BigQuery...")
	err = loadDataToBigQuery(projectID, datasetID, tableID, structuredData)
	if err != nil {
		log.Fatalf("Error loading data to BigQuery: %v", err)
	}
	log.Println("Data loading to BigQuery completed successfully!")
}