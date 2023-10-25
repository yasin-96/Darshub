using System.Net;
using System.Xml.Schema;
using Azure.Identity;
using Azure.Storage.Blobs;
using Microsoft.Extensions.Logging;

namespace service.azure.BlobStorage;
public class BlobStorageService
{
    private readonly ILogger<BlobStorageService> _logger;
    private readonly BlobServiceClient _storageServiceClient;
    private readonly BlobContainerClient _storageContainerClient;
    
    public BlobStorageService(ILogger<BlobStorageService> logger)
    {
        _logger = logger;
        
        var containerName = ""; //TODO settings
        _storageServiceClient = new BlobServiceClient(
            new Uri("https://<storage-account-name>.blob.core.windows.net"),
            new DefaultAzureCredential()
        );
        
        _storageContainerClient =  _storageServiceClient.CreateBlobContainer(containerName);
        
    }

    public async Task<bool> IsCoursePdfAvailable(string courseId)
    {
        try
        {
            var blobClient = _storageContainerClient.GetBlobClient(courseId);
            var check = await blobClient.ExistsAsync();
            
            if (check)
            {
                return true;
            }
        }
        catch (Exception e)
        {
            _logger.LogError(e,"IsCoursePdfAvailable Exception: {emsg}", e.Message);
        }
        return false;
    }
   
    public async Task<bool> UploadCoursePdf(string courseId, byte[] pdfContent)
    {
        try
        {
            var blobClient = _storageContainerClient.GetBlobClient(courseId);
            var uploadResult = await blobClient.UploadAsync(new BinaryData(pdfContent));

            if (uploadResult is null)
            {
                return false;
            }
            
            if (!uploadResult.HasValue)
            {
                return false;
            }

            return true;
            //todo: infos ablegen? eigene collection mit id und dem aktuellen stand 
        }
        catch (Exception e)
        {
            _logger.LogError(e,"IsCoursePdfAvailable Exception: {emsg}", e.Message);
            return false;
        }
    }


    public async Task DownloadCoursePdf(string courseId)
    {
        // Download the blob's contents and save it to a file
        try
        {
            var blobClient = _storageContainerClient.GetBlobClient(courseId);
            var downloadResult = await blobClient.DownloadContentAsync();
            if (downloadResult is null)
            {
            }
            
            if (!downloadResult.HasValue)
            {
            } 
        }
        catch (Exception e)
        {
            _logger.LogError(e,"IsCoursePdfAvailable Exception: {emsg}", e.Message);
        }
    }
}