namespace pdfGen.services;
public class PDFService
{
    private readonly ILogger<PDFService> _log;

    public PDFService(ILogger<PDFService> logger)
    {
        _log = logger;
    }

    public iAction CreatePdfById(string pdfId)
    {
        if(string.isNullOrEmpty(pdfId)){
            return BadRequest()
        }
    }
}