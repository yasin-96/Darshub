using System.ComponentModel.DataAnnotations;
using LanguageExt.SomeHelp;
using Microsoft.AspNetCore.Mvc;
using pdfGen.service.pdf;
using service.mongo;

namespace api.Controllers;
[ApiController]
[Route("[controller]")]
public class PDFController : ControllerBase
{
    private readonly ILogger<PDFController> _logger;
    private readonly MongoService _dbService;
    private readonly PdfService _pdfService;

    public PDFController(ILogger<PDFController> logger, MongoService dbService,
        PdfService pdfService)
    {
        _logger = logger;
        _dbService = dbService;
        _pdfService = pdfService;
    }

    [HttpGet(Name = "GetPdf")]
    public async Task<IActionResult> GetFullPDFByIdAsync(
        [FromQuery] [StringLength(24, ErrorMessage = "Id is invalid")]
        string id)
    {
        try
        {
            var checkContent = await _dbService.IsCourseAvailable(id);
            if (!checkContent)
            {
                return NotFound("Course not exist");
            }

            var dbResponse = await _dbService.GetCourseById(id);
            if (dbResponse is null)
            {
                return Conflict("Some goes wrong");
            }

            _pdfService.CreatePdf(dbResponse);

            return Ok(dbResponse.ToSome().Value);
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return Conflict("Some goes wrong");
        }

        return BadRequest("Unexpected return");
    }
}