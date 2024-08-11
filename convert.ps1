# Load JSON data from the file
$jsonFilePath = "data/sop-documents.json"
$jsonData = Get-Content -Path $jsonFilePath -Raw | ConvertFrom-Json

# Initialize an empty dictionary
$dict = @{}

# Iterate through each item in the JSON array
foreach ($item in $jsonData) {
    # Use DocumentID as the key and the rest of the data as the value
    $dict[($item.DocumentMetadata.SOPNumber).ToLower()] = @{
        SOPTitle       = $item.DocumentMetadata.SOPTitle
        SOPNumber      = $item.DocumentMetadata.SOPNumber
        SOPVersion     = $item.DocumentMetadata.SOPVersion
        SOPDept        = $item.DocumentMetadata.SOPDept
        EffectiveDate  = $item.DocumentMetadata.EffectiveDate
        NextReviewDate = $item.DocumentMetadata.NextReviewDate
        DocumentURL    = $item.DocumentURL
        SearchIndex    = "$($item.DocumentMetadata.SOPNumber)-($($item.DocumentMetadata.SOPVersion)) $($item.DocumentMetadata.SOPTitle)"
    }
}

# Convert the dictionary to JSON for display or further use
$dictJson = $dict | ConvertTo-Json -Depth 10

# Output or save the result
$dictJson | Out-File -FilePath "data/sop-docs.json"
