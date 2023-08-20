using Microsoft.AspNetCore.Mvc;
using pdfGen.services;

namespace api.Controllers;
[ApiController]
[Route("[controller]")]
public class PDFController : ControllerBase
{
    private readonly 
    private readonly ILogger<PDFController> _logger;

    public PDFController(ILogger<PDFController> logger)
    {
        _logger = logger;
    }

    [HttpGet(Name = "GetFullPDFById")]
    public IEnumerable<WeatherForecast> GetFullPDFByIdAsync([QueryParameter] string pdfId)
    {
        
    }
}
