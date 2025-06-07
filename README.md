![nhs](https://github.com/user-attachments/assets/eef5fd89-4e61-4a04-9d07-8a86eaf8c51c)

# About the NHS
The National Health Service (NHS) is the term for the publicly funded healthcare systems of the United Kingdom: the National Health Service (England), NHS Scotland, NHS Wales, and Health and Social Care (Northern Ireland) which was created separately and is often referred to locally as "the NHS". This repository contains a beginner-level cloud data analyst project focused on collecting, processing, and visualizing public prescription data from the NHS.

# About This Project
This project contains a cloud data analyst project focused on collecting, processing, and visualizing public prescription data from the British NHS. It demonstrates a practical data pipeline leveraging a suite of modern cloud and programming technologies. The goal is to extract meaningful insights from NHS prescription data via the NHS API to understand the medicines that are most prescribed.

# Project Phases & Technologies Used
This project is structed into four main parts, utilizing various tools:

- Go: For efficient API interaction and initial data retrieval from the NHS Business Service Authority (NHSBBA Open Data Portal (CKAN API), along with data preparation for storage.
- Python (Pandas & Plotly): For data cleaning, transformation, and generating insightful, interactive visualizations
- Bigquery: As a scalable, serverless data warehouse to store and query large datasets
- Looker Studio: For creating professional, shareable, and interactive dashboard to present key findings
- Google Cloud Platform (GCP): Leveraging cloud services for infrastructure and data warehousing

# Dataset
This project used data fetched from the [NHS Business Services Authority - NHSBSA CKAN API](https://opendata.nhsbsa.net/dataset/english-prescribing-data-epd/resource/7e7196a0-8fc8-4539-8322-ebd0d6554463) containg various attributes pertaining to prescriptions in English

# Part 1: Fetching the Data via CKAN API
In this part, Go was used to fetch the data from the API and converted into a CSV, as well as being load into Bigquery via and API Key

# Key Activities
- Data Fetching: Created a main.go application to call and fetch the data directly from the NHSBSA Open Data Portal's API
- Created an Google API key in order to load the data into Bigquery
- Converted the Data into a csv file [epd_prescription_data](https://drive.google.com/file/d/1XSiw9UMQRuEgQ_P1yddr8lKBOiVVndjg/view?usp=drive_link)
- Data was also pushed into Bigquery via the go application

# Tools
- Firebase Studio: An online IDE (integreated development environment) developed by Google. It is based on VSCode and the Google Cloud infrastructure
- Go: A general purpose programming language that is statically typed and compiled. It was designed by Google and similar to C

# Part 2: Data Cleaning and Data Visualization with Python
- Data Loading: The epd_prescription_data.csv file was loaded into Pandas DataFrame.
- Data Cleaning & Transformation:
    Handling missing values (e.g Filling missing values with UNKNOWN or the mode)
    Correctied data type Conversion (eg, ACTCost to numeric)
    Standardizing column names

- Exploratory Data Analysis & Visualization (Plotly):

  - Top 10 BNF Names by Total Items Prescribed:
![Top 10 BNF Names](https://github.com/user-attachments/assets/14d1a1df-896c-4b7a-988c-c266734e30df)

 - Net Ingredient Cost vs Items Prescribed:
![NICvsItems](https://github.com/user-attachments/assets/07e9016c-58a8-4b9f-8826-2dfb7891360b)

 - Top 10 Practices by Total Net Ingredient Cost (NIC):
![Top10Practices](https://github.com/user-attachments/assets/87d37181-f020-4106-9fae-5d1b85a422ce)

 - Top 10 Practices by Total Actual Cost (ACTCost):
![totalActualCost](https://github.com/user-attachments/assets/1af570ed-2d22-4f82-a9e4-914c1b932f2c)

 - Average Net Ingredient Cost (NIC) by Region:
![AverageNetIngredient](https://github.com/user-attachments/assets/a6d76ce7-f7e1-433e-82bd-fe49722875e0)

 - Actual Total Cost by Region:
![ActualTotalCostbyRegion](https://github.com/user-attachments/assets/612fa71c-0f00-4d02-ad96-1f7e1c799914)

 - Actual Cost (ACTCost) vs Net Ingredient Cost (NIC):
![actvsnic](https://github.com/user-attachments/assets/054610bf-58da-4034-bdf6-f86e516112c4)

# Tools:
- Pandas: A powerful, open source data analysis and manipulation tool
- Plotly Express: A Python library for creating interactive, publication-quality statistical graphics

# Part 3: Data Management & Analysis with Bigquery
This part involved manipulating the data via Bigquery and performing more analytics queries

# Key Activities:
 - Database setup: Created a projected called 'nhs-data-analysis' within Google Cloud Platform and a dataset called 'epd_data'
 - Created a Bigquery Table called 'march_2025'
 - SQL Analysis: Executed various SQL queries to gain insights:
     - Counting the total amount of records
     - Total Net Ingredient Cost (NIC) by month and year
     - Top 10 BNF Names by Total Items Prescribed
     - Average actual cost per item by practice
     - Calculatinig the percentage of total tiems for each BNF Name
     - Finding practices with an unusually high average actual cost per item(eg. above national average and with significant volume)
     - Rank BNF Names by their total Net Ingredient Cost (NIC) within a specific year and month
     - Identify practices that have prescribed a particular "high-cost" BNF
     - Calculate Average Actual Cost (ACTCost) and Net Ingredient Cost (NIC) per Practice
     - Finding BNF items with the largest difference between Actual Cost (ACTCost) and Net Ingredient Cost (NIC) on average
     - Ranking BNF items by Actual Cost within earch practice and show the previous ranked item's ACTCost
  
# Tools
- Google Cloud Platform: Cloud Computing Services provided by Google
- Bigquery: A fully managed, serverless enterprise data warehouse offered by Google Cloud Platform

# Part 4: [Data Visualization with Looker Studio](https://lookerstudio.google.com/reporting/8f148f69-6d6f-477e-987b-dee927e9a3d8)
The Final part of the NHS Project involved connecting Bigquery to Looker Studio (via Bigquery API) to create an interactive and visually appealing dashboard for NHS use.

# Key Activites
 - Data Source Connection: Established via an Google API key between Bigquery database and Looker Studio
 - Configuration: Adjusted the month data type from location to month
 - Dashboard Creation: Designed and built a comprehensive NHS Healthcare Dashboard featuring:
   
    - Pie Chart in order to see the percentage of Total NIC by Top BNF Names
    - Bar Chart to provide insights on the most frequently prescribed items
    - Viewing detailed prescription data for specific practices or items via a Table
    - A Scorecard to display total Net Ingredient Cost across all data
  
# Tools
- Looker Studio: A free, web-based data visualization and BI tool

# Potential Improvements based on insights
1. Cost Optimization for High-Value Prescriptions:
 - Improvement: Investigate why these items are so costly. Are there cheaper generic alternatives available? Can the NHS negotiate better pricing with suppliers for these high-cost, high-impact items? This insight could lead to targeted procurement strategies.

2. Supply Chain and Dispensing Efficiency:
 - Improvement: Optimize the supply chain for these high-volume items to ensure consistent availability and efficient distribution. Pharmacies could streamline their dispensing processes for these common prescriptions, potentially reducing wait times and operational costs.

3. Targeted Practice-Level Interventions:
 - Improvement: Collaborate with these outlier practices to understand their prescribing patterns. This could involve reviewing clinical guidelines, providing educational resources, or identifying specific patient population needs that contribute to higher costs. Benchmarking between practices could also encourage the adoption of best practices.

4. Policy and Clinical Guideline Review:
 - Improvement: By combining insights from the top cost items and practices, you might identify patterns that suggest a need to review or update existing NHS prescribing policies or clinical guidelines. For instance, if a particularly expensive drug is being widely prescribed where a more cost-effective alternative exists, this could inform policy changes.

# Conclusion
This project successfully demonstrates the end-to-end development of a cloud-based data analytics pipeline, designed to extract meaningful insights from public British NHS prescription data. It provides a robust framework for beginners to gain practical experience with a comprehensive suite of modern data tools and technologies.

# Key Achievements:
 - Diverse Technology Stack Integration: We've effectively integrated Google Cloud Platform (GCP) services with Golang for efficient data fetching, Python (Pandas and Plotly) for sophisticated data cleaning and interactive visualization, Google BigQuery for scalable data warehousing, and Looker Studio for dynamic business intelligence dashboards.

- Automated Data Acquisition: The Golang application efficiently connects to the NHS Business Services Authority (NHSBSA) Open Data Portal's CKAN API, showcasing how to programmatically ingest real-world data and prepare it for further processing.

- Robust Data Cleaning & Local Analysis: The Python script, leveraging Pandas, provided critical data cleaning capabilities, including handling missing values via mode imputation and ensuring data type consistency. The integration of Plotly allowed for the generation of interactive visualizations, offering immediate insights into the cleaned dataset.

- Scalable Data Storage & Querying: By centralizing the structured and cleaned data in BigQuery, the project establishes a foundation for performing complex SQL queries on large datasets, enabling deep analytical exploration.

- Actionable Visualizations & Reporting: The Looker Studio dashboard (with a focus on cost distribution by BNF name, top prescribed items, and practice-level metrics) transforms raw data into understandable and actionable insights, suitable for stakeholder reporting and decision-making.

# Value Proposition:
This project is a testament to how a structured data pipeline can uncover valuable patterns and potential areas for improvement within public health data. The insights gained, such as identifying high-cost prescription categories, top prescribed items, or practices with higher average costs, are crucial for:

- Informing Policy: Guiding decisions related to drug procurement, pricing negotiations, and clinical guideline adjustments.
- Optimizing Resources: Highlighting opportunities for cost efficiencies in prescription and dispensing processes.
- Improving Patient Care: Potentially leading to better management of high-volume or high-cost medications.

This project serves as a strong foundation for me to mix my data analytics skills with cloud computing and my aspiration to work as a Cloud Data Analyst. I believe that Cloud Analytics will become more and more important as wel move into AI and more virtual means of storing data so this project, along with others to come, will help me to build a foundation in Cloud Computing as well as data analytics.



